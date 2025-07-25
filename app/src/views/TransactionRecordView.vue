<template>
  <div class="layout">
    <!-- 上栏：工具栏 -->
    <div class="top-bar">
    </div>

    <!-- 中间栏：消费记录表 -->
    <div class="middle-section">
      <TransactionRecordTable :items="sampleData" :headerHeight="40" :rowHeight="40" />
    </div>

    <!-- 下栏：分页组件 -->
    <div class="bottom-bar">
      <Pagination :pages="15"/>
    </div>
  </div>

</template>

<script setup>
import TransactionRecordTable from '@/components/TransactionRecordTable.vue'
import Pagination from '@/components/Pagination.vue';

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
    "tags": ["三餐", "test"],
    "transaction_at": "2025-07-19T00:15:32.563+08:00",
    "created_at": "2025-07-19T00:15:32.5755985+08:00",
    "updated_at": "2025-07-19T00:15:32.5755985+08:00"
  },
  {
    "transaction_id": "01981e51-c6df-716f-953f-dc8771370af8",
    "ledger_id": "01981e51-b02f-7730-9644-774491b459e4",
    "price": 4.5,
    "transaction_type": "income",
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
    "transaction_type": "transfer",
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
    "tags": ["三餐", "test"],
    "transaction_at": "2025-07-19T00:15:32.563+08:00",
    "created_at": "2025-07-19T00:15:32.5755985+08:00",
    "updated_at": "2025-07-19T00:15:32.5755985+08:00"
  },
  {
    "transaction_id": "01981e51-c6df-716f-953f-dc8771370af8",
    "ledger_id": "01981e51-b02f-7730-9644-774491b459e4",
    "price": 4.5,
    "transaction_type": "income",
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
    "transaction_type": "transfer",
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
</script>

<style scoped>
.layout {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.top-bar {
  background: var(--billadm-color-minor-backgroud-color);
  height: 40px;
  border-bottom: 1px solid var(--billadm-color-window-border-color);
  flex-shrink: 0;
}

.middle-section {
  flex: 1;
  overflow: auto;
}

.bottom-bar {
  background: var(--billadm-color-minor-backgroud-color);
  height: 40px;
  border-top: 1px solid var(--billadm-color-window-border-color);
  flex-shrink: 0;
  display: flex;
  flex-direction: column;
  justify-content: center;
}
</style>