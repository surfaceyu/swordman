import axios from "axios"
import { stringify } from 'qs'

const baseURL = localStorage.getItem("gameUri")
let request = axios.create({
    baseURL: baseURL ? baseURL : "",
    timeout: 5000,
    headers: {
        'content-type': 'application/x-www-form-urlencoded',
        "Access-Control-Allow-Origin": "*"
    },
})

/**
 * 发送Get请求
 * @param u url字符串
 * @param data 参数
 * @param isAuthorization 是否需要验证 验证不通过会跳转到login
 * @returns 
 */
async function Get(u: string, data = {}, isAuthorization = true): Promise<any>{
    const config = isAuthorization ? {
        headers: {
            "Authorization": `Bearer ${localStorage.getItem('token')}`
        }
    } : {}
    try {
        const res = await request.get(u, {params: data, ...config})
        console.log("get url [%s] res is", u, res.data)
        return res ? res.data : null
    } catch (error: any) {
        console.log("Get error", error)
        if (isAuthorization && error.request.status == 401) {
            uni.redirectTo({url: "../login/index"})
        }
    }
}

/**
 * 发送Post请求
 * @param u url字符串
 * @param data 请求数据
 * @param isAuthorization 是否需要验证 验证不通过会跳转到login
 * @return Promise<any>
 */
async function Post(u: string, data = {}, isAuthorization = true) : Promise<any> {
    const config = isAuthorization ? {
        headers: {
            "Authorization": `Bearer ${localStorage.getItem('token')}`
        }
    } : {}
    try {
        const res = await request.post(u, stringify(data), config)
        console.log("post url [%s] res is", u, res.data)
        return res ? res.data : null
    } catch (error: any) {
        if (isAuthorization && error.request.status == 401) {
            uni.redirectTo({url: "../login/index"})
        }
        return error.request
    }
}

async function Put(u: string, data = {}, isAuthorization = true) : Promise<any> {
    const config = isAuthorization ? {
        headers: {
            "Authorization": `Bearer ${localStorage.getItem('token')}`
        }
    } : {}
    try {
        const res = await request.put(u, stringify(data), config)
        console.log("put url [%s] res is", u, res.data)
        return res ? res.data : null
    } catch (error: any) {
        if (isAuthorization && error.request.status == 401) {
            uni.redirectTo({url: "../login/index"})
        }
        return error.request
    }
}

async function Delete(u: string, data = {}, isAuthorization = true) : Promise<any> {
    const config = isAuthorization ? {
        headers: {
            "Authorization": `Bearer ${localStorage.getItem('token')}`
        }
    } : {}
    try {
        const res = await request.delete(u, config)
        console.log("delete url [%s] res is", u, res.data)
        return res ? res.data : null
    } catch (error: any) {
        if (isAuthorization && error.request.status == 401) {
            uni.redirectTo({url: "../login/index"})
        }
        return error.request
    }
}

export default {
    Get,
    Post,
    Put,
    Delete,
}