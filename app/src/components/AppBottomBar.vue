<template>
  <div class="menu-bar">
    <div class="left-groups">
    </div>
    <div class="center-groups">
    </div>
    <div class="right-groups">
      <BilladmIconButton :svg="iconInfo" label="关于软件" :width="uiSizeMenuWidth" :height="uiSizeMenuWidth"
                         :color="iconColor" :hover-bg-color="hoverBgColor"/>
    </div>
  </div>
</template>

<script setup>
import {useCssVariables} from '@/css/css';
import iconInfo from '@/assets/icons/info.svg?raw'
import BilladmIconButton from "@/components/BilladmIconButton.vue";

// css variables
const {iconColor, hoverBgColor, uiSizeMenuWidth} = useCssVariables()

</script>

<style scoped>
.menu-bar {
  display: flex;
  align-items: center;
  position: relative;
}

/* 左边按钮 将它与后面的元素隔开 */
.left-groups {
  margin-right: auto;
  display: flex;
  gap: 4px;
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