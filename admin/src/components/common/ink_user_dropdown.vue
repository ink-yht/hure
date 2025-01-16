<script setup lang="ts">
import router from "@/router";
import {useAdminStore} from "@/stores/modules/admin/admin.ts";

const useStore = useAdminStore()

const handleSelect = (val:any) => {
  if (val === "logout"){
    return
  }
  router.push({
    name: val
  })
}

interface OptionsType {
  name:string
  title:string
}

const options : OptionsType[] = [
  {title: "个人主页",name:"personalHomepage"},
  {title: "退出登录",name:"logout"},
]

const adminInfo =  useStore.adminInfo
console.log(adminInfo)

</script>

<template>
  <a-dropdown  @select="handleSelect" :popup-max-height="false">
    <div class="ink_user_dropdown_com">
      <a-avatar :size="35" image-url="https://avatars.githubusercontent.com/u/182111738?v=4"></a-avatar>
      <span class="user_name">{{adminInfo.nickname}}</span>
      <icon-down></icon-down>
    </div>

    <template #content>
      <a-doption v-for="item in options" :value="item.name">{{item.title}}</a-doption>
    </template>
  </a-dropdown>
</template>

<style lang="less">
.ink_user_dropdown_com{
  cursor: pointer;

  .user_name{
    margin: 0 5px;

    svg{
      margin-right: 0 !important;

    }
  }
}
</style>