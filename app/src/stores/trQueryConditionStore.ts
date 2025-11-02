import {ref} from "vue"
import {defineStore} from 'pinia'
import {getThisMonthRange} from "@/backend/timerange.ts"
import type {RangeValue, TimeRangeTypeValue} from "@/types/billadm"

export const useTrQueryConditionStore = defineStore('trQueryCondition', () => {

    const timeRange = ref<RangeValue>(getThisMonthRange()); // 时间范围
    const timeRangeType = ref('date' as TimeRangeTypeValue); // 时间类型标签

    return {
        timeRange,
        timeRangeType,
    }
})