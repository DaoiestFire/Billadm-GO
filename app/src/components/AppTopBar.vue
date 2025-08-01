<template>
    <div class="menu-bar">
        <div class="left-groups">
            <CommonIcon :icon="iconMenu" label="菜单" width="41" height="30" :color="iconColor" :bgColor="minorBgColor"
                :hoverBgColor="hoverBgColor" hoverStyle="rect" @click="billadmMenu.toggleMenu()" />
            <MultiLevelMenu ref="billadmMenu" />
            <CustomSelect v-model="currentLedger" :options="ledgers" height="24px" width="120px" placeholder="选择账本" />
        </div>
        <div class="center-groups">
            Billadm-{{ route.name }}
        </div>
        <div class="right-groups">
            <CommonIcon :icon="iconZoomOut" label="最小化" width="41" height="30" :color="iconColor"
                :bgColor="minorBgColor" :hoverBgColor="hoverBgColor" hoverStyle="rect" tooltipPlacement="bottom-left" />
            <CommonIcon :icon="iconZoomIn" label="最大化" width="41" height="30" :color="iconColor" :bgColor="minorBgColor"
                :hoverBgColor="hoverBgColor" hoverStyle="rect" tooltipPlacement="bottom-left" />
            <CommonIcon :icon="iconClose" label="关闭" width="41" height="30" :color="iconColor" :bgColor="minorBgColor"
                :hoverBgColor="hoverBgColor" hoverStyle="rect" tooltipPlacement="bottom-left" />
        </div>
    </div>
</template>

<script setup>
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import { useCssVariables } from '@/css/css'
import iconMenu from '@/assets/icons/menu.svg?raw'
import iconZoomOut from '@/assets/icons/zoom-out.svg?raw'
import iconZoomIn from '@/assets/icons/zoom-in.svg?raw'
import iconClose from '@/assets/icons/close.svg?raw'
import CommonIcon from '@/components/CommonIcon.vue'
import CustomSelect from '@/components/CustomSelect.vue'
import MultiLevelMenu from '@/components/MultiLevelMenu.vue'

// 当前视图
const route = useRoute()

// css variables
const { minorBgColor, hoverBgColor, iconColor } = useCssVariables()

// 菜单
const billadmMenu = ref(null)

// 账本选择框
const currentLedger = ref('')

watch(() => currentLedger.value,
    (newValue) => {
        console.log(newValue)
    },
    { immediate: true })


const ledgers = [
    { label: '刘敬威的账本', value: '刘敬威的账本' },
    { label: '证券账本', value: '证券账本' },
]


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