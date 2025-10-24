<template>
  <div class="layout">
    <FileDirSelect v-model:visible="showWorkspaceSelect"
                   title="新建工作目录或打开已存在的工作目录"
                   :cancel-color="negativeColor"
                   :confirm-color="positiveColor"
                   @confirm="handleOpenWorkspace"
    />
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
      <AppBottomBar/>
    </div>
  </div>
</template>

<script setup lang="ts">
import {onMounted, ref} from "vue"
import AppTopBar from '@/components/app_top/AppTopBar.vue'
import AppLeftBar from '@/components/AppLeftBar.vue'
import AppBottomBar from "@/components/app_bottom/AppBottomBar.vue"
import FileDirSelect from "@/components/FileDirSelect.vue"
import {useLedgerStore} from "@/stores/ledgerStore.ts"
import {useTrViewStore} from "@/stores/trViewStore.ts"
import {useCategoryStore} from "@/stores/categoryStore.ts"
import {useTagStore} from "@/stores/tagStore.ts"
import {useCssVariables} from "@/css/css.ts"
import {openWorkspace} from "@/backend/api/workspace.ts"
import NotificationUtil from "@/backend/notification.ts"


const ledgerStore = useLedgerStore()
const trViewStore = useTrViewStore()
const categoryStore = useCategoryStore()
const tagStore = useTagStore()

// 引用颜色
const {positiveColor, negativeColor} = useCssVariables()

const showWorkspaceSelect = ref(false)

const handleOpenWorkspace = async (workspace: string) => {
  try {
    await openWorkspace(workspace)
    await initWorkspace()
  } catch (error) {
    NotificationUtil.error(`打开工作空间失败 ${error}`)
    showWorkspaceSelect.value = true
  }
}

const initWorkspace = async () => {
  await ledgerStore.refreshWorkspaceStatus()
  if (!ledgerStore.workspaceStatus.isOpened) {
    showWorkspaceSelect.value = true
    return
  } else {
    showWorkspaceSelect.value = false
    window.electronAPI.setWorkspace(ledgerStore.workspaceStatus.workspaceDir)
  }
  await ledgerStore.init()
  await trViewStore.init()
  await categoryStore.refreshCategory()
  await tagStore.refreshTag()
}

onMounted(initWorkspace)
</script>

<style scoped>
.layout {
  display: flex;
  flex-direction: column;
  height: 100vh;
  background-color: var(--billadm-color-minor-background-color);
}

.top-bar {
  background: var(--billadm-color-minor-background-color);
  height: var(--billadm-ui-size-menu-width);
}

.middle-section {
  flex: 1;
  display: flex;
  overflow: hidden;
}

.left-panel,
.right-panel {
  width: var(--billadm-ui-size-menu-width);
  background: var(--billadm-color-minor-background-color);
}

.center-panel {
  border-radius: 8px;
  border: 1px solid var(--billadm-color-window-border-color);
  padding: 8px;
  background-color: var(--billadm-color-major-background-color);
  flex: 1;
}

.bottom-bar {
  background: var(--billadm-color-minor-backgroud-color);
  height: var(--billadm-ui-size-menu-width);
}
</style>