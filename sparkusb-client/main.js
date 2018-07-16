const {app, BrowserWindow, ipcMain} = require("electron");

require("electron-debug")({showDevTools: true, enabled: true});

let win;

function createWindow () {
    // Create the browser window.
    win = new BrowserWindow({width: 480, height: 500, show: false});

    // and load the index.html of the app.
    win.loadURL("http://localhost:3000/");

    win.webContents.on("did-finish-load", () => {
       win.show();
    });

    require("./main/sparkusb-communication");

    win.on('closed', () => {
        setTimeout(() => {
            ipcMain.send("kill-server");
        }, 500);
        win = null;
    });
}

app.on('ready', createWindow);

// Quit when all windows are closed.
app.on('window-all-closed', () => {
    if (process.platform !== 'darwin') {
        app.quit()
    }
});

app.on('activate', () => {
    if (win === null) {
        createWindow();
    }
});
