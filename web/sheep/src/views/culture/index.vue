<template>
  <el-container>
    <!--    头 -->
    <el-header>
      <el-menu
        default-active="culture"
        mode="horizontal"
        @select="handleSelect"
        router
        style="border-bottom: none;"
        background-color="rgba(255, 255, 255, 0)"
        text-color="#A0A0A0"
        active-text-color="#51EE48"
      >
        <el-menu-item index="home">首页</el-menu-item>
        <el-menu-item index="warehouse">货仓</el-menu-item>
        <el-menu-item index="culture"> 文化墙</el-menu-item>
        <el-submenu index="self" style="text-align: center; float: right;">
          <template slot="title">{{ this.$store.state.userName }}</template>
          <el-menu-item index="self" style="text-align: left"><i class="el-icon-user"></i>个人中心</el-menu-item>
          <el-menu-item index="manageShop" style="text-align: left" v-if="isSuperAdmin"
            ><i class="el-icon-s-grid"></i>商品管理
          </el-menu-item>
          <el-menu-item index="manageCulture" style="text-align: left" v-if="isSuperAdmin"
            ><i class="el-icon-picture-outline"></i>文化管理
          </el-menu-item>
          <el-menu-item index="login" style="text-align: left"><i class="el-icon-top-left"></i> 退出 </el-menu-item>
        </el-submenu>
      </el-menu>
    </el-header>
    <div style="height: 20px; width: 100%;"></div>
    <el-timeline>
      <el-timeline-item
        v-for="itemArray in culturePictures"
        :key="itemArray.time"
        :timestamp="itemArray.time"
        placement="top"
      >
        <div
          style="display: flex; flex-direction: row; flex-wrap: wrap; justify-content: flex-start; align-items: flex-start; align-content: flex-start"
        >
          <div v-for="item in itemArray.pictures" :key="item.time" style="flex: none; margin: 10px;">
            <img
              :src="item.src"
              style="margin: 0px; padding: 0px; height: 100%; width: 100%; box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1)"
            />
            <div style="width: 100%; text-align: center">{{ item.name }}</div>
          </div>
        </div>
      </el-timeline-item>
    </el-timeline>
    <p v-if="loading" style="text-align: center; color: #8c939d;">加载中...</p>
    <p v-if="noMore" style="text-align: center; color: #8c939d;">没有更多了</p>
    <!--      尾 -->
    <el-footer
      style="color: rgba(100, 100, 100, 0.5);text-align: center;line-height: 60px;width: 100%;
    background: rgba(53, 73, 94, 0.0);bottom: 0px;"
      >你是最棒的！ Little Sheep!!!
    </el-footer>
  </el-container>
</template>

<script>
import {getCulturePictures} from '../../api/culture'
// import { getShops } from '../../api/getData'

export default {
  name: 'Culture',
  data() {
    return {
      loading: false,
      noMore: false,
      culturePictures: []
    }
  },
  methods: {
    // element 导航
    handleSelect(key, keyPath) {
      this.activeIndex = key
    },
    getPictures() {
      this.loading = true
      const pageNumber = 12
      const page = this.culturePictures.length / pageNumber + 1
      setTimeout(() => {
        const res = getCulturePictures({page: page, pageNumber: pageNumber})
        this.culturePictures.push.apply(this.culturePictures, res)
        if (res.length < pageNumber) {
          this.noMore = true
        }
        this.loading = false
      }, 2000)
    },
    // 无限加载
    handleScroll() {
      // 变量windowHeight是可视区的高度
      const scrollTop = document.documentElement.scrollTop || document.body.scrollTop
      // 变量scrollHeight是滚动条的总高度
      const windowHeight = document.documentElement.clientHeight || document.body.clientHeight
      const scrollHeight = document.documentElement.scrollHeight || document.body.scrollHeight
      if (scrollTop + windowHeight === scrollHeight) {
        if (this.loading === false && this.noMore === false) {
          this.getPictures()
        }
      }
    }
  },
  mounted: function() {
    window.addEventListener('scroll', this.handleScroll, true)
  },
  created() {
    this.getPictures()
  }
}
</script>

<style scoped></style>
