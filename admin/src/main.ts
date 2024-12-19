import "@/assets/base.css"

import { createApp } from 'vue'
import pinia  from "@/stores";

import App from './App.vue'
import router from './router'

// 引入 arco-design
import ArcoVue from '@arco-design/web-vue';
import '@arco-design/web-vue/dist/arco.css';
// 额外引入图标库
import ArcoVueIcon from '@arco-design/web-vue/es/icon';
import "@/assets/public.less"


const app = createApp(App)

app.use(pinia)
app.use(router)

app.use(ArcoVue)
app.use(ArcoVueIcon)

app.mount('#app')
