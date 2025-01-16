import {type baseResponse, type listResponse, type paramsType, useAxios} from "@/api";
import type {JobSeekerListType} from "@/api/job_seeker/type.ts";

// 统一管理接口
enum Job_seekerAPI {
    // 获取所有求职用户信息
    Job_seekerList = '/api/job_seeker/get_infos',
}

// 获取所有求职用户信息接口
export const job_SeekerListApi = (params?: paramsType) => useAxios.get<any, baseResponse<listResponse<JobSeekerListType>>>(Job_seekerAPI.Job_seekerList,{params});

