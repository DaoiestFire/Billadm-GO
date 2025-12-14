<template>
  <div class="menu-bar">
    <div class="top-groups">
      <a-badge :count="trQueryConditionStore.conditionLen" dot>
        <a-button :type="trQueryConditionStore.conditionLen?'primary':'text'" @click="openFilterDrawer=true">
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
      @close="closeDrawer"
      :body-style="{ paddingBottom: '80px' }"
      :footer-style="{ textAlign: 'right' }"
      :closable="false"
  >
    <a-form layout="vertical">
      <!-- 交易类型 -->
      <a-divider orientation="left">交易类型</a-divider>
      <a-form-item>
        <a-select
            v-model:value="transactionTypes"
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
      <a-form-item v-if="cateTagsConditions.length>0">
        <a-list size="small" bordered :data-source="cateTagsConditions">
          <template #renderItem="{ item }">
            <a-list-item>
              <template #actions>
                <a @click="deleteCondition(item.category)">删除</a>
              </template>
              {{ item.category }} : [{{ item.tags.join(",") }}]
            </a-list-item>
          </template>
        </a-list>
      </a-form-item>
    </a-form>
  </a-drawer>
</template>

<script setup lang="ts">
import {ref, watch} from 'vue';
import {FilterOutlined} from "@ant-design/icons-vue";
import type {Category, categoryTagsCondition} from "@/types/billadm";
import type {DefaultOptionType} from "ant-design-vue/es/vc-cascader";
import {getCategoryByType, getTagsByCategory} from "@/backend/functions.ts";
import {useTrQueryConditionStore} from "@/stores/trQueryConditionStore.ts";

const trQueryConditionStore = useTrQueryConditionStore();

const openFilterDrawer = ref(false);
const transactionTypes = ref<string[]>([]);
const cateTagsConditions = ref<categoryTagsCondition[]>([]);
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
watch(() => transactionTypes.value, async () => {
      if (!transactionTypes.value) return;
      const promises = transactionTypes.value.map(type => getCategoryByType(type));
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
 * 监听打开操作，赋值表单
 */
watch(openFilterDrawer, (newVal) => {
  if (newVal) {
    // 每次打开抽屉时，从 store 加载最新筛选条件
    transactionTypes.value = [...(trQueryConditionStore.transactionTypes || [])];
    cateTagsConditions.value = [...(trQueryConditionStore.cateTagsConditions || [])];
  }
});

function closeDrawer() {
  trQueryConditionStore.transactionTypes = transactionTypes.value;
  trQueryConditionStore.cateTagsConditions = cateTagsConditions.value;
  console.log(trQueryConditionStore.categoryTags);
  openFilterDrawer.value = false;
}

// 切换分类时清空标签
function onCategoryChange() {
  tempTags.value = [];
}

// 添加条件
function addCondition() {
  if (!tempCategory.value) return;

  const category = tempCategory.value;
  const tags = [...tempTags.value];

  const exists = cateTagsConditions.value.some(c => c.category === category);
  if (exists) {
    const idx = cateTagsConditions.value.findIndex(c => c.category === category);
    if (cateTagsConditions.value[idx]) {
      cateTagsConditions.value[idx].tags = tags;
    }
  } else {
    cateTagsConditions.value.push({category, tags});
  }

  // 清空输入
  tempCategory.value = undefined;
  tempTags.value = [];
}

// 删除条件
function deleteCondition(category: string) {
  const idx = cateTagsConditions.value.findIndex(c => c.category === category);
  if (cateTagsConditions.value[idx]) {
    cateTagsConditions.value.splice(idx, 1);
  }
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