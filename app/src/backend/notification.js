// src/utils/notification.js

// 默认配置
const defaultConfig = {
    duration: 2500,
    showClose: true,
    offset: 100,
}

/**
 * 通知工具类
 */
class NotificationUtil {
    /**
     * 显示信息通知
     * @param {string|Object} message - 通知消息内容，或包含配置的完整对象
     * @param {Object} [options] - 额外的配置选项（当 message 为字符串时使用）
     */
    static info(message, options = {}) {
        this._showNotification('info', message, options)
    }

    /**
     * 显示成功通知
     * @param {string|Object} message - 通知消息内容，或包含配置的完整对象
     * @param {Object} [options] - 额外的配置选项
     */
    static success(message, options = {}) {
        this._showNotification('success', message, options)
    }

    /**
     * 显示警告通知
     * @param {string|Object} message - 通知消息内容，或包含配置的完整对象
     * @param {Object} [options] - 额外的配置选项
     */
    static warning(message, options = {}) {
        this._showNotification('warning', message, options)
    }

    /**
     * 显示错误通知
     * @param {string|Object} message - 通知消息内容，或包含配置的完整对象
     * @param {Object} [options] - 额外的配置选项
     */
    static error(message, options = {}) {
        this._showNotification('error', message, options)
    }

    /**
     * 显示普通通知（无图标）
     * @param {string|Object} message - 通知消息内容，或包含配置的完整对象
     * @param {Object} [options] - 额外的配置选项
     */
    static normal(message, options = {}) {
        this._showNotification('normal', message, options)
    }

    /**
     * 显示通知的内部方法
     * @private
     */
    static _showNotification(type, message, options) {
        let config = {}

        if (typeof message === 'string') {
            // 如果 message 是字符串，则合并配置
            config = {
                ...defaultConfig,
                message,
                ...options
            }
        } else if (typeof message === 'object' && message !== null) {
            // 如果 message 是对象，则将其视为完整配置
            config = {
                ...defaultConfig,
                ...message
            }
        } else {
            // 非法输入，显示错误
            console.warn('Notification message must be a string or an object.')
            return
        }

        // 根据类型设置类型和图标（如果未在 options 中指定）
        if (!config.type && type !== 'normal') {
            config.type = type
        }
        ElNotification(config)
    }
}

export default NotificationUtil