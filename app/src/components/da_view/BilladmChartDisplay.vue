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
              :type="chart.type"
          />
        </billadm-fullscreen>
      </div>
    </div>
  </el-scrollbar>
</template>

<script setup>
import {computed, ref} from 'vue';
import BilladmChartPanel from "@/components/da_view/BilladmChartPanel.vue";
import BilladmFullscreen from "@/components/common/BilladmFullScreen.vue";

const props = defineProps({
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


// 处理传入的trFrom数据生成表格数据
const generateCharts = () => {
  return [];
}

const charts = ref(generateCharts());
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