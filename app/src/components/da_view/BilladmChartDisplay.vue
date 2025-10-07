<template>
  <el-scrollbar>
    <div class="billadm-chart-display">
      <div class="charts-grid" :style="gridStyle">
        <billadm-fullscreen
            v-for="(chart, index) in charts"
            :key="index"
            v-model="chart.isFullscreen"
            :dblclick="true">
          <billadm-chart-panel
              :title="chart.title"
              :option="chart.option"
              :fullscreen="chart.isFullscreen"
          />
        </billadm-fullscreen>
      </div>
    </div>
  </el-scrollbar>
</template>

<script setup>
import {computed} from 'vue';
import BilladmChartPanel from "@/components/da_view/BilladmChartPanel.vue";
import BilladmFullscreen from "@/components/common/BilladmFullScreen.vue";

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
  trFormList: {
    type: Array,
    required: true
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