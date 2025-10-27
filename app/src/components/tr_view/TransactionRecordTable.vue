<template>
  <a-table :columns="columns" :data-source="items">
    <template #bodyCell="{ column, record }">
      <template v-if="column.dataIndex==='transactionAt'">
        {{ formatTimestamp(record.transactionAt) }}
      </template>

      <template v-else-if="column.dataIndex==='transactionType'">
        {{ formatTransactionType(record.transactionType) }}
      </template>

      <template v-else-if="column.dataIndex === 'tags'">
        <span>
          <a-tag
              v-for="tag in record.tags"
              :key="tag"
              :color="tag === 'loser' ? 'volcano' : tag.length > 5 ? 'geekblue' : 'green'"
          >
            {{ tag }}
          </a-tag>
        </span>
      </template>

      <template v-else-if="column.dataIndex === 'action'">
        <a-button type="text">编辑</a-button>
        <a-button type="text">删除</a-button>
      </template>
    </template>
  </a-table>
</template>

<script setup lang="ts">
import type {TransactionRecord} from '@/types/billadm';
import {formatTimestamp} from "@/backend/functions.ts";

const columns = [
  // {
  //   title: '序号',
  //   dataIndex: 'index'
  // },
  {
    title: '消费时间',
    dataIndex: 'transactionAt'
  },
  {
    title: '交易类型',
    dataIndex: 'transactionType'
  },
  {
    title: '消费类型',
    dataIndex: 'category'
  },
  {
    title: '标签',
    dataIndex: 'tags'
  },
  {
    title: '描述',
    dataIndex: 'description'
  },
  {
    title: '价格',
    dataIndex: 'price'
  },
  {
    title: '操作',
    dataIndex: 'action'
  }
];

interface Props {
  items: TransactionRecord[]
}

defineProps<Props>()

// Emit 事件类型
// interface Emits {
//   (e: 'edit-item', item: TransactionRecord): void
// }


// 类型安全的 emit
// const emit = defineEmits<Emits>()
//
// const trViewStore = useTrViewStore()


// const getShowDeleteTrFunc = (id: string) => {
//   return (): void => {
//     message.value = '确认删除消费记录吗？'
//     confirmLabel.value = '删除'
//     confirmColor.value = negativeColor.value
//     cancelColor.value = positiveColor.value
//     confirmFunc.value = async (): Promise<void> => {
//       try {
//         await deleteTrById(id)
//         await trViewStore.refreshPages()
//         await trViewStore.refreshTableData()
//         await trViewStore.refreshStatistics()
//       } catch (error: unknown) {
//         const msg = error instanceof Error ? error.message : String(error)
//         NotificationUtil.error(`删除消费记录失败: ${msg}`)
//       }
//     }
//     showTrConfirmDialog.value = true
//   }
// }

// 交易类型对应颜色
// const formatTransactionTypeStyle = (type: string): Record<string, string> => {
//   const colors: Record<string, string> = {
//     expense: '#F56C6C',
//     income: '#67C23A',
//     transfer: '#409EFF'
//   }
//   return {color: colors[type] || ''}
// }

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