import { createRouter, createMemoryHistory } from 'vue-router';
import Layout from '@/components/Layout.vue';
import TransactionRecordView from '@/components/tr_view/TransactionRecordView.vue';
import DataAnalysisView from '@/components/da_view/DataAnalysisView.vue';

const routes = [
  {
    path: '/',
    component: Layout,
    children: [
      { path: '', redirect: '/tr_view' },
      { name: '消费记录', path: 'tr_view', component: TransactionRecordView },
      { name: '数据分析', path: 'da_view', component: DataAnalysisView },
    ]
  }
];

const router = createRouter({
  history: createMemoryHistory(),
  routes,
});

export default router;