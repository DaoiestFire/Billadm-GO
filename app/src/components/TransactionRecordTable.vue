<template>
  <div class="tr-table-container" :style="containerStyle">
    <table class="tr-table">
      <thead>
        <tr :style="headerRowStyle">
          <th>序号</th>
          <th>价格</th>
          <th>交易类型</th>
          <th>消费类型</th>
          <th>描述</th>
          <th>标签</th>
          <th>交易时间</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, index) in displayedItems" :key="item.transaction_id" :style="rowStyle">
          <td>{{ index + 1 }}</td>
          <td :style="getPriceStyle(item.transaction_type)">{{ item.price }}</td>
          <td>{{ formatTransactionType(item.transaction_type) }}</td>
          <td>{{ item.category }}</td>
          <td>{{ item.description }}</td>
          <td>{{ formatTags(item.tags) }}</td>
          <td>{{ formatTime(item.transaction_at) }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'

const props = defineProps({
  items: {
    type: Array,
    required: true
  },
   headerHeight: {
    type: Number,
    default: 50 // 默认表头高度
  },
  rowHeight: {
    type: Number,
    default: 40 // 默认行高
  },
  fontSize: {
    type: String,
    default: '14px' // 默认字体大小
  }
})

const displayedItems = ref([])

// 动态计算最大显示行数
const calculateMaxRows = () => {
  const availableHeight = window.innerHeight - 62 - props.headerHeight-1
  const maxRowsCount = Math.floor(availableHeight / props.rowHeight)
  displayedItems.value = props.items.slice(0, maxRowsCount)
}

// 获取价格样式
const getPriceStyle = (type) => {
  let color = ''
  if (type === 'expense') {
    color = 'red'
  } else if (type === 'income') {
    color = 'green'
  } else if (type === 'transfer') {
    color = 'orange'
  }
  return { color, fontSize: props.fontSize }
}

// 格式化交易类型
const formatTransactionType = (type) => {
  return type === 'expense' ? '支出' : type === 'income' ? '收入' : '转账'
}

// 格式化标签数组为字符串
const formatTags = (tags) => {
  return tags && tags.length ? tags.join('，') : ''
}

// 格式化时间（去掉毫秒）
const formatTime = (timeStr) => {
  if (!timeStr) return ''
  const date = new Date(timeStr)
  return date.toLocaleString()
}

// 监听窗口变化
const onResize = () => {
  calculateMaxRows()
}

onMounted(() => {
  calculateMaxRows()
  window.addEventListener('resize', onResize)
})

watch(() => [props.items, props.rowHeight, props.headerHeight], () => {
  calculateMaxRows()
})

const containerStyle = {
  '--row-height': `${props.rowHeight}px`,
  '--header-height': `${props.headerHeight}px`,
  '--font-size': props.fontSize
}

const headerRowStyle = {
  height: 'var(--header-height)',
  lineHeight: 'var(--header-height)',
  fontSize: 'var(--font-size)'
}

const rowStyle = {
  height: 'var(--row-height)',
  lineHeight: 'var(--row-height)',
  fontSize: 'var(--font-size)'
}
</script>

<style scoped>
.tr-table-container {
  overflow-y: auto;
  max-height: calc(100vh - 62px);
}

.tr-table {
  width: 100%;
  border-collapse: collapse;
  table-layout: fixed;
}

.tr-table th,
.tr-table td {
  padding: 0; /* 调整内边距为0，因为line-height已经处理了垂直居中 */
  text-align: center;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  border-bottom: 1px solid var(--billadm-color-window-border-color);
}

.tr-table thead th {
  background-color: var(--billadm-color-table-header-bg-color);
  position: sticky;
  top: 0;
  z-index: 1;
}

/* 斑马纹样式 */
.tr-table tbody tr:nth-child(odd) {
  background-color: var(--billadm-color-table-odd-row-bg-color);
}

.tr-table tbody tr:nth-child(even) {
  background-color: var(--billadm-color-table-even-row-bg-color);
}
</style>