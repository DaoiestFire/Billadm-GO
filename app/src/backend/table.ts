import {TransactionTypeToColor, TransactionTypeToLabel} from '@/backend/constant.ts'
import {formatFloat} from '@/backend/functions.ts'

import type {TrForm} from '@/types/billadm'
import type {EChartsOption} from "echarts"

/**
 * 根据交易记录数据生成 ECharts 折线图配置项，展示按月分类的收入、支出、转账金额趋势。
 * 横轴覆盖从最早交易到最晚交易之间的所有月份，缺失月份自动填充 0。
 */
export function buildOptionForTradingTrend(data: TrForm[], displayTypes: string[]): EChartsOption {

    let filteredData = data.filter(item => displayTypes.includes(item.type));

    if (!filteredData || filteredData.length === 0) {
        // 处理空数据情况
        return {
            tooltip: {trigger: 'axis'},
            legend: {data: displayTypes.map(item => TransactionTypeToLabel.get(item))},
            xAxis: {
                type: 'category',
                data: [],
                name: '月份'
            },
            yAxis: {
                type: 'value',
                name: '金额'
            },
            series: displayTypes.map(type => ({
                name: type === 'income' ? '收入' : type === 'expense' ? '支出' : '转账',
                type: 'line',
                data: []
            }))
        };
    }

    // 获取最早的年月和最晚的年月
    let minYear = Infinity, minMonth = Infinity;
    let maxYear = -Infinity, maxMonth = -Infinity;

    filteredData.forEach(item => {
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
    filteredData.forEach(item => {
        const date = item.time;
        const year = date.getFullYear();
        const month = date.getMonth() + 1;
        const key = `${year}-${month}`;
        const type = item.type;
        const amount = item.price;

        if (monthlyData.has(key)) {
            monthlyData.get(key)[type] += amount;
        }
    });

    // 使用完整的月份序列作为横轴
    const xAxisData = totalMonths;

    // 构建每种类型的序列数据
    const seriesDataMap = new Map(
        displayTypes.map(type => [
            type,
            xAxisData.map(month => monthlyData.get(month)[type])
        ])
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
            data: displayTypes.map(item => TransactionTypeToLabel.get(item))
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
        series: displayTypes.map(item => {
            return {
                name: TransactionTypeToLabel.get(item),
                type: 'line',
                data: seriesDataMap.get(item),
                emphasis: {focus: 'series'},
                color: TransactionTypeToColor.get(item)
            }
        })
    };
}

/**
 * 根据交易记录数据生成 ECharts 饼图配置项，展示指定交易类型下各消费类别的金额分布。
 *
 * @param {TrForm[]} data - 交易记录对象数组
 * @param {string} transactionType - 交易类型，取值为 'income', 'expense', 'transfer' 之一
 * @returns {Object} ECharts 的 option 配置对象
 *
 * @example
 * const option = buildOptionForTransactionDistribution(transactionData, 'expense');
 * chart.setOption(option);
 */
export function buildOptionForTransactionDistribution(data, transactionType) {
    // 过滤出指定类型的交易
    const filteredData = data.filter(item => item.type === transactionType);

    // 统计每个 category 的总金额
    const categoryMap = new Map();

    filteredData.forEach(item => {
        const {category, price} = item;
        const current = categoryMap.get(category) || 0;
        categoryMap.set(category, current + price);
    });

    // 转换为 ECharts 所需的数据格式
    const seriesData = Array.from(categoryMap, ([name, value]) => ({
        name: name,
        value: formatFloat(value)
    }));

    // 如果没有数据，返回空的配置项（避免图表报错）
    if (seriesData.length === 0) {
        return {
            title: {
                text: '暂无数据',
                left: 'center',
                top: 'center',
                textStyle: {
                    color: '#999'
                }
            },
            tooltip: {show: false},
            series: []
        };
    }

    return {
        tooltip: {
            trigger: 'item',
            formatter: '{b}: {c} ({d}%)'
        },
        legend: {
            type: 'scroll',
            orient: 'vertical',
            right: 10,
            top: 20,
            bottom: 20
        },
        series: [
            {
                name: TransactionTypeToLabel.get(transactionType) || transactionType,
                type: 'pie',
                radius: ['40%', '70%'],
                center: ['40%', '50%'],
                avoidLabelOverlap: false,
                label: {
                    show: true,
                    formatter: '{b}\n{c} ({d}%)',
                    fontSize: 12,
                },
                emphasis: {
                    label: {
                        show: true,
                        fontSize: 14,
                        fontWeight: 'bold'
                    }
                },
                labelLine: {
                    show: true,
                    smooth: true,
                    length: 8,
                    length2: 10
                },
                data: seriesData
            }
        ]
    };
}