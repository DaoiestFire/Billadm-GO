<template>
  <div class="container">
    <a-typography-text type="secondary">
      收入
    </a-typography-text>
    <a-typography-text>
      {{ formatFloat(data.income) }}
    </a-typography-text>
    <a-typography-text type="secondary">
      支出
    </a-typography-text>
    <a-typography-text>
      {{ formatFloat(data.expense) }}
    </a-typography-text>
    <a-typography-text type="secondary">
      转账
    </a-typography-text>
    <a-typography-text>
      {{ formatFloat(data.transfer) }}
    </a-typography-text>
  </div>
</template>

<script setup lang="ts">
import {ref, watch} from 'vue';
import type {TrQueryCondition, TrStatistics} from "@/types/billadm";
import {useTrQueryConditionStore} from "@/stores/trQueryConditionStore.ts";
import {useLedgerStore} from "@/stores/ledgerStore.ts";
import {convertToUnixTimeRange} from "@/backend/timerange.ts";
import {formatFloat, getStatisticsOnCondition} from "@/backend/functions.ts";

const ledgerStore = useLedgerStore();
const trQueryConditionStore = useTrQueryConditionStore();

const data = ref<TrStatistics>({
  income: 0,
  expense: 0,
  transfer: 0
})

const refreshStatisticsData = async () => {
  if (!ledgerStore.currentLedgerId) return;
  const statisticsCondition: TrQueryCondition = {
    ledgerId: ledgerStore.currentLedgerId,
    tsRange: convertToUnixTimeRange(trQueryConditionStore.timeRange)
  };
  data.value = await getStatisticsOnCondition(statisticsCondition);
}

watch(() => [trQueryConditionStore.timeRange, ledgerStore.currentLedgerId], async () => {
      await refreshStatisticsData();
    },
    {immediate: true}
);
</script>

<style scoped>
.container {
  gap: 8px;
}
</style>