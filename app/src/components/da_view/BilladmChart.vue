<template>
  <div class="billadm-chart" :class="{ fullscreen: isFullscreen }">
    <div class="chart-header">
      <h3 class="chart-title">{{ title }}</h3>
      <BilladmIconButton
          :svg="isFullscreen ? iconOffScreen : iconFullScreen"
          :width="uiSizeMenuWidth"
          :height="uiSizeMenuWidth"
          :color="iconColor"
          :hover-bg-color="hoverBgColor"
          :active-fg-color="iconActiveFgColor"
          @click="toggleFullscreen"
      />
    </div>
    <div class="chart-container">
      <v-chart
          ref="chartRef"
          :option="option"
          :autoresize="true"
          class="billadm-chart-instance"
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
const {iconColor, hoverBgColor, iconActiveFgColor, uiSizeMenuWidth} = useCssVariables();

const props = defineProps({
  title: {
    type: String,
    default: '图表标题'
  },
  option: {
    type: Object,
    required: true
  },
  fullscreen: {
    type: Boolean,
    default: false
  }
});

const emit = defineEmits(['update:fullscreen']);

const isFullscreen = ref(props.fullscreen);

const chartRef = ref(null);

watch(
    () => props.fullscreen,
    (newVal) => {
      isFullscreen.value = newVal;
    },
    {immediate: true}
);

const toggleFullscreen = () => {
  emit('update:fullscreen', !isFullscreen.value);
};
</script>

<style scoped>
.billadm-chart {
  background: var(--billadm-color-major-backgroud-color);
  border: 1px solid var(--billadm-color-window-border-color);
  border-radius: 8px;
  transition: all 0.3s ease;
}

.billadm-chart.fullscreen {
  position: fixed;
  top: 10%;
  left: 10%;
  width: 80%;
  height: 80%;
  border-radius: 16px;
  z-index: 4;
}

.chart-header {
  height: var(--billadm-ui-size-menu-width);
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 8px;
  border-bottom: 1px solid var(--billadm-color-window-border-color);
}

.chart-title {
  font-size: var(--billadm-size-text-title-size);
  color: var(--billadm-color-text-major-color);
}

.chart-container {
  position: relative;
  width: 100%;
  aspect-ratio: 3 / 2;
  max-height: 100%;
}

.billadm-chart .billadm-chart-instance {
  position: absolute !important;
  width: 100% !important;
  height: 100% !important;
}
</style>