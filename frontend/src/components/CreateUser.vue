<template>
    <p>
        <br>
        <br>
        <br>
        <br>
        <br>
        <br>
        你感到头痛欲裂，睁开眼发现你在一片竹林里，身上只有一件破布麻衣。
        <br>
        你记得你叫
        <br>
    </p>
    <Input place-str="请输入你的名字" @on-chat-button-click="onChatButtonClick" />
</template>

<script setup lang="ts">
import { MessageInfo, MessageError } from '../api/messageApi';
import http from '../http/http';
import Input from './Input.vue';

async function onChatButtonClick(name: string) {
    const res = await http.Put("/user/role", { "Name": name })
    if (res) {
        if (res.code == 200) {
                uni.$emit("sendCmd", {id: "gameContent", data: "Map"})

            MessageInfo("欢迎你！" + name)
        } else {
            MessageError(res.message)
        }
    }
}

</script>