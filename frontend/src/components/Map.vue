<template>
    <p>
        「{{ mapName.Name }}」
        <br>
        「附近地图」
        <br>
        上：<a href="javascript:void(0)" @click="moveToMap(Number(topMapName.Id))">{{ topMapName.Name }}</a>
        <br>
        左：<a href="javascript:void(0)" @click="moveToMap(Number(leftMapName.Id))">{{ leftMapName.Name }}</a>
        <br>
        右：<a href="javascript:void(0)" @click="moveToMap(Number(rightMapName.Id))">{{ rightMapName.Name }}</a>
        <br>
        下：<a href="javascript:void(0)" @click="moveToMap(Number(buttonMapName.Id))">{{ buttonMapName.Name }}</a>
    </p>
</template>

<script setup lang="ts">
import get from "lodash/get"
import { onBeforeMount, ref } from 'vue';
// import { get } from "lodash";
import dictMap from '../assets/config/Map.json';

interface mapNode {
    Id: number | undefined,
    Name: string | undefined,
}

let mapIdCache = localStorage.getItem("mapId") || "105";
let mapId = ref(mapIdCache);

let mapName = ref<mapNode>({ Id: 0, Name: "无" })
let topMapName = ref<mapNode>({ Id: 0, Name: "无" })
let leftMapName = ref<mapNode>({ Id: 0, Name: "无" })
let rightMapName = ref<mapNode>({ Id: 0, Name: "无" })
let buttonMapName = ref<mapNode>({ Id: 0, Name: "无" })

onBeforeMount(() => {
    moveToMap(parseInt(mapId.value))
})

function moveToMap(toMapId: number) {
    let curMap = get(dictMap, toMapId)
    mapName.value = { Id: curMap?.Id, Name: curMap?.Name }
    let topMap =  get(dictMap, curMap?.Top)
    topMapName.value = { Id: topMap?.Id, Name: topMap?.Name }
    let leftMap = get(dictMap, curMap?.Left)
    leftMapName.value = { Id: leftMap?.Id, Name: leftMap?.Name }
    let rightMap = get(dictMap, curMap?.Right)
    rightMapName.value = { Id: rightMap?.Id, Name: rightMap?.Name }
    let bottomMap = get(dictMap, curMap?.Bottom)
    buttonMapName.value = { Id: bottomMap?.Id, Name: bottomMap?.Name }

    localStorage.setItem("mapId", String(toMapId))
    // 设置地图视野
    // 获取地图上的NPC、野怪、玩家
}

</script>