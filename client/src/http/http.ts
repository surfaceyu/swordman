import axios from "axios"
import { stringify } from 'qs'
import config from '../config.json'
import router from "../router"

function url(u : string): string {
    return `${config.urlBase}${u}`
}

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
        const res = await axios.get(url(u), {params: data, ...config})
        return res ? res.data : null
    } catch (error: any) {
        if (isAuthorization && error.request.status == 401) {
            router.push("/login")
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
        const res = await axios.post(url(u), stringify(data), config)
        return res ? res.data : null
    } catch (error: any) {
        if (isAuthorization && error.request.status == 401) {
            router.push("/login")
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
        const res = await axios.put(url(u), stringify(data), config)
        return res ? res.data : null
    } catch (error: any) {
        if (isAuthorization && error.request.status == 401) {
            router.push("/login")
        }
    }
}

async function Delete(u: string, data = {}, isAuthorization = true) : Promise<any> {
    const config = isAuthorization ? {
        headers: {
            "Authorization": `Bearer ${localStorage.getItem('token')}`
        }
    } : {}
    try {
        const res = await axios.delete(url(u), config)
        return res ? res.data : null
    } catch (error: any) {
        if (isAuthorization && error.request.status == 401) {
            router.push("/login")
        }
    }
}

export default {
    Get,
    Post,
    Put,
    Delete,
}