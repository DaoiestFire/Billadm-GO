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

<script setup lang="ts">
import {nextTick, ref, watch} from 'vue'
import {ElInput} from 'element-plus'

interface Props {
  visible?: boolean
  title?: string
  message?: string
  cancelLabel?: string
  confirmLabel?: string
  cancelColor?: string
  confirmColor?: string
  showInput?: boolean
  inputPlaceholder?: string
  showButtons?: boolean
  item?: string
}

const props = withDefaults(defineProps<Props>(), {
  visible: false,
  title: '',
  message: '',
  cancelLabel: '取消',
  confirmLabel: '确认',
  cancelColor: '#666',
  confirmColor: '#007bff',
  showInput: false,
  inputPlaceholder: '请输入...',
  showButtons: true,
  item: ''
})

const emit = defineEmits<{
  (e: 'update:visible', visible: boolean): void
  (e: 'confirm', data: {
    input: string
    item: string
  }): void
}>()

const inputValue = ref('')
const inputRef = ref<InstanceType<typeof ElInput> | null>(null)
watch(
    () => props.visible,
    (newVal) => {
      if (newVal && props.showInput) {
        nextTick(() => {
          inputRef.value?.focus()
        })
      }
    }
)

function closeDialog() {
  emit('update:visible', false)
  inputValue.value = ''
}

function handleConfirm() {
  const data = {
    input: props.showInput ? inputValue.value : '',
    item: props.item
  }
  emit('confirm', data)
  closeDialog()
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
  background-color: var(--billadm-color-major-background-color);
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
  font-size: var(--billadm-size-text-large-size);
  font-weight: 600;
  color: var(--billadm-color-text-major-color);
}

.modal-body {
  text-align: left;
}

.modal-message {
  margin: 0 0 16px 0;
  font-size: var(--billadm-size-text-title-size);
  color: var(--billadm-color-text-minor-color);
  white-space: pre-line;
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