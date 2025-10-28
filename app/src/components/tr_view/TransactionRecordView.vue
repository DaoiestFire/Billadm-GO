<template>
  <a-layout style="height: 100%">
    <a-layout-header class="headerStyle">
      <div class="left-groups">
        <a-segmented v-model:value="timeRangeTypeLabel" :options="TimeRangeTypeLabels"/>
        <a-button type="text" @click="goToPrevious">
          <template #icon>
            <LeftOutlined style="display: flex;justify-content: center;align-items: center;font-size: large"/>
          </template>
        </a-button>
        <a-range-picker
            v-model:value="timeRange"
            :picker="timeRangeType"
            :presets="TimeRangePresets"
            inputReadOnly
        />
        <a-button type="text" @click="goToNext">
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
      <transaction-record-table :items="tableData"/>
    </a-layout-content>
    <a-layout-footer class="footerStyle">
      <a-pagination
          v-model:current="currentPage"
          v-model:pageSize="pageSize"
          :total="trTotal"
          :show-total="total => `共${total}记录`"
          show-size-changer
      />
    </a-layout-footer>
  </a-layout>
  <transaction-record-operation v-model:modelValue="recordData" v-model:visible="showDialog" :title="dialogTitle"/>
</template>

<script setup lang="ts">
import {computed, type CSSProperties, ref, watch} from 'vue';
import TransactionRecordTable from '@/components/tr_view/TransactionRecordTable.vue';
import TransactionRecordOperation from '@/components/tr_view/TransactionRecordOperation.vue';
import {TimeRangePresets, TimeRangeTypeLabels, TimeRangeTypes} from "@/backend/constant.ts";
import type {TransactionRecord, TrQueryCondition} from "@/types/billadm";
import {useCssVariables} from "@/css/css.ts";
import {LeftOutlined, RightOutlined} from "@ant-design/icons-vue";
import type {Dayjs} from 'dayjs';
import {convertToUnixTimeRange, getNextPeriod, getPrevPeriod, getTodayRange} from "@/backend/timerange.ts";
import {getTrOnCondition, getTrTotalOnCondition} from "@/backend/functions.ts";
import {useLedgerStore} from "@/stores/ledgerStore.ts";

const ledgerStore = useLedgerStore();

const {majorBgColor} = useCssVariables();

const contentStyle: CSSProperties = {
  backgroundColor: majorBgColor.value,
  "overflow-y": "auto",
  "margin-bottom": "auto"
};

// 时间区间
type RangeValue = [Dayjs, Dayjs];
type TimeRangeLabel = keyof typeof TimeRangeTypes;
type TimeRangeType = (typeof TimeRangeTypes)[TimeRangeLabel];

const timeRangeTypeLabel = ref<TimeRangeLabel>('日');
const timeRange = ref<RangeValue>(getTodayRange());
const timeRangeType = computed<TimeRangeType>(() => {
  return TimeRangeTypes[timeRangeTypeLabel.value];
})

const goToPrevious = () => {
  timeRange.value = getPrevPeriod(timeRange.value[0], timeRange.value[1], timeRangeType.value)
}

const goToNext = () => {
  timeRange.value = getNextPeriod(timeRange.value[0], timeRange.value[1], timeRangeType.value)
}

// 消费记录
const tableData = ref<TransactionRecord[]>([]);
// 分页
const currentPage = ref<number>(1);
const pageSize = ref<number>(20);
const trTotal = ref<number>(0);

const initTable = async () => {
  const trTotalCondition: TrQueryCondition = {
    ledgerId: ledgerStore.currentLedgerId,
    tsRange: convertToUnixTimeRange(timeRange.value)
  }
  trTotal.value = await getTrTotalOnCondition(trTotalCondition);
  const trCondition: TrQueryCondition = {
    ledgerId: ledgerStore.currentLedgerId,
    tsRange: convertToUnixTimeRange(timeRange.value),
    offset: pageSize.value * (currentPage.value - 1),
    limit: pageSize.value
  }
  tableData.value = await getTrOnCondition(trCondition);
}

watch([timeRange, currentPage, pageSize], () => {
  initTable();
}, {immediate: true});

// 消费记录表单
const showDialog = ref(false);
const recordData = ref({});
const opType = ref('create');
const dialogTitle = ref('');

const onCreate = () => {
  dialogTitle.value = '新建消费记录';
  recordData.value = {};
  opType.value = 'create';
  showDialog.value = true;
}
</script>

<style scoped>
.headerStyle {
  height: auto;
  background-color: var(--billadm-color-major-background);
  padding: 0 0 16px 0;
  display: flex;
  align-items: start;
  justify-content: center;
}

.footerStyle {
  height: auto;
  background-color: var(--billadm-color-major-background);
  padding: 16px 0 0 0;
  display: flex;
  align-items: end;
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