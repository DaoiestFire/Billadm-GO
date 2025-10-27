<template>
  <a-layout style="height: 100vh">
    <a-layout-header class="headerStyle">
      <div class="left-groups">
        <a-segmented v-model:value="timeRangeTypeLabel" :options="TimeRangeTypeLabels"/>
        <a-button type="text">
          <template #icon>
            <LeftOutlined style="display: flex;justify-content: center;align-items: center;font-size: large"/>
          </template>
        </a-button>
        <a-range-picker v-model:value="timeRange" :picker="timeRangeType" :presets="TimeRangePresets"/>
        <a-button type="text">
          <template #icon>
            <RightOutlined style="display: flex;justify-content: center;align-items: center;font-size: large"/>
          </template>
        </a-button>
      </div>
      <div class="center-groups">
      </div>
      <div class="right-groups">
        <a-button type="primary" @click="onCreate">
          新增记录
        </a-button>
      </div>
    </a-layout-header>
    <a-layout-content :style="contentStyle">
      <transaction-record-table :items="trViewStore.tableData" :columnStyles="columnStyles" @edit-item="onEditItem"/>
    </a-layout-content>
    <a-layout-footer>
      <billadm-select v-model="trViewStore.pageSize" :options="PageSizeOptions" direction="up"/>
      <pagination v-model:current-page="trViewStore.currentPage" :pages="trViewStore.pages"/>
    </a-layout-footer>
  </a-layout>
  <transaction-record-operation v-model:modelValue="recordData" v-model:visible="showDialog" :title="dialogTitle"
                                :onConfirm="handleConfirm"/>
</template>

<script setup lang="ts">
import {computed, type CSSProperties, ref} from 'vue';
import TransactionRecordTable from '@/components/tr_view/TransactionRecordTable.vue';
import Pagination from '@/components/tr_view/Pagination.vue';
import TransactionRecordOperation from '@/components/tr_view/TransactionRecordOperation.vue';
import {useLedgerStore} from "@/stores/ledgerStore.ts";
import {useTrViewStore} from "@/stores/trViewStore.ts";
import NotificationUtil from "@/backend/notification.ts";
import {createTrForLedger, deleteTrById} from "@/backend/api/tr.ts";
import {PageSizeOptions, TimeRangePresets, TimeRangeTypeLabels, TimeRangeTypes} from "@/backend/constant.ts";
import {trDtoToTrForm, trFormToTrDto} from "@/backend/dto-utils.ts";
import type {TransactionRecord, TrForm} from "@/types/billadm";
import {useCssVariables} from "@/css/css.ts";
import {LeftOutlined, RightOutlined} from "@ant-design/icons-vue";
import type {Dayjs} from 'dayjs';

type RangeValue = [Dayjs, Dayjs];
type TimeRangeLabel = keyof typeof TimeRangeTypes;
type TimeRangeType = (typeof TimeRangeTypes)[TimeRangeLabel];

const timeRangeTypeLabel = ref<TimeRangeLabel>('日');
const timeRange = ref<RangeValue>();
const timeRangeType = computed<TimeRangeType>(() => {
  return TimeRangeTypes[timeRangeTypeLabel.value];
})

const {majorBgColor} = useCssVariables();

const contentStyle: CSSProperties = {
  backgroundColor: majorBgColor.value,
};

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

const onEditItem = (item: TransactionRecord) => {
  dialogTitle.value = '编辑消费记录';
  recordData.value = trDtoToTrForm(item);
  opType.value = 'edit';
  showDialog.value = true;
}

async function handleConfirm(data: TrForm) {
  try {
    const transactionRecord = trFormToTrDto(data, ledgerStore.currentLedgerId);
    await createTrForLedger(transactionRecord);
    if (opType.value === 'edit') {
      await deleteTrById(data.id);
    }
    await trViewStore.refreshPages();
    await trViewStore.refreshTableData();
    await trViewStore.refreshStatistics();
  } catch (error) {
    NotificationUtil.error(`消费记录操作失败 ${error}`);
  }
}
</script>

<style scoped>
.headerStyle {
  height: var(--billadm-size-header-height);
  background-color: var(--billadm-color-major-background);
  padding: 0;
  display: flex;
  align-items: start;
  justify-content: center;
}

/* 左边按钮 将它与后面的元素隔开 */
.left-groups {
  margin-right: auto;
  display: flex;
  gap: 8px;
  align-items: center;
}

/* 中间按钮 */
.center-groups {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  gap: 8px;
}

/* 右边按钮组 */
.right-groups {
  display: flex;
  gap: 8px;
}
</style>