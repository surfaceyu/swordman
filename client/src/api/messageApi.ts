import { h } from 'vue'
import { ElMessage } from 'element-plus'

const OpenLoginSuccessMessage = () => {
    ElMessage.success("恭喜你！登录成功")
}

const OpenLoginErrorMessage = () => {
    ElMessage.error("账号或密码错误")
}

const OpenLogoutMessage = () => {
    ElMessage.success("恭喜你！登出成功")
}

const OpenRegisterMessage = () => {
    ElMessage.success("恭喜你！注册成功")
}

export {
    OpenLoginSuccessMessage,
    OpenLoginErrorMessage,
    OpenLogoutMessage,
    OpenRegisterMessage,
}