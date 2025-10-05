<template>
  <div class="layout">
    <!-- 上栏：工具栏 -->
    <div class="top-bar">
      <div class="left-groups">
        <billadm-icon-button :svg="iconLeft" label="向前一天" width="30px" height="30px" bg-size="26px"
                             :color="iconColor" bg-color="transparent" :hover-bg-color="hoverBgColor"
                             tooltipPlacement="bottom" @click="goToPreviousDay"/>
        <el-date-picker
            v-model="trViewStore.timeRange"
            type="daterange"
            range-separator="~"
            start-placeholder="起始时间"
            end-placeholder="结束时间"
            size="small"
            :editable="false"
            style="width: 200px;"
            :shortcuts="getShortcuts()"
        />
        <billadm-icon-button :svg="iconRight" label="向后一天" width="30px" height="30px" bg-size="26px"
                             :color="iconColor" bg-color="transparent" :hover-bg-color="hoverBgColor"
                             tooltipPlacement="bottom" @click="goToNextDay"/>
      </div>
      <div class="center-groups">
      </div>
      <div class="right-groups">
        <button class="btn-add" @click="onCreate">新增记录</button>
      </div>
    </div>

    <!-- 中间栏：消费记录表 -->
    <div class="middle-section">
      <transaction-record-table :items="trViewStore.tableData" :columnStyles="columnStyles" :headerHeight="40"
                                :rowHeight="40" @edit-item="onEditItem"/>
    </div>

    <!-- 下栏：分页组件 -->
    <div class="bottom-bar">
      <custom-select v-model="trViewStore.pageSize" :options="options" direction="up"/>
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
import CustomSelect from '@/components/CustomSelect.vue';
import TransactionRecordOperation from '@/components/tr_view/TransactionRecordOperation.vue';
import BilladmIconButton from "@/components/BilladmIconButton.vue";
import iconLeft from '@/assets/icons/left.svg?raw';
import iconRight from '@/assets/icons/right.svg?raw';
import {useLedgerStore} from "@/stores/ledgerStore.js";
import {useTrViewStore} from "@/stores/trViewStore.js";
import NotificationUtil from "@/backend/notification.js";
import {createTrForLedger, deleteTrById, trDtoToTrForm, trFormToTrDto} from "@/backend/tr.js";
import {getNextPeriod, getPrevPeriod, getShortcuts} from "@/backend/timerange.js";
import {useCssVariables} from "@/css/css.js";

// css variables
const {hoverBgColor, iconColor} = useCssVariables();

// store
const ledgerStore = useLedgerStore();
const trViewStore = useTrViewStore();

// 视图常量
const options = [
  {label: '每页10行', value: 10},
  {label: '每页20行', value: 20},
  {label: '每页50行', value: 50}
]

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