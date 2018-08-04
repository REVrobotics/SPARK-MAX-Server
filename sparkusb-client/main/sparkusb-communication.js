const {ipcMain, dialog, BrowserWindow} = require("electron");
const execute = require("child_process").execFile;
const path = require("path");
const fs = require("fs");
const grpc = require("grpc");
const PROTO_BUFFERS = path.join(__dirname, "../sparkusb/commands.proto");
const revCommands = grpc.load(PROTO_BUFFERS).sparkgrpc;
const client = new revCommands.sparkusb('localhost:8001', grpc.credentials.createInsecure());
const SerialPort = require("serialport");

let usbPID = -1;
let heartbeatID = -1;
let connCheckID = -1;
let setpoint = 0;

ipcMain.on("start-server", (event) => {
    const exePath = path.join(__dirname, "../sparkusb/sparkusb.exe");
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
    if (heartbeatID === -1) {
      console.log("Enabling heartbeat for every " + interval + "ms");
      heartbeatID = setInterval(() => {
        client.heartbeat({enable: true}, (err, response) => {
            event.sender.send("enable-heartbeat-response", err, response);
        });
        client.setpoint({setpoint: setpoint}, (err, response) => {

        });
      }, interval);
    }
});

ipcMain.on("disable-heartbeat", (event) => {
   if (heartbeatID !== -1) {
     console.log("Disabling heartbeat");
     clearInterval(heartbeatID);
     client.heartbeat({enable: false}, (err, response) => {
       heartbeatID = -1;
       event.sender.send("disable-heartbeat-response", err, response);
     });
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