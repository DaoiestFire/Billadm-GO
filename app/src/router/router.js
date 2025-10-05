import {createMemoryHistory, createRouter} from 'vue-router';
import Layout from '@/components/Layout.vue';
import TransactionRecordView from '@/components/tr_view/TransactionRecordView.vue';
import DataAnalysisView from '@/components/da_view/DataAnalysisView.vue';
import SettingsView from "@/components/settings_view/SettingsView.vue";

const routes = [
    {
        path: '/',
        component: Layout,
        children: [
            {path: '', redirect: '/tr_view'},
            {name: '消费记录', path: 'tr_view', component: TransactionRecordView},
            {name: '数据分析', path: 'da_view', component: DataAnalysisView},
            {name: '应用设置', path: 'settings_view', component: SettingsView},
        ]
    }
];

const router = createRouter({
    history: createMemoryHistory(),
    routes,
});

export default router;