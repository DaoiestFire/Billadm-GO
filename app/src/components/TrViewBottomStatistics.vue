<template>
  <div class="tr-view-bottom-statistics">
    <div
        class="statistics-item"
        v-for="item in statistics"
        :key="item.label"
    >
      <a-typography-text>{{ item.label }}</a-typography-text>
      <span class="value">{{ formatFloat(item.value) }}</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed} from "vue";
import {useTrViewStore} from "@/stores/trViewStore.ts";
import {formatFloat} from "@/backend/functions.ts";

const trViewStore = useTrViewStore();

const statistics = computed(() => {
  return [
    {
      label: '记录数',
      value: trViewStore.trCount,
    }, {
      label: '收入',
      value: trViewStore.trStatistics.income,
    }, {
      label: '支出',
      value: trViewStore.trStatistics.expense,
    },
    {
      label: '转账',
      value: trViewStore.trStatistics.transfer,
    }
  ];
})
</script>

<style scoped>
.tr-view-bottom-statistics {
  display: flex;
  font-size: var(--billadm-size-text-base-size);
  gap: 8px;
}

.statistics-item {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 4px;
}
</style>