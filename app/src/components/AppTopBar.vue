<template>
  <div class="menu-bar">
    <div class="left-groups">
      <CommonIcon :icon="iconMenu" label="菜单" width="41" height="30" :color="iconColor" :bgColor="minorBgColor"
                  :hoverBgColor="hoverBgColor" hoverStyle="rect" @click="billadmMenu.toggleMenu()"/>
      <MultiLevelMenu ref="billadmMenu"/>
      <CustomSelect v-model="currentLedgerId" :options="ledgers" height="24px" width="100px"
                    placeholder="选择账本"/>
    </div>
    <div class="center-groups">
      Billadm-{{ route.name }}
    </div>
    <div class="right-groups">
      <CommonIcon :icon="iconZoomOut" label="最小化" width="41" height="30" :color="iconColor"
                  :bgColor="minorBgColor" :hoverBgColor="hoverBgColor" hoverStyle="rect"
                  tooltipPlacement="bottom-left" @click="onMinimize"/>
      <CommonIcon :icon="iconZoomIn" label="最大化" width="41" height="30" :color="iconColor" :bgColor="minorBgColor"
                  :hoverBgColor="hoverBgColor" hoverStyle="rect" tooltipPlacement="bottom-left" @click="onMaximize"/>
      <CommonIcon :icon="iconClose" label="关闭" width="41" height="30" :color="iconColor" :bgColor="minorBgColor"
                  :hoverBgColor="hoverBgColor" hoverStyle="rect" tooltipPlacement="bottom-left" @click="onClose"/>
    </div>
  </div>
</template>

<script setup>
import {computed, ref} from 'vue'
import {useRoute} from 'vue-router'
import {useCssVariables} from '@/css/css'
import iconMenu from '@/assets/icons/menu.svg?raw'
import iconZoomOut from '@/assets/icons/zoom-out.svg?raw'
import iconZoomIn from '@/assets/icons/zoom-in.svg?raw'
import iconClose from '@/assets/icons/close.svg?raw'
import CommonIcon from '@/components/CommonIcon.vue'
import CustomSelect from '@/components/CustomSelect.vue'
import MultiLevelMenu from '@/components/MultiLevelMenu.vue'
import {useLedgerStore} from "@/stores/ledgerStore.js";
import {exitApp} from "@/backend/app.js";

// store
const ledgerStore = useLedgerStore()

// 当前视图
const route = useRoute()

// css variables
const {minorBgColor, hoverBgColor, iconColor} = useCssVariables()

// 菜单
const billadmMenu = ref(null)

// 创建可写的 computed 用于 v-model
const currentLedgerId = computed({
  get() {
    return ledgerStore.currentLedgerId
  },
  set(value) {
    ledgerStore.setCurrentLedger(value)
  }
})

const ledgers = computed(() => {
  return ledgerStore.ledgers.map(l => ({
    label: l.name,
    value: l.id,
  }))
})

// 窗口控制
const onMinimize = () => {
  window.electronAPI.minimizeWindow();
}

const onMaximize = () => {
  window.electronAPI.maximizeWindow();
}

const onClose = async () => {
  exitApp()
  window.electronAPI.closeWindow();
}

</script>

<style scoped>
.menu-bar {
  -webkit-app-region: drag;
  display: flex;
  align-items: center;
  position: relative;
}

.menu-bar > * {
  -webkit-app-region: no-drag;
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