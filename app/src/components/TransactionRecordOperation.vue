<!-- TransactionRecordView.vue -->
<template>
    <!-- 使用 Teleport 将模态框插入 body 末尾 -->
    <Teleport to="body">
        <div v-if="visible" class="transaction-overlay" @click="close">
            <div class="transaction-dialog" @click.stop>
                <!-- 标题 -->
                <span class="dialog-title">{{ title }}</span>

                <!-- 表单内容 -->
                <div class="form-container">
                    <!-- 交易时间 -->
                    <div class="form-item">
                        <label>交易时间</label>
                        <el-date-picker v-model="formData.time" type="datetime" placeholder="选择交易时间"
                            style="width: 100%;" />
                    </div>

                    <!-- 交易类型 -->
                    <div class="form-item">
                        <label>交易类型</label>
                        <el-select v-model="formData.type" placeholder="选择交易类型" style="width: 100%">
                            <el-option v-for="type in typeOptions" :key="type.value" :label="type.label"
                                :value="type.value" />
                        </el-select>
                    </div>

                    <!-- 消费类型 -->
                    <div class="form-item">
                        <label>消费类型</label>
                        <el-select v-model="formData.category" placeholder="选择消费类型" style="width: 100%">
                            <el-option v-for="category in categoryOptions" :key="category" :value="category" />
                        </el-select>
                    </div>

                    <!-- 描述 -->
                    <div class="form-item">
                        <label>描述</label>
                        <el-input v-model="formData.description" style="width: 100%" placeholder="简要说明" clearable />
                    </div>

                    <!-- 标签（多选下拉） -->
                    <div class="form-item">
                        <label>标签</label>
                        <el-select v-model="formData.tags" multiple placeholder="添加标签" style="width: 100%">
                            <el-option v-for="tag in tagOptions" :key="tag" :value="tag" />
                        </el-select>
                    </div>

                    <!-- 价格 -->
                    <div class="form-item">
                        <label>价格</label>
                        <el-input-number v-model="formData.amount" :min="0" style="width: 100%" />
                    </div>
                </div>

                <!-- 按钮区域 -->
                <div class="button-container">
                    <button class="dialog-button cancel" @click="close">取消</button>
                    <button class="dialog-button confirm" @click="handleSubmit">确认</button>
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
        default: '自定义标题'
    },
    typeOptions: {
        type: Array,
        default: () => [
            { label: '支出', value: 'expense' },
            { label: '收入', value: 'income' },
            { label: '转账', value: 'transfer' }
        ]
    },
    categoryOptions: {
        type: Array,
        default: () => ['餐饮', '交通', '购物', '娱乐', '工资', '理财', '房租', '医疗']
    },
    tagOptions: {
        type: Array,
        default: () => ['餐饮', '交通', '购物', '娱乐', '工资', '理财', '房租', '医疗']
    },
    modelValue: {
        type: Object,
        default: () => ({})
    },
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
}

.dialog-title {
    font-size: 20px;
    font-weight: bold;
    text-align: left;
    margin-bottom: 24px;
    display: block;
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
    color: var(--billadm-color-icon-color);
    font-weight: 500;
}

.form-input,
.form-select {
    padding: 10px 12px;
    border: 1px solid var(--billadm-color-window-border-color);
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
}

.form-input:focus,
.form-select:focus {
    border-color: var(--billadm-color-positive-color);
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
    font-weight: 500;
    cursor: pointer;
    border: none;
    border-radius: 4px;
    background-color: transparent;
    transition: background-color 0.2s ease;
    outline: none;
}

.dialog-button.cancel {
    color: var(--billadm-color-negative-color);
}

.dialog-button.confirm {
    color: var(--billadm-color-positive-color);
}

/* 按钮 hover 效果 */
.dialog-button:hover {
    background-color: var(--billadm-color-icon-hover-bg-color);
}
</style>