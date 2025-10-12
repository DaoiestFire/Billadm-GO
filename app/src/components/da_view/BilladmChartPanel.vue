<!-- components/da_view/BilladmChartPanel.vue -->
<template>
  <div class="panel-outer">
    <div class="panel-header">
      <h3 class="chart-title">{{ title }}</h3>
    </div>
    <div class="panel-container">
      <BilladmChart :option="option"/>
    </div>
    <div class="panel-footer">
      <el-checkbox-group v-if="title==='交易走势'" v-model="tradingTrendChecked">
        <el-checkbox label="收入" value="income"/>
        <el-checkbox label="支出" value="expense"/>
        <el-checkbox label="转账" value="transfer"/>
      </el-checkbox-group>
      <el-radio-group v-if="title=== '消费分布'" v-model="transactionTypeChecked">
        <el-radio value="income">收入</el-radio>
        <el-radio value="expense">支出</el-radio>
        <el-radio value="transfer">转账</el-radio>
      </el-radio-group>
    </div>
  </div>
</template>

<script setup>
import {computed, ref} from 'vue';
import BilladmChart from "@/components/da_view/BilladmChart.vue";
import {useCssVariables} from "@/css/css.js";
import {buildOptionForTradingTrend, buildOptionForTransactionDistribution} from "@/backend/table.js";

const {uiSizeMenuWidth} = useCssVariables();

const props = defineProps({
  title: {
    type: String,
    default: '图表标题',
    required: true
  },
  data: {
    type: Array,
    required: true
  }
});

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

<style scoped>
.panel-outer {
  background: var(--billadm-color-major-backgroud-color);
  border: 1px solid var(--billadm-color-window-border-color);
  border-radius: 8px;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  width: auto;
  height: auto;
}

/* header & footer */
.panel-header, .panel-footer {
  height: v-bind(uiSizeMenuWidth);
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 8px;
}

.panel-header {
  border-bottom: 1px solid var(--billadm-color-window-border-color);
}

.panel-footer {
  border-top: 1px solid var(--billadm-color-window-border-color);
}

.panel-container {
  flex: 1;
  position: relative;
  width: 100%;
  aspect-ratio: 3 / 2;
  min-height: 0;
}
</style>