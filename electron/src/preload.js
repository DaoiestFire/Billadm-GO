const {contextBridge, ipcRenderer, dialog} = require('electron')

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
    }
})