import {computed, ref, watch} from "vue";
import {defineStore} from 'pinia';
import {queryTrCountOnCondition, queryTrsOnCondition, queryTrStatisticsOnCondition} from "@/backend/tr.js";
import NotificationUtil from "@/backend/notification.js";
import {useLedgerStore} from "@/stores/ledgerStore.js";
import {dateToUnixTimestamp} from "@/backend/functions.js";
import {setToEndOfDay, setToStartOfDay} from "@/backend/timerange.js";

export const useTrViewStore = defineStore('trView', () => {
    // 定义状态
    const tableData = ref([])
    const pages = ref(1) // 总页数
    const currentPage = ref(1) // 当前页数
    const pageSize = ref(20) // 每页记录数
    const trCount = ref(0) // 总记录数
    const timeRange = ref([]) // 时间选择器
    const trStatistics = ref({}) // 时间范围内消费记录的统计数据

    // computed
    const tsRange = computed(() => {
        if (!Array.isArray(timeRange.value) || timeRange.value.length < 2) {
            return null
        }
        /** @type {Date} */
        const startTs = timeRange.value[0];
        /** @type {Date} */
        const endTs = timeRange.value[1];
        return [dateToUnixTimestamp(setToStartOfDay(startTs)), dateToUnixTimestamp(setToEndOfDay(endTs))]
    })

    // ledgerStore
    const ledgerStore = useLedgerStore()

    // 初始化接口
    const init = async () => {
        await refreshPages()
        await refreshTableData()
    }

    const refreshPages = async () => {
        try {
            let condition = {};
            condition['ledger_id'] = ledgerStore.currentLedgerId;
            if (tsRange.value !== null) {
                condition['ts_range'] = tsRange.value;
            }
            const trCnt = await queryTrCountOnCondition(condition);
            const pagesVal = Math.ceil(trCnt / pageSize.value);
            pages.value = pagesVal < 1 ? 1 : pagesVal;
            trCount.value = trCnt;
        } catch (error) {
            NotificationUtil.error(`查询消费记录数量失败 ${error}`);
        }
    }

    const refreshTableData = async () => {
        try {
            const offset = (currentPage.value - 1) * pageSize.value
            let condition = {};
            condition['ledger_id'] = ledgerStore.currentLedgerId;
            condition['offset'] = offset;
            condition['limit'] = pageSize.value;
            if (tsRange.value !== null) {
                condition['ts_range'] = tsRange.value;
            }
            tableData.value = await queryTrsOnCondition(condition)
        } catch (error) {
            NotificationUtil.error(`消费记录数据刷新失败 ${error}`)
        }
    }

    const refreshStatistics = async () => {
        try {
            let condition = {};
            condition['ledger_id'] = ledgerStore.currentLedgerId;
            if (tsRange.value !== null) {
                condition['ts_range'] = tsRange.value;
            }
            trStatistics.value = await queryTrStatisticsOnCondition(condition)
        } catch (error) {
            NotificationUtil.error(`消费记录统计数据刷新失败 ${error}`)
        }
    }

    const resetView = () => {
        tableData.value = [];
        currentPage.value = 1;
        pageSize.value = 20;
    }

    watch(() => [pageSize.value, currentPage.value, tsRange.value], async () => {
        await init();
    })

    watch(() => trCount.value, async () => {
        await refreshStatistics();
    })

    watch(() => ledgerStore.currentLedger, async () => {
        resetView();
        await init();
        await refreshStatistics();
    })

    // 返回 store 的状态和方法
    return {
        tableData,
        pages,
        currentPage,
        pageSize,
        trCount,
        timeRange,
        trStatistics,
        init,
        refreshTableData,
        refreshPages,
        refreshStatistics,
    }
})