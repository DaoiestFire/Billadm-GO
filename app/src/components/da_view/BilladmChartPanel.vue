<!-- components/da_view/BilladmChartPanel.vue -->
<template>
  <div class="panel-outer">
    <div class="panel-header">
      <h3 class="chart-title">{{ title }}</h3>
    </div>
    <div class="panel-container">
      <BilladmChart :option="option"/>
    </div>
    <div class="panel-footer"/>
  </div>
</template>

<script setup>
import {ref} from 'vue';
import BilladmChart from "@/components/da_view/BilladmChart.vue";
import {useCssVariables} from "@/css/css.js";
import {buildOptionForTradingTrend} from "@/backend/table.js";

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

const option = ref({});

const generateOption = (title, data) => {
  switch (title) {
    case '交易走势':
      return buildOptionForTradingTrend(data);
    default:
      return buildOptionForTradingTrend(data);
  }
}

option.value = generateOption(props.title, props.data);
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