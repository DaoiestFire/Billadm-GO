<!-- components/BilladmLabel.vue -->
<template>
  <span class="billadm-label">
    <slot>{{ label }}</slot>
  </span>
</template>

<script setup lang="ts">
withDefaults(defineProps<{ label?: string }>(), {
  label: ''
})
</script>

<style scoped>
.billadm-label {
  color: var(--billadm-color-text-minor-color);
}
</style>