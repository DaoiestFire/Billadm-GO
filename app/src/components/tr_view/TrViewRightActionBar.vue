<template>
  <div class="menu-bar">
    <div class="top-groups">
      <a-button type="primary">
        <template #icon>
          <FilterOutlined style="display: flex;justify-content: center;align-items: center;font-size: medium"/>
        </template>
      </a-button>
    </div>
    <div class="bottom-groups">
    </div>
  </div>
</template>

<script setup lang="ts">
import {FilterOutlined} from "@ant-design/icons-vue";
</script>

<style scoped>
.menu-bar {
  display: flex;
  flex-direction: column;
  width: 100%;
  height: 100%;
  align-items: center;
  justify-content: center;
  padding: 8px 0;
}

.top-groups {
  margin-bottom: auto;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.bottom-groups {
  display: flex;
  flex-direction: column;
  gap: 8px;
}
</style>