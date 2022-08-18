<template>
    <div class="mt-4">
        <uni-easyinput v-model="chatMsg" :placeholder="props.placeStr" style="width: 80vw;" suffixIcon="paperplane" @iconClick="onChatButtonClick"/>
        <!-- <button size="mini" @click="onChatButtonClick">确认</button> -->
        <!-- <el-input v-model="chatMsg" :placeholder="props.placeStr" clearable style="width: 360px;">
            <template #append>
                <el-button @click="onChatButtonClick()">确认</el-button>
            </template>
        </el-input> -->
    </div>
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