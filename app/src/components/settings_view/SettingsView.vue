<template>
  <el-tabs v-model="activeName" class="billadm-tabs">
    <el-tab-pane label="工作空间" name="workspace">工作空间</el-tab-pane>
    <el-tab-pane label="消费类型" name="category">消费类型</el-tab-pane>
    <el-tab-pane label="消费标签" name="tags">消费标签</el-tab-pane>
  </el-tabs>
</template>

<script setup>
import {ref} from 'vue'

const activeName = ref('workspace')
</script>

<style>
</style>