import { createRouter, createMemoryHistory } from 'vue-router';
import Layout from '@/components/Layout.vue';
import TransactionRecordView from '@/views/TransactionRecordView.vue';
import DashboardView from '@/views/DashboardView.vue';

const routes = [
  {
    path: '/',
    component: Layout,
    children: [
      { path: '', redirect: '/tr_view' },
      { name: '消费记录', path: 'tr_view', component: TransactionRecordView },
      { name: '数据分析', path: 'dashboard', component: DashboardView },
    ]
  }
];

const router = createRouter({
  history: createMemoryHistory(),
  routes,
});

export default router;