const {app, BrowserWindow} = require("electron");

let win;

function createWindow () {
    // Create the browser window.
    win = new BrowserWindow({width: 320, height: 500, show: false});

    // and load the index.html of the app.
    win.loadURL("http://localhost:3000/");

    win.webContents.on("did-finish-load", () => {
       win.show();
    });

    // Open the DevTools.
    // win.webContents.openDevTools();

    win.on('closed', () => {
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
