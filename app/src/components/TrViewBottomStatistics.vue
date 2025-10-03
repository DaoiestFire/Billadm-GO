<template>
  <div class="tr-view-bottom-statistics">
    <div
        class="statistics-item"
        v-for="item in statistics"
        :key="item.label"
    >
      <span class="label">{{ item.label }}</span>
      <span class="value">{{ item.value }}</span>
    </div>
  </div>
</template>

<script setup>
import {ref, watch,onMounted} from "vue";
import {useTrViewStore} from "@/stores/trViewStore.js";

const statistics = ref([]);

const trViewStore = useTrViewStore()

function aggregateByTransactionType(records) {
  // 初始化聚合结果对象
  const result = {
    expense: 0,
    income: 0,
    transfer: 0
  }

  // 检查输入是否为有效数组且非空
  if (!Array.isArray(records) || records.length === 0) {
    return result
  }

  // 遍历数组中的每个元素
  for (const record of records) {
    // 确保元素是对象且包含必要的字段
    if (record && typeof record === 'object' && 'transaction_type' in record && 'price' in record) {
      const {transaction_type, price} = record
      if (transaction_type === 'expense' ||
          transaction_type === 'income' ||
          transaction_type === 'transfer') {
        result[transaction_type] += Number(price) || 0
      }
    }
  }

  return result
}

const refreshStatistics = () => {
  statistics.value = [];
  const {expense, income, transfer} = aggregateByTransactionType(trViewStore.tableData);
  statistics.value.push({
    label: '记录数',
    value: trViewStore.trCount,
  });
  statistics.value.push({
    label: '收入',
    value: income,
  })
  ;
  statistics.value.push({
    label: '支出',
    value: expense,
  });
  statistics.value.push({
    label: '转账',
    value: transfer,
  });
}

watch(() => [trViewStore.tableData, trViewStore.trCount], () => {
  refreshStatistics();
})

onMounted(() => {
  refreshStatistics();
})
</script>

<style scoped>
.tr-view-bottom-statistics {
  display: flex;
  font-size: 14px;
  gap: 8px;
}

.statistics-item {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 4px;
}

.label {
  color: var(--billadm-color-text-minor-color);
}
</style>