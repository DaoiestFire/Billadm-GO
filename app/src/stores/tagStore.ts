import {defineStore} from 'pinia'
import {computed, ref} from 'vue'
import {queryTags} from "@/backend/api/tag.ts"
import NotificationUtil from "@/backend/notification.ts"
import type {Tag} from "@/types/billadm"

export const useTagStore = defineStore('tag', () => {
    const tags = ref([] as Tag[])

    const tagNames = computed(() => {
        const map = new Map<string, string[]>()

        tags.value.forEach(tag => {
            const {category, name} = tag
            if (!map.has(category)) {
                map.set(category, [])
            }
            const names = map.get(category)
            if (names) {
                names.push(name)
            }
        })

        return map
    })

    const getTagNamesByCategory = (category: string) => {
        return tagNames.value.get(category) || []
    }

    const refreshTag = async () => {
        try {
            tags.value = await queryTags('all')
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