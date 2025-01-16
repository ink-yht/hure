import {type baseResponse, type paramsType} from "@/api";

// 统一管理接口
import type {listResponse} from "@/api";
import {useAxios} from "@/api";
import type {RecruiteListType} from "@/api/recruite/type.ts";

enum recruiteAPI {
    // 获取所有招聘用户信息
    recruiteList = '/api/recruiter/get_infos',
}

// 获取所有求职用户信息接口
export const recruiteListApi = (params?: paramsType) => useAxios.get<any, baseResponse<listResponse<RecruiteListType>>>(recruiteAPI.recruiteList,{params});

