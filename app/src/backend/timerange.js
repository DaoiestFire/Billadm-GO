/**
 * 获取本月最后一天的时间
 * @param {Date} date - 输入的日期对象
 * @returns {Date} 时间设置为 00:00:00 的新日期对象
 */
export function getLastDayOfMonth(date) {
    let newDate = new Date(date);
    newDate.setMonth(newDate.getMonth() + 1);
    newDate.setDate(0);
    newDate = setToEndOfDay(new Date(newDate.getFullYear(), newDate.getMonth() + 1, 0));
    return newDate;
}

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

/**
 * 获取今年的时间范围 [开始, 结束]
 * @returns {Array<Date>} 包含今年开始和结束时间的数组
 */
export function getThisYearRange() {
    const today = new Date();
    const year = today.getFullYear();

    // 今年1月1日
    const start = new Date(year, 0, 1);
    // 今年12月31日
    const end = new Date(year, 12, 0);

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

/**
 * 判断两个 Date 对象之间的时间区间粒度，并返回向后一个周期的新区间
 * @param {Date} startDate - 开始时间，时分秒毫秒为 0,0,0,0
 * @param {Date} endDate - 结束时间，时分秒毫秒为 23:59:59.999
 * @returns {[Date, Date]} 新的时间区间
 */
export function getNextPeriod(startDate, endDate) {
    startDate = setToStartOfDay(startDate);
    endDate = setToEndOfDay(endDate);
    return shiftPeriod(startDate, endDate, 1);
}

/**
 * 判断两个 Date 对象之间的时间区间粒度，并返回向前一个周期的新区间
 * @param {Date} startDate - 开始时间，时分秒毫秒为 0,0,0,0
 * @param {Date} endDate - 结束时间，时分秒毫秒为 23:59:59.999
 * @returns {[Date, Date]} 新的时间区间
 */
export function getPrevPeriod(startDate, endDate) {
    startDate = setToStartOfDay(startDate);
    endDate = setToEndOfDay(endDate);
    return shiftPeriod(startDate, endDate, -1);
}

function shiftPeriod(startDate, endDate, direction) {
    // 确保输入是 Date 对象
    startDate = new Date(startDate);
    endDate = new Date(endDate);

    // 计算时间差（毫秒）
    const diffMs = endDate.getTime() - startDate.getTime();

    // 一天的毫秒数
    const dayMs = 24 * 60 * 60 * 1000;
    const weekMs = 7 * dayMs;
    const yearMs = 365 * dayMs;

    // 判断粒度
    if (Math.abs(diffMs - (dayMs - 1)) < 2000) {
        // 一天：00:00:00.000 -> 23:59:59.999，差值为 86,399,999 ms
        const newStart = new Date(startDate);
        newStart.setDate(startDate.getDate() + direction);
        const newEnd = new Date(endDate);
        newEnd.setDate(endDate.getDate() + direction);
        return [newStart, newEnd];
    } else if (Math.abs(diffMs - (weekMs - 1)) < 2000) {
        // 一周
        const newStart = new Date(startDate);
        newStart.setDate(startDate.getDate() + 7 * direction);
        const newEnd = new Date(endDate);
        newEnd.setDate(endDate.getDate() + 7 * direction);
        return [newStart, newEnd];
    } else if (isSameMonth(startDate, endDate)) {
        // 同一个月，可能是月粒度,检查是否为整月（例如 1月1日 -> 1月31日）
        // 获取开始月最后一天与结束时间比较
        const expectedEnd = getLastDayOfMonth(startDate)
        if (Math.abs(endDate.getTime() - expectedEnd.getTime()) < 2000) {
            // 是整月区间
            const newStart = new Date(startDate);
            newStart.setMonth(startDate.getMonth() + direction);
            const newEnd = getLastDayOfMonth(newStart);
            return [newStart, newEnd];
        }
    } else if (startDate.getDate() === 1 && endDate.getDate() === 31 && Math.abs(diffMs - (yearMs - 1)) < 2 * dayMs) {
        // 年粒度：通常是 1月1日 到 12月31日
        const newStart = new Date(startDate);
        newStart.setFullYear(startDate.getFullYear() + direction);
        const newEnd = new Date(endDate);
        newEnd.setFullYear(endDate.getFullYear() + direction);
        return [newStart, newEnd];
    }

    // 默认 fallback：按天处理（或无法识别时按一天偏移）
    const newStart = new Date(startDate);
    newStart.setDate(startDate.getDate() + direction);
    const newEnd = new Date(endDate);
    newEnd.setDate(endDate.getDate() + direction);
    return [newStart, newEnd];
}

// 辅助函数：判断两个日期是否在同一月（忽略时分秒）
function isSameMonth(date1, date2) {
    return date1.getFullYear() === date2.getFullYear() && date1.getMonth() === date2.getMonth();
}

export function getShortcuts() {
    return [
        {
            text: '今天',
            value: getTodayRange,
        },
        {
            text: '本周',
            value: getThisWeekRange,
        },
        {
            text: '本月',
            value: getThisMonthRange,
        },
        {
            text: '上周',
            value: getLastWeekRange,
        },
        {
            text: '上月',
            value: getLastMonthRange,
        },
        {
            text: '本年',
            value: getThisYearRange,
        }
    ]
}

export function getTimeRangeTypes() {
    return [
        {label: '年', value: 'yearrange'},
        {label: '月', value: 'monthrange'},
        {label: '日', value: 'daterange'}
    ]
}

export function getPageSizeOptions() {
    return [
        {label: '每页10行', value: 10},
        {label: '每页20行', value: 20},
        {label: '每页50行', value: 50}
    ]
}