<template>
  <!-- 非全屏时正常渲染 -->
  <div v-if="!isFullscreen" class="billadm-chart-panel">
    <div class="panel-header">
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
    <div class="panel-container">
      <BilladmChart :option="option"/>
    </div>
    <div class="panel-footer">
    </div>
  </div>

  <!-- 全屏时使用 teleport 渲染到 body -->
  <teleport to="body" :disabled="!isFullscreen">
    <div v-if="isFullscreen" class="billadm-chart-panel fullscreen">
      <div class="chart-mask" @click.self="toggleFullscreen"></div>
      <div class="panel-wrapper">
        <div class="panel-header">
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
        <div class="panel-container">
          <BilladmChart :option="option"/>
        </div>
        <div class="panel-footer">
        </div>
      </div>
    </div>
  </teleport>
</template>

<script setup>
import {ref, watch} from 'vue';
import BilladmIconButton from "@/components/BilladmIconButton.vue";
import BilladmChart from "@/components/da_view/BilladmChart.vue";
import iconFullScreen from '@/assets/icons/full-screen.svg?raw';
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
.billadm-chart-panel {
  background: var(--billadm-color-major-backgroud-color);
  border: 1px solid var(--billadm-color-window-border-color);
  border-radius: 8px;
  overflow: hidden;
  position: relative;
}

.billadm-chart-panel.fullscreen {
  position: fixed;
  inset: 0;
  z-index: 1000;
  display: flex;
  justify-content: center;
  align-items: center;
}

.panel-wrapper {
  width: 80%;
  height: auto;
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

.panel-header, .panel-footer {
  height: var(--billadm-ui-size-menu-width);
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 8px;
}

.panel-header {
  border-bottom: 1px solid var(--billadm-color-window-border-color);
}

.panel-footer {
  border-top: 1px solid var(--billadm-color-window-border-color);
}

.panel-container {
  position: relative;
  width: 100%;
  aspect-ratio: 3 / 2;
  max-height: 100%;
}
</style>