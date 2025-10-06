<template>
  <div class="ledger-select-container">
    <!-- 主选择框 -->
    <div class="ledger-select" :style="{ height: height }">
      <!-- 左侧显示选中项 -->
      <div class="select-left" @click="toggleDropdown">
        {{ ledgerStore.currentLedgerName ? ledgerStore.currentLedgerName : placeholder }}
      </div>
      <!-- 右侧按钮 -->
      <div class="select-right">
        <BilladmIconButton
            :svg="iconAdd" label="" width="30px" height="30px" :color="iconColor"
            :hover-bg-color="hoverBgColor" bg-size="26px" tooltip-placement="bottom-end"
            @click="createLedgerFunc"
        />
      </div>
    </div>

    <!-- 下拉框 -->
    <div v-if="showDropdown" class="dropdown">
      <div v-for="option in options" :key="option.value" class="dropdown-item" @click="selectOption(option)">
        <!-- 下拉项左侧文本 -->
        <div class="item-left">{{ option.label }}</div>
        <!-- 下拉项右侧按钮 -->
        <div class="item-right">
          <BilladmIconButton
              :svg="iconTrash" label="" width="30px" height="30px" :color="iconColor"
              :hover-bg-color="hoverBgColor" bg-size="26px" tooltip-placement="bottom-end"
              @click="deleteLedgerFunc(option.value,option.label)"
          />
        </div>
      </div>
    </div>
    <billadm-modal
        v-model:visible="showLedgerConfirmDialog"
        :showInput="showLedgerInput"
        :show-buttons="true"
        :title="title"
        :message="message"
        :cancel-color="cancelColor"
        :confirm-label="confirmLabel"
        :confirm-color="confirmColor"
        :item="dialogItem"
        @confirm="onConfirm"
    />
  </div>
</template>

<script setup>
import {computed, ref} from 'vue'
import iconAdd from "@/assets/icons/add.svg?raw";
import iconTrash from "@/assets/icons/trash.svg?raw";
import BilladmModal from "@/components/BilladmModal.vue";
import BilladmIconButton from "@/components/BilladmIconButton.vue";
import {useLedgerStore} from "@/stores/ledgerStore.js";
import {useCssVariables} from "@/css/css.js";

// css variables
const {iconColor, hoverBgColor} = useCssVariables()

// store
const ledgerStore = useLedgerStore()

// 定义组件属性
const props = defineProps({
  height: {
    type: String,
    default: '30px'
  },
  placeholder: {
    type: String,
    default: '选择账本'
  }
})

const options = computed(() => {
  if (!Array.isArray(ledgerStore.ledgers)) {
    return []
  }

  return ledgerStore.ledgers.map(ledger => ({
    label: ledger.name,
    value: ledger.id,
  }))
})

// 是否显示下拉框
const showDropdown = ref(false)

// 切换下拉框显示状态
const toggleDropdown = () => {
  showDropdown.value = !showDropdown.value
}

// 选择一个选项
const selectOption = (option) => {
  ledgerStore.setCurrentLedger(option.value);
  showDropdown.value = false;
}

// 点击外部区域关闭下拉框
document.addEventListener('click', (event) => {
  const target = event.target
  const selectContainer = target.closest('.ledger-select-container')
  if (!selectContainer) {
    showDropdown.value = false
  }
})

// 引用颜色
const {positiveColor, negativeColor} = useCssVariables()

// 各种框的控制变量
const showLedgerConfirmDialog = ref(false);
const showLedgerInput = ref(false);
const title = ref('');
const message = ref('');
const confirmLabel = ref('确认');
const confirmColor = ref('');
const cancelColor = ref('');
const dialogItem = ref(null);

const createLedgerFunc = () => {
  title.value = '新建账本';
  message.value = '输入账本名称';
  confirmLabel.value = '创建';
  confirmColor.value = positiveColor.value;
  cancelColor.value = negativeColor.value;
  showLedgerInput.value = true;
  showLedgerConfirmDialog.value = true;
}

const deleteLedgerFunc = (id, name) => {
  dialogItem.value = id;
  title.value = '删除账本';
  message.value = `确认删除账本<<<${name}>>>吗？`;
  confirmLabel.value = '删除';
  confirmColor.value = negativeColor.value;
  cancelColor.value = positiveColor.value;
  showLedgerInput.value = false;
  showLedgerConfirmDialog.value = true;
}

const onConfirm = async (data) => {
  if (confirmLabel.value === '创建') {
    await ledgerStore.createLedger(data.input);
  }

  if (confirmLabel.value === '删除') {
    await ledgerStore.deleteLedger(data.item);
  }
}
</script>

<style scoped>
.ledger-select-container {
  position: relative;
  display: inline-block;
}

.ledger-select {
  display: flex;
  border: 1px solid var(--billadm-color-window-border-color);
  border-radius: 8px;
  cursor: pointer;
  background-color: var(--billadm-color-major-backgroud-color);
}

.ledger-select:hover {
  border-color: var(--billadm-color-positive-color);
}

.select-left {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 130px;
}

.select-right {
  display: flex;
  align-items: center;
  z-index: 1;
}

.dropdown {
  position: absolute;
  top: 110%;
  left: 0;
  border: 1px solid var(--billadm-color-window-border-color);
  border-radius: 4px;
  background-color: var(--billadm-color-major-backgroud-color);
  overflow-y: auto;
  z-index: 3;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.dropdown-item {
  position: relative;
  display: flex;
  height: 30px;
  cursor: pointer;
}

.item-left {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  width: 130px;
}

.item-left:hover::after {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 126px;
  height: 26px;
  border-radius: 4px;
  transition: background-color 0.3s ease;
  background-color: var(--billadm-color-icon-hover-bg-color);
  z-index: -1;
}

.item-right {
  display: flex;
  align-items: center;
}
</style>