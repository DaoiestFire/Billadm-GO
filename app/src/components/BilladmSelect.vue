<template>
  <div
      class="billadm-select"
      :style="{ width: width + 'px' }"
      ref="selectWrapper"
  >
    <!-- 选择按钮 -->
    <button
        @click="toggleDropdown"
        class="select-button"
        :style="{ height: height + 'px' }"
        :class="{ 'is-open': isOpen }"
    >
      {{ selectedLabel || placeholder }}
    </button>

    <!-- 下拉菜单 -->
    <div
        v-show="isOpen"
        :class="[
        'dropdown',
        direction === 'up' ? 'dropdown-up' : 'dropdown-down'
      ]"
        :style="{ maxHeight: dropdownMaxHeight }"
    >
      <div
          v-for="option in options"
          :key="option.value"
          class="option"
          @click="selectOption(option)"
      >
        {{ option.label }}
      </div>
      <!-- 无选项提示 -->
      <div v-if="options.length === 0" class="option disabled">
        无数据
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, onUnmounted, ref, watch} from 'vue';

// 使用 defineModel 支持 v-model
const selectedValue = defineModel();

// 定义下拉选项的类型
interface Option {
  value: string
  label: string
  // 可选：其他字段
  disabled?: boolean
}

// 定义 Props 类型
interface Props {
  options: Option[]
  placeholder?: string
  direction?: 'down' | 'up'
  width?: number
  height?: number
  dropdownMaxHeight?: string
}

const props = withDefaults(defineProps<Props>(), {
  placeholder: '请选择',
  direction: 'down',
  width: 100,
  height: 28,
  dropdownMaxHeight: '200px',
})

// 数据响应式
const isOpen = ref<boolean>(false);
const selectedLabel = ref<string>('');
const selectWrapper = ref<HTMLDivElement | null>(null);

// 计算内部宽高（用于 hover 背景）
const innerWidth = computed(() => props.width - 4 + 'px');
const innerHeight = computed(() => props.height - 4 + 'px');

// 切换下拉框
function toggleDropdown() {
  isOpen.value = !isOpen.value;
}

// 选择选项
function selectOption(option: Option) {
  selectedValue.value = option.value;
  selectedLabel.value = option.label;
  isOpen.value = false;
}

// 监听 options 变化，同步 label 显示
watch(
    () => props.options,
    (newOptions) => {
      if (newOptions.length === 0) {
        selectedLabel.value = '';
        return;
      }

      const currentValue = selectedValue.value;
      if (
          currentValue === undefined ||
          currentValue === null ||
          currentValue === ''
      ) {
        selectedLabel.value = '';
        return;
      }

      const matched = newOptions.find((opt) => opt.value === currentValue);
      if (matched) {
        selectedLabel.value = matched.label;
      } else {
        selectedValue.value = '';
        selectedLabel.value = '';
      }
    },
    {deep: true, immediate: true}
);

// 点击外部关闭下拉框
function closeDropdown(event: MouseEvent) {
  if (selectWrapper.value && !selectWrapper.value.contains(event.target as Node)) {
    isOpen.value = false;
  }
}

onMounted(() => {
  document.addEventListener('click', closeDropdown);
});

onUnmounted(() => {
  document.removeEventListener('click', closeDropdown);
});
</script>

<style scoped>
.billadm-select {
  position: relative;
  display: inline-block;
}

.select-button {
  width: 100%;
  padding: 0;
  text-align: center;
  border: 1px solid var(--billadm-color-window-border-color);
  border-radius: 4px;
  background-color: var(--billadm-color-major-background-color);
  color: var(--billadm-color-text-major-color);
  cursor: pointer;
  transition: border-color 0.2s ease, box-shadow 0.2s ease;
  outline: none;
}

.select-button:hover {
  border-color: var(--billadm-color-positive-color);
}

.dropdown {
  position: absolute;
  top: 100%;
  left: 0;
  width: 100%;
  max-height: v-bind(dropdownMaxHeight);
  overflow-y: auto;
  border: 1px solid var(--billadm-color-window-border-color);
  border-radius: 4px;
  background-color: white;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  z-index: 1000;
  font-size: 14px;
}

.dropdown-up {
  top: auto;
  bottom: 110%;
}

.dropdown-down {
  top: 110%;
}

.option {
  cursor: pointer;
  white-space: nowrap;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 30px;
  user-select: none;
  position: relative;
}

.option:hover::after {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: v-bind(innerWidth);
  height: v-bind(innerHeight);
  border-radius: 4px;
  background-color: var(--billadm-color-icon-hover-bg-color);
  z-index: -1;
  transition: background-color 0.3s ease;
}

.option.disabled {
  color: #999;
  cursor: not-allowed;
  font-style: italic;
}
</style>