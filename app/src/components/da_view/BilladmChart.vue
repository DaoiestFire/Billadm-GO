<template>
  <div class="billadm-chart" :class="{ fullscreen: isFullscreen }">
    <!-- 图表头部 -->
    <div class="chart-header">
      <h3 class="chart-title">{{ title }}</h3>
      <BilladmIconButton
          :svg="isFullscreen?iconOffScreen:iconFullScreen"
          :width="uiSizeMenuWidth" :height="uiSizeMenuWidth"
          :color="iconColor" :hover-bg-color="hoverBgColor" :active-fg-color="iconActiveFgColor"
          @click="toggleFullscreen"
      />
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
import BilladmIconButton from "@/components/BilladmIconButton.vue";
import iconFullScreen from '@/assets/icons/full-screen.svg?raw';
import iconOffScreen from '@/assets/icons/off-screen.svg?raw';
import VChart from 'vue-echarts';
import {useCssVariables} from "@/css/css.js";

// css variables
const {iconColor, hoverBgColor, iconActiveFgColor, uiSizeMenuWidth} = useCssVariables()

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
  background: var(--billadm-color-major-backgroud-color);
  border: 1px solid var(--billadm-color-window-border-color);
  border-radius: 8px;
  overflow: hidden;
  transition: all 0.3s ease;
}

.billadm-chart.fullscreen {
  position: fixed;
  top: 10%;
  left: 10%;
  width: 80% !important;
  height: 80%;
  border-radius: 16px;
}

.chart-header {
  height: var(--billadm-ui-size-menu-width);
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 8px;
  background-color: var(--billadm-color-minor-backgroud-color);
}

.chart-title {
  font-size: var(--billadm-size-text-title-size);
  color: var(--billadm-color-text-major-color);
}

.chart-container {
  width: 100%;
}
</style>