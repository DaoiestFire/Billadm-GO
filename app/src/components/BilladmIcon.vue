<template>
  <div
      class="billadm-icon"
      :style="iconStyles"
      v-html="processedSvg"
  ></div>
</template>

<script setup>
import {computed} from 'vue'

// 定义组件的 props
const props = defineProps({
  // 必须传入的 SVG 图像字符串
  svg: {
    type: String,
    required: true
  },
  color: {
    type: String,
    default: 'currentColor'
  },
  size: {
    type: [String, Number],
    default: '20px'
  },
  // 图标粗细（控制描边宽度）
  weight: {
    type: [String, Number],
    default: 0.1,
    validator: (value) => {
      // 验证值是否为有效数字或可转换为数字的字符串
      const num = typeof value === 'string' ? parseFloat(value) : value;
      return !isNaN(num) && num >= 0;
    }
  }
})

const processedSvg = computed(() => {
  const tempDiv = document.createElement('div')
  tempDiv.innerHTML = props.svg

  const svgElement = tempDiv.querySelector('svg')

  svgElement.removeAttribute('fill')
  svgElement.removeAttribute('stroke')
  svgElement.removeAttribute('stroke-width')

  // 设置样式，确保颜色和粗细优先级
  const strokeStyle = props.weight > 0 ? `stroke-width: ${props.weight}px;` : ''
  svgElement.setAttribute('style', `fill: ${props.color}; stroke: ${props.color}; ${strokeStyle}`)

  // 移除 width 和 height 属性，由外部容器控制大小
  svgElement.removeAttribute('width')
  svgElement.removeAttribute('height')

  return svgElement.outerHTML
})

const iconStyles = computed(() => ({
  width: typeof props.size === 'number' ? `${props.size}px` : props.size,
  height: typeof props.size === 'number' ? `${props.size}px` : props.size,
  display: 'inline-flex',
  alignItems: 'center',
  justifyContent: 'center',
  flexShrink: 0,
  lineHeight: 0
}))
</script>

<style scoped>
</style>