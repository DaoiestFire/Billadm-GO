<template>
  <div class="layout">
    <!-- 上栏：工具栏 -->
    <div class="top-bar">
      <div class="left-groups">
        <billadm-icon-button :svg="iconLeft" label="向前一天" width="30px" height="30px" bg-size="26px"
                             :color="iconColor" bg-color="transparent" :hover-bg-color="hoverBgColor"
                             tooltipPlacement="bottom" @click="goToPreviousDay"/>
        <el-date-picker
            v-model="trViewStore.timeRange"
            type="daterange"
            range-separator="~"
            start-placeholder="起始时间"
            end-placeholder="结束时间"
            size="small"
            :editable="false"
            style="width: 200px;"
            :shortcuts="getShortcuts()"
        />
        <billadm-icon-button :svg="iconRight" label="向后一天" width="30px" height="30px" bg-size="26px"
                             :color="iconColor" bg-color="transparent" :hover-bg-color="hoverBgColor"
                             tooltipPlacement="bottom" @click="goToNextDay"/>
      </div>
      <div class="center-groups">
      </div>
      <div class="right-groups">
      </div>
    </div>

    <div class="middle-section">
      <billadm-chart-display :charts="chartData"/>
    </div>
  </div>
</template>

<script setup>
import {ref} from 'vue';
import BilladmIconButton from "@/components/BilladmIconButton.vue";
import BilladmChartDisplay from "@/components/da_view/BilladmChartDisplay.vue";
import iconLeft from '@/assets/icons/left.svg?raw';
import iconRight from '@/assets/icons/right.svg?raw';
import {useTrViewStore} from "@/stores/trViewStore.js";
import {getNextPeriod, getPrevPeriod, getShortcuts,} from "@/backend/timerange.js";
import {useCssVariables} from "@/css/css.js";

// css variables
const {hoverBgColor, iconColor} = useCssVariables();

// store
const trViewStore = useTrViewStore();

const goToPreviousDay = () => {
  let range = trViewStore.timeRange;
  if (!Array.isArray(range)) {
    return;
  }
  trViewStore.timeRange = getPrevPeriod(range[0], range[1]);
}

const goToNextDay = () => {
  let range = trViewStore.timeRange;
  if (!Array.isArray(range)) {
    return;
  }
  trViewStore.timeRange = getNextPeriod(range[0], range[1]);
}

// 图表
const chartData = ref([
  {
    title: '销售额趋势',
    option: {
      tooltip: {trigger: 'axis'},
      xAxis: {type: 'category', data: ['周一', '周二', '周三', '周四', '周五', '周六', '周日']},
      yAxis: {type: 'value'},
      series: [{data: [120, 132, 101, 134, 90, 230, 210], type: 'line'}]
    },
    height: '300px',
    isFullscreen: false
  },
  {
    title: '产品分布',
    option: {
      tooltip: {trigger: 'item'},
      series: [{
        type: 'pie',
        data: [
          {value: 335, name: '商品A'},
          {value: 310, name: '商品B'},
          {value: 234, name: '商品C'},
          {value: 135, name: '商品D'}
        ]
      }]
    },
    isFullscreen: false
  },
  {
    title: '用户活跃度',
    option: {
      xAxis: {type: 'value'},
      yAxis: {type: 'category', data: ['北京', '上海', '广州', '深圳']},
      series: [{data: [200, 180, 160, 140], type: 'bar'}]
    },
    isFullscreen: false
  }
])
</script>

<style scoped>
.layout {
  display: flex;
  flex-direction: column;
  height: 100%;
  background-color: var(--billadm-color-major-backgroud-color);
}

.top-bar {
  background: var(--billadm-color-major-backgroud-color);
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