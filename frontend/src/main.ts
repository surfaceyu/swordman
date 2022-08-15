import { createApp } from 'vue'
import App from './App.vue'
import router from "./router"

import AboutGame from './components/AboutGame.vue'
import Map from './components/Map.vue'
import CreateUser from './components/CreateUser.vue';
import ContentChat from './components/ContentChat.vue';

const app = createApp(App)
app.use(router)

app.component("CreateUser", CreateUser)
app.component("AboutGame", AboutGame)
app.component("ContentChat", ContentChat)
app.component("Map", Map)

app.mount("#app")
