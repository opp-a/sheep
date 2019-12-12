<template>
  <el-container>
    <!--    头 -->
    <el-header>
      <el-menu :default-active="activeIndex" mode="horizontal" @select="handleSelect" router
               style="border-bottom: none;">
        <el-menu-item><img style="width: 155px; height:60px" src="../assets/home/logo.png" alt="IMG"></el-menu-item>
        <el-menu-item index="home">首页</el-menu-item>
        <el-menu-item index="warehouse">货仓</el-menu-item>
        <el-menu-item index="culture"> 文化墙</el-menu-item>
        <el-submenu class="el-menu-item-self" index="self">
          <template slot="title">个人中心</template>
          <el-menu-item index="self" style="text-align: left"><i class="el-icon-user"></i>个人中心</el-menu-item>
          <el-menu-item index="login" style="text-align: left"><i class="el-icon-top-left"></i> 退出</el-menu-item>
        </el-submenu>
      </el-menu>
    </el-header>

    <!--    内容区域-->
    <div>
      <el-main class="brand">
        <full-page ref="fullpage" :options="options" id="fullpage">
          <div class="section" style="margin-top: -20px">
            宣传展示
          </div>
          <div class="section">
            代表性产品展示区
          </div>
          <div class="section">
            公司历程简介
          </div>
          <div class="section">
            企业文化照片墙
          </div>
          <div class="section">
            联系方式
          </div>
        </full-page>
      </el-main>
    </div>

    <!--      尾 -->
    <el-footer>你是最棒的！ Little Sheep!!!</el-footer>
  </el-container>
</template>

<style>
@import 'home.css';
</style>

<script>
import '@/plugins/js/fullpage.js'

export default {
  data () {
    return {
      options: {
        menu: '#menu',
        anchors: ['page1', 'page2', 'page3', 'page3', 'page3'],
        sectionsColor: ['#5deeee', '#ef43ef', '#ece74d', '#67ec60', '#535FA2']
      },
      activeIndex: 'home'
    }
  },
  methods: {
    handleSelect (key, keyPath) {
      console.log(key, keyPath)
    }
  },
  beforeCreate () {
    console.log('xxxxx')
  }
}
</script>
