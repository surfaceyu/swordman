<template>
    <el-row type="flex" justify="center" align="middle" style="height: 100%;">
        <el-form ref="ruleFormRef" :model="ruleForm" status-icon :rules="rules" label-width="30%"
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
            <el-form-item label="确认密码" prop="passCheck">
                <el-input v-model="ruleForm.passCheck" type="password" autocomplete="off" />
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="submitForm(ruleFormRef)">注册</el-button>
                <el-button @click="resetForm(ruleFormRef)">重置</el-button>
            </el-form-item>
        </el-form>
    </el-row>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import type { FormInstance } from 'element-plus'
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
        return
    }
    if (ruleForm.passCheck !== "") {
        ruleFormRef.value?.validateField("passCheck", () => null)
    }
    callback()
}

const validatePassCheck = (rule: any, value: string, callback: any) => {
    if (value === '') {
        callback(new Error('请输入密码'))
        return
    }
    if (value !== ruleForm.pass) {
        callback(new Error('密码不一致，请重新输入'))
        return
    }
    callback()
}

const ruleForm = reactive({
    account: "",
    pass: "",
    passCheck: "",
})

const rules = reactive({
    account: [{ validator: validateAccount, trigger: 'blur' }],
    pass: [{ validator: validatePass, trigger: 'blur' }],
    passCheck: [{ validator: validatePassCheck, trigger: 'blur' }],
})

const submitForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.validate(async (valid) => {
        if (valid) {
            const res = await http.Put("/user/login", {
                "ID": ruleForm.account,
                "Passwd": ruleForm.pass
            }, false)
            if (res && res.code == 200) {
                router.replace("/login")
                // OpenLoginSuccessMessage()
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