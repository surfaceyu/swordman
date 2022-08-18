<template>
    <p>
        「传音」
        <br>
    <p style="margin: 0px;" v-for="msg in chatMsgs">
        <a href="javascript:void(0)">{{ msg.UserName }}</a>:{{ msg.Message }}
    </p>
    <Input place-str="请输入聊天内容" @on-chat-button-click="onChatButtonClick" />
    </p>
</template>

<script setup lang="ts">
import { onBeforeMount, ref } from 'vue';
import http from '../http/http';
import Input from './Input.vue';

interface chatMessage {
    Uid: number,
    UserName: string,
    Channel: string,
    Message: string,
    CreateAt: string,
}

const chatMsgs = ref<chatMessage[]>()

onBeforeMount(async () => {
    await refreshPage()
})

async function refreshPage() {
    const msgs = await http.Get("/chat/chat", { "Channel": "w", "Limit": 10 })
    chatMsgs.value = msgs.data
}

async function onChatButtonClick(data: string) {
    const res = await http.Post("/chat/chat", { "Channel": "w", "Data": data })
    if (res && res.code == 200) {
        refreshPage()
        uni.$emit("sendCmd", {
            id: "chatRefresh",
            data: 1
        })
    }
}


</script>