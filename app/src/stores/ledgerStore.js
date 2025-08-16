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

    const currentLedgerIdAction = computed(() => {
        return currentLedger.value ? currentLedger.value.id : ''
    })

    const currentLedgerNameAction = computed(() => {
        return currentLedger.value ? currentLedger.value.name : ''
    })

    // 访问后端更新账本
    const updateLedgers = async () => {
        try {
            const ledgersFromServer = await getAllLedgers()
            if (ledgersFromServer && Array.isArray(ledgersFromServer)) {
                ledgers.value = []
                ledgersFromServer.forEach(ledger => {
                    ledgers.value.push(ledger)
                })
            }
            if (currentLedger.value != null) {
                setCurrentLedger(currentLedger.value.id)
            }
        } catch (error) {
            console.log('更新账本缓存失败', error)
        }
    }

    // 添加账本
    const createLedger = async (name) => {
        try {
            await createLedgerByName(name)
            await updateLedgers()
            NotificationUtil.success(`账本 ${name} 创建成功`)
        } catch (error) {
            console.log('新增账本失败', error)
        }
    }

    // 删除账本
    const deleteLedger = async (id) => {
        console.log('delete', id)
        try {
            await deleteLedgerById(id)
            await updateLedgers()
        } catch (error) {
            console.log('删除账本失败', error)
        }
    }

    // 设置当前账本
    const setCurrentLedger = (id) => {
        const ledger = ledgers.value.find(l => l.id === id)
        if (ledger) {
            currentLedger.value = {...ledger} // 创建副本，避免直接引用
        } else {
            currentLedger.value = null
        }
        console.log('current ledger', currentLedger.value)
    }

    // 清除当前账本
    const clearCurrentLedger = () => {
        currentLedger.value = null
    }

    return {
        ledgers,
        currentLedger,
        currentLedgerIdAction,
        currentLedgerNameAction,
        updateLedgers,
        createLedger,
        deleteLedger,
        setCurrentLedger,
        clearCurrentLedger // 可选
    }
})