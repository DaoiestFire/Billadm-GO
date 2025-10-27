import {
    getLastMonthRange,
    getLastWeekRange,
    getThisMonthRange,
    getThisWeekRange,
    getThisYearRange,
    getTodayRange
} from "@/backend/timerange.ts";


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

export const TimeRangePresets = [
    {
        label: '今天',
        value: getTodayRange(),
    },
    {
        label: '本周',
        value: getThisWeekRange(),
    },
    {
        label: '本月',
        value: getThisMonthRange(),
    },
    {
        label: '上周',
        value: getLastWeekRange(),
    },
    {
        label: '上月',
        value: getLastMonthRange(),
    },
    {
        label: '今年',
        value: getThisYearRange(),
    }
];

export const TimeRangeTypes = {
    '日': 'date',
    '月': 'month',
    '年': 'year'
} as const;

export const TimeRangeTypeLabels = ['年', '月', '日'];

export const PageSizeOptions = [
    {label: '每页10行', value: 10},
    {label: '每页20行', value: 20},
    {label: '每页50行', value: 50}
];
