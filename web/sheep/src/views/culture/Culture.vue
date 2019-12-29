<template>
  <el-container>
    <!--    头 -->
    <el-header>
      <el-menu default-active="culture" mode="horizontal" @select="handleSelect" router
               style="border-bottom: none;" background-color="rgba(255, 255, 255, 0)" text-color="#A0A0A0"
               active-text-color="#51EE48">
        <el-menu-item index="home">首页</el-menu-item>
        <el-menu-item index="warehouse">货仓</el-menu-item>
        <el-menu-item index="culture"> 文化墙</el-menu-item>
        <el-submenu index="self" style="text-align: center; float: right;">
          <template slot="title">{{this.$store.state.userName}}</template>
          <el-menu-item index="self" style="text-align: left"><i class="el-icon-user"></i>个人中心</el-menu-item>
          <el-menu-item index="login" style="text-align: left"><i class="el-icon-top-left"></i> 退出
          </el-menu-item>
        </el-submenu>
      </el-menu>
    </el-header>
    <div style="height: 20px; width: 100%;"></div>
    <el-timeline>
      <el-timeline-item v-for="itemArray in culturePictures" :key="itemArray.time" :timestamp="itemArray.time"
                        placement="top">
        <div
          style="display: flex; flex-direction: row; flex-wrap: wrap; justify-content: flex-start; align-items: flex-start; align-content: flex-start">
          <div  v-for="item in itemArray.pictures" :key="item.time" style="flex: none; margin: 10px;">
            <el-card :body-style="{ padding: '0px' }">
              <el-image fit="scale-down" :src="item.src" lazy></el-image>
            </el-card>
            <div style="width: 100%; text-align: center">{{item.name}}</div>
          </div>
        </div>
      </el-timeline-item>
    </el-timeline>
    <p v-if="loading" style="text-align: center; color: #8c939d;">加载中...</p>
    <p v-if="noMore" style="text-align: center; color: #8c939d;">没有更多了</p>
    <!--      尾 -->
    <el-footer style="color: rgba(100, 100, 100, 0.5);text-align: center;line-height: 60px;width: 100%;
    background: rgba(53, 73, 94, 0.0);bottom: 0px;">你是最棒的！ Little Sheep!!!
    </el-footer>
  </el-container>
</template>

<script>
import { getCulturePictures } from '../../api/culture'
import { getShops } from '../../api/getData'

export default {
  name: 'Culture',
  data () {
    return {
      loading: false,
      noMore: false,
      culturePictures: []
    }
  },
  methods: {
    // element 导航
    handleSelect (key, keyPath) {
      this.activeIndex = key
    },
    getPictures () {
      this.loading = true
      let pageNumber = 12
      let page = this.culturePictures.length / pageNumber + 1
      setTimeout(() => {
        let res = getCulturePictures({ page: page, pageNumber: pageNumber })
        this.culturePictures.push.apply(this.culturePictures, res)
        if (res.length < pageNumber) {
          this.noMore = true
        }
        this.loading = false
      }, 2000)
    },
    // 无限加载
    handleScroll () {
      // 变量windowHeight是可视区的高度
      let scrollTop = document.documentElement.scrollTop || document.body.scrollTop
      // 变量scrollHeight是滚动条的总高度
      let windowHeight = document.documentElement.clientHeight || document.body.clientHeight
      let scrollHeight = document.documentElement.scrollHeight || document.body.scrollHeight
      if (scrollTop + windowHeight === scrollHeight) {
        if (this.loading === false && this.noMore === false) {
          this.getPictures()
        }
      }
    }
  },
  mounted: function () {
    window.addEventListener('scroll', this.handleScroll, true)
  },
  created () {
    this.getPictures()
  }
}
</script>

<style scoped>

</style>
