<!-- components/BilladmChart.vue -->
<template>
  <div class="billadm-chart" :class="{ fullscreen: isFullscreen }">
    <!-- 图表头部 -->
    <div class="chart-header">
      <h3 class="chart-title">{{ title }}</h3>
      <button
          class="fullscreen-btn"
          @click="toggleFullscreen"
          :aria-label="isFullscreen ? '退出全屏' : '全屏显示'"
      >
        {{ isFullscreen ? '×' : '⛶' }}
      </button>
    </div>

    <!-- 图表容器 -->
    <div class="chart-container" :style="{ height: chartHeight }">
      <!-- 使用 vue-echarts 组件 -->
      <v-chart
          :option="option"
          :autoresize="true"
          :style="{ height: '100%' }"
      />
    </div>
  </div>
</template>

<script setup>
import {nextTick, ref, watch} from 'vue';
// 引入 v-chart 组件和主题（可选）
import VChart from 'vue-echarts';
// 如果需要，可以引入 ECharts 的模块（例如地图、主题等）
// import 'echarts/lib/chart/line';
// import 'echarts/lib/component/tooltip';
// import 'echarts/lib/component/title';

// 接收 props
const props = defineProps({
  title: {
    type: String,
    default: '图表标题'
  },
  option: {
    type: Object,
    required: true
  },
  // 可选：自定义图表高度
  height: {
    type: String,
    default: '300px'
  },
  // 是否初始全屏（由父组件控制）
  fullscreen: {
    type: Boolean,
    default: false
  }
});

// 定义 emits 事件
const emit = defineEmits(['update:fullscreen']);

// 是否全屏状态
const isFullscreen = ref(props.fullscreen);

// 实际图表高度（考虑全屏时更高）
const chartHeight = ref(props.height);

// 监听父组件传入的 fullscreen 变化
watch(
    () => props.fullscreen,
    (newVal) => {
      isFullscreen.value = newVal;
      // 根据状态调整高度
      chartHeight.value = newVal ? '600px' : props.height;
    },
    {immediate: true}
);

// 切换全屏状态
const toggleFullscreen = () => {
  isFullscreen.value = !isFullscreen.value;
  // 向父组件同步状态
  emit('update:fullscreen', isFullscreen.value);
  // 调整高度
  chartHeight.value = isFullscreen.value ? '600px' : props.height;
  // 延迟，确保 DOM 更新后图表能正确 resize
  nextTick(() => {
    // vue-echarts 的 autoresize 已经处理了 resize，通常不需要手动触发
    // 但如果你有特殊需求，可以通过 ref 获取实例
  });
};

// 暴露给父组件的方法（可选）
// 如果需要访问底层 ECharts 实例
// const chartRef = ref(null);
// defineExpose({
//   getChart: () => chartRef.value?.$echarts,
//   refresh: () => {
//     // 由于 option 是响应式的，通常只需更新 option 即可
//     // 这里可以触发一些额外逻辑
//   }
// });
</script>

<style scoped>
/* 样式保持不变 */
.billadm-chart {
  background: #fff;
  border: 1px solid #e0e0e0;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  overflow: hidden;
  transition: all 0.3s ease;
}

.billadm-chart.fullscreen {
  position: fixed;
  top: 10%;
  left: 10%;
  width: 80% !important;
  height: 70%;
  z-index: 1000;
  border-radius: 12px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.3);
}

.chart-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background-color: #f8f9fa;
  border-bottom: 1px solid #e0e0e0;
}

.chart-title {
  margin: 0;
  font-size: 16px;
  font-weight: 600;
  color: #333;
}

.fullscreen-btn {
  background: none;
  border: none;
  font-size: 18px;
  cursor: pointer;
  color: #666;
  padding: 4px 8px;
  border-radius: 4px;
  transition: all 0.2s;
}

.fullscreen-btn:hover {
  background-color: #e0e0e0;
  color: #333;
}

.chart-container {
  width: 100%;
  /* 高度由父组件或自身状态控制 */
}
</style>