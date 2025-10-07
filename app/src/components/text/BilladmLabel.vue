<!-- components/BilladmLabel.vue -->
<template>
  <span class="billadm-label">
    <slot>{{ label }}</slot>
  </span>
</template>

<script setup>
defineProps({
  label: {
    type: String,
    default: ''
  }
})
</script>

<style scoped>
.billadm-label {
  color: var(--billadm-color-text-minor-color);
}
</style>