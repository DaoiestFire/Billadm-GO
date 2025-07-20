<!-- GlobalTooltip.vue -->
<template>
    <teleport to="body">
        <span v-if="isVisible" class="global-tooltip" :style="{ left: x + 'px', top: y + 'px' }">
            {{ content }}
        </span>
    </teleport>
</template>

<script setup>
import { ref } from 'vue'

const props = defineProps({
    content: {
        type: String,
        default: ''
    }
})

const x = ref(0)
const y = ref(0)
const isVisible = ref(false)

const showTooltip = (event) => {
    if (!props.content) return
    const rect = event.target.getBoundingClientRect()
    x.value = rect.left + window.scrollX + rect.width
    y.value = rect.top + rect.height / 2
    isVisible.value = true
}

const hideTooltip = () => {
    isVisible.value = false
}

defineExpose({
    showTooltip,
    hideTooltip
})
</script>

<style scoped>
.global-tooltip {
    position: absolute;
    background-color: rgba(0, 0, 0, 0.8);
    color: white;
    padding: 4px 8px;
    border-radius: 4px;
    font-size: 12px;
    white-space: nowrap;
    z-index: 99999;
    pointer-events: none;
    transition: opacity 0.2s ease-in-out;
    font-family: sans-serif;
}
</style>