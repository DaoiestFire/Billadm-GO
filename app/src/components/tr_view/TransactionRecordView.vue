<template>
  <div class="layout">
    <!-- 上栏：工具栏 -->
    <div class="top-bar">
      <div class="left-groups">
        <billadm-time-select
            v-model="trViewStore.timeRange"
            v-model:timeRangeType="trViewStore.timeRangeType"
        />
      </div>
      <div class="center-groups">
      </div>
      <div class="right-groups">
        <button class="btn-add" @click="onCreate">新增记录</button>
      </div>
    </div>

    <!-- 中间栏：消费记录表 -->
    <div class="middle-section">
      <transaction-record-table :items="trViewStore.tableData" :columnStyles="columnStyles" @edit-item="onEditItem"/>
    </div>

    <!-- 下栏：分页组件 -->
    <div class="bottom-bar">
      <billadm-select v-model="trViewStore.pageSize" :options="getPageSizeOptions()" direction="up"/>
      <pagination v-model:current-page="trViewStore.currentPage" :pages="trViewStore.pages"/>
    </div>
  </div>
  <transaction-record-operation v-model:modelValue="recordData" v-model:visible="showDialog" :title="dialogTitle"
                                :onConfirm="handleConfirm"/>
</template>

<script setup>
import {ref} from 'vue';
import TransactionRecordTable from '@/components/tr_view/TransactionRecordTable.vue';
import Pagination from '@/components/tr_view/Pagination.vue';
import BilladmSelect from '@/components/BilladmSelect.vue';
import TransactionRecordOperation from '@/components/tr_view/TransactionRecordOperation.vue';
import BilladmTimeSelect from "@/components/BilladmTimeSelect.vue";
import {useLedgerStore} from "@/stores/ledgerStore.js";
import {useTrViewStore} from "@/stores/trViewStore.js";
import NotificationUtil from "@/backend/notification.js";
import {createTrForLedger, deleteTrById, trDtoToTrForm, trFormToTrDto} from "@/backend/tr.js";
import {getNextPeriod, getPageSizeOptions, getPrevPeriod,} from "@/backend/timerange.js";


// store
const ledgerStore = useLedgerStore();
const trViewStore = useTrViewStore();

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

// 消费记录表单
const showDialog = ref(false);
const recordData = ref({});
const opType = ref('create');
const dialogTitle = ref('');

// 功能函数
const onCreate = () => {
  dialogTitle.value = '新建消费记录';
  recordData.value = {};
  opType.value = 'create';
  showDialog.value = true;
}

const onEditItem = (item) => {
  dialogTitle.value = '编辑消费记录';
  recordData.value = trDtoToTrForm(item);
  opType.value = 'edit';
  showDialog.value = true;
}

async function handleConfirm(data) {
  try {
    const transactionRecord = trFormToTrDto(data, ledgerStore.currentLedgerId);
    await createTrForLedger(transactionRecord);
    if (opType.value === 'edit') {
      await deleteTrById(data.id);
    }
    await trViewStore.init();
    await trViewStore.refreshStatistics();
  } catch (error) {
    NotificationUtil.error(`消费记录操作失败 ${error}`);
  }
}

const goToPreviousDay = () => {
  let range = trViewStore.timeRange;
  if (!Array.isArray(range)) {
    return;
  }
  trViewStore.timeRange = getPrevPeriod(range[0], range[1]);
}

const goToNextDay = () => {
  let range = trViewStore.timeRange;
  if (!Array.isArray(range)) {
    return;
  }
  trViewStore.timeRange = getNextPeriod(range[0], range[1]);
}
</script>

<style scoped>
.layout {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.top-bar {
  background: var(--billadm-color-major-backgroud-color);
  height: 30px;
  flex-shrink: 0;
  display: flex;
  align-items: center;
  position: relative;
  margin-bottom: 10px;
}

/* 左边按钮 将它与后面的元素隔开 */
.left-groups {
  margin-right: auto;
  display: flex;
  gap: 4px;
  align-items: center;
}

/* 中间按钮 */
.center-groups {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  gap: 4px;
}

/* 右边按钮组 */
.right-groups {
  display: flex;
  gap: 4px;
}

.btn-add {
  background-color: transparent;
  color: var(--billadm-color-positive-color);
  padding: 4px 8px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.btn-add:hover {
  background-color: var(--billadm-color-positive-color);
  color: var(--billadm-color-hover-fg-color);
}

.middle-section {
  flex: 1;
  overflow: hidden;
}

.bottom-bar {
  background: var(--billadm-color-major-backgroud-color);
  height: 30px;
  flex-shrink: 0;
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  margin-top: 10px;
}
</style>