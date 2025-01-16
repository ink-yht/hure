<script setup lang="ts">
import {IconClose} from "@arco-design/web-vue/es/icon";
import {useRoute} from "vue-router";
import router from "@/router";
import {onMounted, ref, watch} from "vue";
import {Swiper, SwiperSlide} from 'swiper/vue';

const route = useRoute()

interface TabType {
  name: string
  title: string
}

const tabs = ref<TabType[]>([
  {title: "工作台", name: "workbench"},
])

const check = (item: TabType) => {
  router.push({
    name: item.name
  })
  saveTabs()

}

const saveTabs = () => {
  localStorage.setItem("ink_tabs", JSON.stringify(tabs.value))
}

const removeItem = (item: TabType) => {

  if (item.name === "workbench") {
    return
  }
  const index = tabs.value.findIndex(value => item.name === value.name)
  if (index != -1) {

    if (item.name === route.name) {
      router.push({
        name: tabs.value[index - 1].name
      })
    }

    tabs.value.splice(index, 1)
  }
  saveTabs()
}

const removeAllItem = (item: TabType) => {
  tabs.value = [{title: "工作台", name: "workbench"},]
  router.push({
    name: "workbench"
  })
  saveTabs()
}

const loadTabs = () => {
  const ink_tabs = localStorage.getItem("ink_tabs")
  if (ink_tabs) {
    try {
      tabs.value = JSON.parse(ink_tabs)
    } catch (e) {
      console.log(e)
    }
  }
}
loadTabs()

watch(() => route.name, () => {
  // 判断当前路由的名称在不在tabs中
  const index = tabs.value.findIndex(value => route.name === value.name);
  if (index === -1) {
    const title = route.meta?.title as string || "未知页面"; // 使用类型断言并提供默认值
    tabs.value.push({
      name: route.name as string,
      title
    });
  }
}, {immediate: true});


const slideCount = ref(100)

onMounted(()=>{
  // 显示的总宽度
  const swiperDom = document.querySelector(".ink_tabs_swiper") as HTMLDivElement
  const swiperWidth = swiperDom.clientWidth

  // 实际的总宽度
  const wrapperDom = document.querySelector(".ink_tabs_swiper") as HTMLDivElement
  const wrapperwidth = wrapperDom.scrollWidth

  if (swiperWidth > wrapperwidth){
    return
  }

  // 如果实际总宽度大于了显示的总宽度
  const slideList = document.querySelectorAll(".ink_tabs_swiper .swiper-slide")
  let allWith = 0
  let index  = 1
  for (const slideListElement of slideList) {
    allWith += (slideListElement.clientWidth + 20)
    index++
    if (allWith >= swiperWidth){
     break
    }
  }
  slideCount.value = index

  // 选中高亮的元素
  // const activeSlide = document.querySelector(".ink_tabs_swiper .swiper-slide.active") as HTMLDivElement
  // if (activeSlide.offsetLeft > swiperWidth){
  //   const offsetLeft = swiperWidth - activeSlide.offsetLeft
  //   console.log(offsetLeft)
  //   setTimeout(()=>{
  //     wrapperDom.style.transform = `translate3d(${offsetLeft}px, 0px, 0px)`
  //   })
  //
  // }
})
</script>

<template>
  <div class="ink_tabs">
    <swiper class="ink_tabs_swiper" :slides-per-view="slideCount">
      <swiper-slide v-for="item in tabs"  :class="{active: route.name===item.name}">
        <div class="item" @click="check(item)" @mousedown.middle.stop="removeItem(item)"
             :class="{active: route.name===item.name}" >
          {{ item.title }}
          <span class="close" @click.stop="removeItem(item)" v-if="item.name !== 'workbench'">
        <IconClose></IconClose>
      </span>
        </div>
      </swiper-slide>
    </swiper>
    <div class="item" @click="removeAllItem">
      清除全部
    </div>
  </div>
</template>

<style lang="less">
.ink_tabs {
  display: flex;
  align-items: center;
  padding: 0 10px;
  justify-content: space-between;

  .swiper {
    width: calc(100% - 100px);
    display: flex;
    overflow-x: hidden;
    overflow-y: hidden;

    .swiper-wrapper{
      display: flex;
      align-items: center;

      .swiper-slide{
        width: fit-content !important;
        flex-shrink: 0;
      }
    }
  }

  .item {
    padding: 3px 8px;
    background-color: var(--color-bg-1);
    border: @ink_border;
    margin-right: 10px;
    cursor: pointer;
    border-radius: 5px;


    &:hover {
      background-color: var(--color-fill-1);
    }

    &.active {
      background-color: @primary-6;
      color: white;
    }
  }
}
</style>