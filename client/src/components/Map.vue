<template>
    <p>
        「{{ mapName }}」
        <br>
        「附近地图」
        <br>
        上：{{ topMapName }}
        <br>
        左：{{ leftMapName }}
        <br>
        右： {{ rightMapName }}
        <br>
        下： {{ buttonMapName }}
    </p>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const emits = defineEmits(["gameContentClick"])

let mapName = ref("自由大道")
let topMapName = ref("朱雀大街")
let leftMapName = ref("自由大道西")
let rightMapName = ref("自由大道东")
let buttonMapName = ref("自由村")

function clickAboutGame() {
    emits("gameContentClick", "AboutGame")
}
</script>