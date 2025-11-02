<template>
  <a-layout style="height: 100%">
    <a-layout-header class="headerStyle">
      <div class="left-groups">
        <BilladmTimeRangePicker
            v-model:time-range="trQueryConditionStore.timeRange"
            v-model:time-range-type="trQueryConditionStore.timeRangeType"
        />
      </div>
      <div class="center-groups">
      </div>
      <div class="right-groups">
      </div>
    </a-layout-header>
    <a-layout-content :style="contentStyle">
      <billadm-chart-display :tr-list="trs"/>
    </a-layout-content>
  </a-layout>
</template>

<script setup lang="ts">
import {type CSSProperties, ref, watch} from 'vue';
import BilladmChartDisplay from "@/components/da_view/BilladmChartDisplay.vue";
import {useLedgerStore} from "@/stores/ledgerStore.ts";
import type {TransactionRecord, TrQueryCondition} from "@/types/billadm";
import {useTrQueryConditionStore} from "@/stores/trQueryConditionStore.ts";
import {useCssVariables} from "@/backend/css.ts";
import {convertToUnixTimeRange} from "@/backend/timerange.ts";
import {getTrOnCondition} from "@/backend/functions.ts";

const {majorBgColor} = useCssVariables();

const contentStyle: CSSProperties = {
  backgroundColor: majorBgColor.value,
  "overflow-y": "auto",
  "margin-bottom": "auto"
};

const ledgerStore = useLedgerStore();
const trQueryConditionStore = useTrQueryConditionStore();

const trs = ref<TransactionRecord[]>([]);

const refreshData = async () => {
  const trCondition: TrQueryCondition = {
    ledgerId: ledgerStore.currentLedgerId,
    tsRange: convertToUnixTimeRange(trQueryConditionStore.timeRange),
  }
  trs.value = await getTrOnCondition(trCondition);
};
/**
 * 查询条件，账本改变都需要刷新数据
 */
watch(() => [
      trQueryConditionStore.timeRange,
      ledgerStore.currentLedgerId,
    ],
    async () => {
      await refreshData();
    }, {immediate: true}
);
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