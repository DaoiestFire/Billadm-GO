<!-- components/da_view/BilladmChartPanel.vue -->
<template>
  <a-card :title="title" :body-style="bodyCss">
    <template #extra>
      <a-checkbox-group v-if="title==='交易走势'"
                        v-model:value="tradingTrendChecked"
                        :options="transactionTypeOptions"
      />
      <a-radio-group v-if="title=== '消费分布'"
                     v-model:value="transactionTypeChecked"
                     :options="transactionTypeOptions"
      />
    </template>
    <BilladmChart :option="option"/>
  </a-card>
</template>

<script setup lang="ts">
import {computed, type CSSProperties, ref} from 'vue';
import BilladmChart from "@/components/da_view/BilladmChart.vue";
import {buildOptionForTradingTrend, buildOptionForTransactionDistribution} from "@/backend/table.ts";
import type {TransactionRecord} from "@/types/billadm";

const bodyCss: CSSProperties = {
  aspectRatio: 3 / 2,
  minHeight: 0
}

const props = defineProps<{
  title: string
  data: TransactionRecord[]
}>();

const transactionTypeOptions = [
  {label: '收入', value: 'income'},
  {label: '支出', value: 'expense'},
  {label: '转账', value: 'transfer'},
];

const tradingTrendChecked = ref(['income', 'expense', 'transfer'])
const transactionTypeChecked = ref('expense')

const option = computed(() => {
  switch (props.title) {
    case '交易走势':
      return buildOptionForTradingTrend(props.data, tradingTrendChecked.value);
    case '消费分布':
      return buildOptionForTransactionDistribution(props.data, transactionTypeChecked.value);
    default:
      return buildOptionForTradingTrend(props.data, tradingTrendChecked.value);
  }
});
</script>