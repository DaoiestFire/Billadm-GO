import {ElNotification} from 'element-plus'


class NotificationUtil {
    static info(message: string) {
        this._showNotification('info', message)
    }

    static success(message: string) {
        this._showNotification('success', message)
    }

    static warning(message: string) {
        this._showNotification('warning', message)
    }

    static error(message: string) {
        this._showNotification('error', message)
    }

    /**
     * 内部方法：支持 message 为字符串或配置对象
     */
    private static _showNotification(type: string, message: string) {
        ElNotification({
            type: type,
            message: message,
            duration: 2500,
            showClose: true,
            offset: 100,
        })
    }
}

export default NotificationUtil