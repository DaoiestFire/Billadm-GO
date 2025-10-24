<template>
  <div class="menu-bar">
    <div class="top-groups">
      <BilladmIconButton :svg="iconBills" label="消费记录" tooltip-placement="right"
                         :width="uiSizeMenuWidth" :height="uiSizeMenuWidth"
                         :color="iconColor" :hover-bg-color="hoverBgColor" :active-fg-color="iconActiveFgColor"
                         @click="navigate('tr_view')" :isActive="route.path === '/tr_view'"/>
      <BilladmIconButton :svg="iconStatistic" label="数据分析" tooltip-placement="right"
                         :width="uiSizeMenuWidth" :height="uiSizeMenuWidth"
                         :color="iconColor" :hover-bg-color="hoverBgColor" :active-fg-color="iconActiveFgColor"
                         @click="navigate('da_view')" :isActive="route.path === '/da_view'"/>
    </div>
    <div class="bottom-groups">
      <BilladmIconButton :svg="iconSettings" label="应用设置" tooltip-placement="right"
                         :width="uiSizeMenuWidth" :height="uiSizeMenuWidth"
                         :color="iconColor" :hover-bg-color="hoverBgColor" :active-fg-color="iconActiveFgColor"
                         @click="navigate('settings_view')" :isActive="route.path === '/settings_view'"/>
    </div>
  </div>
</template>

<script setup lang="ts">
import {useRoute, useRouter} from 'vue-router'
import {useCssVariables} from '@/css/css.js'
import iconBills from '@/assets/icons/bills.svg?raw'
import iconStatistic from '@/assets/icons/statistic.svg?raw'
import iconSettings from '@/assets/icons/settings.svg?raw'
import BilladmIconButton from '@/components/BilladmIconButton.vue'

// css variables
const {iconColor, hoverBgColor, iconActiveFgColor, uiSizeMenuWidth} = useCssVariables()

const router = useRouter()
const route = useRoute()

const navigate = (path: string) => {
  router.push(path)
}
</script>

<style scoped>
.menu-bar {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.top-groups {
  margin-bottom: auto;
  display: flex;
  flex-direction: column;
}

.bottom-groups {
  display: flex;
  flex-direction: column;
}
</style>