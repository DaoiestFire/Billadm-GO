<!-- TransactionRecordView.vue -->
<template>
    <!-- 使用 Teleport 将模态框插入 body 末尾 -->
    <Teleport to="body">
        <div v-if="visible" class="transaction-overlay" @click="close">
            <div class="transaction-dialog" @click.stop>
                <!-- 标题 -->
                <h2 class="dialog-title">{{ title }}</h2>

                <!-- 表单内容 -->
                <div class="form-container">
                    <!-- 交易时间 -->
                    <div class="form-item">
                        <label>交易时间</label>
                        <input v-model="formData.time" type="datetime-local" class="form-input" />
                    </div>

                    <!-- 交易类型 -->
                    <div class="form-item">
                        <label>交易类型</label>
                        <select v-model="formData.type" class="form-select">
                            <option v-for="type in typeOptions" :key="type.value" :value="type.value">
                                {{ type.label }}
                            </option>
                        </select>
                    </div>

                    <!-- 消费类型（可结合交易类型动态变化，此处为静态示例） -->
                    <div v-if="formData.type === 'expense'" class="form-item">
                        <label>消费类型</label>
                        <input v-model="formData.category" type="text" class="form-input" placeholder="如：餐饮、交通" />
                    </div>

                    <!-- 描述 -->
                    <div class="form-item">
                        <label>描述</label>
                        <input v-model="formData.description" type="text" class="form-input" placeholder="简要说明" />
                    </div>

                    <!-- 标签（多选下拉） -->
                    <div class="form-item">
                        <label>标签</label>
                        <select v-model="formData.tags" multiple class="form-select">
                            <option v-for="tag in tagOptions" :key="tag" :value="tag">
                                {{ tag }}
                            </option>
                        </select>
                    </div>

                    <!-- 价格 -->
                    <div class="form-item">
                        <label>价格</label>
                        <input v-model.number="formData.amount" type="number" step="0.01" min="0" class="form-input"
                            placeholder="请输入金额" />
                    </div>
                </div>

                <!-- 按钮区域 -->
                <div class="button-container">
                    <button class="dialog-button cancel" @click="close">{{ cancelLabel }}</button>
                    <button class="dialog-button confirm" :style="{ color: confirmColor }" @click="handleSubmit">
                        {{ confirmLabel }}
                    </button>
                </div>
            </div>
        </div>
    </Teleport>
</template>

<script setup>
import { ref, watch, defineProps, defineEmits } from 'vue';

// --- Props ---
const props = defineProps({
    visible: {
        type: Boolean,
        required: true
    },
    title: {
        type: String,
        default: '填写交易记录'
    },
    cancelLabel: {
        type: String,
        default: '取消'
    },
    confirmLabel: {
        type: String,
        default: '确认'
    },
    confirmColor: {
        type: String,
        default: '#007bff'
    },
    // 交易类型选项，可自定义
    typeOptions: {
        type: Array,
        default: () => [
            { label: '支出', value: 'expense' },
            { label: '收入', value: 'income' },
            { label: '转账', value: 'transfer' }
        ]
    },
    // 标签选项
    tagOptions: {
        type: Array,
        default: () => ['餐饮', '交通', '购物', '娱乐', '工资', '理财', '房租', '医疗']
    },
    // 初始表单数据
    modelValue: {
        type: Object,
        default: () => ({})
    },
    // 确认回调
    onConfirm: {
        type: Function
    }
});

// --- Emits ---
const emit = defineEmits(['update:visible', 'confirm', 'update:modelValue']);

// --- 本地状态 ---
const formData = ref({ ...props.modelValue });

// --- 监听器 ---
// 当 visible 变为 true 时，初始化表单数据
watch(
    () => props.visible,
    (newVal) => {
        if (newVal) {
            // 合并默认值与传入的 modelValue
            formData.value = {
                time: new Date().toISOString().slice(0, 16), // 默认当前时间
                type: 'expense',
                category: '',
                description: '',
                tags: [],
                amount: 0,
                ...props.modelValue
            };
        }
    }
);

// --- 方法 ---
function close() {
    emit('update:visible', false);
}

function handleSubmit() {
    // 调用 onConfirm 回调
    if (typeof props.onConfirm === 'function') {
        props.onConfirm(formData.value);
    }

    // 通过 emit 更新 v-model
    emit('update:modelValue', formData.value);
    emit('confirm', formData.value);

    close();
}
</script>

<style scoped>
.transaction-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
}

.transaction-dialog {
    background-color: white;
    border-radius: 8px;
    padding: 24px;
    width: 90%;
    max-width: 500px;
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
}

.dialog-title {
    font-size: 20px;
    font-weight: bold;
    text-align: left;
    margin: 0 0 20px 0;
    color: #333;
}

.form-container {
    display: flex;
    flex-direction: column;
    gap: 16px;
}

.form-item {
    display: flex;
    flex-direction: column;
    gap: 6px;
}

.form-item label {
    font-size: 14px;
    color: #555;
    font-weight: 500;
}

.form-input,
.form-select {
    padding: 10px 12px;
    border: 1px solid #ddd;
    border-radius: 6px;
    font-size: 14px;
    outline: none;
    background-color: #fff;
}

.form-select {
    height: auto;
}

.form-select[multiple] {
    height: 100px;
    /* 多选下拉框高度 */
}

.form-input:focus,
.form-select:focus {
    border-color: v-bind(confirmColor);
    box-shadow: 0 0 0 2px rgba(0, 123, 255, 0.2);
}

.button-container {
    display: flex;
    justify-content: flex-end;
    gap: 12px;
    margin-top: 24px;
}

.dialog-button {
    padding: 8px 16px;
    border: none;
    border-radius: 6px;
    font-size: 14px;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.2s ease;
}

.dialog-button.cancel {
    background-color: #f1f1f1;
    color: #666;
}

.dialog-button.cancel:hover {
    background-color: #e0e0e0;
}

.dialog-button.confirm {
    background-color: transparent;
}

.dialog-button.confirm:hover {
    text-decoration: underline;
}
</style>