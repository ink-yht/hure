<script setup lang="ts">
import type {baseResponse, listResponse, paramsType} from "@/api";
import {reactive, ref} from "vue";
import {Message, type TableColumnData} from "@arco-design/web-vue";
import {timestampToDateTime} from "@/utils/date.ts";



interface Props {
  url: (params?: paramsType)=>Promise<baseResponse<listResponse<any>>>
  columns: TableColumnData[]
}

const props = defineProps<Props>()
const loading = ref(false)

const data = reactive<listResponse<any>>({
  "pagination": {
    page: 0,
    size: 0,
    total: 0,
    has_next: false,
  },
  "list": []
})

const params = reactive<paramsType>({
  page: 1,
  size: 1
})

const getList = async ()=>{
  loading.value = true
  const res = await props.url(params)
  loading.value = false
  console.log(res)
  if (res.code){
    Message.error(res.msg)
    return
  }

  data.list = res.data.list
  data.pagination = res.data.pagination

  console.log(data)
}
getList()

const remove = ()=>{

}

const update = (record: any)=>{

}

const pageChange = () => {
  getList()
}


const search = () =>{
  getList()
}

</script>

<template>
  <div class="ink_list_com">
    <div class="ink_list_head">
      <div class="action_create">
        <a-button type="primary">创建</a-button>
      </div>
      <div class="action_group">
        <a-select placeholder="操作" ></a-select>
      </div>
      <div class="action_search">
        <a-input-search placeholder="搜索" @search="search"></a-input-search>
      </div>
      <div class="action_search_slot">

      </div>
      <div class="action_flush">
        <icon-refresh></icon-refresh>
      </div>
    </div>
    <div class="ink_list_body">
      <a-spin :loading="loading" tip="加载中...">
        <div class="ink_list_table">
          <a-table :data="data.list" :pagination="false">
            <template #columns>
              <template v-for="col in props.columns">
                <a-table-column v-if="col.dataIndex" v-bind="col"></a-table-column>
                <a-table-column v-else-if="col.slotName" v-bind="col">
                  <template #cell="data">
                    <div class="col_actions" v-if="col.slotName === 'action'">
                      <slot v-bind="data" name="action_left"></slot>
                      <a-button type="primary" @click="update(data.record)">编辑</a-button>
                      <a-popconfirm @ok="remove" content="确定要删除该记录吗？">
                        <a-button type="primary" status="danger">删除</a-button>
                      </a-popconfirm>
                      <slot v-bind="data" name="action_right"></slot>
                    </div>
                    <div v-if="col.slotName === 'created_at'">
                      {{timestampToDateTime(data.record[col.slotName])}}
                    </div>
                    <div v-if="col.slotName === 'updated_at'">
                      {{timestampToDateTime(data.record[col.slotName])}}
                    </div>
                    <slot v-else :name="col.slotName" v-bind="data"></slot>
                  </template>
                </a-table-column>
              </template>
            </template>
          </a-table>
        </div>
        <div class="ink_list_page">
          <a-pagination show-total @change="pageChange" :total="data.pagination.total" v-model="params.page" :page-size="params.size"></a-pagination>
        </div>
      </a-spin>
    </div>
  </div>

</template>

<style lang="less">
.ink_list_com{
  .ink_list_head{
    padding: 20px 20px 10px 20px;
    border-bottom: @ink_border;
    display: flex;
    align-items: center;
    position: relative;

    .action_create,.action_group,.action_search,.action_search_slot{
      margin-right: 10px;
    }

    .action_flush{
      position: absolute;
      right: 20px;
      width: 30px;
      height: 30px;
      display: flex;
      align-items: center;
      justify-content: center;
      background-color: var(--color-fill-2);
      border-radius: 5px;
      cursor: pointer;
    }
  }

  .ink_list_body{
    padding: 10px 20px 20px 20px;

    > .arco-spin{
      width: 100%;
    }

    .ink_list_page{
      display: flex;
      justify-content: center;
      margin-top: 10px;
    }

    .col_actions{
      button{
        margin-right: 10px;

        &:last-child{
          margin-right: 0;
        }
      }
    }
  }
}
</style>