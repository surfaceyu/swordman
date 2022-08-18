<template>
    <p>
        <a href="javascript:void(0)">设置</a>·
        <a href="javascript:void(0)" @click="sendCmd('AboutGame')">游戏介绍</a>·
        <a href="javascript:void(0)" @click="logout">退出游戏</a>
    </p>
</template>

<script setup lang="ts">
import http from '../http/http';
import accountHttp from '../http/accountHttp';

async function logout() {
    await accountHttp.Post("/auth/logout")
    localStorage.removeItem("token")
    uni.redirectTo({url: "../login/index"})
}

function sendCmd(name: string) {
	uni.$emit('sendCmd', {
        id: "gameContent",
        data: name
    })
}

</script>