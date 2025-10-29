<template>
  <a-tabs v-model:activeKey="activeKey">
    <a-tab-pane key="category&tag" tab="类型与标签">Content of Tab Pane 1</a-tab-pane>
  </a-tabs>
</template>
<script lang="ts" setup>
import {ref} from 'vue';

const activeKey = ref('category&tag');
</script>