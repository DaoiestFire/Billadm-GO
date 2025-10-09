import {defineStore} from 'pinia';
import {computed, ref} from 'vue';
import {queryCategory} from "@/backend/api/category.js";
import NotificationUtil from "@/backend/notification.js";

export const useCategoryStore = defineStore('category', () => {
    const categories = ref([]);

    const categoryNames = computed(() => {
        const map = {};

        categories.value.forEach(category => {
            const {transaction_type, name} = category;
            if (!map[transaction_type]) {
                map[transaction_type] = [];
            }
            map[transaction_type].push(name);
        });

        return map;
    });

    const getCategoryNamesByType = (transactionType) => {
        return categoryNames.value[transactionType] || [];
    };

    // 更新指定 name 的 category
    const refreshCategory = async () => {
        try {
            categories.value = await queryCategory()
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