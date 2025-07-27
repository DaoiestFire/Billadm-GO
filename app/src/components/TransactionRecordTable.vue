<template>
  <div class="tr-table-container" :style="containerStyle">
    <table class="tr-table">
      <thead>
        <tr :style="headerRowStyle">
          <th v-for="styleItem in columnStyles" :key="styleItem.field" :style="getColumnStyle(styleItem)">
            {{ styleItem.name }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, index) in items" :key="item.transaction_id" :style="rowStyle">
          <td v-for="styleItem in columnStyles" :key="styleItem.field"
            :style="formatCellStyle(styleItem.field, item)">
            {{ styleItem.field === 'index' ? index + 1 : formatCell(styleItem.field, item[styleItem.field]) }}
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
const props = defineProps({
  items: {
    type: Array,
    required: true
  },
  headerHeight: {
    type: Number,
    default: 50
  },
  rowHeight: {
    type: Number,
    default: 40
  },
  columnStyles: {
    type: Array,
    default: []
  },
})

// 获取表头风格
const getColumnStyle = (item) => {
  const width = item.width || 'auto'
  return {
    width,
    minWidth: width === 'auto' ? 'auto' : undefined,
    maxWidth: width === 'auto' ? 'none' : undefined
  }
}

// 获取表格样式
const formatCellStyle = (field, item) => {
  switch (field) {
    case 'price':
      return formatPriceStyle(item.transaction_type)
    default:
      return {}
  }
}

// 获取价格样式
const formatPriceStyle = (type) => {
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

// 格式化数据
const formatCell = (field, value) => {
  switch (field) {
    case 'transaction_at':
      return formatTime(value)
    case 'transaction_type':
      return formatTransactionType(value)
    case 'tags':
      return formatTags(value)
    default:
      return value
  }
}

// 格式化时间
const formatTime = (timeStr) => {
  if (!timeStr) return ''
  const date = new Date(timeStr)
  return date.toLocaleString()
}

// 格式化交易类型
const formatTransactionType = (type) => {
  return type === 'expense' ? '支出' : type === 'income' ? '收入' : '转账'
}

// 格式化标签数组为字符串
const formatTags = (tags) => {
  return tags && tags.length ? tags.join('，') : ''
}

const containerStyle = {
  '--row-height': `${props.rowHeight}px`,
  '--header-height': `${props.headerHeight}px`,
}

const headerRowStyle = {
  height: 'var(--header-height)',
  lineHeight: 'var(--header-height)',
}

const rowStyle = {
  height: 'var(--row-height)',
  lineHeight: 'var(--row-height)',
}
</script>

<style scoped>
.tr-table-container {
  overflow-y: auto;
  height: 100%;
}

.tr-table {
  width: 100%;
  border-collapse: collapse;
  table-layout: fixed;
}

.tr-table th,
.tr-table td {
  padding: 0;
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