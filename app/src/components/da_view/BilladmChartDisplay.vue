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
import BilladmChart from './BilladmChart.vue'; // 注意导入新名称

// 定义图表数据结构
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
  // 可选：自定义列数
  columns: {
    type: Number,
    default: 2
  }
});

// 计算 grid 布局样式
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
  // 更新对应图表的全屏状态
  props.charts[index].isFullscreen = isFullscreen;

  // 可选：关闭其他图表的全屏状态（实现单个全屏）
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