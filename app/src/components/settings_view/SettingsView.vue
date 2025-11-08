<template>
  <a-tabs v-model:activeKey="activeKey" style="height: 100%">
    <a-tab-pane key="category&tag" tab="分类与标签">
      <billadm-category-tag-setting/>
    </a-tab-pane>
    <a-tab-pane key="workspace" tab="工作空间">
    </a-tab-pane>
  </a-tabs>
</template>
<script lang="ts" setup>
import {ref} from 'vue';

const activeKey = ref('category&tag');
</script>

<style scoped>
:deep(.ant-tabs-content-holder) {
  overflow-y: auto;
}
</style>