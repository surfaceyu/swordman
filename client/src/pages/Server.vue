<template>
    <!-- <el-row type="flex" justify="center" align="middle" style="height: 100%;">
        <el-col :span="12" :offset="6" v-for="server in serversRef">
            <el-button>{{ server.Name }}</el-button>
        </el-col>
    </el-row> -->
    <el-row type="flex" justify="center" align="middle" style="height: 100%;">
        <el-form style="text-align: center; position: absolute;">
            <el-row :gutter="10">
                <el-col :span="12" :offset="6">
                    <el-button disabled>服务器列表</el-button>
                </el-col>
            </el-row>
            <el-row :gutter="10" v-for="server in serversRef">
                <el-col :span="12" :offset="6">
                    <el-button @click="onServerSelect(server)">{{ server.Name }}</el-button>
                </el-col>
            </el-row>
        </el-form>
    </el-row>
</template>

<script setup lang="ts">
import { onBeforeMount, ref } from 'vue';
import httpAccount from '../http/accountHttp'
import http from '../http/http'
import router from "../router"

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
    http.SetRequest(`http://${server.Host}:${server.Port}/api`)
    router.replace("/game")
    console.log("==============", server)
}
</script>

<style>
.el-row {
    margin-bottom: 20px;
}

.el-row:last-child {
    margin-bottom: 0;
}

.el-col {
    border-radius: 4px;
}

.grid-content {
    border-radius: 4px;
    min-height: 36px;
}
</style>