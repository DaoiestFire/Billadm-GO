const {app, BrowserWindow, ipcMain, dialog} = require('electron');
const path = require('path');
const fs = require('fs');
const os = require('os')
const {spawn} = require('child_process');

process.noAsar = false;

const isDev = !app.isPackaged
const localServer = isDev ? 'http://localhost' : 'http://127.0.0.1';
const kernelPort = isDev ? 31945 : 31943;
const urlPath = isDev ? '/static' : '/static/index.html';
const appPath = isDev ? path.dirname(__dirname) : app.getAppPath();

const getServer = (port = kernelPort) => {
    return localServer + ':' + port;
};


// 应用日志
const logDir = path.join(appPath, 'logs');
const logFile = path.join(logDir, 'app.log');

if (!fs.existsSync(logDir)) {
    fs.mkdirSync(logDir, {recursive: true});
}
const logStream = fs.createWriteStream(logFile, {flags: 'a'});
const log = (message) => {
    const time = new Date().toISOString();
    const logMessage = `[${time}] ${message}\n`;
    logStream.write(logMessage);
};

let billadmCfg = {
    width: 1600,
    height: 1000,
    workspaceDir: '',
}

function readBilladmFile() {
    const homeDir = os.homedir();
    const filePath = path.join(homeDir, '.billadm');
    let tmpObj;
    try {
        const fileContent = fs.readFileSync(filePath, 'utf8');
        tmpObj = JSON.parse(fileContent);
        billadmCfg = {
            ...billadmCfg,
            ...tmpObj,
        }
    } catch (err) {
        log(`读取 .billadm 文件时发生错误:', ${err.message}`);
    }

    log(`窗口宽度 ${billadmCfg.width} 窗口高度 ${billadmCfg.height} 工作空间路径 ${billadmCfg.workspaceDir}`)
}

function saveBilladmConfig() {
    const homeDir = os.homedir();
    const filePath = path.join(homeDir, '.billadm');

    try {
        if (typeof billadmCfg !== 'object' || billadmCfg === null) {
            log('billadmCfg 不是一个有效的对象，无法保存');
            return false;
        }

        const data = JSON.stringify(billadmCfg, null, 2);

        fs.writeFileSync(filePath, data, 'utf8');

        log(`配置已成功保存至: ${filePath}`);
    } catch (err) {
        log(`保存 .billadm 文件时发生错误 ${err.message}`);
    }
}


// 内核
let kernelProcess = null;

const startKernel = () => {
    const kernelExe = path.join(appPath, 'Billadm-Kernel.exe');
    log(`Starting kernel: ${kernelExe}`);

    kernelProcess = spawn(kernelExe, ['-port', kernelPort, '-mode', 'release', '-workspace', billadmCfg.workspaceDir]);

    // 捕获标准输出
    kernelProcess.stdout.on('data', (data) => {
        log(`[Kernel STDOUT]: ${data.toString()}`);
    });

    // 捕获错误输出
    kernelProcess.stderr.on('data', (data) => {
        log(`[Kernel STDERR]: ${data.toString()}`);
    });

    // 进程关闭
    kernelProcess.on('close', (code) => {
        log(`[Kernel Process] exited with code ${code}`);
        kernelProcess = null;
        if (code !== 0) {
            log(`Backend process exited unexpectedly code ${code}. `);
        }
    });

    kernelProcess.on('error', (err) => {
        log('[Kernel Process] Failed to start:', err);
    });
};

const createWindow = () => {
    const mainWindow = new BrowserWindow({
        width: billadmCfg.width,
        height: billadmCfg.height,
        frame: false,
        webPreferences: {
            nodeIntegration: false,
            contextIsolation: true,
            preload: path.join(__dirname, 'preload.js'),
        },
    });

    mainWindow.loadURL(getServer() + urlPath);

    if (isDev) {
        mainWindow.webContents.openDevTools();
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

    ipcMain.handle('dialog:open', async (event, options) => {
        try {
            return await dialog.showOpenDialog({
                properties: ['openFile'],
                ...options,
            });
        } catch (err) {
            log(`Dialog error: ${err.message}`);
            return {canceled: true, filePaths: [], error: err.message};
        }
    });

    ipcMain.on('workspace:set', (event, workspaceDir) => {
        log(`workspace:set ${workspaceDir}`);
        billadmCfg.workspaceDir = workspaceDir;
    });
};

app.whenReady().then(() => {
    readBilladmFile();
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
        if (kernelProcess && kernelProcess.exitCode === null) {
            kernelProcess.kill()
        }
        saveBilladmConfig()
        app.quit();
    }
});
