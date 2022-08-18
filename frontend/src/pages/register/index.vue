<template>
    <view class="uni-flex uni-column" style="height: 30vh;">
    </view>
    <view class="uni-flex uni-row">
        <view class="flex-item"></view>
        <view class="flex-item">
            <uni-forms ref="ruleFormRef" :modelValue="baseFormData" :rules="rules">
                <uni-forms-item label="账号" required name="account">
                    <uni-easyinput v-model="baseFormData.account" placeholder="请输入账号" style="width: 30vw;" />
                </uni-forms-item>
                <uni-forms-item label="密码" required name="password">
                    <uni-easyinput type=password v-model="baseFormData.password" placeholder="请输入密码"
                        style="width: 30vw;" />
                </uni-forms-item>
                <uni-forms-item label="确认密码" required name="passCheck">
                    <uni-easyinput type=password v-model="baseFormData.passCheck" placeholder="请确认密码"
                        style="width: 30vw;" />
                </uni-forms-item>
            </uni-forms>
            <view class="button-group">
                <button size="mini" @click="resetForm">重置</button>
                <button size="mini" @click="submitForm">提交</button>
            </view>
        </view>
    </view>
    <!-- <el-row type="flex" justify="center" align="middle" style="height: 100%;">
        <el-form ref="ruleFormRef" :model="ruleForm" status-icon :rules="rules" label-width="30%" class="demo-ruleForm"
            style="text-align: center; position: absolute;">
            <el-form-item>
                <p>文字武侠</p>
            </el-form-item>
            <el-form-item label="账户" prop="account">
                <el-input v-model.number="ruleForm.account" />
            </el-form-item>
            <el-form-item label="密码" prop="pass">
                <el-input v-model="ruleForm.pass" type="password" show-password autocomplete="off" />
            </el-form-item>
            <el-form-item label="确认密码" prop="passCheck">
                <el-input v-model="ruleForm.passCheck" type="password" show-password autocomplete="off" />
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="submitForm(ruleFormRef)">注册</el-button>
                <el-button @click="resetForm(ruleFormRef)">重置</el-button>
            </el-form-item>
        </el-form>
    </el-row> -->
</template>

<script setup lang="ts">
import { onReady } from "@dcloudio/uni-app";
import { ref } from 'vue'
import http from '../../http/accountHttp'

const ruleFormRef = ref()

let baseFormData = ref({
    account: '',
    password: '',
    passCheck: '',
})

const validateAccount = (rule: any, value: string, data: any, callback: any) => {
    if (!value) {
        return callback(new Error('请输入账号'))
    } else {
        callback()
    }
}

const validatePass = (rule: any, value: string, data: any, callback: any) => {
    if (value === '') {
        callback(new Error('请输入密码'))
        return
    }
    if (baseFormData.value.passCheck !== "") {
        ruleFormRef.value?.validateField(["passCheck"], () => null)
    }
    callback()
}

const validatePassCheck = (rule: any, value: string, data: any, callback: any) => {
    if (value === '') {
        callback(new Error('请输入密码'))
        return
    }
    if (value !== baseFormData.value.password) {
        callback(new Error('密码不一致，请重新输入'))
        return
    }
    callback()
}

const rules = {
    account: {
        rules: [{
            validateFunction: validateAccount
        }]
    },
    password: {
        rules: [{
            validateFunction: validatePass
        }]
    },
    passCheck: {
        rules: [{
            validateFunction: validatePassCheck
        }]
    },
}

onReady(() => {
    // 需要在onReady中设置规则
    ruleFormRef.value.setRules(rules)
})

const submitForm = () => {
    ruleFormRef.value.validate().then(async () => {
        const res = await http.Put("/auth/login", {
            "Account": baseFormData.value.account,
            "Passwd": baseFormData.value.password
        }, false)
        if (res && res.code == 200) {
            uni.navigateTo({ url: "../login/index" })
        }
    })
}

const resetForm = () => {
    baseFormData.value = {
        account: '',
        password: '',
        passCheck: '',
    }
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

.button-group {
    margin-top: 15px;
    display: flex;
    justify-content: space-around;
}
</style>