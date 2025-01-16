// 创建用户相关的仓库
import { defineStore } from 'pinia'
import {reactive, ref} from "vue";
import {type AdminType} from "./type.ts"
import {AdminInfoApi} from "@/api/admin/admin.ts";
import {Message} from "@arco-design/web-vue";

export const useAdminStore = defineStore('admin', ()=>{
    const token = ref("")
    const adminInfo = reactive<AdminType>({
        id:0,
        email:"",
        phone:"",
        avatar:"",
        nickname:"",
        signature:"",
    })

    const saveAdminInfo = async () =>{

        const res = await AdminInfoApi()
        if (res.code){
            Message.error(res.msg)
            return
        }
        adminInfo.id = res.data.id
        adminInfo.email = res.data.email
        adminInfo.phone = res.data.phone
        adminInfo.avatar = res.data.avatar
        adminInfo.nickname = res.data.nickname
        adminInfo.signature = res.data.signature
    }

    const saveToken = (newToken:string)=>{
        token.value = newToken
        localStorage.setItem("token",newToken)
        saveAdminInfo().then(r => r)
    }

    const removeToken = ()=>{
        token.value = ''
    }
    return {adminInfo,token,saveToken,saveAdminInfo,removeToken}
},{
    persist: true,
})
