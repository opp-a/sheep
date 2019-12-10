<template>
  <el-container>
    <el-header>
      <el-menu :default-active="activeIndex" mode="horizontal" @select="handleSelect" router="true"
               style="border-bottom: none;">
        <el-menu-item><img style="width: 155px; height:60px" src="../assets/home/logo.png" alt="IMG"></el-menu-item>
        <el-menu-item index="home" route="/">首页</el-menu-item>
        <el-menu-item index="warehouse" route="">货仓</el-menu-item>
        <el-menu-item index="culture" route="">文化墙</el-menu-item>
        <el-submenu class="el-menu-item-self" index="self" route="">
          <template slot="title">个人中心</template>
          <el-menu-item index="self">个人中心</el-menu-item>
          <el-menu-item index="logout">退出</el-menu-item>
        </el-submenu>
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
/* 导航头区域 */
.el-header, .el-footer {
  background-color: #FFFFFF;
  color: #333;
  text-align: center;
  line-height: 60px;
}

.el-submenu.el-menu-item-self {
  text-align: center;
  background-color: #ffffff;
  float: right;
}

/* 内容区域 */
.el-main.brand {
  background-color: #a9fff1;
  color: #333;
  text-align: center;
  line-height: 460px;
  margin-top: 5px;
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

</style>

<script>
export default {
  data () {
    return {
      activeIndex: 'home'
    }
  },
  methods: {
    handleSelect (key, keyPath) {
      console.log(key, keyPath)
    }
  }
}
</script>
