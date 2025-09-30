<template>
  <button class="common-icon"
          :style="[
              buttonStyle,
              {
                '--hover-bg-color': hoverBgColor ,
                '--bg-size': `${props.bgSize}px`,
              }
          ]"
          :class="[
              { 'is-hovered': isHovered || isActive },
              { 'is-active': isActive }
          ]"
          @mouseenter="handleMouseEnter"
          @mouseleave="handleMouseLeave">
    <el-tooltip :content="label" :placement="tooltipPlacement">
      <billadm-icon class="icon" :svg="svg" :color="isActive ? activeFgColor : color" :size="iconSize"/>
    </el-tooltip>
  </button>
</template>

<script setup>
import {computed, ref} from 'vue'
import BilladmIcon from "@/components/BilladmIcon.vue";

const props = defineProps({
  svg: {
    type: String,
    required: true
  },
  label: {
    type: String,
    default: ''
  },
  width: {
    type: String,
    default: 30
  },
  height: {
    type: String,
    default: 30
  },
  // icon
  iconSize: {
    type: [String, Number],
    default: '20px'
  },
  color: {
    type: String,
    default: '#7b7b7b'
  },
  bgColor: {
    type: String,
    default: 'transparent'
  },
  bgSize: {
    type: [String, Number],
    default: 30
  },
  hoverBgColor: {
    type: String,
    default: 'white'
  },
  isActive: {
    type: Boolean,
    default: false
  },
  activeFgColor: {
    type: String,
    default: 'white'
  },
  tooltipPlacement: {
    type: String,
    default: 'right',
  },
})

const isHovered = ref(false)

const handleMouseEnter = () => {
  isHovered.value = true
}

const handleMouseLeave = () => {
  isHovered.value = false
}

const buttonStyle = computed(() => {
  return {
    width: typeof props.width === 'number' ? `${props.width}px` : props.width,
    height: typeof props.height === 'number' ? `${props.height}px` : props.height,
    backgroundColor: props.bgColor
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

.common-icon.is-hovered::after,
.common-icon.is-active::after {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: var(--bg-size);
  height: var(--bg-size);
  border-radius: 4px;
  z-index: 1;
  transition: background-color 0.3s ease;
}

.common-icon.is-hovered::after {
  background-color: var(--billadm-color-icon-hover-bg-color);
}

.common-icon.is-active::after {
  background-color: var(--billadm-color-icon-active-color);
}

.icon {
  z-index: 2;
}
</style>