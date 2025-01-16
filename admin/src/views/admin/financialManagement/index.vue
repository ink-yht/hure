<script setup lang="ts">

</script>

<template>
  <div class="scrollbar">
    <router-view class="ink_base_view" v-slot="{Component}">
      <transition name="fade" mode="out-in">
        <component :is="Component"></component>
      </transition>
    </router-view>
  </div>
</template>

<style lang="less">

// 组件刚开始离开
.fade-leave-active{
}
// 组件离开结束
.fade-leave-to {
  transform: translateX(30px);
  opacity: 0;
}

// 组件刚开始进入
.fade-enter-active {
  transform: translateX(-30px);
  opacity: 1;
}

// 组件进入完成
.fade-enter-to {
  transform: translateX(0px);
  opacity: 0;
}

// 正在进入和离开
.fade-leave-active, .fade-enter-active {
  transition: all .3s ease-out;
}
</style>