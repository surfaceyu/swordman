<template>
    <view style="height: 50px;">
        <uni-easyinput v-model="chatMsg" :placeholder="props.placeStr" style="width: 80vw;" suffixIcon="paperplane" @iconClick="onChatButtonClick"/>
    </view>
</template>

<script setup lang="ts">
import { nextTick, ref } from 'vue';

let props = defineProps<{
    placeStr: string
}>()

const chatMsg = ref('')

const emits = defineEmits(["onChatButtonClick"])

function onChatButtonClick() {
    emits("onChatButtonClick", chatMsg.value)
    nextTick(() => {
        chatMsg.value = ""
    })
}
</script>