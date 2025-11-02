import type {TransactionRecord} from '@/types/billadm'
import type {EChartsOption} from "echarts"
import {TransactionTypeToColor, TransactionTypeToLabel} from '@/backend/constant'
import {formatFloat} from '@/backend/functions'
import {trDtoToTrForm} from "@/backend/dto-utils.ts";

/**
 * 根据交易记录数据生成 ECharts 折线图配置项，展示按月分类的收入、支出、转账金额趋势。
 * 横轴覆盖从最早交易到最晚交易之间的所有月份，缺失月份自动填充 0。
 */
export function buildOptionForTradingTrend(trList: TransactionRecord[], displayTypes: string[]): EChartsOption {

    const data = trList.map(tr => trDtoToTrForm(tr));

    const filteredData = data.filter(item => displayTypes.includes(item.type))

    if (!filteredData || filteredData.length === 0) {
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
        }
    }

    let minYear = Infinity, minMonth = Infinity
    let maxYear = -Infinity, maxMonth = -Infinity

    filteredData.forEach(item => {
        const date = item.time
        const year = date.year()
        const month = date.month() + 1

        if (year < minYear || (year === minYear && month < minMonth)) {
            minYear = year
            minMonth = month
        }
        if (year > maxYear || (year === maxYear && month > maxMonth)) {
            maxYear = year
            maxMonth = month
        }
    })

    const totalMonths = []
    let currentYear = minYear
    let currentMonth = minMonth

    while (currentYear < maxYear || (currentYear === maxYear && currentMonth <= maxMonth)) {
        totalMonths.push(`${currentYear}-${String(currentMonth).padStart(2, '0')}`)
        currentMonth++
        if (currentMonth > 12) {
            currentMonth = 1
            currentYear++
        }
    }

    const monthlyData = new Map<string, { income: number; expense: number; transfer: number }>()
    totalMonths.forEach(month => {
        monthlyData.set(month, {income: 0, expense: 0, transfer: 0})
    })

    filteredData.forEach(item => {
        const date = item.time
        const key = `${date.year()}-${String(date.month() + 1).padStart(2, '0')}`
        const type = item.type
        const amount = item.price

        if (monthlyData.has(key)) {
            monthlyData.get(key)[type] += amount
        }
    })

    const xAxisData = totalMonths
    const seriesDataMap = new Map<string, number[]>(
        displayTypes.map(type => [
            type,
            xAxisData.map(month => monthlyData.get(month)[type])
        ])
    )

    return {
        tooltip: {
            trigger: 'axis',
            formatter: (params: any) => {
                const date = params[0].name
                let tooltip = `${date}<br/>`
                params.forEach(param => {
                    tooltip += `${param.marker}${param.seriesName}: ${param.value.toFixed(2)}<br/>`
                })
                return tooltip
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
    }
}

/**
 * 根据交易记录数据生成 ECharts 饼图配置项，展示指定交易类型下各消费类别的金额分布。
 */
export function buildOptionForTransactionDistribution(trList: TransactionRecord[], transactionType: string): EChartsOption {
    const data = trList.map(tr => trDtoToTrForm(tr));
    const filteredData = data.filter(item => item.type === transactionType)

    const categoryMap = new Map<string, number>()

    filteredData.forEach(item => {
        const {category, price} = item
        const current = categoryMap.get(category) || 0
        categoryMap.set(category, current + price)
    })

    const seriesData = Array.from(categoryMap, ([name, value]) => ({
        name,
        value: formatFloat(value)
    }))

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
        }
    }

    return {
        title: {text: ''},
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
        series: [{
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
        }]
    }
}