<template>
  <div class="billadm-chart-display">
    <div class="charts-grid" :style="gridStyle">
      <billadm-fullscreen
          v-for="(chart, index) in charts"
          :key="index"
          v-model="chart.isFullscreen"
          :dblclick="true">
        <billadm-chart-panel
            :title="chart.title"
            :data="chart.data"
        />
      </billadm-fullscreen>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed} from 'vue';
import BilladmChartPanel from "@/components/da_view/BilladmChartPanel.vue";
import BilladmFullscreen from "@/components/common/BilladmFullScreen.vue";
import type {TransactionRecord} from "@/types/billadm";

interface Props {
  trList: TransactionRecord[]
  columns?: number
}

const props = withDefaults(defineProps<Props>(), {
  columns: 2
})

const charts = computed(() => {
  const charts = [];
  charts.push({
    title: "交易走势",
    data: props.trList,
    isFullscreen: false
  });
  charts.push({
    title: "消费分布",
    data: props.trList,
    isFullscreen: false
  });
  return charts;
});

const gridStyle = computed(() => {
  return {
    display: 'grid',
    'grid-template-columns': `repeat(${props.columns}, 1fr)`,
    gap: '16px',
    width: '100%'
  };
});
</script>

<style scoped>
.billadm-chart-display {
  width: 100%;
  box-sizing: border-box;
}

/* 响应式：小屏幕下变为单列 */
@media (max-width: 768px) {
  .charts-grid {
    grid-template-columns: 1fr !important;
  }
}
</style>