<template>
  <div class="menu-bar">
    <div class="left-groups">
      <AppDisplay size="30px"/>
      <LedgerSelect/>
    </div>
    <div class="center-groups">
      Billadm-{{ route.name }}
    </div>
    <div class="right-groups">
      <BilladmIconButton :svg="iconZoomOut" label="最小化" :width="uiSizeMenuWidth" :height="uiSizeMenuWidth"
                         :color="iconColor" :hover-bg-color="hoverBgColor"
                         tooltip-placement="bottom-start" @click="onMinimize"/>
      <BilladmIconButton :svg="iconZoomIn" label="最大化" :width="uiSizeMenuWidth" :height="uiSizeMenuWidth"
                         :color="iconColor" :hover-bg-color="hoverBgColor"
                         tooltip-placement="bottom-start" @click="onMaximize"/>
      <BilladmIconButton :svg="iconClose" label="关闭" :width="uiSizeMenuWidth" :height="uiSizeMenuWidth"
                         :color="iconColor" :hover-bg-color="hoverBgColor"
                         tooltip-placement="bottom-start" @click="onClose"/>
    </div>
  </div>
</template>

<script setup>
import {useRoute} from 'vue-router';
import BilladmIconButton from '@/components/BilladmIconButton.vue';
import LedgerSelect from "@/components/app_top/LedgerSelect.vue";
import AppDisplay from "@/components/app_top/AppDisplay.vue";
import {useCssVariables} from '@/css/css.js';
import iconZoomOut from '@/assets/icons/zoom-out.svg?raw';
import iconZoomIn from '@/assets/icons/zoom-in.svg?raw';
import iconClose from '@/assets/icons/close.svg?raw';
import {exitApp} from "@/backend/api/app.js";

// 当前视图
const route = useRoute()

// css variables
const {hoverBgColor, iconColor, uiSizeMenuWidth} = useCssVariables()

// 窗口控制
const onMinimize = () => {
  window.electronAPI.minimizeWindow();
}

const onMaximize = () => {
  window.electronAPI.maximizeWindow();
}

const onClose = async () => {
  exitApp()
  window.electronAPI.closeWindow();
}

</script>

<style scoped>
.menu-bar {
  -webkit-app-region: drag;
  display: flex;
  align-items: center;
  position: relative;
}

.menu-bar > * {
  -webkit-app-region: no-drag;
}

/* 左边按钮 将它与后面的元素隔开 */
.left-groups {
  margin-right: auto;
  display: flex;
  align-items: center;
}

/* 中间按钮 */
.center-groups {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
}

/* 右边按钮组 */
.right-groups {
  display: flex;
}
</style>