<template>
    <div class="menu-bar">
        <CommonIcon :icon="iconBills" width="40" height="40" :color="iconColor" :bgColor="minorBgColor"
            :hoverBgColor="hoverBgColor" hoverStyle="rect" />
        <CommonIcon :icon="iconStatistic" width="40" height="40" :color="iconColor" :bgColor="minorBgColor"
            :hoverBgColor="hoverBgColor" hoverStyle="rect" />
    </div>
</template>

<script setup>
import { useCssVariables } from '@/css/css'
import iconBills from '@/assets/icons/bills.svg?raw'
import iconStatistic from '@/assets/icons/statistic.svg?raw'
import CommonIcon from '@/components/CommonIcon.vue'

// css variables
const { minorBgColor, hoverBgColor, iconColor } = useCssVariables()
</script>

<style scoped>
.menu-bar {
    display: flex;
    flex-direction: column;
    align-items: center;
    position: relative;
}
</style>