import {defineStore} from 'pinia'
import {computed, ref} from 'vue'
import {createLedgerByName, deleteLedgerById, getAllLedgers} from "@/backend/ledger.js";
import NotificationUtil from "@/backend/notification";

// 定义账本对象的类型 (JavaScript 中主要用于文档和类型提示，TypeScript 中更严格)
/**
 * @typedef {Object} Ledger
 * @property {string} id - 账本唯一标识符
 * @property {string} name - 账本名称
 * @property {string} created_at - 创建时间 (unix timestamp)
 * @property {string} updated_at - 更新时间 (unix timestamp)
 */

export const useLedgerStore = defineStore('ledger', () => {
    const ledgers = ref([])
    const currentLedger = ref(null)

    const refreshTrTable = ref(false)

    const currentLedgerId = computed(() => {
        return currentLedger.value ? currentLedger.value.id : ''
    })

    const currentLedgerName = computed(() => {
        return currentLedger.value ? currentLedger.value.name : ''
    })

    // 访问后端更新账本
    const updateLedgers = async () => {
        try {
            ledgers.value = []
            const ledgersFromServer = await getAllLedgers()
            if (ledgersFromServer && Array.isArray(ledgersFromServer)) {
                ledgersFromServer.forEach(ledger => {
                    ledgers.value.push(ledger)
                })
                ledgers.value.sort((a, b) => a.name.localeCompare(b.name));
            }
            if (currentLedger.value !== null) {
                setCurrentLedger(currentLedger.value.id)
            }
            if (currentLedger.value === null && ledgers.value.length > 0) {
                setCurrentLedger(ledgers.value[0].id)
            }
        } catch (error) {
            NotificationUtil.error(`请求全部账本失败 ${error}`)
        }
    }

    // 添加账本
    const createLedger = async (name) => {
        try {
            await createLedgerByName(name)
            await updateLedgers()
            NotificationUtil.success(`账本 ${name} 创建成功`)
        } catch (error) {
            NotificationUtil.error(`账本 ${name} 创建失败 ${error}`)
        }
    }

    // 删除账本
    const deleteLedger = async (id) => {
        try {
            await deleteLedgerById(id)
            await updateLedgers()
            NotificationUtil.success(`账本删除成功`)
        } catch (error) {
            NotificationUtil.error(`账本删除失败`)
        }
    }

    // 设置当前账本
    const setCurrentLedger = (id) => {
        if (id === null) {
            currentLedger.value = null
            return
        }
        const ledger = ledgers.value.find(l => l.id === id)
        if (ledger) {
            currentLedger.value = {...ledger} // 创建副本，避免直接引用
        } else {
            currentLedger.value = null
        }
    }

    return {
        ledgers,
        currentLedger,
        refreshTrTable,
        currentLedgerId,
        currentLedgerName,
        updateLedgers,
        createLedger,
        deleteLedger,
        setCurrentLedger
    }
})