<template>
  <div
      class="billadm-icon"
      :style="iconStyles"
      v-html="processedSvg"
  ></div>
</template>

<script setup lang="ts">
import {computed} from 'vue'

interface Props {
  svg: string                    // 必填，SVG 字符串
  color?: string                 // 图标颜色，默认 currentColor
  size?: string | number         // 图标尺寸
  weight?: string | number       // 描边粗细，支持数字或可解析字符串（如 '0.5'）
}

const props = withDefaults(defineProps<Props>(), {
  color: 'currentColor',
  size: '20px',
  weight: 0.1,
})

const processedSvg = computed(() => {
  const tempDiv = document.createElement('div')
  tempDiv.innerHTML = props.svg

  const svgElement = tempDiv.querySelector('svg')
  if (!svgElement) {
    console.warn('[BilladmIcon] 提供的 svg 字符串中未找到 <svg> 标签')
    return ''
  }

  // 移除可能干扰样式的属性
  svgElement.removeAttribute('fill')
  svgElement.removeAttribute('stroke')
  svgElement.removeAttribute('stroke-width')

  // 解析 weight 为数字
  const weightValue = parseFloat(
      typeof props.weight === 'string' ? props.weight : String(props.weight)
  )

  // 只有 weight > 0 时才设置 stroke
  const strokeStyle = weightValue > 0 ? `stroke-width: ${weightValue}px;` : ''

  // 注入内联样式，确保颜色优先级
  svgElement.setAttribute(
      'style',
      `fill: ${props.color}; stroke: ${props.color}; ${strokeStyle}`
  )

  // 移除 width/height，由外部控制
  svgElement.removeAttribute('width')
  svgElement.removeAttribute('height')

  return svgElement.outerHTML
})

const iconStyles = computed(() => {
  const size = typeof props.size === 'number' ? `${props.size}px` : props.size
  return {
    width: size,
    height: size,
    display: 'inline-flex',
    alignItems: 'center',
    justifyContent: 'center',
    flexShrink: 0,
    lineHeight: 0,
  }
})
</script>

<style scoped>
</style>