<template>
  <!-- 非全屏时正常渲染 -->
  <div v-if="!isFullscreen" class="billadm-chart">
    <div class="chart-header">
      <h3 class="chart-title">{{ title }}</h3>
      <BilladmIconButton
          :svg="iconFullScreen"
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

  <!-- 全屏时使用 teleport 渲染到 body -->
  <teleport to="body" :disabled="!isFullscreen">
    <div v-if="isFullscreen" class="billadm-chart billadm-chart-teleported fullscreen">
      <div class="chart-mask" @click.self="toggleFullscreen"></div>
      <div class="chart-wrapper">
        <div class="chart-header">
          <h3 class="chart-title">{{ title }}</h3>
          <BilladmIconButton
              :svg="iconOffScreen"
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
    </div>
  </teleport>
</template>

<script setup>
import {ref, watch} from 'vue';
import BilladmIconButton from "@/components/BilladmIconButton.vue";
import iconFullScreen from '@/assets/icons/full-screen.svg?raw';
import iconOffScreen from '@/assets/icons/off-screen.svg?raw';
import VChart from 'vue-echarts';
import {useCssVariables} from "@/css/css.js";

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
  overflow: hidden;
  position: relative;
}

/* 全屏样式（teleported） */
.billadm-chart-teleported {
  position: fixed;
  inset: 0;
  z-index: 1000;
  display: flex;
  justify-content: center;
  align-items: center;
}

.billadm-chart-teleported .chart-wrapper {
  width: 80%;
  height: 80%;
  background: var(--billadm-color-major-backgroud-color);
  border: 1px solid var(--billadm-color-window-border-color);
  border-radius: 16px;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

/* 遮罩层 */
.chart-mask {
  position: absolute;
  inset: 0;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: -1;
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