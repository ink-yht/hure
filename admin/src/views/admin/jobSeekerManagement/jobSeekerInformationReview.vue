<script setup lang="ts">

import Ink_list from "@/components/admin/ink_list.vue";
import {job_SeekerListApi} from "@/api/job_seeker/job_seeker.ts";

const columns = [
  {title: "ID", dataIndex: 'id'},
  {title: "ID", dataIndex: 'name'},
  {title: "ID", dataIndex: 'gender'},
  {title: "ID", dataIndex: 'phone'},
  {title: "ID", dataIndex: 'degree'},
  {title: "ID", dataIndex: 'year_of_birth'},
  {title: "ID", dataIndex: 'working_hours'},
  {title: "ID", dataIndex: 'profile'},
  {title: "修改时间", slotName: 'updated_at'},
  {title: "操作", slotName: 'action'},
]

</script>

<template>
  <div>
    <ink_list :url="job_SeekerListApi" :columns="columns">
    </ink_list>
  </div>
</template>

<style lang="less">

</style>