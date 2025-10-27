import dayjs from "dayjs";

export function dateToUnixTimestamp(date: Date = new Date()): number {
    return Math.floor(new Date(date).getTime() / 1000);
}

export function formatFloat(num: number): number {
    return parseFloat(num.toFixed(2));
}

/**
 * 将秒级时间戳转换为格式化时间字符串
 * @param timestamp 秒级时间戳
 * @param format 格式，默认为 'YYYY-MM-DD HH:mm:ss'
 * @returns 格式化后的时间字符串
 */
export function formatTimestamp(timestamp: number, format: string = 'YYYY年MM月DD日'): string {
    return dayjs(timestamp * 1000).format(format);
}
