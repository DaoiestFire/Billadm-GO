<template>
  <div class="menu-container" ref="menuContainer">
    <ul v-if="showMenu" class="menu" :style="menuStyle">
      <MenuItem v-for="(item, index) in menuItems" :key="index" :item="item" :leftIcon="item.icon" :depth="0"
                @close-menu="closeMenu"/>
    </ul>
  </div>
  <ConfirmDialog v-model:visible="showLedgerConfirmDialog" :showInput="showLedgerInput" :message="message"
                 :cancel-color="cancelColor" :confirm-label="confirmLabel" :confirm-color="confirmColor"
                 :item="dialogItem" @confirm="onConfirm"/>
</template>

<script setup>
import {computed, onMounted, onUnmounted, ref} from 'vue';
import {useCssVariables} from '@/css/css';
import MenuItem from '@/components/MenuItem.vue';
import ConfirmDialog from '@/components/ConfirmDialog.vue';
import iconBook from '@/assets/icons/book.svg?raw';
import iconAdd from '@/assets/icons/add.svg?raw';
import iconTrash from '@/assets/icons/trash.svg?raw';
import {useLedgerStore} from "@/stores/ledgerStore.js";

// store
const ledgerStore = useLedgerStore()

// 引用颜色
const {positiveColor, negativeColor} = useCssVariables()

// 各种框的控制变量
const showLedgerConfirmDialog = ref(false)
const showLedgerInput = ref(false)
const message = ref('')
const confirmLabel = ref('确认')
const confirmColor = ref('')
const cancelColor = ref('')
const dialogItem = ref(null)

const onConfirm = async (data) => {
  if (confirmLabel.value === '创建') {
    await ledgerStore.createLedger(data.input)
  }

  if (confirmLabel.value === '删除') {
    await ledgerStore.deleteLedger(data.item)
  }
}

// 响应式计算删除账本的菜单项
const deleteLedgers = computed(() => {
  return ledgerStore.ledgers.map(l => ({
    label: l.name,
    action: () => {
      dialogItem.value = l.id
      message.value = '确认删除账本吗？'
      confirmLabel.value = '删除'
      confirmColor.value = negativeColor.value
      cancelColor.value = positiveColor.value
      showLedgerInput.value = false
      showLedgerConfirmDialog.value = true
    }
  }))
})

// 定义菜单项类型
const menuItems = ref([
  {
    label: '账本', icon: iconBook, children: [
      {
        label: '新建账本',
        icon: iconAdd,
        action: () => {
          message.value = '输入账本名称'
          confirmLabel.value = '创建'
          confirmColor.value = positiveColor.value
          cancelColor.value = negativeColor.value
          showLedgerInput.value = true
          showLedgerConfirmDialog.value = true
        },
      },
      {
        label: '删除账本',
        icon: iconTrash,
        children: deleteLedgers,
      },
    ]
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