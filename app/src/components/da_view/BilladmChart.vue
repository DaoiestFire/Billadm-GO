<template>
  <v-chart
      :option="option"
      :autoresize="true"
      class="billadm-chart-instance"
  />
</template>

<script setup>
import VChart from 'vue-echarts';

const props = defineProps({
  option: {
    type: Object,
    required: true
  }
});
</script>

<style scoped>
.billadm-chart-instance {
  width: 100%;
  height: 100%;
}
</style>