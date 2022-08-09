<template>
    <el-row type="flex" justify="center" align="middle" style="height: 100%;">
        <el-form ref="ruleFormRef" :model="ruleForm" status-icon :rules="rules" label-width="20%"
            class="demo-ruleForm" style="text-align: center; position: absolute;">
            <el-form-item>
                <p>文字武侠</p>
            </el-form-item>
            <el-form-item label="账户" prop="account">
                <el-input v-model.number="ruleForm.account" />
            </el-form-item>
            <el-form-item label="密码" prop="pass">
                <el-input v-model="ruleForm.pass" type="password" autocomplete="off" />
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="submitForm(ruleFormRef)">登录</el-button>
                <el-button @click="resetForm(ruleFormRef)">重置</el-button>
            </el-form-item>
            <el-form-item>
                <router-link to="/createUser">没有账号？先去注册</router-link>
            </el-form-item>
        </el-form>
    </el-row>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import type { FormInstance } from 'element-plus'
import { OpenLoginSuccessMessage } from '../api/messageApi';
import router from "../router"
import http from '../http/http'

const ruleFormRef = ref<FormInstance>()

const validateAccount = (rule: any, value: string, callback: any) => {
    if (!value) {
        return callback(new Error('请输入账号'))
    } else {
        callback()
    }
}

const validatePass = (rule: any, value: string, callback: any) => {
    if (value === '') {
        callback(new Error('请输入密码'))
    } else {
        if (!ruleFormRef.value) return
        callback()
    }
}

const ruleForm = reactive({
    account: "",
    pass: '',
})

const rules = reactive({
    account: [{ validator: validateAccount, trigger: 'blur' }],
    pass: [{ validator: validatePass, trigger: 'blur' }],
})

const submitForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.validate(async (valid) => {
        if (valid) {
            const res = await http.Post("/user/login", {
                "ID": ruleForm.account,
                "Passwd": ruleForm.pass
            }, false)
            if (res && res.code == 200) {
                localStorage.setItem('token', res.token);
                router.replace("/game")
                OpenLoginSuccessMessage()
            }
        } else {
            console.log('error submit!')
            return false
        }
    })
}

const resetForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.resetFields()
}
</script>