<template>
  <a-row :gutter="16" style="width: 100%;padding: 16px">
    <a-col v-for="(chart, index) in charts" :key="index" :span="colSpan">
      <billadm-fullscreen v-model="chart.isFullscreen" :dblclick="true">
        <billadm-chart-panel :title="chart.title" :data="chart.data"/>
      </billadm-fullscreen>
    </a-col>
  </a-row>
</template>

<script setup lang="ts">
import {computed} from 'vue';
import BilladmChartPanel from '@/components/da_view/BilladmChartPanel.vue';
import BilladmFullscreen from '@/components/common/BilladmFullScreen.vue';
import type {TransactionRecord} from '@/types/billadm';

interface Props {
  trList: TransactionRecord[];
  columns?: number;
}

const props = withDefaults(defineProps<Props>(), {
  columns: 2,
});

const charts = computed(() => [
  {title: '交易走势', data: props.trList, isFullscreen: false},
  {title: '消费分布', data: props.trList, isFullscreen: false},
]);

const colSpan = computed(() => {
  const cols = props.columns;
  if (cols <= 0 || 24 % cols !== 0) {
    console.warn(`columns 应为 24 的约数（如 1,2,3,4,6,8,12,24），当前值: ${cols}，回退到 2 列`);
    return 12;
  }
  return 24 / cols;
});
</script>