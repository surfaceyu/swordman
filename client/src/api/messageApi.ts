import { h } from 'vue'
import { ElMessage } from 'element-plus'

const OpenLoginSuccessMessage = () => {
    ElMessage({
        message: h('p', null, [
            h('span', null, '恭喜你! '),
            h('i', { style: 'color: teal' }, '登录成功'),
        ]),
    })
}

const OpenLogoutMessage = () => {
    ElMessage({
        message: h('p', null, [
            h('span', null, '恭喜你! '),
            h('i', { style: 'color: red' }, '登出成功'),
        ]),
    })
}

export {
    OpenLoginSuccessMessage,
    OpenLogoutMessage
}