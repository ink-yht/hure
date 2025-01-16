import {type baseResponse, useAxios} from "@/api";
import {type AdminRequest, type LoginFormRequest} from "./type.ts"

// 统一管理接口
enum AdminAPI {
    // 用户登录接口
    Login = '/api/admins/login',
    // 获取用户信息
    GetAdminInfo = '/api/admins/info',
    // 获取用户列表
    GetUserList = '/user/list',
    // 删除用户
    DeleteUser = '/user/delete',
}

// 登录接口
export const LoginApi = (data: LoginFormRequest) => useAxios.post<any, baseResponse<any>>(AdminAPI.Login, data);

// 获取管理员个人信息

export const AdminInfoApi = () => useAxios.get<any, baseResponse<AdminRequest>>(AdminAPI.GetAdminInfo);
