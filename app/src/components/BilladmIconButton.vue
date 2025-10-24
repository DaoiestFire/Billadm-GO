<template>
  <button
      class="common-icon"
      :style="buttonStyle"
      :class="{ 'is-active': isActive }">
    <el-tooltip v-if="label" :content="label" :placement="tooltipPlacement">
      <billadm-icon class="icon" :svg="svg" :color="isActive ? activeFgColor : color" :size="iconSize"/>
    </el-tooltip>
    <billadm-icon v-else class="icon" :svg="svg" :color="isActive ? activeFgColor : color" :size="iconSize"/>
  </button>
</template>

<script setup lang="ts">
import {computed} from 'vue'
import BilladmIcon from '@/components/BilladmIcon.vue'

interface Props {
  svg: string                    // 必填，图标 SVG 路径或名称
  label?: string                 // 可选，tooltip 文本
  width?: string                 // 按钮宽度
  height?: string                // 按钮高度
  iconSize?: string | number     // 图标大小
  color?: string                 // 默认图标颜色
  bgColor?: string               // 按钮背景色
  bgSize?: string | number       // 悬浮/激活时背景圆大小
  hoverBgColor?: string          // 悬浮时背景色
  isActive?: boolean             // 是否处于激活状态
  activeFgColor?: string         // 激活时图标颜色
  tooltipPlacement?: 'top' | 'top-start' | 'top-end' | 'bottom' | 'bottom-start' | 'bottom-end' | 'left' | 'left-start' | 'left-end' | 'right' | 'right-start' | 'right-end'
}

const props = withDefaults(defineProps<Props>(), {
  label: '',
  width: '30px',
  height: '30px',
  iconSize: '20px',
  color: '#7b7b7b',
  bgColor: 'transparent',
  bgSize: '30px',
  hoverBgColor: 'white',
  isActive: false,
  activeFgColor: 'white',
  tooltipPlacement: 'bottom',
})

const buttonStyle = computed(() => {
  return {
    width: props.width,
    height: props.height,
    backgroundColor: props.bgColor,
  }
})
</script>

<style scoped>
.common-icon {
  position: relative;
  border: none;
  padding: 0;
  cursor: pointer;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  transition: background-color 0.3s ease;
  overflow: hidden;
  z-index: 1;
}

.common-icon:hover::after,
.common-icon.is-active::after {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: v-bind(bgSize);
  height: v-bind(bgSize);
  border-radius: 4px;
  z-index: 1;
  transition: background-color 0.3s ease;
}

.common-icon:hover::after {
  background-color: v-bind(hoverBgColor);
}

.common-icon.is-active::after {
  background-color: var(--billadm-color-icon-active-color);
}

.icon {
  z-index: 2;
}
</style>