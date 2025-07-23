<template>
  <BillTable ref="billTableRef" />
  <button @click="loadData">加载数据</button>
</template>

<script setup>
import { ref } from 'vue'
import BillTable from '@/components/BillTable.vue'

const billTableRef = ref()

const sampleData = [
  {
    "transaction_id": "01981e51-c6df-716f-953f-dc8771370af8",
    "ledger_id": "01981e51-b02f-7730-9644-774491b459e4",
    "price": 4.5,
    "transaction_type": "expense",
    "category": "餐饮美食",
    "description": "早餐",
    "tags": ["三餐"],
    "transaction_at": "2025-07-19T00:15:32.563+08:00",
    "created_at": "2025-07-19T00:15:32.5755985+08:00",
    "updated_at": "2025-07-19T00:15:32.5755985+08:00"
  },
  {
    "transaction_id": "01981e51-c6df-716f-953f-dc8771370af8",
    "ledger_id": "01981e51-b02f-7730-9644-774491b459e4",
    "price": 4.5,
    "transaction_type": "expense",
    "category": "餐饮美食",
    "description": "早餐",
    "tags": ["三餐"],
    "transaction_at": "2025-07-19T00:15:32.563+08:00",
    "created_at": "2025-07-19T00:15:32.5755985+08:00",
    "updated_at": "2025-07-19T00:15:32.5755985+08:00"
  },
  {
    "transaction_id": "01981e51-c6df-716f-953f-dc8771370af8",
    "ledger_id": "01981e51-b02f-7730-9644-774491b459e4",
    "price": 4.5,
    "transaction_type": "expense",
    "category": "餐饮美食",
    "description": "早餐",
    "tags": ["三餐"],
    "transaction_at": "2025-07-19T00:15:32.563+08:00",
    "created_at": "2025-07-19T00:15:32.5755985+08:00",
    "updated_at": "2025-07-19T00:15:32.5755985+08:00"
  },
  {
    "transaction_id": "01981e51-c6df-716f-953f-dc8771370af8",
    "ledger_id": "01981e51-b02f-7730-9644-774491b459e4",
    "price": 4.5,
    "transaction_type": "expense",
    "category": "餐饮美食",
    "description": "早餐",
    "tags": ["三餐"],
    "transaction_at": "2025-07-19T00:15:32.563+08:00",
    "created_at": "2025-07-19T00:15:32.5755985+08:00",
    "updated_at": "2025-07-19T00:15:32.5755985+08:00"
  },
  {
    "transaction_id": "01981e51-c6df-716f-953f-dc8771370af8",
    "ledger_id": "01981e51-b02f-7730-9644-774491b459e4",
    "price": 4.5,
    "transaction_type": "expense",
    "category": "餐饮美食",
    "description": "早餐",
    "tags": ["三餐"],
    "transaction_at": "2025-07-19T00:15:32.563+08:00",
    "created_at": "2025-07-19T00:15:32.5755985+08:00",
    "updated_at": "2025-07-19T00:15:32.5755985+08:00"
  },
  {
    "transaction_id": "01981e51-c6df-716f-953f-dc8771370af8",
    "ledger_id": "01981e51-b02f-7730-9644-774491b459e4",
    "price": 4.5,
    "transaction_type": "expense",
    "category": "餐饮美食",
    "description": "早餐",
    "tags": ["三餐"],
    "transaction_at": "2025-07-19T00:15:32.563+08:00",
    "created_at": "2025-07-19T00:15:32.5755985+08:00",
    "updated_at": "2025-07-19T00:15:32.5755985+08:00"
  },
  {
    "transaction_id": "01981e51-c6df-716f-953f-dc8771370af8",
    "ledger_id": "01981e51-b02f-7730-9644-774491b459e4",
    "price": 4.5,
    "transaction_type": "expense",
    "category": "餐饮美食",
    "description": "早餐",
    "tags": ["三餐"],
    "transaction_at": "2025-07-19T00:15:32.563+08:00",
    "created_at": "2025-07-19T00:15:32.5755985+08:00",
    "updated_at": "2025-07-19T00:15:32.5755985+08:00"
  },
  {
    "transaction_id": "01981e51-c6df-716f-953f-dc8771370af8",
    "ledger_id": "01981e51-b02f-7730-9644-774491b459e4",
    "price": 4.5,
    "transaction_type": "expense",
    "category": "餐饮美食",
    "description": "早餐",
    "tags": ["三餐"],
    "transaction_at": "2025-07-19T00:15:32.563+08:00",
    "created_at": "2025-07-19T00:15:32.5755985+08:00",
    "updated_at": "2025-07-19T00:15:32.5755985+08:00"
  },
  {
    "transaction_id": "01981e51-c6df-716f-953f-dc8771370af8",
    "ledger_id": "01981e51-b02f-7730-9644-774491b459e4",
    "price": 4.5,
    "transaction_type": "expense",
    "category": "餐饮美食",
    "description": "早餐",
    "tags": ["三餐"],
    "transaction_at": "2025-07-19T00:15:32.563+08:00",
    "created_at": "2025-07-19T00:15:32.5755985+08:00",
    "updated_at": "2025-07-19T00:15:32.5755985+08:00"
  },
]

const loadData = () => {
  billTableRef.value.setBillList(sampleData)
}
</script>