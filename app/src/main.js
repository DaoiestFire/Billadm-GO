import {createApp} from 'vue';
import {createPinia} from "pinia";
import router from '@/router/router';
import App from '@/App.vue';
import VueECharts from 'vue-echarts';
import * as echarts from 'echarts/core';
import {CanvasRenderer} from 'echarts/renderers';
import {GridComponent, LegendComponent, TooltipComponent} from "echarts/components";
import {BarChart, LineChart, PieChart} from "echarts/charts";
import 'normalize.css';
import '@/style.css';

const pinia = createPinia();
const app = createApp(App);

app.use(pinia);
app.use(router);
echarts.use([CanvasRenderer, TooltipComponent, GridComponent, LegendComponent, LineChart, PieChart, BarChart]);
app.component('v-chart', VueECharts);
app.mount('#app');
