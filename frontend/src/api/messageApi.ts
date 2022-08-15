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

function MessageSuccess(msg:string) {
    ElMessage.success(msg)
}
function MessageError(msg:string) {
    ElMessage.error(msg)
}
function MessageWarn(msg:string) {
    ElMessage.warning(msg)
}
function MessageInfo(msg:string) {
    ElMessage.info(msg)
}

export {
    OpenLoginSuccessMessage,
    OpenLoginErrorMessage,
    OpenLogoutMessage,
    OpenRegisterMessage,
    MessageSuccess,
    MessageError,
    MessageWarn,
    MessageInfo,
}