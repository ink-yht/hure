<script setup lang="ts">
import {reactive, ref} from "vue";
import type {LoginFormRequest} from "@/api/admin/type.ts";
import {LoginApi} from "@/api/admin/admin.ts";
import {Message} from "@arco-design/web-vue";
import router from "@/router";
import {useAdminStore} from "@/stores/modules/admin/admin.ts";

const useStore = useAdminStore()

const formRef = ref()

const form = reactive<LoginFormRequest>({
  email: "",
  password:"",
})

const login = async () =>{

  const val = await  formRef.value.validate()
  if (val){
    return
  }
  const res = await LoginApi(form)
  console.log(res)
  if (res.code){
    Message.error(res.msg)
    return
  }
  await useStore.saveAdminInfo()
  await router.push({
    name: "workbench"
  })
  Message.success("登录成功")
}
</script>

<template>
  <div class="login_view">
    <div class="login_mask">
      <a-form ref="formRef" :model="form" :label-col-props="{span:0}" :wrapper-col-props="{span:24}">
        <div class="title">用户登录</div>
        <a-form-item label="邮箱" field="email" :rules="[{required:true,message:'请输入邮箱'}]">
          <a-input placeholder="邮箱" v-model="form.email">
            <template #prefix>
              <icon-email></icon-email>
            </template>
          </a-input>
        </a-form-item>
        <a-form-item label="密码" field="password" :rules="[{required:true,message:'请输入密码'}]">
          <a-input type="password" placeholder="密码" v-model="form.password">
            <template #prefix>
              <icon-lock></icon-lock>
            </template>
          </a-input>
        </a-form-item>
        <a-form-item>
          <a-button type="primary" @click="login" long>登录</a-button>
        </a-form-item>
      </a-form>
    </div>
  </div>
</template>

<style lang="less">
.login_view{
  width: 100%;
  min-height: 100vh;
  background-image: url('../../../public/登录背景图.jpg'); // 替换为实际的图片路径
  background-repeat: no-repeat;
  background-size: cover;
  background-position: center;
  position: relative;

  .login_mask{
    width: 400px;
    height: 100vh;
    background-color: rgba(white,0.6);
    position: absolute;
    right: 0;
    top: 0;
    display: flex;
    justify-content: center;
    align-items: center;

    .arco-form{
      padding: 40px;
    }

    .title{
      font-size: 26px;
      font-weight: 600;
      text-align: center;
      color: @primary-6;
      margin-bottom: 20px;
    }
  }
}
</style>