<template>
  <el-container>
    <!--    头 -->
    <el-header>
      <el-menu default-active="warehouse" mode="horizontal" @select="handleSelect" router
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
    <el-main>

    </el-main>
    <!--      尾 -->
    <el-footer style="color: rgba(100, 100, 100, 0.5);
    text-align: center;
    line-height: 60px;
    width: 100%;
    background: rgba(53, 73, 94, 0.0);
    bottom: 0px;">你是最棒的！ Little Sheep!!!
    </el-footer>
  </el-container>
</template>

<script>
import VueGridLayout from 'vue-grid-layout'

export default {
  name: 'wareHouse',
  data () {
    return {
      layout: {
        layout: [
          { 'x': 0, 'y': 0, 'w': 4, 'h': 6, 'i': '0', 'title': 'attitude' }, // 态度
          { 'x': 4, 'y': 0, 'w': 6, 'h': 4, 'i': '1' }, // 留白
          { 'x': 10, 'y': 0, 'w': 4, 'h': 6, 'i': '2', 'title': 'quality' }, // 品质
          { 'x': 0, 'y': 4, 'w': 4, 'h': 2, 'i': '3' }, // 留白
          { 'x': 4, 'y': 4, 'w': 6, 'h': 6, 'i': '4' }, // 观点
          { 'x': 10, 'y': 4, 'w': 4, 'h': 2, 'i': '5' }, // 留白
          { 'x': 0, 'y': 10, 'w': 4, 'h': 6, 'i': '6', 'title': 'details' }, // 细心
          { 'x': 4, 'y': 10, 'w': 6, 'h': 4, 'i': '7' }, // 留白
          { 'x': 10, 'y': 10, 'w': 4, 'h': 6, 'i': '8', 'title': 'responsibility' } // 责任
        ],
        colNum: 14,
        maxRows: 14,
        rowHeight: 60,
        isDraggable: false,
        isResizable: false,
        isMirrored: false,
        autoSize: false,
        verticalCompact: true,
        margin: [0, 0],
        useCssTransforms: true
      }
    }
  },
  methods: {
    // element 导航
    handleSelect (key, keyPath) {
      this.activeIndex = key
    }
  }
}
</script>

<style scoped>

</style>
