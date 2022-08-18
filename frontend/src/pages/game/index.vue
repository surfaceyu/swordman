<template>
    <template v-if="gameContentId == 'AboutGame'">
        <AboutGame />
    </template>
    <template v-if="gameContentId == 'Map'">
        <Map />
    </template>
        <template v-if="gameContentId == 'ContentChat'">
        <ContentChat />
    </template>
    <Functions childDisabled />
    <FooterChat ref="footerChatRef" />
    <Footer />
</template>

<script setup lang="ts">
import { onLoad, onUnload } from "@dcloudio/uni-app";
import { onBeforeMount, ref } from 'vue'
import http from '../../http/http.js';
import Functions from '../../components/Functions.vue';
import Footer from '../../components/Footer.vue';
import FooterChat from '../../components/FooterChat.vue';
import AboutGame from '../../components/AboutGame.vue';
import Map from "../../components/Map.vue";
import ContentChat from "../../components/ContentChat.vue";

const footerChatRef = ref()
const gameContentIdCache = uni.getStorageSync("gameContentId") || "Map" 
let gameContentId = ref(gameContentIdCache)
let childDisabled = ref(true)

onBeforeMount(async () => {
    const res = await http.Get("/user/role")
    if (!(res && res.data && res.data.role.ID > 0)) {
        gameContentId.value = "CreateUser"
        childDisabled.value = false
    }
})

function gameContentClick(id: string) {
    uni.setStorage({
        key: "gameContentId",
        data: id,
    })
    gameContentId.value = id
}

function sendCmd(cmd: any) {
    console.log("sendCmd ==", cmd)
    switch (cmd.id) {
        case "gameContent":
            gameContentClick(cmd.data)
            break;
        default:
            footerChatRef.value.onRefreshPage()
            break;
    }
}

onLoad(() => {
    uni.$on('sendCmd', sendCmd);
});

onUnload(() => {
    uni.$off('sendCmd');
});

</script>

<style scoped>
.read-the-docs {
    color: #888;
}
</style>
