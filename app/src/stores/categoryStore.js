import {defineStore} from 'pinia'
import {computed, ref} from 'vue'
import {queryAllCategory} from "@/backend/category.js";
import NotificationUtil from "@/backend/notification.js";

export const useCategoryStore = defineStore('category', () => {
    const categories = ref([])

    // 计算属性：获取 category 名称列表，并按字典降序排序
    const categoryNames = computed(() => {
        return [...categories.value]
            .map(category => category.name)
            .sort((a, b) => b.localeCompare(a))
    })

    const init = async () => {
        await refreshCategory()
    }

    // 更新指定 name 的 category
    const refreshCategory = async () => {
        try {
            categories.value = await queryAllCategory()
        } catch (error) {
            NotificationUtil.error(`获取全部消费类型失败 ${error}`)
        }
    }

    // 返回需要在组件中使用的状态、计算属性和方法
    return {
        categories,
        categoryNames,
        init,
        refreshCategory,
    }
})