/**
 * 设置日期对象的时间部分为 00:00:00
 * @param {Date} date - 输入的日期对象
 * @returns {Date} 时间设置为 00:00:00 的新日期对象
 */
export function setToStartOfDay(date) {
    const newDate = new Date(date);
    newDate.setHours(0, 0, 0, 0);
    return newDate;
}

/**
 * 设置日期对象的时间部分为 23:59:59
 * @param {Date} date - 输入的日期对象
 * @returns {Date} 时间设置为 23:59:59 的新日期对象
 */
export function setToEndOfDay(date) {
    const newDate = new Date(date);
    newDate.setHours(23, 59, 59, 999);
    return newDate;
}

/**
 * 获取今天的时间范围 [开始, 结束]
 * @returns {Array<Date>} 包含今天开始和结束时间的数组
 */
export function getTodayRange() {
    const today = new Date();
    return [setToStartOfDay(today), setToEndOfDay(today)];
}

/**
 * 获取本周的时间范围 [开始, 结束]
 * 假设一周从星期一开始
 * @returns {Array<Date>} 包含本周开始和结束时间的数组
 */
export function getThisWeekRange() {
    const today = new Date();
    const day = today.getDay(); // 0 (Sunday) to 6 (Saturday)

    // 计算距离周一的天数差
    // 如果是周日(day=0)，则距离周一有6天；其他情况减1
    const daysToMonday = day === 0 ? 6 : day - 1;

    const monday = new Date(today);
    monday.setDate(today.getDate() - daysToMonday);
    const start = setToStartOfDay(monday);

    const sunday = new Date(monday);
    sunday.setDate(monday.getDate() + 6);
    const end = setToEndOfDay(sunday);

    return [start, end];
}

/**
 * 获取本月的时间范围 [开始, 结束]
 * @returns {Array<Date>} 包含本月开始和结束时间的数组
 */
export function getThisMonthRange() {
    const today = new Date();
    const year = today.getFullYear();
    const month = today.getMonth();

    const start = new Date(year, month, 1);
    const end = new Date(year, month + 1, 0); // 下个月的第0天即为本月最后一天

    return [setToStartOfDay(start), setToEndOfDay(end)];
}

/**
 * 获取上周的时间范围 [开始, 结束]
 * 假设一周从星期一开始
 * @returns {Array<Date>} 包含上周开始和结束时间的数组
 */
export function getLastWeekRange() {
    const today = new Date();
    const day = today.getDay(); // 0 (Sunday) to 6 (Saturday)

    // 计算距离上周一的天数差
    // 如果是周日(day=0)，则距离上周一有13天；其他情况加6
    const daysToLastMonday = day === 0 ? 13 : day + 6;

    const lastMonday = new Date(today);
    lastMonday.setDate(today.getDate() - daysToLastMonday);
    const start = setToStartOfDay(lastMonday);

    const lastSunday = new Date(lastMonday);
    lastSunday.setDate(lastMonday.getDate() + 6);
    const end = setToEndOfDay(lastSunday);

    return [start, end];
}

/**
 * 获取上月的时间范围 [开始, 结束]
 * @returns {Array<Date>} 包含上月开始和结束时间的数组
 */
export function getLastMonthRange() {
    const today = new Date();
    const year = today.getFullYear();
    const month = today.getMonth();

    // 上个月
    let prevMonth, prevYear;
    if (month === 0) {
        prevMonth = 11;
        prevYear = year - 1;
    } else {
        prevMonth = month - 1;
        prevYear = year;
    }

    const start = new Date(prevYear, prevMonth, 1);
    const end = new Date(prevYear, prevMonth + 1, 0); // 当月的第0天即为上月最后一天

    return [setToStartOfDay(start), setToEndOfDay(end)];
}

export function getPreviousDay(date) {
    const result = new Date(date);
    result.setDate(result.getDate() - 1);
    return result;
}

export function getNextDay(date) {
    const result = new Date(date);
    result.setDate(result.getDate() + 1);
    return result;
}