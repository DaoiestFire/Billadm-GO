import {notification} from "ant-design-vue";

class NotificationUtil {
    static success(message: string, description?: string) {
        notification.success({
            message,
            description
        })
    }

    static error(message: string, description?: string) {
        notification.error({
            message,
            description
        })
    }
}

export default NotificationUtil