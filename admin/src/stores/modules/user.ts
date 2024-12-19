// 创建用户相关的仓库
import { defineStore } from 'pinia'
import {ref} from "vue";

export const useUserStore = defineStore('user', ()=>{
    const token = ref('')
    const  setToken = (newToken:string)=>{
        token.value = newToken
    }
    const removeToken = ()=>{
        token.value = ''
    }
    return {token,setToken,removeToken}
})
