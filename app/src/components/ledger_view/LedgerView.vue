<template>
  <a-layout style="height: 100%">
    <a-layout-header class="headerStyle">
      <div class="left-groups">
      </div>
      <div class="center-groups">
      </div>
      <div class="right-groups">
        <a-button type="primary" @click="createLedger">新增账本</a-button>
      </div>
    </a-layout-header>
    <a-layout-content :style="contentStyle">
      <div class="card-grid">
        <a-card v-for="ledger in ledgerStore.ledgers"
                :key:="ledger.id"
                hoverable
        >
          <a-descriptions :title="ledger.name" layout="vertical">
            <template #extra>
              <a-button type="text" :style="editButtonStyle" @click="modifyLedgerName(ledger.name)">编辑</a-button>
              <a-popconfirm title="确认删除吗"
                            ok-text="确认"
                            :showCancel="false"
                            @confirm="ledgerStore.deleteLedger(ledger.id)"
              >
                <a-button type="text" :style="deleteButtonStyle">删除</a-button>
              </a-popconfirm>
            </template>
            <a-descriptions-item label="创建时间">{{ formatTimestamp(ledger.createdAt) }}</a-descriptions-item>
            <a-descriptions-item label="更新时间">{{ formatTimestamp(ledger.updatedAt) }}</a-descriptions-item>
          </a-descriptions>
        </a-card>
      </div>
    </a-layout-content>
  </a-layout>
  <a-modal
      v-model:open="showModal"
      :title="modalTitle"
      ok-text="确认"
      cancel-text="取消"
      style="top: 250px"
      @ok="handleOk"
  >
    <a-input v-model:value.lazy="ledgerName" placeholder="输入账本名称"/>
  </a-modal>
</template>

<script setup lang="ts">
import type {CSSProperties} from "vue";
import {ref} from 'vue';
import {useLedgerStore} from "@/stores/ledgerStore.ts";
import dayjs from "dayjs";
import {useCssVariables} from "@/backend/css.ts";

const {majorBgColor, positiveColor, negativeColor} = useCssVariables();

const editButtonStyle: CSSProperties = {
  color: positiveColor.value,
};

const deleteButtonStyle: CSSProperties = {
  color: negativeColor.value,
};

const contentStyle: CSSProperties = {
  backgroundColor: majorBgColor.value,
  "overflow-y": "auto",
  "margin-bottom": "auto"
};

const ledgerStore = useLedgerStore();

const formatTimestamp = (ts: number) => {
  return dayjs(ts * 1000).format('YYYY-MM-DD HH:mm:ss');
}

const showModal = ref<boolean>(false);
const modalTitle = ref<string>("");
const ledgerName = ref<string>("");

const createLedger = () => {
  modalTitle.value = "创建账本";
  ledgerName.value = "";
  showModal.value = true;
};

const modifyLedgerName = (name: string) => {
  modalTitle.value = "修改账本名称";
  ledgerName.value = name;
  showModal.value = true;
};

const handleOk = async () => {
  if (!ledgerName) return;
  if (modalTitle.value === "创建账本") {
    await ledgerStore.createLedger(ledgerName.value);
  } else if (modalTitle.value === "修改账本名称") {

  }
  showModal.value = false;
};
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

.card-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(400px, 1fr));
  gap: 16px;
  padding: 16px;
}
</style>