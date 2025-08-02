<template>
    <!-- 使用 Teleport 将内容传送到 body 标签末尾 -->
    <Teleport to="body">
        <!-- 遮罩层，点击可关闭（可配置） -->
        <div v-if="visible" class="confirm-dialog-overlay" @click="closeDialog()">
            <!-- 确认框主体 -->
            <div class="confirm-dialog" @click.stop>
                <!-- 提示信息 -->
                <p class="confirm-message">{{ message }}</p>

                <!-- 可选输入框 -->
                <div v-if="showInput" class="input-container">
                    <el-input v-model="inputValue" style="width: 100%" placeholder="输入账本名称" clearable
                        @keydown.enter="handleConfirm" />
                </div>

                <!-- 按钮容器 -->
                <div class="button-container">
                    <!-- 取消按钮 -->
                    <button class="dialog-button" :style="{ color: cancelColor }" @click="closeDialog">
                        {{ cancelLabel }}
                    </button>

                    <!-- 确认按钮 -->
                    <button class="dialog-button " :style="{ color: confirmColor }" @click="handleConfirm">
                        {{ confirmLabel }}
                    </button>
                </div>
            </div>
        </div>
    </Teleport>
</template>

<script setup>
import { ref, watch, nextTick, defineProps } from 'vue';

// --- Props 定义 ---
const props = defineProps({
    visible: {
        type: Boolean,
        required: true
    },
    message: {
        type: String,
        required: true
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
    onConfirm: {
        type: Function,
    }
});

// --- Emits 定义 ---
const emit = defineEmits(['update:visible', 'confirm']);

// --- 本地状态 ---
const inputValue = ref('');
const inputRef = ref(null);

// --- 监听器 ---
watch(() => props.visible, (newVal) => {
    if (newVal && props.showInput) {
        nextTick(() => {
            if (inputRef.value) {
                inputRef.value.focus();
            }
        });
    }
});

// --- 方法 ---
function closeDialog() {
    emit('update:visible', false);
    inputValue.value = '';
}

function handleConfirm() {
    if (typeof props.onConfirm === 'function') {
        props.onConfirm(props.showInput ? inputValue.value : undefined);
    }
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

/* 确认框主体样式 */
.confirm-dialog {
    background-color: white;
    border-radius: 8px;
    padding: 20px;
    margin-top: calc(100vh / 4);
    min-width: 300px;
}

/* 提示信息样式 */
.confirm-message {
    margin: 0 0 20px 0;
    font-size: 16px;
    text-align: left;
}

/* 输入框容器 */
.input-container {
    margin-bottom: 20px;
}

/* 输入框样式 */
.confirm-input {
    width: 100%;
    padding: 10px;
    border: 1px solid #ddd;
    border-radius: 4px;
    font-size: 14px;
    outline: none;
    box-sizing: border-box;
}

/* 浅色 placeholder */
.confirm-input::placeholder {
    color: #aaa;
    font-style: italic;
}

/* 按钮容器 */
.button-container {
    display: flex;
    justify-content: center;
    gap: 12px;
}

/* 通用按钮样式 */
.dialog-button {
    padding: 8px 16px;
    font-weight: 500;
    cursor: pointer;
    border: none;
    border-radius: 4px;
    background-color: transparent;
    transition: background-color 0.2s ease;
    outline: none;
    width: 100%;
}

/* 按钮 hover 效果 */
.dialog-button:hover {
    background-color: var(--billadm-color-icon-hover-bg-color);
}
</style>