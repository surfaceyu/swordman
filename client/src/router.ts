import {createRouter, createWebHashHistory} from 'vue-router'
import Game from './pages/Game.vue'
import Login from './pages/Login.vue'
import CreateUser from './pages/CreateUser.vue'

const router = createRouter({
    history: createWebHashHistory(),
    routes: [
        {path: "/", component: Game},
        {path: "/login", component: Login},
        {path: "/createUser", component: CreateUser},
        {path: "/game", component: Game},
    ]
})

router.beforeEach((to, from) => {
    let token = localStorage.getItem('token');
    if (!token && to.path !== "/login" && to.path !== "/createUser") {        
        return { path: "/login" }
    }
})

export default router