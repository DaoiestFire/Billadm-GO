const {app, BrowserWindow} = require('electron');
const path = require('node:path');

if (require('electron-squirrel-startup')) {
    app.quit();
}

const localServer = "http://127.0.0.1";
const kernelPort = 31943;

const getServer = (port = kernelPort) => {
    return localServer + ":" + port;
};

const createWindow = () => {
    const mainWindow = new BrowserWindow({
        width: 1600,
        height: 800,
        frame: false,
        webPreferences: {
            nodeIntegration: false,
            contextIsolation: true,
            preload: path.join(__dirname, 'preload.js'),
        },
    });

    // mainWindow.loadURL(getServer() + '/static/index.html');
    mainWindow.loadURL('http://localhost:5173/static/index.html');
};

app.whenReady().then(() => {
    createWindow();

    app.on('activate', () => {
        if (BrowserWindow.getAllWindows().length === 0) {
            createWindow();
        }
    });
});

// Quit when all windows are closed, except on macOS. There, it's common
// for applications and their menu bar to stay active until the user quits
// explicitly with Cmd + Q.
app.on('window-all-closed', () => {
    if (process.platform !== 'darwin') {
        app.quit();
    }
});
