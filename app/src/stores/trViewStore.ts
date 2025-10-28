import {computed, ref, watch} from "vue"
import {defineStore} from 'pinia'
import {queryTrCountOnCondition, queryTrOnCondition, queryTrStatisticsOnCondition} from "@/backend/api/tr.ts"
import NotificationUtil from "@/backend/notification.ts"
import {useLedgerStore} from "@/stores/ledgerStore.ts"
import {getTodayRange, setToEndOfDay, setToStartOfDay} from "@/backend/timerange.ts"
import type {TimeRangeType, TransactionRecord, TrQueryCondition, TrStatistics} from "@/types/billadm"

export const useTrViewStore = defineStore('trView', () => {
    // 定义状态
    const tableData = ref([] as TransactionRecord[])
    const pages = ref(1) // 总页数
    const currentPage = ref(1) // 当前页数
    const pageSize = ref(20) // 每页记录数
    const trCount = ref(0) // 总记录数
    const timeRange = ref(getTodayRange()) // 时间选择器
    const timeRangeType = ref('daterange' as TimeRangeType) // 时间选择器
    const trStatistics = ref({income: 0, expense: 0, transfer: 0} as TrStatistics) // 时间范围内消费记录的统计数据

    // computed
    const tsRange = computed(() => {
        if (!Array.isArray(timeRange.value) || timeRange.value.length < 2) {
            return null
        }
        const startTs = timeRange.value[0]
        const endTs = timeRange.value[1]
        return [setToStartOfDay(startTs).unix(), setToEndOfDay(endTs).unix()]
    })

    // ledgerStore
    const ledgerStore = useLedgerStore()

    // 初始化接口
    const init = async () => {
        resetView()
        await refreshPages()
        await refreshTableData()
        await refreshStatistics()
    }

    const refreshPages = async () => {
        try {
            let condition = {} as TrQueryCondition
            condition.ledgerId = ledgerStore.currentLedgerId
            if (tsRange.value !== null) {
                condition.tsRange = tsRange.value
            }
            const trCnt = await queryTrCountOnCondition(condition)
            const pagesVal = Math.ceil(trCnt / pageSize.value)
            pages.value = pagesVal < 1 ? 1 : pagesVal
            if (pages.value < currentPage.value) {
                currentPage.value = pages.value
            }
            trCount.value = trCnt
        } catch (error) {
            NotificationUtil.error(`查询消费记录数量失败 ${error}`)
        }
    }

    const refreshTableData = async () => {
        try {
            const offset = (currentPage.value - 1) * pageSize.value
            let condition = {} as TrQueryCondition
            condition.ledgerId = ledgerStore.currentLedgerId
            condition.offset = offset
            condition.limit = pageSize.value
            if (tsRange.value !== null) {
                condition.tsRange = tsRange.value
            }
            tableData.value = await queryTrOnCondition(condition)
        } catch (error) {
            NotificationUtil.error(`消费记录数据刷新失败 ${error}`)
        }
    }

    const refreshStatistics = async () => {
        try {
            let condition = {} as TrQueryCondition
            condition.ledgerId = ledgerStore.currentLedgerId
            if (tsRange.value !== null) {
                condition.tsRange = tsRange.value
            }
            trStatistics.value = await queryTrStatisticsOnCondition(condition)
        } catch (error) {
            NotificationUtil.error(`消费记录统计数据刷新失败 ${error}`)
        }
    }

    const resetView = () => {
        tableData.value = []
        currentPage.value = 1
        pageSize.value = 20
    }

    watch(() => currentPage.value, async () => {
        await refreshTableData()
    })

    watch(() => [ledgerStore.currentLedger, pageSize.value, tsRange.value], async () => {
        resetView()
        await refreshPages()
        await refreshTableData()
        await refreshStatistics()
    })

    return {
        tableData,
        pages,
        currentPage,
        pageSize,
        trCount,
        timeRange,
        timeRangeType,
        trStatistics,
        init,
        refreshTableData,
        refreshPages,
        refreshStatistics,
    }
})