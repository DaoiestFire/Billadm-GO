import { createRouter, createMemoryHistory } from 'vue-router';
import Layout from '@/components/Layout.vue';
import HomeView from '@/views/HomeView.vue';
import DashboardView from '@/views/DashboardView.vue';

const routes = [
  {
    path: '/',
    component: Layout,
    children: [
      { path: '', redirect: '/home' },
      { path: 'home', component: HomeView },
      { path: 'dashboard', component: DashboardView },
    ]
  }
];

const router = createRouter({
  history: createMemoryHistory(),
  routes,
});

export default router;