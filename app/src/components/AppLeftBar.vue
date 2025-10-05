<template>
  <div class="menu-bar">
    <BilladmIconButton :svg="iconBills" label="消费记录" :width="uiSizeMenuWidth" :height="uiSizeMenuWidth"
                       :color="iconColor" :hover-bg-color="hoverBgColor" :active-fg-color="iconActiveFgColor"
                       @click="navigate('tr_view')" :isActive="route.path === '/tr_view'"/>
    <BilladmIconButton :svg="iconStatistic" label="数据分析" :width="uiSizeMenuWidth" :height="uiSizeMenuWidth"
                       :color="iconColor" :hover-bg-color="hoverBgColor" :active-fg-color="iconActiveFgColor"
                       @click="navigate('da_view')" :isActive="route.path === '/da_view'"/>
    <BilladmIconButton :svg="iconSettings" label="应用设置" :width="uiSizeMenuWidth" :height="uiSizeMenuWidth"
                       :color="iconColor" :hover-bg-color="hoverBgColor" :active-fg-color="iconActiveFgColor"
                       @click="navigate('settings_view')" :isActive="route.path === '/settings_view'"/>
  </div>
</template>

<script setup>
import {useRoute, useRouter} from 'vue-router'
import {useCssVariables} from '@/css/css'
import iconBills from '@/assets/icons/bills.svg?raw'
import iconStatistic from '@/assets/icons/statistic.svg?raw'
import iconSettings from '@/assets/icons/settings.svg?raw'
import BilladmIconButton from '@/components/BilladmIconButton.vue'

// css variables
const {iconColor, hoverBgColor, iconActiveFgColor, uiSizeMenuWidth} = useCssVariables()

const router = useRouter()
const route = useRoute()

const navigate = (path) => {
  router.push(path);
}
</script>

<style scoped>
.menu-bar {
  display: flex;
  flex-direction: column;
  align-items: center;
  position: relative;
}
</style>