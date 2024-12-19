import {createRouter, createWebHistory} from 'vue-router'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        // 首页
        {
            name: "web",
            path: "/",
            // component: () =>import("@/views/web/index.vue"),
            redirect: "/admin", // 重定向
        },
        // 登录
        {
            name: "login",
            path: "/login",
            component: () => import("@/views/login/login.vue"),
        },
        // admin
        {
            name: "admin",
            path: "/admin",
            meta: {
                title: "首页"
            },
            component: () => import("@/views/admin/index.vue"),
            children: [
                // 首页
                {
                    name: "home",
                    path: "home",
                    meta: {
                        title: "首页"
                    },
                    component: () => import("@/views/admin/home/index.vue"),
                    children: [
                        {
                            name: "workbench",
                            path: "workbench",
                            meta: {
                                title: "工作台"
                            },
                            component: () => import("@/views/admin/home/workbench.vue"),
                        },
                        {
                            name: "systemNotice",
                            path: "systemNotice",
                            meta: {
                                title: "系统通知"
                            },
                            component: () => import("@/views/admin/home/systemNotice.vue"),
                        }
                    ]
                },
                // 招聘者管理
                {
                    name: "recruiterManagement",
                    path: "recruiterManagement",
                    meta: {
                        title: "招聘者管理"
                    },
                    component: () => import("@/views/admin/recruiterManagement/index.vue"),
                    children: [
                        {
                            name: "enterpriseInformationReview",
                            path: "enterpriseInformationReview",
                            meta: {
                                title: "企业信息审核"
                            },
                            component: () => import("@/views/admin/recruiterManagement/enterpriseInformationReview.vue"),
                        },
                        {
                            name: "publishOrderManagement",
                            path: "publishOrderManagement",
                            meta: {
                                title: "发布订单管理"
                            },
                            component: () => import("@/views/admin/recruiterManagement/publishOrderManagement.vue"),
                        },
                        {
                            name: "recruiterEvaluationManagement",
                            path: "recruiterEvaluationManagement",
                            meta: {
                                title: "招聘者评价管理"
                            },
                            component: () => import("@/views/admin/recruiterManagement/recruiterEvaluationManagement.vue"),
                        },
                        {
                            name: "recruiterCreditScore",
                            path: "recruiterCreditScore",
                            meta: {
                                title: "招聘者信用评分"
                            },
                            component: () => import("@/views/admin/recruiterManagement/recruiterCreditScore.vue"),
                        }
                    ]
                },
                // 求职者管理
                {
                    name: "jobSeekerManagement",
                    path: "jobSeekerManagement",
                    meta: {
                        title: "求职者管理"
                    },
                    component: () => import("@/views/admin/jobSeekerManagement/index.vue"),
                    children: [
                        {
                            name: "jobSeekerInformationReview",
                            path: "jobSeekerInformationReview",
                            meta: {
                                title: "求职者信息审核"
                            },
                            component: () => import("@/views/admin/jobSeekerManagement/jobSeekerInformationReview.vue"),
                        },
                        {
                            name: "jobApplicationRecords",
                            path: "jobApplicationRecords",
                            meta: {
                                title: "求职申请记录"
                            },
                            component: () => import("@/views/admin/jobSeekerManagement/jobApplicationRecords.vue"),
                        },
                        {
                            name: "jobSeekerEvaluationManagement",
                            path: "jobSeekerEvaluationManagement",
                            meta: {
                                title: "求职者评价管理"
                            },
                            component: () => import("@/views/admin/jobSeekerManagement/jobSeekerEvaluationManagement.vue"),
                        },
                        {
                            name: "jobSeekerCreditScore",
                            path: "jobSeekerCreditScore",
                            meta: {
                                title: "求职者信用评分"
                            },
                            component: () => import("@/views/admin/jobSeekerManagement/jobSeekerCreditScore.vue"),
                        }
                    ]
                },
                // 订单管理
                {
                    name: "orderManagement",
                    path: "orderManagement",
                    meta: {
                        title: "订单管理"
                    },
                    component: () => import("@/views/admin/orderManagement/index.vue"),
                    children: [
                        {
                            name: "allOrders",
                            path: "allOrders",
                            meta: {
                                title: "全部订单"
                            },
                            component: () => import("@/views/admin/orderManagement/allOrders.vue"),
                        },
                        {
                            name: "pendingOrders",
                            path: "pendingOrders",
                            meta: {
                                title: "待审核订单"
                            },
                            component: () => import("@/views/admin/orderManagement/pendingOrders.vue"),
                        },
                        {
                            name: "theOrderHasBeenApproved",
                            path: "theOrderHasBeenApproved",
                            meta: {
                                title: "已通过订单"
                            },
                            component: () => import("@/views/admin/orderManagement/theOrderHasBeenApproved.vue"),
                        },
                        {
                            name: "theOrderHasBeenCompleted",
                            path: "theOrderHasBeenCompleted",
                            meta: {
                                title: "已完成订单"
                            },
                            component: () => import("@/views/admin/orderManagement/theOrderHasBeenCompleted.vue"),
                        },
                        {
                            name: "theOrderWasCancelled",
                            path: "theOrderWasCancelled",
                            meta: {
                                title: "被取消订单"
                            },
                            component: () => import("@/views/admin/orderManagement/theOrderWasCancelled.vue"),
                        }
                    ]
                },
                // 信用体系管理
                {
                    name: "creditSystemManagement",
                    path: "creditSystemManagement",
                    meta: {
                        title: "信用体系管理"
                    },
                    component: () => import("@/views/admin/creditSystemManagement/index.vue"),
                    children: [
                        {
                            name: "userCreditScore",
                            path: "userCreditScore",
                            meta: {
                                title: "用户信用评分"
                            },
                            component: () => import("@/views/admin/creditSystemManagement/userCreditScore.vue"),
                        },
                        {
                            name: "statisticsAndAnalysis",
                            path: "statisticsAndAnalysis",
                            meta: {
                                title: "数据统计与分析"
                            },
                            component: () => import("@/views/admin/creditSystemManagement/statisticsAndAnalysis.vue"),
                        },
                        {
                            name: "riskRecordingAndHandling",
                            path: "riskRecordingAndHandling",
                            meta: {
                                title: "风险记录与处理"
                            },
                            component: () => import("@/views/admin/creditSystemManagement/riskRecordingAndHandling.vue"),
                        }
                    ]
                },
                // 财务管理
                {
                    name: "financialManagement",
                    path: "financialManagement",
                    meta: {
                        title: "财务管理"
                    },
                    component: () => import("@/views/admin/financialManagement/index.vue"),
                    children: [
                        {
                            name: "orderRevenueStatistics",
                            path: "orderRevenueStatistics",
                            meta: {
                                title: "订单收入统计"
                            },
                            component: () => import("@/views/admin/financialManagement/orderRevenueStatistics.vue"),
                        },
                        {
                            name: "platformExpenseManagement",
                            path: "platformExpenseManagement",
                            meta: {
                                title: "平台费用管理"
                            },
                            component: () => import("@/views/admin/financialManagement/platformExpenseManagement.vue"),
                        },
                        {
                            name: "reconciliationRecords",
                            path: "reconciliationRecords",
                            meta: {
                                title: "对账记录"
                            },
                            component: () => import("@/views/admin/financialManagement/reconciliationRecords.vue"),
                        }
                    ]
                },
                // 系统设置
                {
                    name: "settingsMessage",
                    path: "settingsMessage",
                    meta: {
                        title: "系统设置"
                    },
                    component: () => import("@/views/admin/settingsMessage/index.vue"),
                    children: [
                        {
                            name: "adminSettings",
                            path: "adminSettings",
                            meta: {
                                title: "管理员设置"
                            },
                            component: () => import("@/views/admin/settingsMessage/adminSettings.vue"),
                        },
                        {
                            name: "personalHomepage",
                            path: "personalHomepage",
                            meta: {
                                title: "个人主页"
                            },
                            component: () => import("@/views/admin/settingsMessage/self_info.vue"),
                        },
                        {
                            name: "operationLogs",
                            path: "operationLogs",
                            meta: {
                                title: "操作日志"
                            },
                            component: () => import("@/views/admin/settingsMessage/operationLogs.vue"),
                        },
                        {
                            name: "systemInformation",
                            path: "systemInformation",
                            meta: {
                                title: "系统信息"
                            },
                            component: () => import("@/views/admin/settingsMessage/systemInformation.vue"),
                        }
                    ]
                }
            ]
        }
    ],
})

export default router
