const {ipcMain, dialog, BrowserWindow} = require("electron");
const execute = require("child_process").execFile;
const path = require("path");
const fs = require("fs");
//const grpc = require("grpc");
const PROTO_BUFFERS = path.join(__dirname, "../sparkusb/commands.proto");
//const revCommands = grpc.load(PROTO_BUFFERS).sparkgrpc;
//const client = new revCommands.sparkusb('localhost:8001', grpc.credentials.createInsecure());
const protobuf = require("protobufjs");
const zmq = require('zeromq');
const SerialPort = require("serialport");
const Queue = require('better-queue');

let usbPID = -1;
let heartbeatID = -1;
let connCheckID = -1;
let setpoint = 0;
let isWin = process.platform === "win32";
let isEnabled = false

ipcMain.on("start-server", (event) => {
    const relPath = "../sparkusb/sparkusb" + (isWin ? ".exe" : "");
    const exePath = path.join(__dirname, relPath);
    if (fs.existsSync(exePath)) {
        try {
            usbPID = execute(exePath, ["-r"], (error, data) => {
                if (error) {
                    event.sender.send("start-server-error", "There was en error while running the sparkusb TCP/IP server: " + error);
                } else {
                    event.sender.send("start-server-success");
                }
            });
            event.sender.send("start-server-success");
        } catch (e) {
            event.sender.send("start-server-error", "There was en error while trying to execute the sparkusb binary. " + e);
        }
    } else {
        event.sender.send("start-server-error", "The sparkusb executable was not found.");
    }
});

ipcMain.on("kill-server", () => {
    process.kill(usbPID);
});

ipcMain.on("connect", (event, device) => {
    console.log("Attempting to connect on " + device + "...");
    client.connect({device: device, keepalive: true}, (err, response) => {
      if (connCheckID === -1) {
        connCheckID = setInterval(() => {
          SerialPort.list().then((ports) => {
            let found = false;
            for (let port of ports) {
              if (port.comName === device) {
                found = true;
              }
            }
            if (!found) {
              console.log("Disconnection on " + device + ".");
              event.sender.send("disconnection", device);
              clearInterval(connCheckID);
              connCheckID = -1;
            }
          });
        }, 1000);
      }
      event.sender.send("connect-response", err, response);
    });
});

ipcMain.on("disconnect", (event, device) => {
  console.log("Disconnecting on " + device + "...");
  client.disconnect({device: device, keepalive: true}, (err, response) => {
    if (connCheckID !== -1) {
      clearInterval(connCheckID);
      connCheckID = -1;
    }
    event.sender.send("disconnect-response", err, response);
  });
});

ipcMain.on("set-param", (event, parameter, value) => {
    client.setParameter({value: value, parameter: parameter}, (err, response) => {
        setTimeout(() => {
          event.sender.send("set-param-response", err, response);
        });
    });
});

ipcMain.on("get-param", (event, parameter) => {
    client.getParameter({parameter: parameter}, (err, response) => {
        setTimeout(() => {
          event.sender.send("get-param-response", err, response);
        }, 50);
    });
});

ipcMain.on("list-devices", (event) => {
    client.list({all: true}, (err, response) => {
        event.sender.send("list-devices-response", err, response);
    });
});

ipcMain.on("enable-heartbeat", (event, interval) => {
  isEnabled = true
    if (heartbeatID === -1) {
      console.log("Enabling heartbeat for every " + interval + "ms");
      heartbeatID = setInterval(() => {
        client.setpoint({setpoint: setpoint}, (err, response) => {

        });
      }, interval);
    }
});

ipcMain.on("disable-heartbeat", (event) => {
  isEnabled = false
   if (heartbeatID !== -1) {
     console.log("Disabling heartbeat");
     clearInterval(heartbeatID);
     heartbeatID = -1;
     err = ""
     response = ""
     event.sender.send("disable-heartbeat-response", err, response);
   }
});

ipcMain.on("set-setpoint", (event, newSetpoint) => {
    setpoint = newSetpoint;
});

ipcMain.on("save-config", (event, device) => {
  console.log("Saving configuration to " + device + "...");
  client.burnFlash({device: device}, (error, response) => {
    setTimeout(() => {
      event.sender.send("save-config-response", error, response);
    });
  });
});

ipcMain.on("request-firmware", (event) => {
    dialog.showOpenDialog(BrowserWindow.getFocusedWindow(), {
      title: "Firmware Loading",
      filters: [{name: "Firmware Files (*.bin)", extensions: ["bin"]}],
      properties: ["openFile"]
    }, (filePaths) => {
      if (filePaths && filePaths[0]) {
        event.sender.send("request-firmware-response", filePaths[0]);
      }
    });
});

class sparkusb {
  constructor(port){
    this.port = port;
    //sock.bindSync('tcp://localhost:' + port);
    this.sock = zmq.socket('req');
    //this.sock.bindSync('tcp://127.0.0.1:' + port);
    this.sock.connect('tcp://127.0.0.1:' + port);
    console.log('Producer bound to port ' + port);
    this.root;

    let self = this;

    //Commands run one at a time in priority order
    //Leave default of 1 task concurrently
    this.cmdQueue = new Queue(function (input, cb) {
      //console.log("Running on queue item:");
      //console.log(input);

      // Init message loads pb befor any other call
      if (input.id === "init") {
        protobuf.load(PROTO_BUFFERS)
          .then(function(root) {
            self.root = root;
            cb(null,null);
          });
      } else {
        //All calls here should have a 'msg' field set with a message that
        //can be part of the 'Oneof' file of the 'RequestWire'        
        let wire = self.root.lookupType("sparkusb.RequestWire");
        let wireMsg = wire.create({req: input.id});
        wireMsg[input.id] = input.msg;
        let wireBuf = wire.encode(wireMsg).finish();

        //console.log(wireMsg);
        //console.log(wireBuf);

        self.sock.send(wireBuf);
        
        // Send a message and wait for response before triggering callback
        self.sock.on('message', function(msg) {
          cb(null, msg);
        });
      }
    }, {
      priority: function (input, cb) {
        if (input.id === "init") return cb(null,100);
        if (input.id === "control") return cb(null, 10);
        if (input.id === "setpoint") return cb(null, 5);
        if (input.id === "heartbeat") return cb(null, 5);
        cb(null, 1);
      }
    })

    //First task is to initialize the queue
    this.cmdQueue.push({id: "init"})
  }

  sendCommand(lookupType,msg,cb) {
    /*Queue the attached request
    * Priority:
    *   - Connect (also flush queue)
    *   - Disconnect (also flush queue)
    *     - Setpoint (named so only 1 is ever in the queue)
    *     - Heartbeat (named so only 1 is ever in the queue)
    *       - Get Param
    *       - Set Param
    *       - Burn Param
    *       - List
    *       - Burn Flash
    *       - All others   
    * 
    * All commands will be part of the requestWire
    * and need to be encoded as such
    */
    let req = {id: lookupType};
    let self = this;

    if (lookupType === "setpoint" || lookupType === "heartbeat") {
          req.count = 1;
    }

    req.msg = msg;
    this.cmdQueue.push(req, function (err, result) {
      // Decode message
      let cmd = self.root.lookupType("sparkusb.ResponseWire");
      let message = cmd.decode(result);
      //console.log(message[lookupType]);
      cb(err,message[lookupType]);
    });
  }

  // Convenience functions for each message type

//Connect(rootCommand) returns (rootResponse) {}
//Disconnect(rootCommand) returns (rootResponse) {}
//List(listRequest) returns (listResponse) {}
// NOTIMP Firmware(firmwareRequest) returns (firmwareResponse) {}
// NOTIMP Heartbeat(heartbeat) returns (rootResponse) {}
// NOTIMP Address(addressRequest) returns (addressResponse) {}
//SetParameter(parameterRequest) returns (parameterResponse) {}
//GetParameter(parameterRequest) returns (parameterResponse) {}
//BurnFlash(rootCommand) returns (rootResponse) {}
// NOTIMP ListParameters(parameterListRequest) returns (parameterListResponse) {}
//Setpoint(setpointRequest) returns (setpointResponse) {}

  connect(controlCommand,cb) {
    controlCommand.ctrl = 1;
    this.sendCommand("control",controlCommand,cb)
  }
  disconnect(controlCommand,cb) {    
    controlCommand.ctrl = 2;
    this.sendCommand("control",controlCommand,cb)
  }
  list(listCommand,cb) {
    this.sendCommand("list",listCommand,cb)
  }
  getParameter(paramCommand,cb) {
    this.sendCommand("parameter",paramCommand, function (err, result){
      result.value = Number(result.value);
      cb(err,result);
    })
  }
  setParameter(paramCommand,cb) {
    //Make sure 'paramCommand' is a string
    paramCommand.value += '';
    this.sendCommand("parameter",paramCommand,cb)
  }
  setpoint(setpointCommand,cb) {
    setpointCommand.enable = isEnabled;
    setpointCommand.setpoint = setpointCommand.setpoint / 1024;
    this.sendCommand("setpoint",setpointCommand,cb)
  }
  burnFlash(burnCommand,cb) {
    burnCommand.verify = true
    this.sendCommand("burn",burnCommand,cb)
    //cb(null,null);
  }
  heartbeat(heartbeatRequest,cb) {
    //this.sendCommand("heartbeat",heartbeatRequest,cb)
    cb(null,null)
  }
}

const client = new sparkusb(8001);
