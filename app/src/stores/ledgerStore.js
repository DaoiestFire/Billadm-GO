import {defineStore} from 'pinia'
import {computed, ref} from 'vue'
import {createLedgerByName, getAllLedgers} from "@/backend/ledger.js";

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

    const allLedgersAction = computed(() => {
        return ledgers.value
    })

    const currentLedgerAction = computed(() => {
        return currentLedger.value
    })

    // 访问后端更新账本
    const updateLedgers = async () => {
        try {
            const ledgersFromServer = await getAllLedgers()
            if (ledgersFromServer && Array.isArray(ledgersFromServer) && ledgersFromServer.length > 0) {
                ledgersFromServer.forEach(ledger => {
                    ledgers.value = []
                    ledgers.value.push(ledger)
                })
            }
            setCurrentLedger(currentLedger.value.id)
        } catch (error) {
            console.log('更新账本缓存失败')
        }
    }

    // 添加账本
    const createLedger = async (name) => {
        try {
            await createLedgerByName(name)
            await updateLedgers()
        } catch (error) {
            console.log('新增账本失败')
        }
    }

    // 设置当前账本
    const setCurrentLedger = (id) => {
        const ledger = ledgers.value.find(l => l.id === id)
        if (ledger) {
            currentLedger.value = {...ledger} // 创建副本，避免直接引用
        } else {
            console.warn(`未找到 ID 为 ${id} 的账本`)
            currentLedger.value = null
        }
    }

    // 清除当前账本
    const clearCurrentLedger = () => {
        currentLedger.value = null
    }

    return {
        ledgers,
        currentLedger,
        allLedgersAction,
        currentLedgerAction,
        updateLedgers,
        createLedger,
        setCurrentLedger,
        clearCurrentLedger // 可选
    }
})