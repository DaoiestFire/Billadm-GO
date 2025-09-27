<template>
  <div class="pagination" :style="{ width }">
    <button class="page-item" :disabled="currentPage === 1" @click="handlePageChange(currentPage - 1)">
      上一页
    </button>

    <div class="page-list">
      <template v-for="page in displayedPages" :key="page">
        <span v-if="page === '...'" class="page-ellipsis">...</span>
        <button v-else class="page-item number" :class="{ active: currentPage === page }"
                @click="handlePageChange(page)">
          {{ page }}
        </button>
      </template>
    </div>

    <button class="page-item" :disabled="currentPage === pages" @click="handlePageChange(currentPage + 1)">
      下一页
    </button>
  </div>
</template>

<script setup>
import {computed, ref} from 'vue'

const props = defineProps({
  pages: {
    type: Number,
    required: true,
    default: 1,
    validator: value => value >= 1
  },
  width: {
    type: String,
    default: 'auto'
  }
})

const emit = defineEmits(['update:current-page'])

const currentPage = ref(1)

const displayedPages = computed(() => {
  const total = props.pages
  const current = currentPage.value
  const maxVisible = 10

  if (total <= maxVisible) {
    // 页数少于等于10，直接显示所有页码
    return Array.from({length: total}, (_, i) => i + 1)
  }

  const pages = []

  // 前6个
  if (current <= 5) {
    for (let i = 1; i <= 6; i++) pages.push(i)
    pages.push('...')
    pages.push(total)
  }
  // 后6个
  else if (current > total - 5) {
    pages.push(1)
    pages.push('...')
    for (let i = total - 5; i <= total; i++) pages.push(i)
  }
  // 中间5个
  else {
    pages.push(1)
    pages.push('...')
    for (let i = current - 2; i <= current + 2; i++) pages.push(i)
    pages.push('...')
    pages.push(total)
  }

  return pages
})

function handlePageChange(page) {
  if (page < 1 || page > props.pages || page === currentPage.value) return
  currentPage.value = page
  emit('update:current-page', page)
}
</script>

<style scoped>
.pagination {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  padding: 0 10px;
  user-select: none;
  font-size: 14px;
}

.page-list {
  display: flex;
  gap: 4px;
  width: auto;
  justify-content: center;
}

.page-item {
  padding: 4px 8px;
  border: 1px solid var(--billadm-color-window-border-color);
  background-color: #f9f9f9;
  cursor: pointer;
  border-radius: 4px;
  transition: all 0.2s ease;
}

.page-item.number {
  width: 30px;
}

.page-item:hover:not([disabled]) {
  background-color: var(--billadm-color-icon-hover-bg-color);
  border-color: var(--billadm-color-window-border-color);
  color: black;
}

.page-item.active {
  background-color: var(--billadm-color-icon-active-color);
  color: var(--billadm-color-icon-active-fg-color);
  border-color: var(--billadm-color-icon-active-color);
}

.page-item:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.page-ellipsis {
  width: 14px;
  padding: 4px 8px;
  color: #666;
}
</style>