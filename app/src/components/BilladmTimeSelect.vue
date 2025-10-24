<template>
  <div class="billadm-time-select">
    <!-- 时间范围类型选择 -->
    <div class="group">
      <billadm-label>时间粒度</billadm-label>
      <billadm-select
          v-model="timeRangeType"
          :options="TimeRangeTypes"
          :height="24"
          :width="60"
          direction="down"
      />
    </div>

    <div class="group">
      <billadm-label>时间范围</billadm-label>
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
          style="width: 200px"
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
  </div>
</template>

<script setup lang="ts">
import {computed} from 'vue'
import BilladmSelect from '@/components/BilladmSelect.vue'
import BilladmIconButton from '@/components/BilladmIconButton.vue'
import iconLeft from '@/assets/icons/left.svg?raw'
import iconRight from '@/assets/icons/right.svg?raw'
import {getNextPeriod, getPrevPeriod, normalizeTimeRange} from '@/backend/timerange.ts'
import {useCssVariables} from '@/css/css.ts'
import {TimeRangeShortcuts, TimeRangeTypes} from "@/backend/constant.ts"
import BilladmLabel from "@/components/text/BilladmLabel.vue"
import type {TimeRangeType} from "@/types/billadm";

// 👇 定义 Props 类型
interface Props {
  timeRange: Date[]
  timeRangeType: TimeRangeType
}

// 使用类型声明 defineProps
const props = withDefaults(defineProps<Props>(), {})

// 定义事件
const emit = defineEmits(['update:timeRange', 'update:timeRangeType', 'change'])

// 样式变量
const {hoverBgColor, iconColor} = useCssVariables()

// 双向绑定时间范围和类型
const timeRange = computed({
  get: () => props.timeRange,
  set: (val) => {
    val = normalizeTimeRange(val[0], val[1], props.timeRangeType)
    emit('update:timeRange', val)
    emit('change', val, props.timeRangeType)
  }
})

const timeRangeType = computed({
  get: () => props.timeRangeType,
  set: (val) => {
    let timeRange = normalizeTimeRange(props.timeRange[0], props.timeRange[1], val)
    emit('update:timeRange', timeRange)
    emit('update:timeRangeType', val)
    emit('change', props.timeRange, val)
  }
})

// 按钮点击事件
const goToPrevious = () => {
  if (!Array.isArray(props.timeRange)) return
  const [start, end] = props.timeRange
  const newRange = getPrevPeriod(start, end, props.timeRangeType)
  emit('update:timeRange', newRange)
  emit('change', newRange, props.timeRangeType)
}

const goToNext = () => {
  if (!Array.isArray(props.timeRange)) return
  const [start, end] = props.timeRange
  const newRange = getNextPeriod(start, end, props.timeRangeType)
  emit('update:timeRange', newRange)
  emit('change', newRange, props.timeRangeType)
}
</script>

<style scoped>
.billadm-time-select {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 16px;
}

.group {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
}
</style>