<template>
    <view>
        <!-- 提示信息弹窗 -->
        <uni-popup ref="messageRef" type="message">
            <uni-popup-message :type="msgType" :message="messageText" :duration="2000"></uni-popup-message>
        </uni-popup>
    </view>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const messageRef = ref();

let msgType = ref('success');
let messageText = ref('这是一条成功提示');


function PopMessage(t: string, msg: string) {
    msgType.value = t
    messageText.value = msg
    messageRef.value.open()
}

defineExpose({
    PopMessage,
});
</script>