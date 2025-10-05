<!-- components/BilladmChartDisplay.vue -->
<template>
  <div class="billadm-chart-display">
    <div class="charts-grid" :style="gridStyle">
      <BilladmChart
          v-for="(chart, index) in charts"
          :key="index"
          :title="chart.title"
          :option="chart.option"
          :height="chart.height || '300px'"
          :fullscreen="chart.isFullscreen"
          @update:fullscreen="handleFullscreenChange(index, $event)"
          class="chart-grid-item"
      />
    </div>
  </div>
</template>

<script setup>
import {computed} from 'vue';
import BilladmChart from "@/components/da_view/BilladmChart.vue";

const props = defineProps({
  charts: {
    type: Array,
    required: true,
    validator: (charts) => {
      return charts.every(chart =>
          typeof chart.title === 'string' &&
          typeof chart.option === 'object'
      );
    }
  },
  columns: {
    type: Number,
    default: 2
  }
});

const gridStyle = computed(() => {
  return {
    display: 'grid',
    'grid-template-columns': `repeat(${props.columns}, 1fr)`,
    gap: '20px',
    width: '100%'
  };
});

// 处理全屏切换
const handleFullscreenChange = (index, isFullscreen) => {
  props.charts[index].isFullscreen = isFullscreen;
  props.charts.forEach((chart, i) => {
    if (i !== index) {
      chart.isFullscreen = false;
    }
  });
};
</script>

<style scoped>
.billadm-chart-display {
  width: 100%;
  padding: 20px;
  box-sizing: border-box;
}

.charts-grid {
  width: 100%;
}

.chart-grid-item {
  /* 可以在这里添加额外的边距或样式 */
}

/* 响应式：小屏幕下变为单列 */
@media (max-width: 768px) {
  .charts-grid {
    grid-template-columns: 1fr !important;
  }
}
</style>