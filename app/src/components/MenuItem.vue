<template>
    <li class="menu-item" :class="{ 'has-children': item.children }" @mouseenter="handleMouseEnter"
        @mouseleave="handleMouseLeave" @click="handleClick" ref="menuItem">
        <DisplayIcon :rawSvg="leftIcon" />
        <span class="item-label">{{ item.label }}</span>
        <DisplayIcon v-if="item.children" :rawSvg="iconChevronRight" />

        <!-- 子菜单 -->
        <ul v-if="item.children && showChildren" class="submenu" :style="submenuStyle">
            <MenuItem v-for="(child, index) in item.children" :key="index" :item="child" :leftIcon="child.icon"
                :depth="depth + 1" @close-menu="$emit('close-menu')" />
        </ul>
    </li>
</template>

<script setup>
import { ref, computed, defineEmits, defineProps } from 'vue'
import DisplayIcon from '@/components/DisplayIcon.vue'
import iconChevronRight from '@/assets/icons/chevron-right.svg?raw'

// 定义 props
const props = defineProps({
    item: {
        type: Object,
        required: true,
    },
    depth: {
        type: Number,
        default: 0,
    },
    leftIcon: {
        type: String,
        default: '',
    },
});

// 定义 emit
const emit = defineEmits(['close-menu']);

// 子菜单显示状态
const showChildren = ref(false);

// 计算子菜单的位置（在父菜单项右侧，顶部对齐）
const menuItem = ref(null);
const submenuStyle = computed(() => {
    if (!showChildren.value || !menuItem.value) return {};
    const parentRect = menuItem.value.getBoundingClientRect();
    return {
        position: 'absolute',
        top: `0px`,
        left: `${parentRect.width}px`,
        zIndex: 1000 + props.depth,
    };
});

// 鼠标进入（Hover）
const handleMouseEnter = () => {
    showChildren.value = true;
};

// 鼠标离开
const handleMouseLeave = () => {
    showChildren.value = false;
};

// 点击菜单项
const handleClick = (event) => {
    event.stopPropagation(); // 阻止事件冒泡到父级
    if (props.item.action) {
        props.item.action();
        emit('close-menu'); // 执行动作后关闭整个菜单
    }
};
</script>

<style scoped>
.menu-item {
    position: relative;
    cursor: pointer;
    display: flex;
    justify-content: space-between;
    align-items: center;
    height: 40px;
}

.menu-item:hover {
    background-color: var(--billadm-color-icon-hover-bg-color);
}

/* 子菜单样式 */
.submenu {
    list-style: none;
    margin: 0;
    padding: 0;
    background-color: white;
    border: 1px solid var(--billadm-color-window-border-color);
    border-radius: 4px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
    min-width: 160px;
}
</style>