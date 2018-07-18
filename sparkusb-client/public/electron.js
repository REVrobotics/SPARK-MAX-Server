const {app, BrowserWindow, ipcMain} = require("electron");
const url = require("url");
const path = require("path");

const production = false;

if (!production) {
  require("electron-debug")({showDevTools: true, enabled: true});
}

let win;

function createWindow () {
  // Create the browser window.
  win = new BrowserWindow({width: 600, height: 550, show: false, resizable: false, icon: "./favicon.ico"});

  if (production) {
    win.loadURL(url.format({
      pathname: path.join(__dirname, "./index.html"),
      protocol: "file:",
      slashes: true
    }));
  } else {
    win.loadURL("http://localhost:3000/");
  }

  win.webContents.on("did-finish-load", () => {
    win.show();
  });

  win.setMenu(null);

  if (production) {
    require("./main/sparkusb-communication");
  } else {
    require("../main/sparkusb-communication");
  }

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
