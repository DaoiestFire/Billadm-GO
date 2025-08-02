<template>
  <div class="layout">
    <!-- 上栏：工具栏 -->
    <div class="top-bar">
      <div class="left-groups">
      </div>
      <div class="center-groups">
      </div>
      <div class="right-groups">
        <CommonIcon :icon="iconAdd" label="新增消费记录" width="40" height="40" :color="iconColor" :bgColor="minorBgColor"
          :hoverBgColor="hoverBgColor" hoverStyle="circle" tooltipPlacement="bottom-left" @click="showDialog = true" />
      </div>
    </div>

    <!-- 中间栏：消费记录表 -->
    <div class="middle-section">
      <TransactionRecordTable :items="sampleData" :columnStyles="columnStyles" :headerHeight="40" :rowHeight="40" />
    </div>

    <!-- 下栏：分页组件 -->
    <div class="bottom-bar">
      <CustomSelect v-model="maxRows" :options="options" direction="up" />
      <Pagination :pages="15" />
    </div>
  </div>
  <TransactionRecordOperation v-model="recordData" v-model:visible="showDialog" title="新增消费记录" :onConfirm="handleConfirm" />
</template>

<script setup>
import { ref, watch } from 'vue'
import { useCssVariables } from '@/css/css'
import iconAdd from '@/assets/icons/add.svg?raw'
import TransactionRecordTable from '@/components/TransactionRecordTable.vue'
import Pagination from '@/components/Pagination.vue'
import CustomSelect from '@/components/CustomSelect.vue'
import CommonIcon from '@/components/CommonIcon.vue'
import TransactionRecordOperation from '@/components/TransactionRecordOperation.vue'

// css variables
const { minorBgColor, hoverBgColor, iconColor } = useCssVariables()

// 表格最大行数
const maxRows = ref(10)

watch(() => maxRows.value,
  (newValue) => {
    console.log(newValue)
  },
  { immediate: true }
)



const options = [
  { label: '每页10行', value: 10 },
  { label: '每页20行', value: 20 },
  { label: '每页50行', value: 50 }
]

// 表格风格设置
const columnStyles = [
  {
    field: 'index',
    name: '序号',
    width: '100px',
  },
  {
    field: 'transaction_at',
    name: '交易时间',
    width: '200px',
  },
  {
    field: 'transaction_type',
    name: '交易类型',
    width: '100px',
  },
  {
    field: 'category',
    name: '消费类型',
    width: '100px',
  },
  {
    field: 'description',
    name: '描述',
    width: 'auto',
  },
  {
    field: 'tags',
    name: '标签',
    width: 'auto',
  },
  {
    field: 'price',
    name: '价格',
  },
  {
    field: 'actions',
    name: '操作',
    width: '120px',
  }
]

// 消费记录创建表单
const showDialog = ref(false);
const recordData = ref({});

function handleConfirm(data) {
  console.log('提交的数据：', data);
  // 发送请求等
}

// 表格数据
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
  display: flex;
  align-items: center;
  position: relative;
}

/* 左边按钮 将它与后面的元素隔开 */
.left-groups {
  margin-right: auto;
  display: flex;
  gap: 4px;
}

/* 中间按钮 */
.center-groups {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
}

/* 右边按钮组 */
.right-groups {
  display: flex;
}

.middle-section {
  flex: 1;
  overflow: hidden;
}

.bottom-bar {
  background: var(--billadm-color-minor-backgroud-color);
  height: 40px;
  border-top: 1px solid var(--billadm-color-window-border-color);
  flex-shrink: 0;
  display: flex;
  flex-direction: row;
  justify-content: center;
}
</style>