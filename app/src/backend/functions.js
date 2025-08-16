/**
 * 获取指定日期的秒级时间戳
 * @param {Date|string|number} date - 可接受 Date 对象、时间字符串或毫秒时间戳
 * @returns {number} 对应时间的秒级时间戳
 */
export function getUnixTimestampInSeconds(date = new Date()) {
    return Math.floor(new Date(date).getTime() / 1000);
}

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
    return response.code === 0;
}