<script setup lang="ts">
import { onBeforeMount, ref } from 'vue'
import http from '../http/http.js';
import Functions from '../components/Functions.vue';
import Footer from '../components/Footer.vue';
import FooterChat from '../components/FooterChat.vue';

const footerChatRef = ref<any>()
const gameContentIdCache = localStorage.getItem("gameContentId")
let gameContentId = ref(gameContentIdCache || "Map")
let childDisabled = ref(true)

onBeforeMount(async () => {
    const res = await http.Get("/user/role")
    if (!(res && res.data && res.data.role.ID > 0)) {
        gameContentId.value = "CreateUser"
        childDisabled.value = false
    }
})

function gameContentClick(id: string) {
    localStorage.setItem("gameContentId", id)
    gameContentId.value = id
}

function sendCmd(cmd: string, data: string) {
    console.log("sendCmd ==", cmd, data)
    footerChatRef.value.onRefreshPage()
}
</script>

<template>
    <component :is="gameContentId" @game-content-click="gameContentClick" @send-cmd="sendCmd"></component>
    <el-divider />
    <Functions @game-content-click="gameContentClick" childDisabled/>
    <el-divider />
    <Footer @game-content-click="gameContentClick" />
    <FooterChat ref="footerChatRef" @game-content-click="gameContentClick"/>
</template>

<style scoped>
.read-the-docs {
    color: #888;
}
</style>
