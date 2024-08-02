import { createApp } from 'vue'

import App from './App.vue'
import './style.css';
import router from './router/index';
import store from './store';
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

// import VueI18n from 'vue-i18n'
import i18n from '@/lang/index' // 多语言引入



const app = createApp(App)
app.use(ElementPlus)
app.use(router)
app.use(store)
app.use(i18n)
app.mount('#app')
