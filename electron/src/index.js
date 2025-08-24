const {app, BrowserWindow, ipcMain} = require('electron');
const path = require('node:path');

if (require('electron-squirrel-startup')) {
    app.quit();
}

const isDev = !app.isPackaged
const localServer = "http://127.0.0.1";
const kernelPort = 31943;

const getServer = (port = kernelPort) => {
    return localServer + ":" + port;
};

const createWindow = () => {
    const mainWindow = new BrowserWindow({
        width: 1600, height: 800, frame: false, webPreferences: {
            nodeIntegration: false, contextIsolation: true, preload: path.join(__dirname, 'preload.js'),
        },
    });

    if (isDev) {
        mainWindow.loadURL('http://localhost:5173/static/index.html');
        mainWindow.webContents.openDevTools();
    } else {
        mainWindow.loadURL(getServer() + '/static/index.html');
    }

    ipcMain.on('window-control', (event, command) => {
        switch (command) {
            case 'minimize':
                mainWindow.minimize();
                break;
            case 'maximize':
                mainWindow.isMaximized() ? mainWindow.unmaximize() : mainWindow.maximize();
                break;
            case 'close':
                mainWindow.close();
                break;
        }
    });
};

app.whenReady().then(() => {
    createWindow();

    app.on('activate', () => {
        if (BrowserWindow.getAllWindows().length === 0) {
            createWindow();
        }
    });
});

app.on('window-all-closed', () => {
    if (process.platform !== 'darwin') {
        app.quit();
    }
});
