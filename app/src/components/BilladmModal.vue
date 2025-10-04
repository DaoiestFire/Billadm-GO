<template>
  <!-- 使用 Teleport 将模态框传送到 body 标签末尾 -->
  <Teleport to="body">
    <!-- 遮罩层，点击可关闭（可配置） -->
    <div v-if="visible" class="billadm-modal-overlay" @click="closeDialog()">
      <!-- 模态框主体 -->
      <div class="billadm-modal" @click.stop>
        <!-- 头部标题（可选） -->
        <div v-if="title" class="modal-header">
          <h3 class="modal-title">{{ title }}</h3>
        </div>

        <!-- 主体内容 -->
        <div class="modal-body">
          <p class="modal-message">{{ message }}</p>
          <!-- 可选输入框 -->
          <div v-if="showInput" class="input-container">
            <el-input
                ref="inputRef"
                v-model="inputValue"
                style="width: 100%"
                :placeholder="inputPlaceholder"
                clearable
                @keydown.enter="handleConfirm"
            />
          </div>
        </div>

        <!-- 底部按钮区域（可选显示） -->
        <div v-if="showButtons" class="modal-footer">
          <div class="button-container">
            <!-- 取消按钮 -->
            <button class="dialog-button" :style="{ color: cancelColor }" @click="closeDialog">
              {{ cancelLabel }}
            </button>

            <!-- 确认按钮 -->
            <button class="dialog-button" :style="{ color: confirmColor }" @click="handleConfirm">
              {{ confirmLabel }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </Teleport>
</template>

<script setup>
import {nextTick, ref, watch} from 'vue';

// --- Props 定义 ---
const props = defineProps({
  visible: {
    type: Boolean,
    required: false
  },
  title: {
    type: String,
    default: ''
  },
  message: {
    type: String,
    default: ''
  },
  cancelLabel: {
    type: String,
    default: '取消'
  },
  confirmLabel: {
    type: String,
    default: '确认'
  },
  cancelColor: {
    type: String,
    default: '#666'
  },
  confirmColor: {
    type: String,
    default: '#007bff'
  },
  showInput: {
    type: Boolean,
    default: false
  },
  inputPlaceholder: {
    type: String,
    default: '请输入...'
  },
  showButtons: {
    type: Boolean,
    default: false
  },
  item: {
    type: [Object, null],
    default: null
  }
});

// --- Emits 定义 ---
const emit = defineEmits(['update:visible', 'confirm']);

// --- 本地状态 ---
const inputValue = ref('');
const inputRef = ref(null);

// --- 监听器 ---
watch(
    () => props.visible,
    (newVal) => {
      if (newVal && props.showInput) {
        nextTick(() => {
          if (inputRef.value) {
            inputRef.value.focus();
          }
        });
      }
    }
);

// --- 方法 ---
function closeDialog() {
  emit('update:visible', false);
  inputValue.value = '';
}

function handleConfirm() {
  const data = {
    input: props.showInput ? inputValue.value : null,
    item: props.item
  };
  emit('confirm', data);
  closeDialog();
}
</script>

<style scoped>
.billadm-modal-overlay {
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

.billadm-modal {
  background-color: var(--billadm-color-major-backgroud-color);
  border-radius: 8px;
  padding: 16px;
  margin-top: calc(100vh / 4);
  min-width: 300px;
  max-width: 90vw;
}

/* 头部样式 */
.modal-header {
  margin-bottom: 16px;
}

.modal-title {
  margin: 0;
  font-size: 18px;
  font-weight: 600;
  color: var(--billadm-color-text-major-color);
}

.modal-body {
  text-align: left;
}

.modal-message {
  margin: 0 0 16px 0;
  font-size: 16px;
  color: var(--billadm-color-text-minor-color);
}

.input-container {
  margin-bottom: 16px;
}

.modal-footer {
  margin-top: 20px;
}

.button-container {
  display: flex;
  justify-content: center;
  gap: 8px;
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

.dialog-button:hover {
  background-color: var(--billadm-color-icon-hover-bg-color);
}
</style>