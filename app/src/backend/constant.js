import {
    getLastMonthRange,
    getLastWeekRange,
    getThisMonthRange,
    getThisWeekRange,
    getThisYearRange,
    getTodayRange
} from "@/backend/timerange.js";


export const TransactionTypeToLabel = new Map([
    ['income', '收入'],
    ['expense', '支出'],
    ['transfer', '转账']
]);

export const TransactionTypeToColor = new Map([
    ['income', '#67C23A'],
    ['expense', '#f56c6c'],
    ['transfer', '#409eff']
]);

export const TimeRangeShortcuts = [
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
        text: '今年',
        value: getThisYearRange,
    }
];

export const TimeRangeTypes = [
    {label: '年', value: 'yearrange'},
    {label: '月', value: 'monthrange'},
    {label: '日', value: 'daterange'}
];

export const PageSizeOptions = [
    {label: '每页10行', value: 10},
    {label: '每页20行', value: 20},
    {label: '每页50行', value: 50}
];
