<template>
  <div class="menu-bar">
    <BilladmButton :icon="iconBills" label="消费记录" :width="uiSizeMenuWidth" :height="uiSizeMenuWidth"
                   :color="iconColor" :bgColor="minorBgColor" :hoverBgColor="hoverBgColor" hoverStyle="circle"
                   @click="navigate('tr_view')" :isActive="route.path === '/tr_view'"/>
    <BilladmButton :icon="iconStatistic" label="数据分析" :width="uiSizeMenuWidth" :height="uiSizeMenuWidth"
                   :color="iconColor" :bgColor="minorBgColor" :hoverBgColor="hoverBgColor" hoverStyle="circle"
                   @click="navigate('da_view')" :isActive="route.path === '/da_view'"/>
  </div>
</template>

<script setup>
import {useRoute, useRouter} from 'vue-router'
import {useCssVariables} from '@/css/css'
import iconBills from '@/assets/icons/bills.svg?raw'
import iconStatistic from '@/assets/icons/statistic.svg?raw'
import BilladmButton from '@/components/BilladmButton.vue'

// css variables
const {minorBgColor, hoverBgColor, iconColor, uiSizeMenuWidth} = useCssVariables()

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