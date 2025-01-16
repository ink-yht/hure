<script setup lang="ts">
import {collapsed} from "@/components/admin/ink_menu.ts";
</script>

<template>
  <div class="ink_logo" :class="{collapsed:collapsed}">
    <img src="https://www.fengfengzhidao.com/image/logo.png" alt="logo" class="logo">
    <div class="slogan">
      <div>淡漠</div>
      <div>ink-Light</div>
    </div>

  </div>
</template>

<style lang="less">
.ink_logo{
  display: flex;
  justify-content: center;
  align-items: center;

  &.collapsed{
    .slogan{
      transform: scale(0);
      transform-origin: left;
      opacity: 0;
    }
    .logo{
      transform: translateX(20px);
      width: 40px;
      height: 40px;
    }
  }

  .logo{
    width: 50px;
    height: 50px;
    transition: all .3s;
  }

  .slogan{
    margin-left: 10px;
    transition: all .3s;
    div:nth-child(1){
      font-size: 22px;
      font-weight: 600;
    }
    div:nth-child(2){
      font-size: 12px;
      margin-top: 1px;
    }
  }
}
</style>