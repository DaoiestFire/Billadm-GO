export function dateToUnixTimestamp(date: Date = new Date()): number {
    return Math.floor(new Date(date).getTime() / 1000);
}

export function formatFloat(num: number): number {
    return parseFloat(num.toFixed(2));
}