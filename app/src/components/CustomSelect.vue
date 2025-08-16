<template>
  <div class="custom-select" :style="{ width: width }" ref="selectWrapper">
    <button @click="toggleDropdown" class="select-button" :style="{ height: height }">
      {{ selectedLabel || placeholder }}
    </button>
    <div v-show="isOpen" :class="['dropdown', direction === 'up' ? 'dropdown-up' : 'dropdown-down']">
      <div v-for="option in options" :key="option.value" class="option" @click="selectOption(option)">
        {{ option.label }}
      </div>
      <div v-if="options.length===0" class="option">
        无账本
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
    type: [String, Number],
    default: '100px'
  },
  height: {
    type: [String, Number],
    default: '28px'
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
  bottom: 120%;
}

.option {
  padding: 5px;
  cursor: pointer;
  text-align: center;
  white-space: nowrap;
}

.option:hover {
  background-color: var(--billadm-color-icon-hover-bg-color);
}
</style>