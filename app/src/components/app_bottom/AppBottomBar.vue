<template>
  <div class="menu-bar">
    <div class="left-groups">
    </div>
    <div class="center-groups">
    </div>
    <div class="right-groups">
      <tr-view-bottom-statistics v-if="shouldShowStatistics"/>
      <billadm-modal
          v-model:visible="showAboutApp"
          title="关于"
          :message="`应用名称 ${appName}\n应用版本 ${appVersion}`"
      />
      <billadm-icon-button
          :svg="iconInfo"
          label="关于软件"
          :width="uiSizeMenuWidth"
          :height="uiSizeMenuWidth"
          :color="iconColor"
          :hover-bg-color="hoverBgColor"
          @click="showAboutApp=!showAboutApp"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, ref} from "vue"
import {useRoute} from "vue-router"
import BilladmIconButton from "@/components/BilladmIconButton.vue"
import TrViewBottomStatistics from "@/components/app_bottom/TrViewBottomStatistics.vue"
import BilladmModal from "@/components/BilladmModal.vue"
import iconInfo from '@/assets/icons/info.svg?raw'
import {useCssVariables} from '@/css/css.ts'


// css variables
const {iconColor, hoverBgColor, uiSizeMenuWidth} = useCssVariables()

const route = useRoute()

const shouldShowStatistics = computed(() => {
  return route.path === '/tr_view'
})

const showAboutApp = ref(false)
let appName = ''
let appVersion = ''

onMounted(async () => {
  appName = await window.electronAPI.getAppInfo('name')
  appVersion = await window.electronAPI.getAppInfo('version')
})
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
  gap: 8px;
}
</style>