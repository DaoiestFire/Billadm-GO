<template>
    <div class="menu-container" ref="menuContainer">
        <ul v-if="showMenu" class="menu" :style="menuStyle">
            <MenuItem v-for="(item, index) in menuItems" :key="index" :item="item" :leftIcon="item.icon" :depth="0"
                @close-menu="closeMenu" />
        </ul>
    </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue';
import MenuItem from '@/components/MenuItem.vue';
import iconBook from '@/assets/icons/book.svg?raw'
import iconInfo from '@/assets/icons/info.svg?raw'
import iconAdd from '@/assets/icons/add.svg?raw'
import iconTrash from '@/assets/icons/trash.svg?raw'

// 定义菜单项类型
const menuItems = ref([
    {
        label: '账本', icon: iconBook, children: [
            {
                label: '新建账本', icon: iconAdd
            },
            {
                label: '删除账本', icon: iconTrash, children: [
                    {
                        label: '刘敬威的账本',
                        action: () => console.log('删除刘敬威的账本')
                    },
                    {
                        label: '默认账本',
                        action: () => console.log('删除默认账本')
                    }
                ]
            },
        ]
    },
    {
        label: '关于', icon: iconInfo, action: () => console.log('关于')
    },
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
        left: `-50px`,
        zIndex: 1000,
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