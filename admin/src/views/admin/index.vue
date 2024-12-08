<script setup lang="ts">


import Ink_theme from "@/components/common/ink_theme.vue";
</script>

<template>
  <!-- 网页布局 -->
  <div class="ink_admin">
    <!--  左边栏  -->
    <div class="ink_aside">
      <!--   logo   -->
      <div class="ink_logo">

      </div>
      <!--   菜单栏   -->
      <div class="ink_menu">

      </div>
    </div>
    <!--  主体  -->
    <div class="ink_main">
      <!--   头部   -->
      <div class="ink_head">
        <!--   面包屑     -->
        <div class="ink_breadcrumbs"></div>
        <!--   右侧信息     -->
        <div class="ink_actions">
          <icon-home/>
          <ink_theme></ink_theme>
          <icon-fullscreen/>
          <icon-fullscreen-exit/>
          <!--    右侧个人信息      -->
          <div class="ink_user_info_action">

          </div>
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
  background-color: @color-fill-1;
  color: @color-text-1;

  .ink_aside {
    width: 240px;
    height: 100vh;
    overflow-y: auto; // 纵向滚动条
    overflow-x: hidden; // 超出隐藏
    border-right: @ink_border;

    .ink_logo{
      width: 100%;
      height: 90px;
      border-bottom: @ink_border;
    }
  }

  .ink_main {
    width: calc(100% - 240px);

    .ink_head{
      height: 60px;
      width: 100%;
      border-bottom: @ink_border;
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: 0 20px;
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
    }
  }
}
</style>