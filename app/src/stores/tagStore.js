import {defineStore} from 'pinia'
import {computed, ref} from 'vue'
import {queryAllTag} from "@/backend/tag.js";
import NotificationUtil from "@/backend/notification.js";

export const useTagStore = defineStore('tag', () => {
    const tags = ref([])

    // 计算属性：获取 tag 名称列表
    const tagNames = computed(() => {
        return [...tags.value].map(tag => tag.name)
    })

    const init = async () => {
        await refreshTag()
    }

    // 更新指定 name 的 tag
    const refreshTag = async () => {
        try {
            tags.value = await queryAllTag()
        } catch (error) {
            NotificationUtil.error(`获取全部消费标签失败 ${error}`)
        }
    }

    // 返回需要在组件中使用的状态、计算属性和方法
    return {
        tags,
        tagNames,
        init,
        refreshTag,
    }
})