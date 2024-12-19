<script setup lang="ts">
import Ink_theme from "@/components/common/ink_theme.vue";
import Ink_screen from "@/components/common/ink_screen.vue";
import Ink_menu from "@/components/admin/ink_menu.vue";
import { collapsed } from "@/components/admin/ink_menu";
import Ink_breadcrumb from "@/components/admin/ink_breadcrumb.vue";
import Ink_user_dropdown from "@/components/common/ink_user_dropdown.vue";
</script>

<template>
  <!-- 网页布局 -->
  <div class="ink_admin">
    <!--  左边栏  -->
    <div class="ink_aside" :class="{collapsed: collapsed}">
      <!--   logo   -->
      <div class="ink_logo">

      </div>
      <!--   菜单栏   -->
      <ink_menu></ink_menu>
    </div>
    <!--  主体  -->
    <div class="ink_main">
      <!--   头部   -->
      <div class="ink_head">
        <!--   面包屑     -->
        <ink_breadcrumb></ink_breadcrumb>
        <!--   右侧信息     -->
        <div class="ink_actions">
          <icon-home/>
          <ink_theme></ink_theme>
          <ink_screen></ink_screen>
          <!--    右侧个人信息      -->
          <ink_user_dropdown></ink_user_dropdown>
        </div>
      </div>
      <!--  导航    -->
      <div class="ink_tabs">

      </div>
      <!--  主题容器    -->
      <div class="ink_container">
        <router-view></router-view>
      </div>
    </div>
  </div>
</template>

<style lang="less">

.ink_admin {
  //  左右结构
  display: flex;
  background-color: var(--color-bg-1);
  color: @color-text-1;

  .ink_aside {
    width: 240px;
    height: 100vh;
    border-right: @ink_border;
    transition: width .3s;

    &.collapsed{
      width: 48px;

      &~.ink_main{
        width: calc(100% - 48px);
      }
    }

    .ink_logo{
      width: 100%;
      height: 90px;
      border-bottom: @ink_border;
    }
  }

  .ink_main {
    width: calc(100% - 240px);
    transition: width .3s;

    .ink_head{
      height: 60px;
      width: 100%;
      border-bottom: @ink_border;
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: 0 20px;

      .ink_actions{
        display: flex;
        align-items: center;
        svg{
          font-size: 18px;
          cursor: pointer;
          margin-right: 10px;
        }
      }
    }

    .ink_tabs{
      height: 30px;
      width: 100%;
      border-bottom: @ink_border;
    }

    .ink_container{
      height: calc(100% - 90px);
      width: 100%;
      overflow-y: auto;
      overflow-x: hidden;
      background-color: @color-fill-2;
    }
  }
}
</style>