<script setup lang="ts">
import { onBeforeMount, ref } from 'vue'
import http from '../http/http.js';
import Functions from '../components/Functions.vue';
import Footer from '../components/Footer.vue';

const gameContentIdCache = localStorage.getItem("gameContentId")
let gameContentId = ref(gameContentIdCache || "Map")

onBeforeMount(async () => {
    const res = await http.Get("/user/hello")
    console.log("onBeforeMount = ", res)
})

function gameContentClick(id: string) {
    localStorage.setItem("gameContentId", id)
    gameContentId.value = id
}
</script>

<template>
    <component :is="gameContentId" @game-content-click="gameContentClick"></component>
    <el-divider />
    <Functions @game-content-click="gameContentClick" />
    <el-divider />
    <Footer @game-content-click="gameContentClick" />
</template>

<style scoped>
.read-the-docs {
    color: #888;
}
</style>
