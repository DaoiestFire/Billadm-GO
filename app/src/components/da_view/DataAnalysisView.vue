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
      <billadm-chart-display :charts="chartData"/>
    </div>
  </div>
</template>

<script setup>
import {ref} from 'vue';
import BilladmChartDisplay from "@/components/da_view/BilladmChartDisplay.vue";
import BilladmTimeSelect from "@/components/BilladmTimeSelect.vue";
import {useTrViewStore} from "@/stores/trViewStore.js";

// store
const trViewStore = useTrViewStore();

// 组件响应式变量
const timeRangeType = ref(trViewStore.timeRangeType);
const timeRange = ref([new Date(trViewStore.timeRange[0]), new Date(trViewStore.timeRange[1])]);

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