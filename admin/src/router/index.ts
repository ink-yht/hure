import { createRouter, createWebHistory } from 'vue-router'

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
      component: () =>import("@/views/login/login.vue"),
    },
      // admin
    {
      name: "admin",
      path: "/admin",
      component: () =>import("@/views/admin/index.vue"),
    }
  ],
})

export default router
