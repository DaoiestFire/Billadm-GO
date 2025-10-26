<template>
  <div class="tr-table-container">
    <el-scrollbar>
      <table class="tr-table">
        <thead>
        <tr class="tr-table-header-row">
          <th v-for="styleItem in columnStyles" :key="styleItem.field" :style="getColumnStyle(styleItem)">
            {{ styleItem.name }}
          </th>
        </tr>
        </thead>
        <tbody>
        <tr class="tr-table-body-row" v-for="(item, index) in items" :key="item.transactionId">
          <td v-for="styleItem in columnStyles" :key="styleItem.field"
              :style="formatCellStyle(styleItem.field, item)">
            <template v-if="styleItem.field === 'actions'">
              <div class="action-buttons">
                <button class="btn-edit" @click="handleEdit(item)">编辑</button>
                <button class="btn-delete" @click="getShowDeleteTrFunc(item.transactionId)()">删除</button>
              </div>
            </template>
            <template v-if="styleItem.field === 'tags'">
              <el-space size="small"
                        style="display: flex; align-items: center; height: 100%; width: 100%; justify-content: center;">
                <el-tag v-for="tag in item[styleItem.field]" type="primary">{{ tag }}</el-tag>
              </el-space>
            </template>
            <template v-else>
              {{
                styleItem.field === 'index' ? (trViewStore.pageSize * (trViewStore.currentPage - 1)) + index + 1 : formatCell(styleItem.field, item[styleItem.field])
              }}
            </template>
          </td>
        </tr>
        </tbody>
      </table>
    </el-scrollbar>
    <billadm-modal
        v-model:visible="showTrConfirmDialog"
        :show-buttons="true"
        title="删除消费记录"
        :message="message"
        :cancel-color="cancelColor"
        :confirm-label="confirmLabel"
        :confirm-color="confirmColor"
        @confirm="handleConfirm"
    />
  </div>
</template>

<script setup lang="ts">
import {ref} from 'vue'
import {useCssVariables} from '@/css/css.ts'
import {useTrViewStore} from '@/stores/trViewStore.ts'
import {deleteTrById} from '@/backend/api/tr.ts'
import NotificationUtil from '@/backend/notification.ts'
import BilladmModal from '@/components/BilladmModal.vue'
import type {TransactionRecord} from '@/types/billadm'

// =======================
// 🔹 类型定义
// =======================

interface ColumnStyle {
  field: keyof TransactionRecord | 'index' | 'actions'
  name: string
  width: string
}

// Props 类型
interface Props {
  items: TransactionRecord[]
  headerHeight?: string
  rowHeight?: string
  columnStyles?: ColumnStyle[]
}

// Emit 事件类型
interface Emits {
  (e: 'edit-item', item: TransactionRecord): void
}

// =======================
// 🔹 Props & Emits
// =======================

// 使用 withDefaults 定义带默认值的 props（TS 安全）
withDefaults(defineProps<Props>(), {
  headerHeight: '50px',
  rowHeight: '40px',
  columnStyles: () => []
})

// 类型安全的 emit
const emit = defineEmits<Emits>()

// =======================
// 🔹 Store & 工具
// =======================

const trViewStore = useTrViewStore()
const {positiveColor, negativeColor} = useCssVariables()

// =======================
// 🔹 模态框控制
// =======================

const showTrConfirmDialog = ref(false)
const message = ref<string>('')
const confirmLabel = ref<string>('确认')
const confirmColor = ref<string>('')
const cancelColor = ref<string>('')
const confirmFunc = ref<(() => Promise<void>) | null>(null)

// =======================
// 🔹 事件处理
// =======================

const handleEdit = (item: TransactionRecord): void => {
  emit('edit-item', item)
}

const handleConfirm = async () => {
  if (confirmFunc.value) {
    await confirmFunc.value()
  }
}

const getShowDeleteTrFunc = (id: string) => {
  return (): void => {
    message.value = '确认删除消费记录吗？'
    confirmLabel.value = '删除'
    confirmColor.value = negativeColor.value
    cancelColor.value = positiveColor.value
    confirmFunc.value = async (): Promise<void> => {
      try {
        await deleteTrById(id)
        await trViewStore.refreshPages()
        await trViewStore.refreshTableData()
        await trViewStore.refreshStatistics()
      } catch (error: unknown) {
        const msg = error instanceof Error ? error.message : String(error)
        NotificationUtil.error(`删除消费记录失败: ${msg}`)
      }
    }
    showTrConfirmDialog.value = true
  }
}

// =======================
// 🔹 样式与格式化函数
// =======================

// 表头列样式
const getColumnStyle = (styleItem: ColumnStyle): Record<string, string | undefined> => {
  const width = styleItem.width || 'auto'
  return {
    width,
    minWidth: width === 'auto' ? 'auto' : undefined,
    maxWidth: width === 'auto' ? 'none' : undefined
  }
}

// 单元格样式（字段级）
const formatCellStyle = (field: string, item: TransactionRecord): Record<string, string> => {
  if (field === 'transaction_type') {
    return formatTransactionTypeStyle(item.transactionType)
  }
  return {}
}

// 交易类型对应颜色
const formatTransactionTypeStyle = (type: string): Record<string, string> => {
  const colors: Record<string, string> = {
    expense: '#F56C6C',
    income: '#67C23A',
    transfer: '#409EFF'
  }
  return {color: colors[type] || ''}
}

// 格式化单元格内容
const formatCell = (field: string, value: unknown): string | number => {
  switch (field) {
    case 'transaction_at':
      return formatTime(value as number)
    case 'transaction_type':
      return formatTransactionType(value as string)
    default:
      return value as string | number
  }
}

// 时间格式化
const formatTime = (timestamp: number): string => {
  if (!timestamp) return ''
  const date = new Date(timestamp * 1000)
  return date.toLocaleString()
}

// 交易类型中文映射
const formatTransactionType = (type: string): string => {
  const map: Record<string, string> = {
    expense: '支出',
    income: '收入',
    transfer: '转账'
  }
  return map[type] || type
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
  background-color: transparent;
  color: var(--billadm-color-positive-color);
}

.btn-edit:hover {
  background-color: var(--billadm-color-positive-color);
  color: var(--billadm-color-hover-fg-color);
}

.btn-delete {
  background-color: transparent;
  color: var(--billadm-color-negative-color);
}

.btn-delete:hover {
  background-color: var(--billadm-color-negative-color);
  color: var(--billadm-color-hover-fg-color);
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

.tr-table-header-row {
  height: v-bind(headerHeight);
  line-height: v-bind(headerHeight);
}

.tr-table-body-row {
  height: v-bind(rowHeight);
  line-height: v-bind(rowHeight);
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
}

.tr-table tbody tr:nth-child(odd) {
  background-color: var(--billadm-color-table-odd-row-bg-color);
}

.tr-table tbody tr:nth-child(even) {
  background-color: var(--billadm-color-table-even-row-bg-color);
}
</style>