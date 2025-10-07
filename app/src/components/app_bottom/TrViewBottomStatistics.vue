<template>
  <div class="tr-view-bottom-statistics">
    <div
        class="statistics-item"
        v-for="item in statistics"
        :key="item.label"
    >
      <billadm-label>{{ item.label }}</billadm-label>
      <span class="value">{{ formatFloat(item.value) }}</span>
    </div>
  </div>
</template>

<script setup>
import {computed} from "vue";
import BilladmLabel from "@/components/text/BilladmLabel.vue";
import {useTrViewStore} from "@/stores/trViewStore.js";
import {formatFloat} from "@/backend/functions.js";

const trViewStore = useTrViewStore();

const statistics = computed(() => {
  return [
    {
      label: '记录数',
      value: trViewStore.trCount,
    }, {
      label: '收入',
      value: trViewStore.trStatistics?.income ?? 0,
    }, {
      label: '支出',
      value: trViewStore.trStatistics?.expense ?? 0,
    },
    {
      label: '转账',
      value: trViewStore.trStatistics?.transfer ?? 0,
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