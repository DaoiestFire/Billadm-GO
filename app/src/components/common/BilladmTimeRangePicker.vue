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
import {
  TimeRangeLabelToValue,
  TimeRangePresets,
  TimeRangeTypeLabels,
  TimeRangeValueToLabel
} from '@/backend/constant.ts';
import {getNextPeriod, getPrevPeriod, setToEndOfDay, setToStartOfDay} from '@/backend/timerange.ts';
import {LeftOutlined, RightOutlined} from '@ant-design/icons-vue';
import {Dayjs} from 'dayjs';
import type {RangeValue, TimeRangeTypeValue} from '@/types/billadm';
import type {SegmentedValue} from "ant-design-vue/es/segmented/src/segmented";

const timeRange = defineModel<RangeValue>('timeRange', {required: true});
const timeRangeTypeValue = defineModel<TimeRangeTypeValue>('timeRangeType', {required: true});

const timeRangeTypeLabel = computed({
  get() {
    return TimeRangeValueToLabel[timeRangeTypeValue.value];
  },
  set(val) {
    timeRangeTypeValue.value = TimeRangeLabelToValue[val];
  }
});

// 上一周期
const goToPrevious = () => {
  timeRange.value = getPrevPeriod(timeRange.value[0], timeRange.value[1], timeRangeTypeValue.value);
};

// 下一周期
const goToNext = () => {
  timeRange.value = getNextPeriod(timeRange.value[0], timeRange.value[1], timeRangeTypeValue.value);
};

// 切换时间类型时 修改时间范围
const handleSegmentChange = (val: SegmentedValue) => {
  let start: Dayjs = timeRange.value[0], end: Dayjs = timeRange.value[1];
  switch (val) {
    case '日':
      start = start.startOf('day');
      end = end.endOf('day');
      break;
    case '月':
      start = start.startOf('month');
      end = end.endOf('month');
      break;
    case '年':
      start = start.startOf('year');
      end = end.endOf('year');
      break;
  }
  timeRange.value = [setToStartOfDay(start), setToEndOfDay(end)];
};
</script>

<style scoped>
.billadm-time-range-picker {
  display: flex;
  align-items: center;
  gap: 8px;
}
</style>