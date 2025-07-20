<template>
    <button class="common-icon" :style="[
        buttonStyle,
        { '--hover-bg-color': hoverBgColor }
    ]" @mouseenter="handleMouseEnter" @mouseleave="handleMouseLeave" :class="[
        { 'common-icon--rect': hoverStyle === 'rect' },
        { 'common-icon--circle': hoverStyle === 'circle' },
        { 'is-hovered': isHovered || isActive },
        { 'is-active': isActive }
    ]">
        <span class="icon" :style="iconStyle" v-html="icon"></span>
        <teleport to="body">
            <span v-if="isTooltipVisible" class="global-tooltip"
                :style="{ left: tooltipX + 'px', top: tooltipY + 'px' }">
                {{ label }}
            </span>
        </teleport>
    </button>
</template>

<script setup>
import { computed, ref } from 'vue'

const props = defineProps({
    icon: {
        type: String,
        required: true
    },
    label: {
        type: String,
        default: ''
    },
    width: {
        type: [String, Number],
        default: 24
    },
    height: {
        type: [String, Number],
        default: 24
    },
    iconWidth: {
        type: [String, Number],
        default: 18
    },
    iconHeight: {
        type: [String, Number],
        default: 18
    },
    strokeWidth: {
        type: [String, Number],
        default: 0.1
    },
    color: {
        type: String,
        default: '#7a7a7a'
    },
    bgColor: {
        type: String,
        default: '#f5f5f5'
    },
    hoverBgColor: {
        type: String,
        default: '#e0e0e0'
    },
    hoverStyle: {
        type: String,
        default: 'rect',
        validator: value => ['rect', 'circle'].includes(value)
    },
    isActive: {
        type: Boolean,
        default: false
    }
})

const isHovered = ref(false)
const isTooltipVisible = ref(false)
const tooltipX = ref(0)
const tooltipY = ref(0)

const handleMouseEnter = (event) => {
    if (!props.isActive) {
        isHovered.value = true
    }
    if (!props.label || props.isActive) return
    const rect = event.target.getBoundingClientRect()
    tooltipX.value = rect.left + window.scrollX + rect.width / 2
    tooltipY.value = rect.top + window.scrollY - 10
    isTooltipVisible.value = true
}

const handleMouseLeave = () => {
    if (!props.isActive) {
        isHovered.value = false
    }
    isTooltipVisible.value = false
}

const buttonStyle = computed(() => {
    return {
        width: `${props.width}px`,
        height: `${props.height}px`,
        backgroundColor: props.bgColor
    }
})

const iconStyle = computed(() => {
    return {
        color: props.color,
        width: `${props.iconWidth}px`,
        height: `${props.iconHeight}px`,
        strokeWidth: `${props.strokeWidth}px`
    }
})
</script>

<style scoped>
.common-icon {
    position: relative;
    border: none;
    padding: 0;
    cursor: pointer;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    transition: background-color 0.3s ease;
    overflow: hidden;

    --hover-bg-color: rgba(0, 0, 0, 0.2);
}

/* rect 模式：悬停或 active 时全屏暗色背景 */
.common-icon--rect.is-hovered::after,
.common-icon--rect.is-active::after {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: var(--hover-bg-color);
    z-index: 1;
}

/* circle 模式：悬停或 active 时圆形背景 */
.common-icon--circle.is-hovered::after,
.common-icon--circle.is-active::after {
    content: '';
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 70%;
    height: 70%;
    background-color: var(--hover-bg-color);
    border-radius: 20%;
    z-index: 1;
    transition: background-color 0.3s ease;
}

.icon {
    position: relative;
    z-index: 2;
    display: block;
    fill: currentColor;
    stroke: currentColor;
}

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