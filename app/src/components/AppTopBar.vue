<template>
    <div class="menu-bar">
        <div class="left-groups">
            <CommonIcon :icon="homeIcon" width="41" height="30" :color="iconColor" :bgColor="minorBgColor"
                :hoverBgColor="hoverBgColor" hoverStyle="rect" />
        </div>
        <div class="center-groups">
            <CommonIcon :icon="homeIcon" width="40" height="30" :color="iconColor" :bgColor="minorBgColor"
                :hoverBgColor="hoverBgColor" hoverStyle="rect" />
        </div>
        <div class="right-groups">
            <CommonIcon :icon="homeIcon" width="41" height="30" :color="iconColor" :bgColor="minorBgColor"
                :hoverBgColor="hoverBgColor" hoverStyle="rect" />
            <CommonIcon :icon="homeIcon" width="41" height="30" :color="iconColor" :bgColor="minorBgColor"
                :hoverBgColor="hoverBgColor" hoverStyle="rect" />
            <CommonIcon :icon="homeIcon" width="41" height="30" :color="iconColor" :bgColor="minorBgColor"
                :hoverBgColor="hoverBgColor" hoverStyle="rect" />
        </div>
    </div>
</template>

<script setup>
import { useCssVariables } from '@/css/css'
import CommonIcon from '@/components/CommonIcon.vue'

// css variables
const { minorBgColor, hoverBgColor, iconColor } = useCssVariables()

const homeIcon = `
  <svg viewBox="0 0 24 24" width="24" height="24">
    <path d="M10 20v-6h4v6h5v-8h3L12 3 2 12h3v8z" fill="currentColor"/>
  </svg>
`

// props
const props = defineProps({
    title: {
        type: String,
        default: '默认标题'
    }
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