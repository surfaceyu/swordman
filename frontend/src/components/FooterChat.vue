<template>
    <p>
        「传音」
        <br>
    <p style="margin: 0px;" v-for="msg in chatMsgs">
        <text>{{ msg.UserName }}</text>:{{ msg.Message }}
    </p>
    <text @click="onMoreChatCliCk()">查看更多</text>
    </p>
    <p>
        「江湖传闻」
        <br>
    <p style="margin: 0px;" v-for="msg in chatMsgs">
        <text>{{ msg.UserName }}</text>:{{ msg.Message }}
    </p>
    <text>查看更多</text>
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

function onMoreChatCliCk() {
    uni.$emit("sendCmd", {
        id: "gameContent",
        data: "ContentChat"
    })
}

defineExpose({onRefreshPage})
</script>