<template>
  <div class="layout">
    <!-- 上栏：工具栏 -->
    <div class="top-bar">
      <div class="left-groups">
        <el-date-picker
            v-model="trViewStore.timeRange"
            type="daterange"
            range-separator="~"
            start-placeholder="起始时间"
            end-placeholder="结束时间"
            size="small"
            :editable="false"
            style="width: 200px;"
            :shortcuts="shortcuts"
        />
      </div>
      <div class="center-groups">
      </div>
      <div class="right-groups">
        <button class="btn-add" @click="onCreate">新增记录</button>
      </div>
    </div>

    <!-- 中间栏：消费记录表 -->
    <div class="middle-section">
      <TransactionRecordTable :items="trViewStore.tableData" :columnStyles="columnStyles" :headerHeight="40"
                              :rowHeight="40" @edit-item="onEditItem"/>
    </div>

    <!-- 下栏：分页组件 -->
    <div class="bottom-bar">
      <CustomSelect v-model="trViewStore.pageSize" :options="options" direction="up"/>
      <Pagination v-model:current-page="trViewStore.currentPage" :pages="trViewStore.pages"/>
    </div>
  </div>
  <TransactionRecordOperation v-model:modelValue="recordData" v-model:visible="showDialog" :title="dialogTitle"
                              :onConfirm="handleConfirm"/>
</template>

<script setup>
import {ref} from 'vue'
import {useCssVariables} from '@/css/css'
import TransactionRecordTable from '@/components/TransactionRecordTable.vue'
import Pagination from '@/components/Pagination.vue'
import CustomSelect from '@/components/CustomSelect.vue'
import TransactionRecordOperation from '@/components/TransactionRecordOperation.vue'
import {useLedgerStore} from "@/stores/ledgerStore.js";
import {useTrViewStore} from "@/stores/trViewStore.js";
import NotificationUtil from "@/backend/notification.js";
import {createTrForLedger, deleteTrById, trDtoToTrForm, trFormToTrDto} from "@/backend/tr.js";
import {
  getLastMonthRange,
  getLastWeekRange,
  getThisMonthRange,
  getThisWeekRange,
  getTodayRange
} from "@/backend/timerange.js";

// store
const ledgerStore = useLedgerStore()
const trViewStore = useTrViewStore()

// 视图常量
const options = [
  {label: '每页10行', value: 10},
  {label: '每页20行', value: 20},
  {label: '每页50行', value: 50}
]

const shortcuts = [
  {
    text: '今天',
    value: getTodayRange,
  },
  {
    text: '本周',
    value: getThisWeekRange,
  },
  {
    text: '本月',
    value: getThisMonthRange,
  },
  {
    text: '上周',
    value: getLastWeekRange,
  },
  {
    text: '上月',
    value: getLastMonthRange,
  }
]

const columnStyles = [
  {
    field: 'index',
    name: '序号',
    width: '100px',
  },
  {
    field: 'transaction_at',
    name: '交易时间',
    width: '200px',
  },
  {
    field: 'transaction_type',
    name: '交易类型',
    width: '100px',
  },
  {
    field: 'category',
    name: '消费类型',
    width: '100px',
  },
  {
    field: 'description',
    name: '描述',
    width: 'auto',
  },
  {
    field: 'tags',
    name: '标签',
    width: 'auto',
  },
  {
    field: 'price',
    name: '价格',
  },
  {
    field: 'actions',
    name: '操作',
    width: '120px',
  }
]

// css variables
const {minorBgColor, hoverBgColor, iconColor} = useCssVariables()


// 消费记录表单
const showDialog = ref(false);
const recordData = ref({});
const opType = ref('create')
const dialogTitle = ref('')

// 功能函数
const onCreate = () => {
  dialogTitle.value = '创建消费记录'
  recordData.value = {}
  opType.value = 'create'
  showDialog.value = true
}

const onEditItem = (item) => {
  dialogTitle.value = '编辑消费记录'
  recordData.value = trDtoToTrForm(item)
  opType.value = 'edit'
  showDialog.value = true
}

async function handleConfirm(data) {
  try {
    const transactionRecord = trFormToTrDto(data, ledgerStore.currentLedgerId)
    await createTrForLedger(transactionRecord)
    if (opType.value === 'edit') {
      await deleteTrById(data.id)
    }
    await trViewStore.init()
  } catch (error) {
    NotificationUtil.error(`消费记录操作失败 ${error}`)
  }
}
</script>

<style scoped>
.layout {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.top-bar {
  background: var(--billadm-color-minor-backgroud-color);
  height: 40px;
  border-bottom: 1px solid var(--billadm-color-window-border-color);
  flex-shrink: 0;
  display: flex;
  align-items: center;
  position: relative;
}

/* 左边按钮 将它与后面的元素隔开 */
.left-groups {
  margin-left: 20px;
  margin-right: auto;
  display: flex;
  gap: 4px;
}

/* 中间按钮 */
.center-groups {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  gap: 4px;
}

/* 右边按钮组 */
.right-groups {
  margin-right: 20px;
  display: flex;
  gap: 4px;
}

.btn-add {
  background-color: transparent;
  color: var(--billadm-color-positive-color);
  padding: 4px 8px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.btn-add:hover {
  background-color: var(--billadm-color-positive-color);
  color: var(--billadm-color-hover-fg-color);
}

.middle-section {
  flex: 1;
  overflow: hidden;
}

.bottom-bar {
  background: var(--billadm-color-minor-backgroud-color);
  height: 40px;
  border-top: 1px solid var(--billadm-color-window-border-color);
  flex-shrink: 0;
  display: flex;
  flex-direction: row;
  justify-content: center;
}
</style>