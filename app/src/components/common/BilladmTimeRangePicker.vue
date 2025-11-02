<!-- @/components/BilladmTimeRangePicker.vue -->
<template>
  <div class="billadm-time-range-picker">
    <a-segmented
        v-model:value="timeRangeTypeLabel"
        :options="TimeRangeTypeLabels"
        @change="handleSegmentChange"
    />
    <a-button type="text" @click="goToPrevious">
      <template #icon>
        <LeftOutlined style="display: flex; justify-content: center; align-items: center; font-size: large"/>
      </template>
    </a-button>
    <a-range-picker
        v-model:value="timeRange"
        :picker="timeRangeTypeValue"
        :presets="TimeRangePresets"
        inputReadOnly
    />
    <a-button type="text" @click="goToNext">
      <template #icon>
        <RightOutlined style="display: flex; justify-content: center; align-items: center; font-size: large"/>
      </template>
    </a-button>
  </div>
</template>

<script setup lang="ts">
import {computed} from 'vue';
import {TimeRangePresets, TimeRangeTypeLabels, TimeRangeTypes} from '@/backend/constant.ts';
import {getNextPeriod, getPrevPeriod} from '@/backend/timerange.ts';
import {LeftOutlined, RightOutlined} from '@ant-design/icons-vue';
import dayjs, {Dayjs} from 'dayjs';
import type {TimeRangeTypeLabel} from '@/types/billadm';

const timeRange = defineModel<[Dayjs, Dayjs]>({required: true});
const timeRangeTypeLabel = defineModel<TimeRangeTypeLabel>({required: true});

// 根据当前 label 获取对应的 picker 类型 ('date' | 'month' | 'year')
const timeRangeTypeValue = computed(() => {
  return TimeRangeTypes[timeRangeTypeLabel.value];
});

// 上一周期
const goToPrevious = () => {
  timeRange.value = getPrevPeriod(timeRange.value[0], timeRange.value[1], timeRangeTypeValue.value);
};

// 下一周期
const goToNext = () => {
  timeRange.value = getNextPeriod(timeRange.value[0], timeRange.value[1], timeRangeTypeValue.value);
};

// 切换时间类型时，重置为当前周期
const handleSegmentChange = () => {
  const now = dayjs();
  let start: Dayjs, end: Dayjs;

  switch (timeRangeTypeLabel.value) {
    case '日':
      start = now.startOf('day');
      end = now.endOf('day');
      break;
    case '月':
      start = now.startOf('month');
      end = now.endOf('month');
      break;
    case '年':
      start = now.startOf('year');
      end = now.endOf('year');
      break;
    default:
      start = now.startOf('month');
      end = now.endOf('month');
  }

  timeRange.value = [start, end];
};
</script>

<style scoped>
.billadm-time-range-picker {
  display: flex;
  align-items: center;
  gap: 8px;
}
</style>