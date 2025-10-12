export function isResponseSuccess(response) {
    // 检查输入是否为有效对象
    if (!response || typeof response !== 'object') {
        return false;
    }

    // 检查code字段是否存在且为数字
    if (typeof response.code !== 'number') {
        return false;
    }

    // 根据code判断响应是否正常
    return true;
}

export function dateToUnixTimestamp(date = new Date()) {
    return Math.floor(new Date(date).getTime() / 1000);
}

export function formatFloat(num) {
    // 检查输入是否为有效数字
    if (typeof num !== 'number' || isNaN(num)) {
        throw new Error('输入必须是一个有效的数字');
    }

    // toFixed 返回字符串，所以需要转换为数字
    return parseFloat(num.toFixed(2));
}