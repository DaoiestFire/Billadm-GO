const {contextBridge, ipcRenderer} = require('electron')

contextBridge.exposeInMainWorld('electronAPI', {
    minimizeWindow: () => {
        ipcRenderer.send('window-control', 'minimize');
    },
    maximizeWindow: () => {
        ipcRenderer.send('window-control', 'maximize');
    },
    closeWindow: () => {
        ipcRenderer.send('window-control', 'close');
    },
    openDialog: async (options) => {
        return await ipcRenderer.invoke('dialog:open', options);
    },
    setWorkspace: (workspaceDir) => {
        ipcRenderer.send('workspace:set', workspaceDir);
    },
    getAppInfo: async (field) => {
        return await ipcRenderer.invoke('app', field);
    },
})