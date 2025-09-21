<!-- components/MenuIcon.vue -->
<template>
    <span v-if="rawSvg" class="menu-icon-container" :style="{ width: containerSize, height: containerSize }">
        <span class="menu-icon" :style="iconStyles" v-html="rawSvg"></span>
    </span>
</template>

<script setup>
import {computed} from 'vue';

const props = defineProps({
  // 原始 SVG 字符串
  rawSvg: {
    type: String,
    default: '',
  },
  // 图标大小，如 '16px', '1.2em'
  iconSize: {
    type: String,
    default: '18px',
  },
  // 容器大小，如 '24px', '2em'
  containerSize: {
    type: String,
    default: '40px',
  },
  // 图标颜色（会作用于 fill 和 stroke）
  color: {
    type: String,
    default: '#7b7b7b',
  },
});

const iconStyles = computed(() => ({
  width: props.iconSize,
  height: props.iconSize,
  display: 'block',
  margin: 'auto',
  color: props.color,
  fill: props.color,
  stroke: props.color,
  strokeWidth: '0.1px'
}));
</script>

<style scoped>
.menu-icon-container {
  display: inline-flex;
  justify-content: center;
  align-items: center;
}

.menu-icon {
  max-width: 100%;
  max-height: 100%;
}
</style>