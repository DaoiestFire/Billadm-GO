<template>
  <div class="ledger-select-container">
    <!-- 主选择框 -->
    <div
        class="ledger-select"
        :style="{ height: height }"
        @click="toggleDropdown"
    >
      <!-- 左侧显示选中项 -->
      <div class="select-left">
        {{ selectedOption ? selectedOption.label : placeholder }}
      </div>
      <!-- 右侧按钮 -->
      <div class="select-right">
        <BilladmButton :icon="iconAdd" label="新建账本" :width="height" :height="height" :color="iconColor"
                       :bgColor="minorBgColor" :hoverBgColor="hoverBgColor" hoverStyle="circle"
                       tooltipPlacement="bottom-end"/>
      </div>
    </div>

    <!-- 下拉框 -->
    <div v-if="showDropdown" class="dropdown">
      <div
          v-for="option in options"
          :key="option.value"
          class="dropdown-item"
          @click="selectOption(option)"
      >
        <!-- 下拉项左侧文本 -->
        <div class="item-left">
          {{ option.label }}
        </div>
        <!-- 下拉项右侧按钮 -->
        <div class="item-right">
          <BilladmButton :icon="iconAdd"/>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import {defineProps, ref} from 'vue'
import BilladmButton from "@/components/BilladmButton.vue";
import iconAdd from "@/assets/icons/add.svg?raw";
import {useCssVariables} from "@/css/css.js";

// css variables
const {minorBgColor, hoverBgColor, iconColor} = useCssVariables()

// 定义组件属性
const props = defineProps({
  // 选择框宽度
  width: {
    type: String,
    default: '150px'
  },
  // 选择框高度
  height: {
    type: String,
    default: '30px'
  },
  // 可选项数组
  options: {
    type: Array,
    default: () => [{value: '1', label: '现金账'},
      {value: '2', label: '银行账'},
      {value: '3', label: '应收账款'}]
  },
  // 占位符文本
  placeholder: {
    type: String,
    default: '请选择...'
  }
})

// 是否显示下拉框
const showDropdown = ref(false)

// 当前选中的选项
const selectedOption = ref(null)

// 切换下拉框显示状态
const toggleDropdown = (event) => {
  // 如果点击的是主选择框，则切换下拉框
  if (event.target.closest('.select-left')) {
    showDropdown.value = !showDropdown.value
  }
}

// 选择一个选项
const selectOption = (option) => {
  selectedOption.value = option
  showDropdown.value = false
}

// 点击外部区域关闭下拉框
document.addEventListener('click', (event) => {
  const target = event.target
  const selectContainer = target.closest('.ledger-select-container')
  if (!selectContainer) {
    showDropdown.value = false
  }
})
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
  overflow: hidden;
}

.select-left {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: var(--billadm-color-major-backgroud-color);
  width: 130px;
}

.select-right {
  display: flex;
  align-items: center;
  background-color: #e9e9e9;
}

.dropdown {
  position: absolute;
  top: 110%;
  left: 0;
  z-index: 1000;
  border: 1px solid var(--billadm-color-window-border-color);
  border-radius: 4px;
  background-color: var(--billadm-color-major-backgroud-color);
  overflow-y: auto;
}

.dropdown-item {
  position: relative;
  overflow: hidden;
  display: flex;
  height: 30px;
  border-bottom: 1px solid var(--billadm-color-window-border-color);
  cursor: pointer;
}

.dropdown-item:last-child {
  border-bottom: none;
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
  z-index: -1;
  transition: background-color 0.3s ease;
  background-color: var(--billadm-color-icon-hover-bg-color);
}

.item-right {
  display: flex;
  align-items: center;
}

/* 点击外部关闭下拉框的样式 */
</style>