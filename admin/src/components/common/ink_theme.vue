<script setup lang="ts">
import {theme, setTheme, loadTheme} from "@/components/common/ink_theme.ts";

loadTheme()

</script>

<template>
  <span v-if="theme === 'dark'" title="白天模式"><icon-sun-fill @click="setTheme('')"/></span>
  <span v-if="theme === ''" title="黑夜模式"><icon-moon-fill @click="setTheme('dark')"/></span>
</template>

<style scoped>

</style>