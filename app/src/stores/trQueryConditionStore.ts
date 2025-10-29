import {computed, ref} from "vue"
import {defineStore} from 'pinia'
import {getThisMonthRange} from "@/backend/timerange.ts"
import type {RangeValue, TimeRangeTypeLabel, TimeRangeTypeValue} from "@/types/billadm"
import {TimeRangeTypes} from "@/backend/constant.ts";

export const useTrQueryConditionStore = defineStore('trQueryCondition', () => {

    const timeRange = ref<RangeValue>(getThisMonthRange()); // 时间范围
    const timeRangeTypeLabel = ref('日' as TimeRangeTypeLabel); // 时间类型标签


    const timeRangeTypeValue = computed<TimeRangeTypeValue>(() => {
        return TimeRangeTypes[timeRangeTypeLabel.value];
    })

    return {
        timeRange,
        timeRangeTypeLabel,
        timeRangeTypeValue,
    }
})