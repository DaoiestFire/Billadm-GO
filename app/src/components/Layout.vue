<template>
  <div class="layout">
    <!-- 上栏：菜单栏 -->
    <div class="top-bar">
      <AppTopBar/>
    </div>

    <!-- 中间栏：左中右 -->
    <div class="middle-section">
      <!-- 左侧功能栏 -->
      <div class="left-panel">
        <AppLeftBar/>
      </div>

      <!-- 中间阅览区 -->
      <div class="center-panel">
        <router-view/>
      </div>

      <!-- 右侧功能栏 -->
      <div class="right-panel">
      </div>
    </div>

    <!-- 下栏：详情栏 -->
    <div class="bottom-bar">
    </div>
  </div>
</template>

<script setup>
import {onMounted} from "vue";
import AppTopBar from '@/components/AppTopBar.vue'
import AppLeftBar from '@/components/AppLeftBar.vue'
import {useLedgerStore} from "@/stores/ledgerStore.js";

const ledgerStore = useLedgerStore()

onMounted(async () => {
  await ledgerStore.updateLedgers()
})
</script>

<style scoped>
.layout {
  display: flex;
  flex-direction: column;
  height: 100vh;
}

.top-bar {
  background: var(--billadm-color-minor-backgroud-color);
  height: 30px;
  border-bottom: 1px solid var(--billadm-color-window-border-color);
}

.middle-section {
  flex: 1;
  display: flex;
  overflow: hidden;
}

.left-panel,
.right-panel {
  width: 40px;
  background: var(--billadm-color-minor-backgroud-color);
}

.left-panel {
  border-right: 1px solid var(--billadm-color-window-border-color);
}

.right-panel {
  border-left: 1px solid var(--billadm-color-window-border-color);
}

.center-panel {
  flex: 1;
}

.bottom-bar {
  background: var(--billadm-color-minor-backgroud-color);
  height: 30px;
  border-top: 1px solid var(--billadm-color-window-border-color);
}
</style>