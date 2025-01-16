<script setup lang="ts">

import Ink_list from "@/components/admin/ink_list.vue";
import {job_SeekerListApi} from "@/api/job_seeker/job_seeker.ts";
import {recruiteListApi} from "@/api/recruite/recruite.ts";

const columns = [
  {title: "ID", dataIndex: 'id'},
  {title: "公司名称", dataIndex: 'company'},
  {title: "公司地址", dataIndex: 'address'},
  {title: "联系人姓名", dataIndex: 'contact'},
  {title: "联系人电话", dataIndex: 'phone'},
  {title: "创建时间", dataIndex: 'created_at'},
]

</script>

<template>
  <div>
    <ink_list :url="recruiteListApi" :columns="columns"></ink_list>
  </div>
</template>

<style lang="less">

</style>