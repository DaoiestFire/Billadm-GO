import {createApp} from 'vue';
import {createPinia} from "pinia";
import router from '@/router/router';
import App from '@/App.vue';
import VueECharts from 'vue-echarts';
import 'normalize.css';
import '@/style.css';

const pinia = createPinia();
const app = createApp(App);

app.use(pinia);
app.use(router);
app.component('v-chart', VueECharts);
app.mount('#app');
