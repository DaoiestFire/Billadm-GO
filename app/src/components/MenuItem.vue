<template>
    <li class="menu-item" :class="{ 'has-children': item.children }" @mouseenter="handleMouseEnter"
        @mouseleave="handleMouseLeave" @click="handleClick" ref="menuItem">
        <span class="item-label">{{ item.label }}</span>

        <!-- 子菜单 -->
        <ul v-if="item.children && showChildren" class="submenu" :style="submenuStyle">
            <MenuItem v-for="(child, index) in item.children" :key="index" :item="child" :depth="depth + 1"
                @close-menu="$emit('close-menu')" />
        </ul>
    </li>
</template>

<script setup>
import { ref, computed, defineEmits, defineProps } from 'vue';

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
});

// 定义 emit
const emit = defineEmits(['close-menu']);

// 子菜单显示状态
const showChildren = ref(false);

// 当前菜单项的 DOM 元素引用 (修正：使用 ref 绑定到 <li>)
const menuItem = ref(null);

// 计算子菜单的位置（在父菜单项右侧，顶部对齐）
const submenuStyle = computed(() => {
    if (!showChildren.value || !menuItem.value) return {};

    // 获取当前（父）菜单项的位置
    const parentRect = menuItem.value.getBoundingClientRect();

    return {
        position: 'absolute',
        // 关键修正：子菜单的 top 与父菜单项的 top 对齐
        top: `${parentRect.top}px`,
        // 子菜单的 left 在父菜单项的 right 边缘
        left: `${parentRect.right}px`,
        zIndex: 1000 + props.depth, // 确保层级正确
    };
});

// 鼠标进入（Hover）
const handleMouseEnter = () => {
    showChildren.value = true;
};

// 鼠标离开
const handleMouseLeave = () => {
    // 延迟关闭，避免鼠标快速移动时菜单闪烁
    setTimeout(() => {
        showChildren.value = false;
    }, 150);
};

// 点击菜单项
const handleClick = (event) => {
    event.stopPropagation(); // 阻止事件冒泡到父级
    if (props.item.action) {
        props.item.action();
        emit('close-menu'); // 执行动作后关闭整个菜单
    }
    // 如果有子菜单，点击只展开/收起，不关闭主菜单
};
</script>

<style scoped>
.menu-item {
    position: relative;
    padding: 8px 12px;
    cursor: pointer;
    transition: background-color 0.2s ease;
    display: flex;
    justify-content: space-between;
    align-items: center;
    /* 为箭头预留空间 */
    padding-right: 20px;
}

.menu-item:hover {
    background-color: #f5f5f5;
}

.menu-item.has-children:hover {
    background-color: #e9ecef;
}

/* 使用 ::after 伪元素创建箭头 */
.menu-item.has-children::after {
    content: '▶';
    /* 箭头符号 */
    position: absolute;
    right: 8px;
    /* 调整箭头位置 */
    top: 50%;
    transform: translateY(-50%);
    font-size: 0.7em;
    color: #6c757d;
    transition: transform 0.2s ease;
    pointer-events: none;
    /* 确保箭头不干扰鼠标事件 */
}

/* 可选：为箭头添加 hover 效果 */
.menu-item.has-children:hover::after {
    color: #495057;
}

/* 子菜单样式 */
.submenu {
    list-style: none;
    margin: 0;
    padding: 0;
    background-color: #fff;
    border: 1px solid #ddd;
    border-radius: 4px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
    min-width: 140px;
    /* position: absolute; 由 style 绑定 */
}
</style>