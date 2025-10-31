<template>
  <a-table :columns="columns" :data-source="items" :pagination="false" :sticky="true" size="small">
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
        <a-button type="text" @click="handleEdit(record as TransactionRecord)" style="color: #1677ff">编辑</a-button>
        <a-popconfirm title="确认删除吗"
                      ok-text="确认"
                      @confirm="handleDelete(record as TransactionRecord)"
                      :showCancel="false">
          <a-button type="text" style="color: #f5222d">删除</a-button>
        </a-popconfirm>
      </template>
    </template>
  </a-table>
</template>

<script setup lang="ts">
import type {TransactionRecord} from '@/types/billadm';
import {formatTimestamp} from "@/backend/functions.ts";

const columns = [
  {
    title: '消费时间',
    dataIndex: 'transactionAt',
    width: 150,
    align: 'center'
  },
  {
    title: '交易类型',
    dataIndex: 'transactionType',
    width: 150,
    align: 'center'
  },
  {
    title: '消费类型',
    dataIndex: 'category',
    width: 150,
    align: 'center'
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
    dataIndex: 'price',
    width: 150,
    align: 'center'
  },
  {
    title: '操作',
    dataIndex: 'action',
    width: 150,
    align: 'center'
  }
];

interface Props {
  items: TransactionRecord[]
}

defineProps<Props>()

// 定义可触发的事件
const emit = defineEmits<{
  (e: 'edit', record: TransactionRecord): void;
  (e: 'delete', record: TransactionRecord): void;
}>();

// 处理编辑操作
const handleEdit = (record: TransactionRecord) => {
  emit('edit', record);
};

// 处理删除操作
const handleDelete = (record: TransactionRecord) => {
  emit('delete', record);
};

const formatTransactionType = (type: string): string => {
  const map: Record<string, string> = {
    expense: '支出',
    income: '收入',
    transfer: '转账'
  }
  return map[type] || type
}
</script>