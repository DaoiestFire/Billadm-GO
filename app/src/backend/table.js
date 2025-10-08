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
 *
 * @example
 * const transaction = {
 *   id: 'txn_123',
 *   time: new Date(1730562600000),
 *   type: 'expense',
 *   category: 'groceries',
 *   description: 'Weekly supermarket shopping',
 *   tags: ['food', 'weekly'],
 *   price: 89.99
 * };
 */

import {TransactionType} from "@/backend/constant.js";

/**
 * 根据交易记录数据生成 ECharts 折线图配置项，展示按月分类的收入、支出、转账金额趋势。
 *
 * @param {TrForm[]} data - 交易记录对象数组
 * @returns {Object} ECharts 的 option 配置对象
 *
 * @example
 * const option = buildOptionForTradingTrend(transactionData);
 * chart.setOption(option);
 */
export function buildOptionForTradingTrend(data) {
    // 用于存储每个月每种类型的总金额
    const monthlyData = new Map();

    // 遍历所有交易数据
    data.forEach(item => {
        const date = item.time;
        const year = date.getFullYear();
        const month = date.getMonth() + 1; // getMonth() 返回 0-11，所以加1
        const key = `${year}-${month}`; // 使用 "YYYY-M" 格式作为月份键
        const type = item.type;
        const amount = item.price;

        // 初始化该月份的数据（如果不存在）
        if (!monthlyData.has(key)) {
            monthlyData.set(key, {
                income: 0,
                expense: 0,
                transfer: 0
            });
        }

        const monthObj = monthlyData.get(key);
        if (Object.values(TransactionType).includes(type)) {
            monthObj[type] += amount;
        }
    });

    // 按时间顺序排序并生成横轴标签和数据序列
    const sortedMonths = Array.from(monthlyData.keys()).sort();
    const xAxisData = sortedMonths; // 横轴显示月份

    // 为每种类型生成数据数组
    const seriesData = Object.values(TransactionType).map(type =>
        sortedMonths.map(month => monthlyData.get(month)[type])
    );

    return {
        tooltip: {
            trigger: 'axis',
            formatter: (params) => {
                const date = params[0].name;
                let tooltip = `${date}月<br/>`;
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