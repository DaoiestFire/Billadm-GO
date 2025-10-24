<template>
  <div class="layout">
    <!-- 上栏：工具栏 -->
    <div class="top-bar">
      <div class="left-groups">
        <billadm-time-select
            v-model:time-range="timeRange"
            v-model:time-range-type="timeRangeType"
        />
      </div>

      <div class="center-groups">
      </div>

      <div class="right-groups">
      </div>
    </div>

    <div class="middle-section">
      <billadm-chart-display :tr-form-list="trs"/>
    </div>
  </div>
</template>

<script setup lang="ts">
import {onMounted, ref, watch} from 'vue'
import BilladmChartDisplay from "@/components/da_view/BilladmChartDisplay.vue"
import BilladmTimeSelect from "@/components/BilladmTimeSelect.vue"
import {useLedgerStore} from "@/stores/ledgerStore.ts"
import {useTrViewStore} from "@/stores/trViewStore.ts"
import {dateToUnixTimestamp} from "@/backend/functions.ts"
import {queryTrsOnCondition} from "@/backend/api/tr.ts"
import NotificationUtil from "@/backend/notification.ts"
import {trDtoToTrForm} from "@/backend/dto-utils.ts"
import type {TrForm, TrQueryCondition} from "@/types/billadm"

// store
const ledgerStore = useLedgerStore()
const trViewStore = useTrViewStore()

// 组件响应式变量
const timeRangeType = ref(trViewStore.timeRangeType)
const timeRange = ref([new Date(trViewStore.timeRange[0]), new Date(trViewStore.timeRange[1])])
const trs = ref([] as TrForm[])

const refreshTrs = async () => {
  try {
    let condition = {} as TrQueryCondition
    condition.ledgerId = ledgerStore.currentLedgerId
    condition.tsRange = [dateToUnixTimestamp(timeRange.value[0]), dateToUnixTimestamp(timeRange.value[1])]
    let trDtos = await queryTrsOnCondition(condition)
    trs.value = trDtos.map(item => {
      return trDtoToTrForm(item)
    });
  } catch (error) {
    NotificationUtil.error(`数据分析请求 ${error}`)
  }
}

watch(() => [timeRange.value, ledgerStore.currentLedger], refreshTrs)

onMounted(refreshTrs);
</script>

<style scoped>
.layout {
  display: flex;
  flex-direction: column;
  height: 100%;
  background-color: var(--billadm-color-major-background-color);
}

.top-bar {
  background: var(--billadm-color-major-background-color);
  height: 30px;
  flex-shrink: 0;
  display: flex;
  align-items: center;
  position: relative;
  margin-bottom: 10px;
}

.left-groups {
  margin-right: auto;
  display: flex;
  gap: 4px;
  align-items: center;
}

.center-groups {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  gap: 4px;
}

.right-groups {
  display: flex;
  gap: 4px;
}

.middle-section {
  flex: 1;
  overflow: hidden;
}
</style>