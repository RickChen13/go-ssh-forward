import { createApp } from 'vue';
import App from '@/app/App.vue';
const app = createApp(App);

import router from '@/app/router';
app.use(router);

import { createPinia } from 'pinia';
app.use(createPinia());

import ElementPlus from 'element-plus';
import zhCn from 'element-plus/es/locale/lang/zh-cn';
import 'element-plus/dist/index.css';
app.use(ElementPlus, {
    locale: zhCn
});
import * as ElementPlusIconsVue from '@element-plus/icons-vue';
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component);
}

import { EventBus } from '@/app/plugins/mitt/EventBus';
window.wailsApi.log = (data: string) => {
    EventBus.emit('log', data);
};

app.mount('#app');

