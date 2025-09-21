export function getTodayMiddleDate() {
    const noonToday = new Date();
    noonToday.setHours(12, 0, 0, 0); // 小时, 分钟, 秒, 毫秒
    return noonToday
}