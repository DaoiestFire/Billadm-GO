<template>
    <div class="menu-bar">
        <CommonIcon :icon="iconBills" label="消费记录" width="40" height="40" :color="iconColor" :bgColor="minorBgColor"
            :hoverBgColor="hoverBgColor" hoverStyle="circle" @click="navigate('home')" />
        <CommonIcon :icon="iconStatistic" label="数据分析" width="40" height="40" :color="iconColor"
            :bgColor="minorBgColor" :hoverBgColor="hoverBgColor" hoverStyle="circle" @click="navigate('dashboard')" />
    </div>
</template>

<script setup>
import { useRouter } from 'vue-router';
import { useCssVariables } from '@/css/css'
import iconBills from '@/assets/icons/bills.svg?raw'
import iconStatistic from '@/assets/icons/statistic.svg?raw'
import CommonIcon from '@/components/CommonIcon.vue'

// css variables
const { minorBgColor, hoverBgColor, iconColor } = useCssVariables()

const router = useRouter();

const navigate = (path) => {
    router.push(path);
};
</script>

<style scoped>
.menu-bar {
    display: flex;
    flex-direction: column;
    align-items: center;
    position: relative;
}
</style>