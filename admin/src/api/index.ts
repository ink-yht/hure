import axios from "axios"
import {Message} from "@arco-design/web-vue";
import {useAdminStore} from "../stores/modules/admin/admin.ts"



export const useAxios = axios.create({
    timeout: 6000,
    baseURL: "", // 在使用前端代理的情况下，这里必须留空，不然会跨域
})

useAxios.interceptors.request.use((config) => {
    const token = localStorage.getItem("token")
    config.headers.setAuthorization("Bearer " + token, true)
    console.log("token1",token)
    return config
})

useAxios.interceptors.response.use((res) => {

    const newToken = res.headers["x-jwt-token"]
    if (newToken) {
        const useStore = useAdminStore()
        useStore.saveToken(newToken)
    }

    if (res.status === 200){
        return res.data
    }
    return res
}, (res)=>{
    Message.error(res.message)
})

export interface baseResponse<T> {
    code: number;
    msg: string;
    data: T
}

export interface listResponse<T> {
        pagination: {
            page: number
            size: number
            total: number
            has_next: boolean
        },
        list: T[]
}

export interface paramsType {
    page?: number
    size?:number
}