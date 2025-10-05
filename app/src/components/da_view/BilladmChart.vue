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
      <v-chart
          :option="option"
          :autoresize="true"
          :style="{ height: '100%' }"
      />
    </div>
  </div>
</template>

<script setup>
import {ref, watch} from 'vue';
import VChart from 'vue-echarts';

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
  height: {
    type: String,
    default: '300px'
  },
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
      chartHeight.value = newVal ? '600px' : props.height;
    },
    {immediate: true}
);

// 切换全屏状态
const toggleFullscreen = () => {
  isFullscreen.value = !isFullscreen.value;
  emit('update:fullscreen', isFullscreen.value);
  chartHeight.value = isFullscreen.value ? '600px' : props.height;
};
</script>

<style scoped>
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
}
</style>