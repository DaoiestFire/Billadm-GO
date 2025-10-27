<template>
  <a-layout style="height: 100vh">
    <billadm-file-select v-model="showWorkspaceSelect"
                         title="新建工作目录或打开已存在的工作目录"
                         @confirm="handleOpenWorkspace"
    />
    <a-layout-header class="headerStyle">
      <app-top-bar/>
    </a-layout-header>
    <a-layout style="height: 100%">
      <a-layout-sider :style="siderStyle" :width="siderWidthSize">
        <app-left-bar/>
      </a-layout-sider>
      <a-layout-content :style="contentStyle">
        <a-card style="height: 100%;padding: 16px" :body-style="{padding:'0px',height:'100%'}" :bordered="false">
          <router-view/>
        </a-card>
      </a-layout-content>
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


const {minorBgColor, siderWidthSize} = useCssVariables();

const siderStyle: CSSProperties = {
  backgroundColor: minorBgColor.value,
};

const contentStyle: CSSProperties = {
  backgroundColor: minorBgColor.value,
};

const ledgerStore = useLedgerStore();
const trViewStore = useTrViewStore();
const categoryStore = useCategoryStore();
const tagStore = useTagStore();

const showWorkspaceSelect = ref(true);

const handleOpenWorkspace = async (workspaceDir: string) => {
  try {
    await openWorkspace(workspaceDir);
    await initWorkspace();
  } catch (error) {
    NotificationUtil.error(`打开工作空间失败 ${error}`);
    showWorkspaceSelect.value = true;
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