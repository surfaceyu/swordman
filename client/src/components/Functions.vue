<template>
    <p>
        <a href="javascript:void(0)" @click="refreshPage">刷新</a>·
        <a href="javascript:void(0)">消息</a>·
        <a href="javascript:void(0)" @click="gameContentClick('Map')">地图</a>
        <br>
        <a href="javascript:void(0)">角色</a>·
        <a href="javascript:void(0)">包袋</a>·
        <a href="javascript:void(0)">私聊</a>
        <br>
        <a href="javascript:void(0)">社交</a>·
        <a href="javascript:void(0)">洞府</a>·
        <a href="javascript:void(0)">门派</a>
        <br>
        <a href="javascript:void(0)">设置</a>
        <br>
        <a href="javascript:void(0)" @click="gameContentClick('AboutGame')">游戏介绍</a>
        <br>
    </p>
</template>

<script setup lang="ts">
const emits = defineEmits(["gameContentClick"])

function gameContentClick(name: string) {
    emits("gameContentClick", name)
}

function refreshPage() {
    location.reload()
}
</script>