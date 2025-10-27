<template>
  <div class="billadm-time-select">
    <div class="group">
      <a-typography-text>时间范围</a-typography-text>
      <!-- 向前按钮 -->
      <a-button type="text">
        <template #icon>
          <LeftOutlined style="display: flex;justify-content: center;align-items: center;font-size: large"/>
        </template>
      </a-button>
      <a-range-picker v-model:value="timeRange" :picker="timeRangeType"/>
      <a-button type="text">
        <template #icon>
          <RightOutlined style="display: flex;justify-content: center;align-items: center;font-size: large"/>
        </template>
      </a-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import {ref} from 'vue';
import {LeftOutlined, RightOutlined} from "@ant-design/icons-vue";
import type {Dayjs} from 'dayjs';
import type {TimeRangeType} from "@/types/billadm";

type RangeValue = [Dayjs, Dayjs];

const timeRange = ref<RangeValue>();
const timeRangeType = ref<TimeRangeType>('date');

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