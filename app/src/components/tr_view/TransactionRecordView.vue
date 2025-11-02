<template>
  <a-layout style="height: 100%">
    <a-layout-header class="headerStyle">
      <div class="left-groups">
        <BilladmTimeRangePicker
            v-model:time-range="trQueryConditionStore.timeRange"
            v-model:time-range-type="trQueryConditionStore.timeRangeType"
        />
      </div>
      <div class="center-groups">
      </div>
      <div class="right-groups">
        <a-button type="primary" @click="createTr">
          新增记录
        </a-button>
      </div>
    </a-layout-header>
    <a-layout-content :style="contentStyle">
      <transaction-record-table :items="tableData" @edit="updateTr" @delete="deleteTr"/>
    </a-layout-content>
    <a-layout-footer class="footerStyle">
      <a-pagination
          v-model:current="currentPage"
          v-model:pageSize="pageSize"
          :total="trTotal"
          :show-total="total => `共${total}记录`"
          show-size-changer
      />
    </a-layout-footer>
    <a-drawer
        :title="drawerTitle"
        :open="openTrDrawer"
        :body-style="{ paddingBottom: '80px' }"
        :footer-style="{ textAlign: 'right' }"
        @close="closeTrDrawer"
        :closable="false"
    >
      <a-form :model="trForm">
        <a-form-item label="时间" name="time">
          <a-date-picker v-model:value="trForm.time" style="width: 100%"/>
        </a-form-item>

        <a-form-item label="类型" name="type">
          <a-radio-group v-model:value="trForm.type" button-style="solid">
            <a-radio-button value="income">收入</a-radio-button>
            <a-radio-button value="expense">支出</a-radio-button>
            <a-radio-button value="transfer">转账</a-radio-button>
          </a-radio-group>
        </a-form-item>

        <a-form-item label="分类" name="category">
          <a-select v-model:value="trForm.category" :options="categories"/>
        </a-form-item>

        <a-form-item label="标签" name="tags">
          <a-select v-model:value="trForm.tags" :options="tags" mode="multiple" placeholder="选择一个或多个标签"/>
        </a-form-item>

        <a-form-item label="描述" name="description">
          <a-input v-model:value="trForm.description" placeholder="描述消费内容" allowClear/>
        </a-form-item>

        <a-form-item label="金额" name="price">
          <a-input-number v-model:value="trForm.price" prefix="￥" :controls="false" :min="0" style="width: 100%"/>
        </a-form-item>
      </a-form>
      <template #extra>
        <a-space>
          <a-button @click="closeTrDrawer">取消</a-button>
          <a-button type="primary" @click="onConfirm">确认</a-button>
        </a-space>
      </template>
    </a-drawer>
  </a-layout>
</template>

<script setup lang="ts">
import {type CSSProperties, ref, watch} from 'vue';
import TransactionRecordTable from '@/components/tr_view/TransactionRecordTable.vue';
import type {TransactionRecord, TrForm, TrQueryCondition} from "@/types/billadm";
import {useCssVariables} from "@/backend/css.ts";
import {convertToUnixTimeRange} from "@/backend/timerange.ts";
import {
  createTransactionRecord,
  deleteTransactionRecord,
  getCategoryByType,
  getTagsByCategory,
  getTrOnCondition,
  getTrTotalOnCondition,
  updateTransactionRecord
} from "@/backend/functions.ts";
import {useLedgerStore} from "@/stores/ledgerStore.ts";
import {useTrQueryConditionStore} from "@/stores/trQueryConditionStore.ts";
import dayjs from "dayjs";
import {trDtoToTrForm, trFormToTrDto} from "@/backend/dto-utils.ts";
import type {DefaultOptionType} from "ant-design-vue/es/vc-cascader";

const {majorBgColor} = useCssVariables();

const contentStyle: CSSProperties = {
  backgroundColor: majorBgColor.value,
  "overflow-y": "auto",
  "margin-bottom": "auto"
};

const ledgerStore = useLedgerStore();
const trQueryConditionStore = useTrQueryConditionStore();

// 消费记录
const tableData = ref<TransactionRecord[]>([]);
// 分页
const currentPage = ref<number>(1);
const pageSize = ref<number>(20);
const trTotal = ref<number>(0);

const refreshTable = async () => {
  const trTotalCondition: TrQueryCondition = {
    ledgerId: ledgerStore.currentLedgerId,
    tsRange: convertToUnixTimeRange(trQueryConditionStore.timeRange)
  };
  trTotal.value = await getTrTotalOnCondition(trTotalCondition);
  const trCondition: TrQueryCondition = {
    ledgerId: ledgerStore.currentLedgerId,
    tsRange: convertToUnixTimeRange(trQueryConditionStore.timeRange),
    offset: pageSize.value * (currentPage.value - 1),
    limit: pageSize.value
  };
  tableData.value = await getTrOnCondition(trCondition);
}

const defaultTrForm: TrForm = {
  id: '',
  price: 0,
  type: '',
  category: '',
  description: '-',
  tags: [],
  time: dayjs()
};
const openTrDrawer = ref(false);
const drawerTitle = ref('');
const trForm = ref<TrForm>({} as TrForm);
const categories = ref<DefaultOptionType[]>([]);
const tags = ref<DefaultOptionType[]>([]);

const createTr = () => {
  trForm.value.type = 'expense';
  if (trQueryConditionStore.timeRange) {
    trForm.value.time = trQueryConditionStore.timeRange[1];
  }
  drawerTitle.value = '新增消费记录';
  openTrDrawer.value = true;
}

const updateTr = (tr: TransactionRecord) => {
  drawerTitle.value = '编辑消费记录';
  trForm.value = trDtoToTrForm(tr);
  openTrDrawer.value = true;
}

const deleteTr = async (tr: TransactionRecord) => {
  await deleteTransactionRecord(tr.transactionId);
  await refreshTable();
}

const closeTrDrawer = () => {
  trForm.value = defaultTrForm;
  openTrDrawer.value = false;
}

const onConfirm = async () => {
  trForm.value.time = trForm.value.time.hour(12).minute(0).second(0);
  const tr = trFormToTrDto(trForm.value, ledgerStore.currentLedgerId);
  if (tr.transactionId === '') {
    // 新建
    await createTransactionRecord(tr);
  } else {
    // 更新
    await updateTransactionRecord(tr);
  }
  await refreshTable();
  closeTrDrawer();
}

// 查询条件变化 → 重置分页 + 刷新
watch(() => [trQueryConditionStore.timeRange, ledgerStore.currentLedgerId], async () => {
      if (!ledgerStore.currentLedgerId) return;
      if (currentPage.value !== 1) {
        currentPage.value = 1;
        return;
      }
      await refreshTable();
    }
);

// 分页变化 → 仅刷新
watch(() => [currentPage.value, pageSize.value], async () => {
  await refreshTable();
});
/**
 * 交易类型变化时要重新刷新分类列表并如果当前分类为空则选择第一个分类作为分类
 * 如果当前分类不为空则选择查看分类是否在列表中，不在列表中则需要选择第一个分类作为分类
 */
watch(() => trForm.value.type, async () => {
      if (trForm.value.type === '') return;
      const categoryList = await getCategoryByType(trForm.value.type);
      categories.value = categoryList.map(category => {
        return {
          value: category.name,
        };
      });
      const categoryNames = categoryList.map(category => category.name);
      if (categoryNames.length > 0) {
        if (!trForm.value.category || !categoryNames.includes(trForm.value.category)) {
          trForm.value.category = categoryNames[0] as string;
        }
      } else {
        trForm.value.category = '';
      }
    }
);
/**
 * 分类变化时要重新刷新标签列表清除不在候选中的标签
 */
watch(() => trForm.value.category, async () => {
      if (trForm.value.category === '') return;
      const tagList = await getTagsByCategory(trForm.value.category);
      tags.value = tagList.map(tag => {
        return {
          value: tag.name,
        };
      });
      const tagNames = tagList.map(tag => tag.name);
      if (tagNames.length > 0) {
        let newTags: string[] = [];
        trForm.value.tags.forEach(tag => {
          if (tag && tagNames.includes(tag)) {
            newTags.push(tag);
          }
        });
        trForm.value.tags = newTags;
      } else {
        trForm.value.tags = [];
      }
    }
);
</script>

<style scoped>
.headerStyle {
  height: auto;
  background-color: var(--billadm-color-major-background);
  padding: 0 0 16px 0;
  display: flex;
  align-items: start;
  justify-content: center;
}

.footerStyle {
  height: auto;
  background-color: var(--billadm-color-major-background);
  padding: 16px 0 0 0;
  display: flex;
  align-items: end;
  justify-content: center;
}

/* 左边按钮 将它与后面的元素隔开 */
.left-groups {
  margin-right: auto;
  display: flex;
  gap: 8px;
  align-items: center;
}

/* 中间按钮 */
.center-groups {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  gap: 8px;
}

/* 右边按钮组 */
.right-groups {
  display: flex;
  gap: 8px;
}
</style>