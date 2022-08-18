import get from "lodash/get"

let baseURL = uni.getStorageSync("gameUri")

function RefreshUrl() {
    baseURL = uni.getStorageSync("gameUri")    
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
            "Authorization": `Bearer ${uni.getStorageSync("token")}`
        }
    } : {}
    try {
        const res = await uni.request({
            url: baseURL + u,
            method: "GET",
            data: data,
            timeout: 5000,
            header: {
                'content-type': 'application/x-www-form-urlencoded',
                "Access-Control-Allow-Origin": "*",
                ...config.headers
            }
        })
        console.log("get url [%s] res is", u, get(res, "data"))
        return get(res, "data")
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
            "Authorization": `Bearer ${uni.getStorageSync("token")}`
        }
    } : {}
    try {
        const res = await uni.request({
            url: baseURL + u,
            method: "POST",
            data: data,
            timeout: 5000,
            header: {
                'content-type': 'application/x-www-form-urlencoded',
                "Access-Control-Allow-Origin": "*",
                ...config.headers
            }
        })
        console.log("post url [%s] res is", u, get(res, "data"))
        return get(res, "data")
    } catch (error: any) {
        if (isAuthorization) {
            uni.redirectTo({url: "../login/index"})
        }
        return error.request
    }
}

async function Put(u: string, data = {}, isAuthorization = true) : Promise<any> {
    const config = isAuthorization ? {
        headers: {
            "Authorization": `Bearer ${uni.getStorageSync("token")}`
        }
    } : {}
    try {
        const res = await uni.request({
            url: baseURL + u,
            method: "PUT",
            data: data,
            timeout: 5000,
            header: {
                'content-type': 'application/x-www-form-urlencoded',
                "Access-Control-Allow-Origin": "*",
                ...config.headers
            }
        })
        console.log("put url [%s] res is", u, get(res, "data"))
        return get(res, "data")
    } catch (error: any) {
        if (isAuthorization) {
            uni.redirectTo({url: "../login/index"})
        }
        return error.request
    }
}

async function Delete(u: string, data = {}, isAuthorization = true) : Promise<any> {
    const config = isAuthorization ? {
        headers: {
            "Authorization": `Bearer ${uni.getStorageSync("token")}`
        }
    } : {}
    try {
        const res = await uni.request({
            url: baseURL + u,
            method: "DELETE",
            data: data,
            timeout: 5000,
            header: {
                'content-type': 'application/x-www-form-urlencoded',
                "Access-Control-Allow-Origin": "*",
                ...config.headers
            }
        })
        console.log("delete url [%s] res is", u, get(res, "data"))
        return get(res, "data")
    } catch (error: any) {
        if (isAuthorization) {
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
    RefreshUrl,
}