<template>
  <div class="menu-bar">
    <div class="left-groups">
    </div>
    <div class="center-groups">
    </div>
    <div class="right-groups">
      <BilladmButton :icon="iconInfo" label="关于软件" :width="uiSizeMenuWidth" :height="uiSizeMenuWidth"
                     :color="iconColor" :bgColor="minorBgColor" :hoverBgColor="hoverBgColor" hoverStyle="circle"/>
    </div>
  </div>
</template>

<script setup>
import {useCssVariables} from '@/css/css';
import iconInfo from '@/assets/icons/info.svg?raw'
import BilladmButton from "@/components/BilladmButton.vue";

// css variables
const {minorBgColor, hoverBgColor, iconColor, uiSizeMenuWidth} = useCssVariables()

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