<template>
  <div class="menu-bar">
    <div class="top-groups">
      <a-badge :count="conditionLen" dot>
        <a-button :type="conditionLen?'primary':'text'" @click="openFilterDrawer=true">
          <template #icon>
            <FilterOutlined style="display: flex;justify-content: center;align-items: center;font-size: medium"/>
          </template>
        </a-button>
      </a-badge>
    </div>
    <div class="bottom-groups">
    </div>
  </div>
  <a-drawer
      title="筛选消费记录"
      :open="openFilterDrawer"
      @close="openFilterDrawer=false"
      :body-style="{ paddingBottom: '80px' }"
      :footer-style="{ textAlign: 'right' }"
      :closable="false"
  >
    <a-form :model="trFilterForm" layout="vertical">
      <!-- 交易类型 -->
      <a-divider orientation="left">交易类型</a-divider>
      <a-form-item>
        <a-select
            v-model:value="trFilterForm.transactionTypes"
            :options="transactionTypeOptions"
            mode="multiple"
            placeholder="请选择交易类型"
        />
      </a-form-item>
      <!-- 分类标签 -->
      <a-divider orientation="left">分类标签</a-divider>
      <a-form-item>
        <a-row :gutter="8" justify="space-between">
          <a-col :span="18">
            <a-select
                v-model:value="tempCategory"
                :options="categories"
                placeholder="选择分类"
                @change="onCategoryChange"
                style="width: 100%"
                allow-clear
            />
          </a-col>
          <a-col :span="6">
            <a-button style="width: 100%" type="primary" @click="addCondition" :disabled="!tempCategory">
              添加
            </a-button>
          </a-col>
        </a-row>
      </a-form-item>
      <a-form-item>
        <a-select
            v-model:value="tempTags"
            :options="tags"
            placeholder="选择标签（可选）"
            mode="multiple"
            :disabled="!tempCategory"
            style="width: 100%"
        />
      </a-form-item>
    </a-form>
  </a-drawer>
</template>

<script setup lang="ts">
import {computed, ref, watch} from 'vue';
import {FilterOutlined} from "@ant-design/icons-vue";
import type {Category, TrFilterCondition} from "@/types/billadm";
import type {DefaultOptionType} from "ant-design-vue/es/vc-cascader";
import {getCategoryByType, getTagsByCategory} from "@/backend/functions.ts";

interface Condition {
  category: string,
  tags: string[]
}

const conditions = ref<Condition[]>([])
const openFilterDrawer = ref(false);
const trFilterForm = ref<TrFilterCondition>({
  transactionTypes: [],
  categoryTags: {}
});
const conditionLen = computed(() => {
  let cnt = 0;
  if (trFilterForm.value.transactionTypes) {
    cnt += trFilterForm.value.transactionTypes.length;
  }
  cnt += conditions.value.length;
  return cnt;
});

// ===== 临时输入状态 =====
const tempCategory = ref(undefined);
const tempTags = ref([]);

// 交易类型选项
const transactionTypeOptions = [
  {label: '收入', value: 'income'},
  {label: '支出', value: 'expense'},
  {label: '转账', value: 'transfer'}
]

// 根据选中的交易类型
const categories = ref<DefaultOptionType[]>([]);
const tags = ref<DefaultOptionType[]>([]);
/**
 * 交易类型变化时要重新刷新分类列表
 * 如果当前分类不为空则选择查看分类是否在列表中，不在列表中则需要选择第一个分类作为分类
 */
watch(() => trFilterForm.value.transactionTypes, async () => {
      if (!trFilterForm.value.transactionTypes) return;
      const types: string[] = trFilterForm.value.transactionTypes;
      const promises = types.map(type => getCategoryByType(type));
      const results = await Promise.all(promises);
      const categoryList: Category[] = results.flat();
      categories.value = categoryList.map(category => {
        return {value: category.name};
      });
    }
);
/**
 * 分类变化时要重新刷新标签列表
 */
watch(() => tempCategory.value, async () => {
      if (!tempCategory.value) return;
      const tagList = await getTagsByCategory(tempCategory.value);
      tags.value = tagList.map(tag => {
        return {value: tag.name};
      });
    }
);
/**
 * 监听 conditions 变化，同步到过滤条件中
 */
watch(conditions, () => {
      const newRecord: Record<string, string[]> = {};
      conditions.value.forEach(cond => {
        newRecord[cond.category] = [...cond.tags];
      })
      trFilterForm.value.categoryTags = newRecord;
    }, {deep: true}
);

// 切换分类时清空标签
function onCategoryChange() {
  tempTags.value = [];
}

// 添加条件
function addCondition() {
  if (!tempCategory.value) return;

  const category = tempCategory.value;
  const tags = [...tempTags.value];

  const exists = conditions.value.some(c => c.category === category);
  if (exists) {
    const idx = conditions.value.findIndex(c => c.category === category);
    if (conditions.value[idx]) {
      conditions.value[idx].tags = tags;
    }
  } else {
    conditions.value.push({category, tags});
  }

  // 清空输入
  tempCategory.value = undefined;
  tempTags.value = [];
}
</script>

<style scoped>
.menu-bar {
  display: flex;
  flex-direction: column;
  width: 100%;
  height: 100%;
  align-items: center;
  justify-content: center;
  padding: 8px 0;
}

.top-groups {
  margin-bottom: auto;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.bottom-groups {
  display: flex;
  flex-direction: column;
  gap: 8px;
}
</style>