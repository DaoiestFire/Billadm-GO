<template>
  <div class="custom-select" :style="{ width: width+'px' }" ref="selectWrapper">
    <button @click="toggleDropdown" class="select-button" :style="{ height: height+'px' }">
      {{ selectedLabel || placeholder }}
    </button>
    <div v-show="isOpen" :class="['dropdown', direction === 'up' ? 'dropdown-up' : 'dropdown-down']">
      <div v-for="option in options" :key="option.value" class="option" @click="selectOption(option)"
           :style="[{
             '--bg-width': (props.width - 4) + 'px',
             '--bg-height': (props.height - 4) + 'px'
              }]">
        {{ option.label }}
      </div>
    </div>
  </div>
</template>

<script setup>
import {onMounted, onUnmounted, ref, watch} from 'vue';

const selectedValue = defineModel()

// Props
const props = defineProps({
  options: {
    type: Array,
    required: true
  },
  placeholder: {
    type: String,
    default: '请选择'
  },
  direction: {
    type: String,
    default: 'down',
    validator: value => ['down', 'up'].includes(value)
  },
  width: {
    type: Number,
    default: 100
  },
  height: {
    type: Number,
    default: 28
  }
});

// Data
const isOpen = ref(false);
const selectedLabel = ref('');
const selectWrapper = ref(null);

// Toggle dropdown
function toggleDropdown() {
  isOpen.value = !isOpen.value;
}

// Select option
function selectOption(option) {
  selectedValue.value = option.value;
  selectedLabel.value = option.label;
  isOpen.value = false;
}

watch(
    () => props.options,
    (newOptions) => {
      if (newOptions.length === 0) {
        selectedLabel.value = ''
        return;
      }
      // 如果没有选中值，无需处理
      if (selectedValue.value === undefined || selectedValue.value === null || !selectedValue.value) {
        selectedLabel.value = ''
        return
      }

      const option = newOptions.find(option => option.value === selectedValue.value)
      if (option) {
        selectedLabel.value = option.label
      } else {
        selectedValue.value = ''
      }
    },
    {deep: true, immediate: true}
)

// 关闭下拉框的函数
function closeDropdown(event) {
  // 如果点击在组件外部，则关闭
  if (selectWrapper.value && !selectWrapper.value.contains(event.target)) {
    isOpen.value = false;
  }
}

// 组件挂载后添加事件监听
onMounted(() => {
  document.addEventListener('click', closeDropdown);
});

// 组件卸载前移除事件监听，防止内存泄漏
onUnmounted(() => {
  document.removeEventListener('click', closeDropdown);
});
</script>

<style scoped>
.custom-select {
  position: relative;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.select-button {
  width: 100%;
  text-align: center;
  border: 1px solid var(--billadm-color-window-border-color);
  border-radius: 4px;
  background: white;
  cursor: pointer;
}

.dropdown {
  position: absolute;
  width: 100%;
  overflow-y: auto;
  border: 1px solid var(--billadm-color-window-border-color);
  border-radius: 4px;
  background: white;
  z-index: 10;
}

.dropdown-down {
  top: 110%;
}

.dropdown-up {
  bottom: 110%;
}

.option {
  cursor: pointer;
  white-space: nowrap;
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 30px;
}

.option:hover::after {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: var(--bg-width);
  height: var(--bg-height);
  border-radius: 4px;
  z-index: -1;
  transition: background-color 0.3s ease;
  background-color: var(--billadm-color-icon-hover-bg-color);
}
</style>