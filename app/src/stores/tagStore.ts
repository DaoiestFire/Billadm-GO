import {defineStore} from 'pinia'
import {computed, ref} from 'vue'
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
        return tagNames.value.get(category)?.map(value => {
            return {value};
        }) || []
    }


    return {
        tags,
        tagNames,
        getTagNamesByCategory,
    }
})