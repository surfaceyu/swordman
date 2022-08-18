<template>
    <p>
        <text @click="refreshPage">刷新</text>·
        <text>消息</text>·
        <text @click="sendCmd('Map')">地图</text>
        <br>
        <text>角色</text>·
        <text>包袋</text>·
        <text>私聊</text>
        <br>
        <text>社交</text>·
        <text>洞府</text>·
        <text>门派</text>
        <br>
    </p>
</template>

<script setup lang="ts">

function sendCmd(name: string) {
    uni.$emit("sendCmd", {
        id: "gameContent",
        data: name
    })
}

function refreshPage() {
    location.reload()
}
</script>