<template>
    <p>
        「传音」
        <br>
    <p style="margin: 0px;" v-for="msg in chatMsgs">
        <a href="javascript:void(0)">{{ msg.UserName }}</a>:{{ msg.Message }}
    </p>
    <a href="javascript:void(0)" @click="onMoreChatCliCk()">查看更多</a>
    </p>
    <p>
        「江湖传闻」
        <br>
    <p style="margin: 0px;" v-for="msg in chatMsgs">
        <a href="javascript:void(0)">{{ msg.UserName }}</a>:{{ msg.Message }}
    </p>
    <a href="javascript:void(0)">查看更多</a>
    </p>
</template>

<script setup lang="ts">
import { onBeforeMount, ref } from 'vue';
import http from '../http/http';

interface chatMessage {
    Uid: number,
    UserName: string,
    Channel: string,
    Message: string,
    CreateAt: string,
}

const chatMsgs = ref<chatMessage[]>()

onBeforeMount(async () => {
    await onRefreshPage()
})

async function onRefreshPage() {
    const msgs = await http.Get("/chat/chat", { "Channel": "w", "Limit": 2 })
    chatMsgs.value = msgs.data
}

const emits = defineEmits(["gameContentClick"])

function gameContentClick(name: string) {
    emits("gameContentClick", name)
}

function onMoreChatCliCk() {
    gameContentClick("ContentChat")
}

defineExpose({onRefreshPage})
</script>