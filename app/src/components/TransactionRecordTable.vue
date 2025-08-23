<template>
  <div class="tr-table-container" :style="containerStyle">
    <el-scrollbar>
      <table class="tr-table">
        <thead>
        <tr :style="headerRowStyle">
          <th v-for="styleItem in columnStyles" :key="styleItem.field" :style="getColumnStyle(styleItem)">
            {{ styleItem.name }}
          </th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="(item, index) in items" :key="item.transaction_id" :style="rowStyle">
          <td v-for="styleItem in columnStyles" :key="styleItem.field"
              :style="formatCellStyle(styleItem.field, item)">
            <template v-if="styleItem.field === 'actions'">
              <div class="action-buttons">
                <button class="btn-edit" @click="handleEdit(item)">编辑</button>
                <button class="btn-delete" @click="getShowDeleteTrFunc(item.transaction_id)()">删除</button>
              </div>
            </template>
            <template v-if="styleItem.field === 'tags'">
              <el-space size="small"
                        style="display: flex; align-items: center; height: 100%; width: 100%; justify-content: center;">
                <el-tag v-for="tag in item[styleItem.field]" type="primary">{{ tag }}</el-tag>
              </el-space>
            </template>
            <template v-else>
              {{ styleItem.field === 'index' ? index + 1 : formatCell(styleItem.field, item[styleItem.field]) }}
            </template>
          </td>
        </tr>
        </tbody>
      </table>
    </el-scrollbar>
    <ConfirmDialog v-model:visible="showTrConfirmDialog" :message="message" :cancel-color="cancelColor"
                   :confirm-label="confirmLabel" :confirm-color="confirmColor" @confirm="confirmFunc"/>
  </div>
</template>

<script setup>
import {ref} from 'vue'
import {useCssVariables} from '@/css/css'
import {useTrViewStore} from "@/stores/trViewStore.js";
import {deleteTrById} from "@/backend/tr.js";
import NotificationUtil from "@/backend/notification.js";
import ConfirmDialog from "@/components/ConfirmDialog.vue";

// emit
const emit = defineEmits(['edit-item'])

// store
const trViewStore = useTrViewStore()

// 引用颜色
const {positiveColor, negativeColor} = useCssVariables()

const props = defineProps({
  items: {
    type: Array,
    required: true
  },
  headerHeight: {
    type: Number,
    default: 50
  },
  rowHeight: {
    type: Number,
    default: 40
  },
  columnStyles: {
    type: Array,
    default: []
  },
})

// 各种框的控制变量
const showTrConfirmDialog = ref(false)
const message = ref('')
const confirmLabel = ref('确认')
const confirmColor = ref('')
const cancelColor = ref('')
const confirmFunc = ref(null)

const handleEdit = (item) => {
  emit('edit-item', item)
}

const getShowDeleteTrFunc = (id) => {
  return () => {
    message.value = '确认删除消费记录吗？'
    confirmLabel.value = '删除'
    confirmColor.value = negativeColor.value
    cancelColor.value = positiveColor.value
    confirmFunc.value = async () => {
      try {
        await deleteTrById(id)
        await trViewStore.init()
      } catch (error) {
        NotificationUtil.error(`删除消费记录失败 ${error}`)
      }
    }
    showTrConfirmDialog.value = true
  }
}

// 获取表头风格
const getColumnStyle = (item) => {
  const width = item.width || 'auto'
  return {
    width,
    minWidth: width === 'auto' ? 'auto' : undefined,
    maxWidth: width === 'auto' ? 'none' : undefined
  }
}

// 获取表格样式
const formatCellStyle = (field, item) => {
  switch (field) {
    case 'transaction_type':
      return formatTransactionTypeStyle(item.transaction_type)
    default:
      return {}
  }
}

// 获取价格样式
const formatTransactionTypeStyle = (type) => {
  let color = ''
  if (type === 'expense') {
    color = '#F56C6C'
  } else if (type === 'income') {
    color = '#67C23A'
  } else if (type === 'transfer') {
    color = '#409EFF'
  }
  return {color}
}

// 格式化数据
const formatCell = (field, value) => {
  switch (field) {
    case 'transaction_at':
      return formatTime(value)
    case 'transaction_type':
      return formatTransactionType(value)
    default:
      return value
  }
}

// 格式化时间
const formatTime = (timestamp) => {
  if (!timestamp) return ''
  const date = new Date(timestamp * 1000)
  return date.toLocaleString()
}

// 格式化交易类型
const formatTransactionType = (type) => {
  return type === 'expense' ? '支出' : type === 'income' ? '收入' : '转账'
}

const containerStyle = {
  '--row-height': `${props.rowHeight}px`,
  '--header-height': `${props.headerHeight}px`,
}

const headerRowStyle = {
  height: 'var(--header-height)',
  lineHeight: 'var(--header-height)',
}

const rowStyle = {
  height: 'var(--row-height)',
  lineHeight: 'var(--row-height)',
}
</script>

<style scoped>
.action-buttons {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 4px;
  height: 100%;
}

.action-buttons button {
  padding: 4px 8px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.btn-edit {
  background-color: var(--billadm-color-icon-edit-bg-color);
  color: var(--billadm-color-icon-active-fg-color);
}

.btn-edit:hover {
  /* 假设 edit-bg-color 是浅色，hover 时加深 */
  background-color: color-mix(in srgb, var(--billadm-color-icon-edit-bg-color) 80%, #000);
  /* 或者你可以定义一个专门的 hover 变量，更可控 */
  transform: scale(1.05);
  /* 轻微放大，可选 */
}

.btn-delete {
  background-color: var(--billadm-color-icon-delete-bg-color);
  color: var(--billadm-color-icon-active-fg-color);
}

.btn-delete:hover {
  /* 删除按钮 hover 时加深红色 */
  background-color: color-mix(in srgb, var(--billadm-color-icon-delete-bg-color) 80%, #900);
  transform: scale(1.05);
}

.tr-table-container {
  overflow-y: auto;
  height: 100%;
}

.tr-table {
  width: 100%;
  border-collapse: collapse;
  table-layout: fixed;
}

.tr-table th,
.tr-table td {
  padding: 0;
  text-align: center;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  border-bottom: 1px solid var(--billadm-color-window-border-color);
}

.tr-table thead th {
  background-color: var(--billadm-color-table-header-bg-color);
  position: sticky;
  top: 0;
  z-index: 1;
}

/* 斑马纹样式 */
.tr-table tbody tr:nth-child(odd) {
  background-color: var(--billadm-color-table-odd-row-bg-color);
}

.tr-table tbody tr:nth-child(even) {
  background-color: var(--billadm-color-table-even-row-bg-color);
}
</style>