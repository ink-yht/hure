// 登录接口请求参数类型
export interface LoginFormRequest {
    email: string;
    password: string;
}

// 管理员个人信息
export interface AdminRequest {
    id:number
    email:string
    phone:string
    avatar:string
    nickname:string
    signature:string
    created_at:number
    updated_at:number
}