<script setup lang="ts">

import {ref} from "vue";

const isScreen = ref(false)

// 全屏
const fullScreen = () => {
  document.documentElement?.requestFullscreen()
  isScreen.value = true
}

// 退出全屏
const exitFullScreen = () => {
  document?.exitFullscreen()
  isScreen.value = false
}

</script>

<template>
  <div>
    <span v-if="!isScreen" title="全屏"><icon-fullscreen @click="fullScreen"/></span>
    <span v-else title="退出全屏"><icon-fullscreen-exit @click="exitFullScreen"/></span>
  </div>
</template>

<style scoped>

</style>