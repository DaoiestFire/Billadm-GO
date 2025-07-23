<template>
  <div class="tr-table-container">
    <table class="tr-table">
      <thead>
        <tr>
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
        <tr v-for="(item, index) in displayedItems" :key="item.transaction_id">
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
import { ref, onMounted, onUnmounted, watch } from 'vue'

const props = defineProps({
  items: {
    type: Array,
    required: true
  },
  rowHeight: {
    type: Number,
    default: 40 // 默认行高
  }
})

const displayedItems = ref([])
const maxRows = ref(10)

// 计算当前窗口高度下能显示的行数
const calculateMaxRows = () => {
  const availableHeight = window.innerHeight - 62 - props.rowHeight // 预留头部、底部空间
  maxRows.value = Math.floor(availableHeight / props.rowHeight)
  displayedItems.value = props.items.slice(0, maxRows.value)
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
  return { color }
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

onUnmounted(() => {
  window.removeEventListener('resize', onResize)
})
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
  padding: 8px;
  text-align: center;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.tr-table th {
  background-color: #f5f5f5;
}
</style>