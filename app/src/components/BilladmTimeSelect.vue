<template>
  <div class="billadm-time-select">
    <!-- 时间范围类型选择 -->
    <billadm-select
        v-model="timeRangeType"
        :options="TimeRangeTypes"
        :height="24"
        :width="30"
        direction="down"
    />

    <!-- 向前按钮 -->
    <billadm-icon-button
        :svg="iconLeft"
        label="向前一天"
        width="30px"
        height="30px"
        bg-size="26px"
        :color="iconColor"
        bg-color="transparent"
        :hover-bg-color="hoverBgColor"
        tooltipPlacement="bottom"
        @click="goToPrevious"
    />

    <!-- 日期选择器 -->
    <el-date-picker
        v-model="timeRange"
        :type="timeRangeType"
        range-separator="~"
        start-placeholder="起始时间"
        end-placeholder="结束时间"
        size="small"
        :editable="false"
        style="width: 200px;"
        :shortcuts="TimeRangeShortcuts"
    />

    <!-- 向后按钮 -->
    <billadm-icon-button
        :svg="iconRight"
        label="向后一天"
        width="30px"
        height="30px"
        bg-size="26px"
        :color="iconColor"
        bg-color="transparent"
        :hover-bg-color="hoverBgColor"
        tooltipPlacement="bottom"
        @click="goToNext"
    />
  </div>
</template>

<script setup>
import {computed, defineEmits, defineProps} from 'vue';
import BilladmSelect from '@/components/BilladmSelect.vue';
import BilladmIconButton from '@/components/BilladmIconButton.vue';
import iconLeft from '@/assets/icons/left.svg?raw';
import iconRight from '@/assets/icons/right.svg?raw';
import {getNextPeriod, getPrevPeriod} from '@/backend/timerange.js';
import {useCssVariables} from '@/css/css.js';
import {TimeRangeShortcuts, TimeRangeTypes} from "@/backend/constant.js";

// 接收父组件传入的 v-model 绑定的时间范围和类型
const props = defineProps({
  timeRange: {
    type: [Array, String, Date],
    required: true
  },
  timeRangeType: {
    type: String,
    required: true
  }
});

// 定义事件
const emit = defineEmits(['update:timeRange', 'update:timeRangeType', 'change']);

// 样式变量
const {hoverBgColor, iconColor} = useCssVariables();

// 双向绑定时间范围和类型
const timeRange = computed({
  get: () => props.timeRange,
  set: (val) => {
    emit('update:timeRange', val);
    emit('change', val, props.timeRangeType);
  }
});

const timeRangeType = computed({
  get: () => props.timeRangeType,
  set: (val) => {
    emit('update:timeRangeType', val);
    emit('change', props.timeRange, val);
  }
});

// 按钮点击事件
const goToPrevious = () => {
  if (!Array.isArray(props.timeRange)) return;
  const [start, end] = props.timeRange;
  const newRange = getPrevPeriod(start, end, props.timeRangeType);
  emit('update:timeRange', newRange);
  emit('change', newRange, props.timeRangeType);
};

const goToNext = () => {
  if (!Array.isArray(props.timeRange)) return;
  const [start, end] = props.timeRange;
  const newRange = getNextPeriod(start, end, props.timeRangeType);
  emit('update:timeRange', newRange);
  emit('change', newRange, props.timeRangeType);
};
</script>

<style scoped>
.billadm-time-select {
  display: flex;
  align-items: center;
  gap: 4px;
}
</style>