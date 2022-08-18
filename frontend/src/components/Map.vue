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
    <p>
        「附近有」
        <br>
        <p style="margin: 0px;" v-for="it in mapName.Npc">
            npc: <a href="javascript:void(0)">{{it.Name}}</a>
        </p>
    </p>
</template>

<script setup lang="ts">
import get from "lodash/get"
import forEach from "lodash/forEach"
import { onBeforeMount, ref } from 'vue';
import dictMap from '../assets/config/Map.json'
import dictNpc from '../assets/config/Npc.json';
interface mapNode {
    Id: number | undefined,
    Name: string | undefined,
    Npc: npcNode[],
}

interface npcNode {
    Id: number,
    Name: string,
    Avatar: string,
}

let mapIdCache = localStorage.getItem("mapId") || "105";
let mapId = ref(mapIdCache);

let mapName = ref<mapNode>({ Id: 0, Name: "无", Npc: [] })
let topMapName = ref<mapNode>({ Id: 0, Name: "无", Npc: [] })
let leftMapName = ref<mapNode>({ Id: 0, Name: "无", Npc: [] })
let rightMapName = ref<mapNode>({ Id: 0, Name: "无", Npc: [] })
let buttonMapName = ref<mapNode>({ Id: 0, Name: "无", Npc: [] })

onBeforeMount(() => {
    moveToMap(parseInt(mapId.value))
})

function moveToMap(toMapId: number) {
    let curMap = get(dictMap, toMapId)
    mapName.value = setMapConfig(curMap?.Id)

    topMapName.value = setMapConfig(curMap?.Top)
    leftMapName.value = setMapConfig(curMap?.Left)
    rightMapName.value = setMapConfig(curMap?.Right)
    buttonMapName.value = setMapConfig(curMap?.Bottom)

    localStorage.setItem("mapId", String(toMapId))
    // 设置地图视野
    // 获取地图上的NPC、野怪、玩家

}

function setMapConfig(mapId: number): mapNode {
    const m = get(dictMap, mapId)
    let res: mapNode = {Id: m?.Id, Name: m?.Name, Npc: []}
    forEach(m?.Npc, (npcId) => {
        const n = get(dictNpc, npcId)
        if (!n) return
        res.Npc.push({Id: n.Id, Name: n.Name, Avatar: n.Avatar})
    })
    return res
}
</script>