<template>
  <el-container>
    <el-header>
      <el-menu :default-active="activeIndex" class="el-menu" mode="horizontal" @select="handleSelect"  router="true">
        <el-menu-item index="1">首页</el-menu-item>
        <el-menu-item index="2">货仓</el-menu-item>
        <el-menu-item index="3">文化墙</el-menu-item>
        <el-menu-item index="4" route="">个人中心</el-menu-item>
      </el-menu>
    </el-header>

    <el-main class="brand">品牌形象宣传展示区</el-main>
    <el-divider></el-divider>
    <el-main class="product">代表性产品展示区</el-main>
    <el-divider></el-divider>
    <el-main class="introduction">公司历程简介</el-main>
    <el-divider></el-divider>
    <el-main class="culture">企业文化照片墙</el-main>
    <el-divider></el-divider>
    <el-main class="contact">联系方式</el-main>
    <el-footer>Footer</el-footer>
  </el-container>
</template>

<style>
.el-header, .el-footer {
  background-color: #FFFFFF;
  color: #333;
  text-align: center;
  line-height: 60px;
}

.el-aside {
  background-color: #D3DCE6;
  color: #333;
  text-align: center;
  line-height: 200px;
}

.el-divider {
  background-color: #CCCCCC;
  color: #333;
  text-align: center;
}

.el-main.brand {
  background-color: #a9fff1;
  color: #333;
  text-align: center;
  line-height: 460px;
}

.el-main.product {
  background-color: #DDDDDD;
  color: #333;
  text-align: center;
  line-height: 460px;
}

.el-main.introduction {
  background-color: #FF00FF;
  color: #333;
  text-align: center;
  line-height: 460px;
}

.el-main.culture {
  background-color: #FFFF00;
  color: #333;
  text-align: center;
  line-height: 460px;
}

.el-main.contact {
  background-color: #00FFFF;
  color: #333;
  text-align: center;
  line-height: 460px;
}

body > .el-container {
  margin-bottom: 40px;
}

.el-container:nth-child(5) .el-aside,
.el-container:nth-child(6) .el-aside {
  line-height: 260px;
}

.el-container:nth-child(7) .el-aside {
  line-height: 320px;
}
</style>

<script>
export default {
  data () {
    return {
      activeIndex: '1',
      activeIndex2: '1'
    }
  },
  methods: {
    handleSelect (key, keyPath) {
      console.log(key, keyPath)
    }
  }
}
</script>
