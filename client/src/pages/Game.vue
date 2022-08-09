<script setup lang="ts">
import { onBeforeMount, ref } from 'vue'
import { OpenLogoutMessage } from '../api/messageApi';
import http from '../http/http.js';
import router from '../router.js';
import Functions from '../components/Functions.vue';

const gameContentIdCache = localStorage.getItem("gameContentId")
let gameContentId = ref(gameContentIdCache || "Map")

onBeforeMount(async () => {
    const res = await http.Get("/hello/hello")
    console.log("onBeforeMount = ", res)
})

async function logout() {
    await http.Post("/user/logout")
    OpenLogoutMessage()
    localStorage.removeItem("token")
    router.replace("/login")
}

function gameContentClick(id: string) {
    localStorage.setItem("gameContentId", id)
    gameContentId.value = id
}
</script>

<template>
    <component :is="gameContentId" @game-content-click="gameContentClick"></component>
    <Functions @game-content-click="gameContentClick" />
    <p>
        <a href="javascript:void(0)" @click="logout">退出游戏</a>
    </p>
</template>

<style scoped>
.read-the-docs {
    color: #888;
}
</style>
