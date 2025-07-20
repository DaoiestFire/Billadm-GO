<template>
  <div class="layout">
    <!-- 上栏：菜单栏 -->
    <div class="top-bar">
      <h2>菜单栏</h2>
    </div>

    <!-- 中间栏：左中右 -->
    <div class="middle-section">
      <!-- 左侧功能栏 -->
      <div class="left-panel">
        <button v-for="item in navItems" :key="item.path" @click="navigate(item.path)">
          {{ item.title }}
        </button>
      </div>

      <!-- 中间阅览区 -->
      <div class="center-panel">
        <router-view />
      </div>

      <!-- 右侧功能栏 -->
      <div class="right-panel">
        <button v-for="btn in rightButtons" :key="btn.name" @click="handleRightClick(btn.action)">
          {{ btn.name }}
        </button>
      </div>
    </div>

    <!-- 下栏：详情栏 -->
    <div class="bottom-bar">
      <p>这里是详情栏，显示一些数据。</p>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router';

const router = useRouter();

const navItems = [
  { title: '主页', path: '/home' },
  { title: '仪表盘', path: '/dashboard' },
  { title: '设置', path: '/settings' },
];

const rightButtons = [
  { name: '刷新', action: 'refresh' },
  { name: '保存', action: 'save' },
];

const navigate = (path) => {
  router.push(path);
};

const handleRightClick = (action) => {
  alert(`执行操作: ${action}`);
};
</script>

<style scoped>
.layout {
  display: flex;
  flex-direction: column;
  height: 100vh;
}

.top-bar {
  background: #333;
  color: white;
  padding: 10px;
  text-align: center;
}

.middle-section {
  flex: 1;
  display: flex;
  overflow: hidden;
}

.left-panel, .right-panel {
  width: 150px;
  background: #f4f4f4;
  padding: 10px;
  border-right: 1px solid #ccc;
}

.left-panel {
  border-right: 1px solid #ccc;
}

.right-panel {
  border-left: 1px solid #ccc;
}

.center-panel {
  flex: 1;
  padding: 10px;
  overflow-y: auto;
}

.bottom-bar {
  background: #eee;
  padding: 10px;
  text-align: center;
  border-top: 1px solid #ccc;
}
</style>