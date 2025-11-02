import {defineStore} from 'pinia'
import {computed, ref} from 'vue'
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
        return categoryNames.value.get(transactionType)?.map(value => {
            return {value};
        }) || [];
    };


    // 返回需要在组件中使用的状态、计算属性和方法
    return {
        categories,
        categoryNames,
        getCategoryNamesByType,
    }
})