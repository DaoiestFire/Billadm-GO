<template>
    <div class="menu-container" ref="menuContainer">
        <ul v-if="showMenu" class="menu" :style="menuStyle">
            <MenuItem v-for="(item, index) in menuItems" :key="index" :item="item" :depth="0" @close-menu="closeMenu" />
        </ul>
    </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue';
import MenuItem from '@/components/MenuItem.vue';
// 定义菜单项类型
const menuItems = ref([
    {
        label: '文件', children: [
            { label: '新建', action: () => console.log('新建文件') },
            { label: '打开', action: () => console.log('打开文件') },
            { label: '保存', action: () => console.log('保存文件') },
            { label: '另存为...', action: () => console.log('另存为') },
            {
                label: '导出', children: [
                    { label: 'PDF', action: () => console.log('导出为 PDF') },
                    { label: 'Word', action: () => console.log('导出为 Word') },
                    {
                        label: '图片', children: [
                            { label: 'PNG', action: () => console.log('导出为 PNG') },
                            { label: 'JPG', action: () => console.log('导出为 JPG') },
                        ]
                    }
                ]
            }
        ]
    },
    {
        label: '编辑', children: [
            { label: '撤销', action: () => console.log('撤销') },
            { label: '重做', action: () => console.log('重做') },
            { label: '剪切', action: () => console.log('剪切') },
            { label: '复制', action: () => console.log('复制') },
            { label: '粘贴', action: () => console.log('粘贴') },
        ]
    },
    {
        label: '视图', children: [
            { label: '放大', action: () => console.log('放大') },
            { label: '缩小', action: () => console.log('缩小') },
            { label: '全屏', action: () => console.log('全屏') },
        ]
    },
    { label: '帮助', action: () => console.log('帮助') },
]);

// 菜单显示状态
const showMenu = ref(false);

// 菜单容器和触发按钮的引用
const menuContainer = ref(null);

// 计算菜单位置（默认在触发按钮右侧）
const menuStyle = computed(() => {
    if (!showMenu.value) return {};
    // 获取触发按钮的位置
    const triggerRect = menuContainer.value?.getBoundingClientRect();
    if (!triggerRect) return {};

    return {
        position: 'absolute',
        top: `${triggerRect.bottom}px`,
        left: `${triggerRect.right}px`,
        zIndex: 1000, // 确保在最上层
    };
});

// 切换菜单显示/隐藏
const toggleMenu = () => {
    showMenu.value = !showMenu.value;
};

// 关闭菜单
const closeMenu = () => {
    showMenu.value = false;
};

// 处理点击外部关闭菜单
const handleClickOutside = (event) => {
    if (menuContainer.value && !menuContainer.value.contains(event.target)) {
        closeMenu();
    }
};

// 监听点击外部事件
onMounted(() => {
    document.addEventListener('mousedown', handleClickOutside);
});

// 移除事件监听器
onUnmounted(() => {
    document.removeEventListener('mousedown', handleClickOutside);
});

defineExpose({
    toggleMenu,
})
</script>

<style scoped>
.menu-container {
    position: relative;
    display: inline-block;
}

.menu {
    list-style: none;
    margin: 0;
    padding: 0;
    background-color: white;
    border: 1px solid var(--billadm-color-window-border-color);
    border-radius: 4px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
    min-width: 160px;
}

/* 确保菜单在小屏幕下不被截断，可以添加滚动或调整位置逻辑，这里简化处理 */
</style>