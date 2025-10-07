<!-- components/da_view/BilladmChartPanel.vue -->
<template>
  <BilladmFullscreen v-model="isFullscreen" :dblclick="true">
    <template #default="{ toggleFullscreen }">
      <div class="panel-outer">
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
        <div class="panel-footer"/>
      </div>
    </template>
  </BilladmFullscreen>
</template>

<script setup>
import {ref, watch} from 'vue';
import BilladmIconButton from "@/components/BilladmIconButton.vue";
import BilladmChart from "@/components/da_view/BilladmChart.vue";
import BilladmFullscreen from "@/components/common/BilladmFullscreen.vue"; // 引入新组件
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

watch(() => props.fullscreen, (newVal) => {
  isFullscreen.value = newVal;
}, {immediate: true});

watch(isFullscreen, (val) => {
  emit('update:fullscreen', val);
});
</script>

<style scoped>
.panel-outer {
  background: var(--billadm-color-major-backgroud-color);
  border: 1px solid var(--billadm-color-window-border-color);
  border-radius: 8px;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  width: auto;
  height: auto;
}

/* header & footer */
.panel-header, .panel-footer {
  height: v-bind(uiSizeMenuWidth);
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
  flex: 1;
  position: relative;
  width: 100%;
  aspect-ratio: 3 / 2;
  min-height: 0;
}
</style>