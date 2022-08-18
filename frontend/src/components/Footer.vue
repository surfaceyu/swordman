<template>
    <p>
        <text>设置</text>·
        <text @click="sendCmd('AboutGame')">游戏介绍</text>·
        <text @click="logout">退出游戏</text>
    </p>
</template>

<script setup lang="ts">
import http from '../http/http';
import accountHttp from '../http/accountHttp';

async function logout() {
    await accountHttp.Post("/auth/logout")
    uni.removeStorage({
        key:"token"
    })
    uni.redirectTo({url: "../login/index"})
}

function sendCmd(name: string) {
	uni.$emit('sendCmd', {
        id: "gameContent",
        data: name
    })
}

</script>