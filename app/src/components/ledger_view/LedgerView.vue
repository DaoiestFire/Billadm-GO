<template>
  <a-layout style="height: 100%">
    <a-layout-header class="headerStyle">
      <div class="left-groups">
      </div>
      <div class="center-groups">
      </div>
      <div class="right-groups">
        <a-button type="primary">
          新增账本
        </a-button>
      </div>
    </a-layout-header>
    <a-layout-content :style="contentStyle">
      <a-space direction="vertical" style="padding: 16px">
        <a-card v-for="ledger in ledgerStore.ledgers"
                :key:="ledger.id"
                title="账本信息"
                hoverable
        >
          <template #extra><a href="#">more</a></template>
          <a-descriptions :title="ledger.name">
            <a-descriptions-item label="账本ID">{{ ledger.id }}</a-descriptions-item>
            <a-descriptions-item label="创建时间">{{ formatTimestamp(ledger.createdAt) }}</a-descriptions-item>
            <a-descriptions-item label="更新时间">{{ formatTimestamp(ledger.updatedAt) }}</a-descriptions-item>
          </a-descriptions>
        </a-card>
      </a-space>
    </a-layout-content>
  </a-layout>
</template>

<script setup lang="ts">
import {useLedgerStore} from "@/stores/ledgerStore.ts";
import dayjs from "dayjs";
import type {CSSProperties} from "vue";
import {useCssVariables} from "@/backend/css.ts";

const {majorBgColor} = useCssVariables();

const contentStyle: CSSProperties = {
  backgroundColor: majorBgColor.value,
  "overflow-y": "auto",
  "margin-bottom": "auto"
};

const ledgerStore = useLedgerStore();


const formatTimestamp = (ts: number) => {
  return dayjs(ts * 1000).format('YYYY-MM-DD HH:mm:ss');
}


</script>

<style scoped>
.headerStyle {
  height: auto;
  background-color: var(--billadm-color-major-background);
  padding: 0 0 16px 0;
  display: flex;
  align-items: start;
  justify-content: center;
}

/* 左边按钮 将它与后面的元素隔开 */
.left-groups {
  margin-right: auto;
  display: flex;
  gap: 8px;
  align-items: center;
}

/* 中间按钮 */
.center-groups {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  gap: 8px;
}

/* 右边按钮组 */
.right-groups {
  display: flex;
  gap: 8px;
}
</style>