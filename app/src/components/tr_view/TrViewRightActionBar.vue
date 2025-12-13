<template>
  <div class="menu-bar">
    <div class="top-groups">
      <a-button type="primary" @click="openFilterDrawer=true">
        <template #icon>
          <FilterOutlined style="display: flex;justify-content: center;align-items: center;font-size: medium"/>
        </template>
      </a-button>
    </div>
    <div class="bottom-groups">
    </div>
  </div>
  <a-drawer
      title="筛选消费记录"
      :open="openFilterDrawer"
      @close="openFilterDrawer=false"
      :body-style="{ paddingBottom: '80px' }"
      :footer-style="{ textAlign: 'right' }"
      :closable="false"
  >
  </a-drawer>
</template>

<script setup lang="ts">
import {ref} from 'vue';
import {FilterOutlined} from "@ant-design/icons-vue";

const openFilterDrawer = ref(false);

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