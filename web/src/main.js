import './assets/main.css'
import './assets/reset.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

// 引入element-ui
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(ElementPlus)
app.use(createPinia())
app.use(router)

app.mount('#app')
