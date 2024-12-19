<script setup lang="ts">

import {type Component, onMounted, ref, watch} from "vue";
import {
  IconHome,
  IconUser,
  IconUserGroup,
  IconSend,
  IconBookmark,
  IconBarChart,
  IconSettings
} from "@arco-design/web-vue/es/icon";
import {collapsed} from "./ink_menu";
import router from "@/router";
import {useRoute} from "vue-router";

const route = useRoute()

interface MenuType {
  title: string
  name: string
  icon?: string | Component
  children?: MenuType[]
}


const menuList: MenuType[] = [
  {
    title: "首页", name: "home", icon: IconHome, children: [
      {title: "工作台", name: "workbench"},
      {title: "系统通知", name: "systemNotice"}
    ]
  },
  {
    title: "招聘者管理", name: "recruiterManagement", icon: IconUser, children: [
      {title: "企业信息审核", name: "enterpriseInformationReview"},
      {title: "发布订单管理", name: "publishOrderManagement"},
      {title: "招聘者评价管理", name: "recruiterEvaluationManagement"},
      {title: "招聘者信用评分", name: "recruiterCreditScore"}
    ]
  },
  {
    title: "求职者管理", name: "jobSeekerManagement", icon: IconUserGroup, children: [
      {title: "求职者信息审核", name: "jobSeekerInformationReview"},
      {title: "求职申请记录", name: "jobApplicationRecords"},
      {title: "求职者评价管理", name: "jobSeekerEvaluationManagement"},
      {title: "求职者信用评分", name: "jobSeekerCreditScore"}
    ]
  },
  {
    title: "订单管理", name: "orderManagement", icon: IconSend, children: [
      {title: "全部订单", name: "allOrders"},
      {title: "待审核订单", name: "pendingOrders"},
      {title: "已通过订单", name: "theOrderHasBeenApproved"},
      {title: "已完成订单", name: "theOrderHasBeenCompleted"},
      {title: "被取消订单", name: "theOrderWasCancelled"}
    ]
  },
  {
    title: "信用体系管理", name: "creditSystemManagement", icon: IconBookmark, children: [
      {title: "用户信用评分", name: "userCreditScore"},
      {title: "数据统计与分析", name: "statisticsAndAnalysis"},
      {title: "风险记录与处理", name: "riskRecordingAndHandling"}
    ]
  },
  {
    title: "财务管理", name: "financialManagement", icon: IconBarChart, children: [
      {title: "订单收入统计", name: "orderRevenueStatistics"},
      {title: "平台费用管理", name: "platformExpenseManagement"},
      {title: "对账记录", name: "reconciliationRecords"}
    ]
  },
  {
    title: "系统设置", name: "settingsMessage", icon: IconSettings, children: [
      {title: "管理员设置", name: "adminSettings"},
      {title: "个人主页", name: "personalHomepage"},
      {title: "操作日志", name: "operationLogs"},
      {title: "系统信息", name: "systemInformation"}
    ]
  }
]

const selectedKeys = ref<string[]>([]);
const openKeys = ref<string[]>([]);

const menuItemClick = (key: string) => {
  localStorage.setItem('currentMenuKey', key); // 存储当前选中菜单项
  router.push({ name: key }); // 跳转路由
  selectedKeys.value = [key]; // 更新选中项
};


const initRoute = () => {
  // 获取当前存储的选中菜单项
  const storedKey = localStorage.getItem('currentMenuKey');

  // 查找当前路由对应的父菜单
  const matchedParent = menuList.find(menu =>
      menu.children?.some(child => child.name === route.name)
  );

  // 设置选中项
  if (storedKey && route.name === storedKey) {
    selectedKeys.value = [storedKey]; // 如果当前路由和存储的路由一致，直接选中
  } else if (route.name) {
    selectedKeys.value = [route.name as string]; // 当前路由名作为选中项
  }

  // 设置展开的父菜单
  if (matchedParent) {
    openKeys.value = [matchedParent.name];
  } else {
    openKeys.value = []; // 没有父菜单时清空展开项
  }
};


onMounted(() => {
  initRoute();
});

watch(
    () => route.name, // 监听路由的 name
    () => {
      initRoute(); // 每当路由变化，重新初始化菜单状态
    },
    { immediate: true } // 确保初始化时立即执行
);

</script>

<template>
  <div class="ink_menu" :class="{collapsed: collapsed}">
    <div class="ink_menu_inner scrollbar">
      <a-menu
          @menu-item-click="menuItemClick"
          v-model:open-keys="openKeys"
          v-model:selected-keys="selectedKeys"
          v-model:collapsed="collapsed"
          show-collapse-button
      >
        <template v-for="menu in menuList" :key="menu.name">
          <a-menu-item :key="menu.name" v-if="!menu.children">
            <template #icon>
              <component :is="menu.icon" style="font-size: 18px;"></component>
            </template>
            {{ menu.title }}
          </a-menu-item>
          <a-sub-menu :key="menu.name" v-else :title="menu.title">
            <template #icon>
              <component :is="menu.icon" style="font-size: 18px;"></component>
            </template>
            <a-menu-item :key="sub.name" v-for="sub in menu.children" :style="{ paddingLeft: '27px' }">{{
                sub.title
              }}
            </a-menu-item>
          </a-sub-menu>
        </template>
      </a-menu>
    </div>

  </div>
</template>

<style lang="less">
.ink_menu {
  height: calc(100vh - 90px);
  position: relative;

  &.collapsed {
    .arco-menu-collapse-button {
      left: 48px !important;
    }
  }

  &:hover {
    .arco-menu-collapse-button {
      opacity: 1 !important;
    }
  }

  .ink_menu_inner {
    height: 100%;
    overflow-y: auto;
    overflow-x: hidden;

    .arco-menu {
      position: inherit;

      .arco-menu-collapse-button {
        top: 50%;
        transform: translate(-50%, -50%);
        left: 240px;
        transition: all .3s;
        opacity: 0;
      }
    }
  }

}
</style>