<template>
    <div class="mt-4">
        <el-input v-model="chatMsg" :placeholder="props.placeStr" clearable>
            <template #append>
                <el-button @click="onChatButtonClick()">确认</el-button>
            </template>
        </el-input>
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