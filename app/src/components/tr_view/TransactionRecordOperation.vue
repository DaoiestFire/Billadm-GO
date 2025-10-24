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
                            style="width: 100%;"/>
          </div>

          <!-- 交易类型 -->
          <div class="form-item">
            <label>交易类型</label>
            <el-radio-group v-model="formData.type">
              <el-radio v-for="type in typeOptions" :key="type.value" :value="type.value">
                {{ type.label }}
              </el-radio>
            </el-radio-group>
          </div>

          <!-- 消费类型 -->
          <div class="form-item">
            <label>消费类型</label>
            <el-select v-model="formData.category" placeholder="选择消费类型" style="width: 100%">
              <el-option v-for="category in categories" :key="category" :value="category"/>
            </el-select>
          </div>

          <!-- 描述 -->
          <div class="form-item">
            <label>描述</label>
            <el-input v-model="formData.description" style="width: 100%" placeholder="简要说明" clearable/>
          </div>

          <!-- 标签（多选下拉） -->
          <div class="form-item">
            <label>标签</label>
            <el-select v-model="formData.tags" multiple placeholder="添加标签" style="width: 100%">
              <el-option v-for="tag in tags" :key="tag" :value="tag"/>
            </el-select>
          </div>

          <!-- 价格 -->
          <div class="form-item">
            <label>价格</label>
            <el-input-number v-model="formData.price" :min="0" style="width: 100%" :controls="false">
              <template #suffix>
                <span>RMB</span>
              </template>
            </el-input-number>
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

<script setup lang="ts">
import {computed, ref, watch} from 'vue';
import {useCategoryStore} from "@/stores/categoryStore.ts";
import {useTagStore} from "@/stores/tagStore.ts";
import {useTrViewStore} from "@/stores/trViewStore.ts";

// store
const categoryStore = useCategoryStore();
const tagStore = useTagStore();
const trViewStore = useTrViewStore();

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
      {label: '支出', value: 'expense'},
      {label: '收入', value: 'income'},
      {label: '转账', value: 'transfer'}
    ]
  },
  modelValue: {
    type: Object,
    default: () => ({})
  },
  onConfirm: {
    type: Function,
    default: null
  }
});

// --- Emits ---
const emit = defineEmits(['update:visible']);

// --- 本地状态 ---

const emptyData = {
  id: '',
  time: new Date(),
  type: '',
  category: '',
  description: '-',
  tags: [],
  price: 0,
}
const formData = ref(emptyData);

const categories = computed(() => {
  return categoryStore.getCategoryNamesByType(formData.value.type);
})

const tags = computed(() => {
  return tagStore.getTagNamesByCategory(formData.value.category);
})

// --- 监听器 ---
// 当 visible 变为 true 时，初始化表单数据
watch(
    () => props.visible,
    (newVal) => {
      if (newVal) {
        formData.value = {
          id: '',
          time: getFormDate(),
          type: 'expense',
          category: '',
          description: '-',
          tags: [],
          price: 0,
          ...props.modelValue
        };
      }
    },
    {immediate: false}
);

// 监听 formData.type 变化，自动设置 category（如果当前为空）
watch(
    () => formData.value.type,
    (newType) => {
      const availableCategories = categoryStore.getCategoryNamesByType(newType);
      if (availableCategories.length > 0) {
        if (!formData.value.category || !availableCategories.includes(formData.value.category)) {
          formData.value.category = availableCategories[0];
        }
      } else {
        formData.value.category = '';
      }
    },
    {immediate: true}
);

// 监听 formData.category 变化，自动清空无效 tags
watch(
    () => formData.value.category,
    (newCategory) => {
      const availableTags = tagStore.getTagNamesByCategory(newCategory);
      // 清理掉当前 tags 中不属于新 category 的标签
      formData.value.tags = formData.value.tags.filter(tag => availableTags.includes(tag));
    },
    {immediate: true}
);

// 如果选了时间范围以结束时间的12点作为消费时间，否则以当天12点作为消费时间
const getFormDate = () => {
  if (Array.isArray(trViewStore.timeRange) && trViewStore.timeRange.length === 2) {
    let ts = trViewStore.timeRange[1];
    let now = new Date();
    if (ts > now) {
      now.setHours(12, 0, 0, 0);
      return now;
    }
    ts.setHours(12, 0, 0, 0);
    return ts;
  }
  const noonToday = new Date();
  noonToday.setHours(12, 0, 0, 0);
  return noonToday;
}

// --- 方法 ---
function close() {
  formData.value = emptyData;
  emit('update:visible', false);
}

function handleSubmit() {
  if (typeof props.onConfirm === 'function') {
    props.onConfirm(formData.value);
  }
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
  font-size: var(--billadm-size-text-large-size);
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
  gap: 4px;
}

.form-item label {
  font-size: var(--billadm-size-text-base-size);
  color: var(--billadm-color-icon-color);
  font-weight: 500;
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