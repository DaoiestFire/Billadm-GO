import {ref} from "vue";
import {defineStore} from 'pinia';
import {queryTrsOnCondition} from "@/backend/tr.js";
import NotificationUtil from "@/backend/notification.js";
import {useLedgerStore} from "@/stores/ledgerStore.js";

export const useTrViewStore = defineStore('trView', () => {
    // 定义状态
    const tableData = ref([])
    const pages = ref(1)
    const currentPage = ref(1)
    const pageSize = ref(10)

    // ledgerStore
    const ledgerStore = useLedgerStore()

    // 初始化接口
    const init = async () => {
        await refreshPages()
        await refreshTableData()
    }

    // set 函数
    const refreshTableData = async () => {
        try {
            const offset = (currentPage - 1) * pageSize
            tableData.value = await queryTrsOnCondition(ledgerStore.currentLedgerId, offset, pageSize)
        } catch (error) {
            NotificationUtil.error(`消费记录数据刷新失败 ${error}`)
        }
    }

    const refreshPages = () => {
        // 函数体为空
    }

    const setCurrentPage = (page) => {
        // 函数体为空
    }

    const setPageSize = (size) => {
        // 函数体为空
    }

    // 返回 store 的状态和方法
    return {
        tableData,
        pages,
        currentPage,
        pageSize,
        init,
        refreshTableData,
        refreshPages,
        setCurrentPage,
        setPageSize
    }
})