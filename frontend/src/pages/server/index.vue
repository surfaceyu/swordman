<template>
    <view class="uni-flex uni-column">
        <view class="flex-item flex-item-V" style="height: 40vh;"></view>
        <view class="flex-item flex-item-V">
            <view class="uni-flex uni-row">
                <view class="flex-item" style="width: 40vw;"></view>
                <view class="flex-item">
                    <p>服务器列表</p>
                    <button v-for="server in serversRef" @click="onServerSelect(server)">{{ server.Name }}</button>
                </view>
                <view class="flex-item" style="width: 40vw;"></view>
            </view>
        </view>
    </view>
</template>

<script setup lang="ts">
import { onBeforeMount, ref } from 'vue';
import httpAccount from '../../http/accountHttp'

interface iServer {
    ID: number,
    Name: string,
    Host: string,
    Port: string,
}

let serversRef = ref<iServer[]>()

onBeforeMount(async () => {
    const res = await httpAccount.Get("/auth/server")
    if (res && res.data) {
        serversRef.value = res.data
    }
})

function onServerSelect(server: iServer) {
    uni.setStorage({
        key: "gameUri",
        data: `http://${server.Host}:${server.Port}/api`,
    })
    uni.redirectTo({
        url: "../game/index", success: () => {
            location.reload()
        }
    })
}
</script>

<style>
.flex-item {
    width: 100%;
    height: 200rpx;
    text-align: center;
}
</style>