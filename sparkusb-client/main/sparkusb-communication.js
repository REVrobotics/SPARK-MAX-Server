const {ipcMain, dialog, BrowserWindow} = require("electron");
const execute = require("child_process").execFile;
const path = require("path");
const fs = require("fs");
const grpc = require("grpc");
const PROTO_BUFFERS = path.join(__dirname, "../sparkusb/commands.proto");
const revCommands = grpc.load(PROTO_BUFFERS).sparkgrpc;
const client = new revCommands.sparkusb('localhost:8001', grpc.credentials.createInsecure());

let usbPID = -1;

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
    client.connect({device: device}, (err, response) => {
        event.sender.send("connect-response", err, response);
    });
});

ipcMain.on("set-param", (event, parameter, value) => {
    console.log("SET PARAM");
    client.setParameter({value: value, parameter: parameter}, (err, response) => {
      console.log("SET PARAM RESPONSE");
        event.sender.send("set-param-response", err, response);
    });
});

ipcMain.on("get-param", (event, parameter) => {
    client.getParameter({parameter: parameter}, (err, response) => {
        event.sender.send("get-param-response", err, response);
    });
});

ipcMain.on("list-devices", (event) => {
    client.list({all: true}, (err, response) => {
        event.sender.send("list-devices-response", err, response);
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