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
    <Input place-str="请输入你的名字" @on-chat-button-click="onChatButtonClick"/>
    <a @click="gameContentClick('Map')"></a>
</template>

<script setup lang="ts">
import { MessageInfo } from '../api/messageApi';
import http from '../http/http';
import Input from './Input.vue';

async function onChatButtonClick(name: string) {
    const res = await http.Put("/user/role", {"Name": name})
    if (res && res.code == 200) {
        gameContentClick("Map")
        MessageInfo("欢迎你！"+name)
    }
}

const emits = defineEmits(["gameContentClick"])

function gameContentClick(name: string) {
    emits("gameContentClick", name)
}
</script>