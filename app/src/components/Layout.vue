<template>
  <div class="layout">
    <!-- 上栏：菜单栏 -->
    <div class="top-bar">
      <CommonIcon :icon="homeIcon" width="40" height="30" :color="iconColor" :bgColor="bgColor"
        :hoverBgColor="hoverBgColor" hoverStyle="rect" />
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
      </div>
    </div>

    <!-- 下栏：详情栏 -->
    <div class="bottom-bar">
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import CommonIcon from '@/components/CommonIcon.vue'

const homeIcon = `
  <svg viewBox="0 0 24 24" width="24" height="24">
    <path d="M10 20v-6h4v6h5v-8h3L12 3 2 12h3v8z" fill="currentColor"/>
  </svg>
`

const router = useRouter();

const navItems = [
  { title: '主页', path: '/home' },
  { title: '仪表盘', path: '/dashboard' },
  { title: '设置', path: '/settings' },
];

const navigate = (path) => {
  router.push(path);
};

const bgColor = ref(
  getComputedStyle(document.documentElement)
    .getPropertyValue('--billadm-color-minor-backgroud-color')
    .trim()
);

const hoverBgColor = ref(
  getComputedStyle(document.documentElement)
    .getPropertyValue('--billadm-color-icon-hover-bg-color')
    .trim()
);

const iconColor = ref(
  getComputedStyle(document.documentElement)
    .getPropertyValue('--billadm-color-icon-color')
    .trim()
);
</script>

<style scoped>
.layout {
  display: flex;
  flex-direction: column;
  height: 100vh;
}

.top-bar {
  background: var(--billadm-color-minor-backgroud-color);
  height: 30px;
  border-bottom: 1px solid var(--billadm-color-window-border-color);
}

.middle-section {
  flex: 1;
  display: flex;
  overflow: hidden;
}

.left-panel,
.right-panel {
  width: 40px;
  background: var(--billadm-color-minor-backgroud-color);
}

.left-panel {
  border-right: 1px solid var(--billadm-color-window-border-color);
}

.right-panel {
  border-left: 1px solid var(--billadm-color-window-border-color);
}

.center-panel {
  flex: 1;
  padding: 10px;
  overflow-y: auto;
}

.bottom-bar {
  background: var(--billadm-color-minor-backgroud-color);
  height: 30px;
  border-top: 1px solid var(--billadm-color-window-border-color);
}
</style>