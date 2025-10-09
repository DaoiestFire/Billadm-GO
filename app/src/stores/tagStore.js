import {defineStore} from 'pinia';
import {computed, ref} from 'vue';
import {queryTag} from "@/backend/api/tag.js";
import NotificationUtil from "@/backend/notification.js";

export const useTagStore = defineStore('tag', () => {
    const tags = ref([])

    const tagNames = computed(() => {
        const map = {}

        tags.value.forEach(tag => {
            const {category, name} = tag
            if (!map[category]) {
                map[category] = []
            }
            map[category].push(name)
        })

        return map
    })

    const getTagNamesByCategory = (category) => {
        return tagNames.value[category] || []
    }

    const refreshTag = async () => {
        try {
            tags.value = await queryTag()
        } catch (error) {
            NotificationUtil.error(`获取全部消费标签失败 ${error}`)
        }
    }

    return {
        tags,
        tagNames,
        getTagNamesByCategory,
        refreshTag,
    }
})