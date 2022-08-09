import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import router from "./router"

import AboutGame from './components/AboutGame.vue'
import Map from './components/Map.vue'
import CreateUser from './components/CreateUser.vue';

const app = createApp(App)
app.use(ElementPlus)
app.use(router)

app.component("CreateUser", CreateUser)
app.component("AboutGame", AboutGame)
app.component("Map", Map)

app.mount("#app")
