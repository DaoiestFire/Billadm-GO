<template>
  <a-layout style="height: 100vh">
    <a-modal v-model:open="showWorkspaceSelect"
             title="新建工作目录或打开已存在的工作目录"
             ok-text="确认"
             cancel-text="取消"
             @ok="handleOpenWorkspace">
      <a-input-search
          v-model:value="workspaceDir"
          placeholder="选择工作目录"
          enter-button="打开目录"
          @search="handleBrowse"
      />
    </a-modal>
    <a-layout-header class="headerStyle">
      <app-top-bar/>
    </a-layout-header>
    <a-layout style="height: 100%">
      <a-layout-sider :style="siderStyle" :width="siderWidthSize">
        <app-left-bar/>
      </a-layout-sider>
      <a-layout-content :style="contentStyle">Content</a-layout-content>
    </a-layout>
  </a-layout>
</template>

<script setup lang="ts">
import {type CSSProperties, onMounted, ref} from "vue";
import {useCssVariables} from "@/css/css.ts";
import {useLedgerStore} from "@/stores/ledgerStore.ts";
import {useTrViewStore} from "@/stores/trViewStore.ts";
import {useCategoryStore} from "@/stores/categoryStore.ts";
import {useTagStore} from "@/stores/tagStore.ts";
import {openWorkspace} from "@/backend/api/workspace.ts";
import NotificationUtil from "@/backend/notification.ts";


const {majorBgColor, minorBgColor, siderWidthSize} = useCssVariables();

const siderStyle: CSSProperties = {
  backgroundColor: minorBgColor.value,
};

const contentStyle: CSSProperties = {
  backgroundColor: majorBgColor.value,
};

const ledgerStore = useLedgerStore();
const trViewStore = useTrViewStore();
const categoryStore = useCategoryStore();
const tagStore = useTagStore();

const showWorkspaceSelect = ref(true);
const workspaceDir = ref('');
const browseMode = ref('directory');

const handleOpenWorkspace = async () => {
  try {
    await openWorkspace(workspaceDir.value);
    await initWorkspace();
  } catch (error) {
    NotificationUtil.error(`打开工作空间失败 ${error}`);
    showWorkspaceSelect.value = true;
  }
}

async function handleBrowse() {
  try {
    let result
    result = await window.electronAPI.openDialog({
      properties: browseMode.value === 'directory' ? ['openDirectory'] : ['openFile'],
    })

    if (result.canceled || !result.filePaths || result.filePaths.length === 0) {
      return
    }
    workspaceDir.value = result.filePaths[0]
  } catch (err) {
    NotificationUtil.error(`选择目录失败 ${err}`)
  }
}

const initWorkspace = async () => {
  await ledgerStore.refreshWorkspaceStatus();
  if (!ledgerStore.workspaceStatus.isOpened) {
    showWorkspaceSelect.value = true;
    return;
  } else {
    showWorkspaceSelect.value = false;
    window.electronAPI.setWorkspace(ledgerStore.workspaceStatus.workspaceDir);
  }
  await ledgerStore.init();
  await trViewStore.init();
  await categoryStore.refreshCategory();
  await tagStore.refreshTag();
}

onMounted(initWorkspace)
</script>

<style scoped>
.headerStyle {
  height: var(--billadm-size-header-height);
  background-color: var(--billadm-color-minor-background);
  padding: 0;
}
</style>