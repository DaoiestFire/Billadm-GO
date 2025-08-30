const {app, BrowserWindow, ipcMain} = require('electron');
const path = require('node:path');
const {spawn} = require('child_process');


const isDev = !app.isPackaged
const localServer = "http://127.0.0.1";
const kernelPort = 31943;

const getServer = (port = kernelPort) => {
    return localServer + ":" + port;
};

let kernelProcess = null;

const startKernel = () => {
    const kernelExe = path.join(app.getAppPath(), 'Billadm-Kernel.exe');
    console.log('Starting kernel:', kernelExe);

    kernelProcess = spawn(kernelExe, ['-port', kernelPort, '-log_file', 'billadm.log', '-mode', 'release']);

    // 捕获标准输出
    kernelProcess.stdout.on('data', (data) => {
        console.log(`[Kernel STDOUT]: ${data.toString()}`);
    });

    // 捕获错误输出
    kernelProcess.stderr.on('data', (data) => {
        console.error(`[Kernel STDERR]: ${data.toString()}`);
    });

    // 进程关闭
    kernelProcess.on('close', (code) => {
        console.log(`[Kernel Process] exited with code ${code}`);
        kernelProcess = null;
        if (code !== 0) {
            console.warn(`Backend process exited unexpectedly code ${code}. `);
        }
    });

    kernelProcess.on('error', (err) => {
        console.error('[Kernel Process] Failed to start:', err);
    });
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
    startKernel();
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
