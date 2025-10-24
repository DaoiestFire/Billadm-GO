import {defineStore} from 'pinia'
import {computed, ref} from 'vue'
import {queryCategory} from "@/backend/api/category.ts"
import NotificationUtil from "@/backend/notification.ts"
import type {Category} from "@/types/billadm"

export const useCategoryStore = defineStore('category', () => {
    const categories = ref([] as Category[]);

    const categoryNames = computed(() => {
        const map = new Map<string, string[]>();

        categories.value.forEach((category: Category) => {
            const {transactionType, name} = category;
            if (!map.has(transactionType)) {
                map.set(transactionType, [])
            }
            const names = map.get(transactionType);
            if (names) {
                names.push(name);
            }
        });

        return map;
    });

    const getCategoryNamesByType = (transactionType: string) => {
        return categoryNames.value.get(transactionType) || [];
    };

    // 更新指定 name 的 category
    const refreshCategory = async () => {
        try {
            categories.value = await queryCategory('all')
        } catch (error) {
            NotificationUtil.error(`获取全部消费类型失败 ${error}`)
        }
    }

    // 返回需要在组件中使用的状态、计算属性和方法
    return {
        categories,
        categoryNames,
        getCategoryNamesByType,
        refreshCategory,
    }
})