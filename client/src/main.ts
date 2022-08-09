import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import router from "./router"

import Functions from './components/Functions.vue';
import AboutGame from './components/AboutGame.vue'
import Map from './components/Map.vue'

const app = createApp(App)
app.use(ElementPlus)
app.use(router)

app.component("Functions", Functions)
app.component("AboutGame", AboutGame)
app.component("Map", Map)

app.mount("#app")
