const {ipcMain} = require("electron");
const execute = require("child_process").execFile;
const path = require("path");
const fs = require("fs");
const grpc = require("grpc");
const PROTO_BUFFERS = path.join(__dirname, "../../sparkgrpc/commands.proto");
const revCommands = grpc.load(PROTO_BUFFERS).sparkgrpc;
const client = new revCommands.sparkusb('localhost:8001', grpc.credentials.createInsecure());

let usbPID = -1;

ipcMain.on("start-server", (event) => {
    const exePath = path.join(__dirname, "../sparkusb/sparkusb.exe");
    if (fs.existsSync(exePath)) {
        try {
            usbPID = execute(exePath, ["-r"], (error, data) => {
                if (error) {
                    event.sender.send("start-server-error", "There was en error while trying to start the sparkusb TCP/IP server: " + error);
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

ipcMain.on("kill-server", (event) => {
    process.kill(usbPID);
});

ipcMain.on("test", (event) => {
    client.connect({device: "COM5"}, (err, response) => {
        event.sender.send("test-response", err, response);
    });
});

ipcMain.on("set-can-id", (event, id) => {
    client.setParameter({value : id, parameter: 0}, (err, response) => {
        event.sender.send("set-can-id-response", err, response);
    });
});

ipcMain.on("get-can-id", (event, parameter) => {
    client.getParameter({value: parameter, parameter: 0}, (err, response) => {
        event.sender.send("get-can-id-response", err, response);
    });
});

ipcMain.on("list-devices", (event) => {
    client.list({all: true}, (err, response) => {
        event.sender.send("list-devices-response", err, response);
    });
});