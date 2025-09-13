<template>
  <!-- 使用 Teleport 将模态框传送到 body -->
  <Teleport to="body">
    <div v-if="visible" class="confirm-dialog-overlay" @click="closeDialog">
      <div class="confirm-dialog" @click.stop>
        <!-- 标题 -->
        <h3 class="confirm-title">{{ title }}</h3>

        <!-- 路径输入框 + 浏览按钮 -->
        <div class="input-with-button">
          <el-input
              v-model="selectedPath"
              :placeholder="placeholder"
              :readonly="true"
              style="flex: 1; margin-right: 8px;"
          />
          <el-button type="primary" @click.stop="handleBrowse">
            {{ browseLabel }}
          </el-button>
        </div>

        <!-- 错误提示 -->
        <p v-if="error" class="error-message">{{ error }}</p>

        <!-- 按钮组 -->
        <div class="button-container">
          <button class="dialog-button" :style="{ color: cancelColor }" @click="closeDialog">
            {{ cancelLabel }}
          </button>
          <button
              class="dialog-button"
              :style="{ color: confirmColor }"
              @click="handleConfirm"
              :disabled="!selectedPath"
          >
            {{ confirmLabel }}
          </button>
        </div>
      </div>
    </div>
  </Teleport>
</template>

<script setup>
import {ref, watch} from 'vue';

// --- Props ---
const props = defineProps({
  visible: {
    type: Boolean,
    required: true,
  },
  title: {
    type: String,
    default: '选择文件夹',
  },
  placeholder: {
    type: String,
    default: '请选择目录...',
  },
  browseLabel: {
    type: String,
    default: '浏览',
  },
  cancelLabel: {
    type: String,
    default: '取消',
  },
  confirmLabel: {
    type: String,
    default: '确定',
  },
  cancelColor: {
    type: String,
    default: '#666',
  },
  confirmColor: {
    type: String,
    default: '#007bff',
  },
  mode: {
    type: String,
    validator: (val) => ['directory', 'file'].includes(val),
    default: 'directory',
  },
  filters: {
    type: Array,
    default: () => [],
  },
});

// --- Emits ---
const emit = defineEmits(['update:visible', 'confirm']);

// --- 本地状态 ---
const selectedPath = ref('');
const error = ref('');

// --- 监听 visible 变化，重置状态 ---
watch(
    () => props.visible,
    (newVal) => {
      if (newVal) {
        selectedPath.value = '';
        error.value = '';
      }
    }
);

// --- 方法 ---

// 关闭对话框
function closeDialog() {
  emit('update:visible', false);
}

// 打开系统文件/目录选择器（需要平台支持）
async function handleBrowse() {
  try {
    error.value = '';
    let result;
    result = await window.electronAPI.openDialog({
      properties: props.mode === 'directory' ? ['openDirectory'] : ['openFile'],
      filters: props.mode === 'file' && props.filters.length > 0 ? props.filters : undefined,
    });

    if (result.canceled || !result.filePaths || result.filePaths.length === 0) {
      return;
    }
    selectedPath.value = result.filePaths[0];
  } catch (err) {
    error.value = `选择失败: ${err.message}`;
  }
}

function handleConfirm() {
  if (!selectedPath.value) {
    error.value = '请先选择一个路径';
    return;
  }

  emit('confirm', selectedPath.value);
  closeDialog();
}
</script>

<style scoped>
.confirm-dialog-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: flex-start;
  z-index: 1000;
}

.confirm-dialog {
  background-color: white;
  border-radius: 8px;
  padding: 20px;
  margin-top: calc(100vh / 4);
  min-width: 380px;
  max-width: 500px;
}

.confirm-title {
  margin: 0 0 16px 0;
  font-size: 18px;
  font-weight: 500;
  text-align: left;
}

.input-with-button {
  display: flex;
  align-items: center;
  margin-bottom: 16px;
}

.error-message {
  color: #e74c3c;
  font-size: 14px;
  margin: 0 0 16px 0;
  text-align: left;
  min-height: 20px;
}

.button-container {
  display: flex;
  justify-content: center;
  gap: 12px;
}

.dialog-button {
  padding: 8px 16px;
  font-weight: 500;
  cursor: pointer;
  border: none;
  border-radius: 4px;
  background-color: transparent;
  transition: background-color 0.2s ease;
  outline: none;
  min-width: 80px;
}

.dialog-button:hover:not(:disabled) {
  background-color: var(--billadm-color-icon-hover-bg-color, #f5f5f5);
}

.dialog-button:disabled {
  color: #ccc !important;
  cursor: not-allowed;
}
</style>