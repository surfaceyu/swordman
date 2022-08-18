<template>
    <MessagePopup ref="msgPop"/>
    <view class="uni-flex uni-column" style="height: 40vh;">
    </view>
    <view class="uni-flex uni-row">
        <view class="flex-item"></view>
        <view class="flex-item">
            <uni-forms ref="baseForm" :modelValue="baseFormData" :rules="rules">
                <uni-forms-item label="账号" required>
                    <uni-easyinput v-model="baseFormData.account" placeholder="请输入账号" style="width: 30vw;" />
                </uni-forms-item>
                <uni-forms-item label="密码" required>
                    <uni-easyinput type=password v-model="baseFormData.password" placeholder="请输入密码"
                        style="width: 30vw;" />
                </uni-forms-item>
            </uni-forms>
            <view class="button-group">
                <button size="mini" @click="resetForm">重置</button>
                <button size="mini" @click="submitForm">提交</button>
            </view>
            <view>
                <navigator url="../register/index" hover-class="navigator-hover" open-type="redirect" style="font-size: 14px; margin-top: 10rpx;">
					没有账号？先去注册
				</navigator>
            </view>
        </view>
    </view>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import http from '../../http/accountHttp'
import MessagePopup from '../../components/MessagePopup.vue'

let baseForm = ref()
const msgPop = ref()

const resetForm = () => {
    baseFormData.value = {
        account: '',
        password: '',
    }
}
// 基础表单数据
let baseFormData = ref({
    account: '',
    password: '',
})

const rules = {
    account: {
        rules:
            [{
                required: true,
                format: "string",
                errorMessage: '请输入姓名',
            }, {
                minLength: 1,
                maxLength: 4,
                errorMessage: '姓名长度在 {minLength} 到 {maxLength} 个字符',
            }],
    },
    password: {
        rules: [{
            required: true,
            minLength: 1,
            format: "string",
            errorMessage: '请输入密码',
        }],
    },
}

function submitForm() {
    baseForm.value.validate().then(async () => {
        const res = await http.Post("/auth/login", {
            "Account": baseFormData.value.account,
            "Passwd": baseFormData.value.password
        }, false)
        if (res && res.code == 200) {
            localStorage.setItem('token', res.token);
            uni.redirectTo({ url: "../server/index" })
        }
    })
}
</script>

<style>
.flex-item {
    width: 33.3%;
    height: 200rpx;
    text-align: center;
}

.flex-pc {
    display: flex;
    justify-content: center;
}

.uni-form-item .title {
    padding: 20rpx 0;
}
</style>

<style lang="scss">
.segmented-control {
    margin-bottom: 15px;
}

.button-group {
    margin-top: 15px;
    display: flex;
    justify-content: space-around;
}

.form-item {
    display: flex;
    align-items: center;
}

.button {
    display: flex;
    align-items: center;
    height: 35px;
    margin-left: 10px;
}
</style>