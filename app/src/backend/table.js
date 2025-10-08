import {TransactionType} from "@/backend/constant.js";

/**
 * 表示一个前端使用的消费记录
 *
 * @typedef {Object} TrForm
 * @property {string} id - 交易ID. 对应transaction_id.
 * @property {Date} time - 交易时间. 对应transaction_at. 时间戳转Date对象.
 * @property {string} type - 交易类型 (e.g., 'income', 'expense', 'transfer'). 对应transaction_type.
 * @property {string} category - 消费类型. 对应category.
 * @property {string} description - 描述. 对应description.
 * @property {Array<string>} tags - 标签. 对应tags.
 * @property {number} price - 金额. 对应price.
 */

/**
 * 根据交易记录数据生成 ECharts 折线图配置项，展示按月分类的收入、支出、转账金额趋势。
 * 横轴覆盖从最早交易到最晚交易之间的所有月份，缺失月份自动填充 0。
 *
 * @param {TrForm[]} data - 交易记录对象数组
 * @returns {Object} ECharts 的 option 配置对象
 *
 * @example
 * const option = buildOptionForTradingTrend(transactionData);
 * chart.setOption(option);
 */
export function buildOptionForTradingTrend(data) {
    if (!data || data.length === 0) {
        // 处理空数据情况
        return {
            tooltip: {trigger: 'axis'},
            legend: {data: ['收入', '支出', '转账']},
            xAxis: {
                type: 'category',
                data: [],
                name: '月份'
            },
            yAxis: {
                type: 'value',
                name: '金额'
            },
            series: Object.values(TransactionType).map(type => ({
                name: type === 'income' ? '收入' : type === 'expense' ? '支出' : '转账',
                type: 'line',
                data: []
            }))
        };
    }

    // 获取最早的年月和最晚的年月
    let minYear = Infinity, minMonth = Infinity;
    let maxYear = -Infinity, maxMonth = -Infinity;

    data.forEach(item => {
        const date = item.time;
        const year = date.getFullYear();
        const month = date.getMonth() + 1;

        if (year < minYear || (year === minYear && month < minMonth)) {
            minYear = year;
            minMonth = month;
        }
        if (year > maxYear || (year === maxYear && month > maxMonth)) {
            maxYear = year;
            maxMonth = month;
        }
    });

    // 生成从 minYear-minMonth 到 maxYear-maxMonth 的所有 YYYY-M 月份列表
    const totalMonths = [];
    let currentYear = minYear;
    let currentMonth = minMonth;

    while (currentYear < maxYear || (currentYear === maxYear && currentMonth <= maxMonth)) {
        totalMonths.push(`${currentYear}-${currentMonth}`);
        currentMonth++;
        if (currentMonth > 12) {
            currentMonth = 1;
            currentYear++;
        }
    }

    // 初始化所有月份的数据为 0
    const monthlyData = new Map();
    totalMonths.forEach(month => {
        monthlyData.set(month, {
            income: 0,
            expense: 0,
            transfer: 0
        });
    });

    // 遍历交易数据，累加金额
    data.forEach(item => {
        const date = item.time;
        const year = date.getFullYear();
        const month = date.getMonth() + 1;
        const key = `${year}-${month}`;
        const type = item.type;
        const amount = item.price;

        if (monthlyData.has(key) && Object.values(TransactionType).includes(type)) {
            monthlyData.get(key)[type] += amount;
        }
    });

    // 使用完整的月份序列作为横轴
    const xAxisData = totalMonths;

    // 构建每种类型的序列数据
    const seriesData = Object.values(TransactionType).map(type =>
        xAxisData.map(month => monthlyData.get(month)[type])
    );

    return {
        tooltip: {
            trigger: 'axis',
            formatter: (params) => {
                const date = params[0].name;
                let tooltip = `${date}<br/>`;
                params.forEach(param => {
                    tooltip += `${param.marker}${param.seriesName}: ${param.value.toFixed(2)}<br/>`;
                });
                return tooltip;
            }
        },
        legend: {
            data: ['收入', '支出', '转账']
        },
        xAxis: {
            type: 'category',
            data: xAxisData,
            name: '月份'
        },
        yAxis: {
            type: 'value',
            name: '金额',
            axisLabel: {
                formatter: '{value}'
            }
        },
        series: [
            {
                name: '收入',
                type: 'line',
                data: seriesData[0],
                emphasis: {focus: 'series'},
                color: '#67C23A'
            },
            {
                name: '支出',
                type: 'line',
                data: seriesData[1],
                emphasis: {focus: 'series'},
                color: '#f56c6c'
            },
            {
                name: '转账',
                type: 'line',
                data: seriesData[2],
                emphasis: {focus: 'series'},
                color: '#409eff'
            }
        ]
    };
}