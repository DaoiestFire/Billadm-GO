<!-- GlobalTooltip.vue -->
<template>
    <teleport to="body">
        <span v-if="isVisible" class="global-tooltip" :style="loc">
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
    },
    placement: {
        type: String,
        default: 'right-bottom',
    }
})


const loc = ref({})
const isVisible = ref(false)

const showTooltip = (event) => {
    if (!props.content) return

    const target = event.target
    const rect = target.getBoundingClientRect()

    console.log(rect)

    // 获取视口宽度（viewport width）
    const windowWidth = window.innerWidth

    // 获取视口高度（viewport height）
    const windowHeight = window.innerHeight

    switch (props.placement) {
        case 'bottom-left':
            loc.value = {
                right: (windowWidth - rect.right) + 'px',
                top: (rect.top + rect.height) + 'px',
            }
            break

        case 'right-bottom':
            loc.value = {
                left: (rect.left + rect.width) + 'px',
                top: (rect.top + rect.height / 2) + 'px',
            }
            break
    }

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
}
</style>