/**
 * 获取指定日期的秒级时间戳
 * @param {Date|string|number} date - 可接受 Date 对象、时间字符串或毫秒时间戳
 * @returns {number} 对应时间的秒级时间戳
 */
export function getUnixTimestampInSeconds(date = new Date()) {
    return Math.floor(new Date(date).getTime() / 1000);
}