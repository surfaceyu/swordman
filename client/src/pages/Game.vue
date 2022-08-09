<script setup lang="ts">
import { onBeforeMount, ref } from 'vue'
import http from '../http/http.js';
import Functions from '../components/Functions.vue';
import Footer from '../components/Footer.vue';

const gameContentIdCache = localStorage.getItem("gameContentId")
let gameContentId = ref(gameContentIdCache || "Map")
let childDisabled = ref(true)

onBeforeMount(async () => {
    const res = await http.Get("/user/role")
    console.log("/user/role = ", res)
    if (!(res && res.data && res.data.role.ID > 0)) {
        gameContentId.value = "CreateUser"
        childDisabled.value = false
    }
})

function gameContentClick(id: string) {
    localStorage.setItem("gameContentId", id)
    gameContentId.value = id
}
</script>

<template>
    <component :is="gameContentId" @game-content-click="gameContentClick"></component>
    <el-divider />
    <Functions @game-content-click="gameContentClick" childDisabled/>
    <el-divider />
    <Footer @game-content-click="gameContentClick" />
</template>

<style scoped>
.read-the-docs {
    color: #888;
}
</style>
