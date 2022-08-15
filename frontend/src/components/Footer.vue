<template>
    <p>
        <a href="javascript:void(0)">设置</a>·
        <a href="javascript:void(0)" @click="gameContentClick('AboutGame')">游戏介绍</a>·
        <a href="javascript:void(0)" @click="logout">退出游戏</a>
    </p>
</template>

<script setup lang="ts">
import { OpenLogoutMessage } from '../api/messageApi';
import http from '../http/http.js';
import router from '../router.js';

async function logout() {
    await http.Post("/auth/logout")
    OpenLogoutMessage()
    localStorage.removeItem("token")
    router.replace("/login")
}

const emits = defineEmits(["gameContentClick"])

function gameContentClick(name: string) {
    emits("gameContentClick", name)
}

</script>