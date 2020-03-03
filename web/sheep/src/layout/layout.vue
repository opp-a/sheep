<template>
  <el-container>
    <!--    头 -->
    <el-header>
      <el-menu
        :default-active="activeIndex"
        mode="horizontal"
        @select="handleSelect"
        style="border-bottom: none;"
        background-color="rgba(255, 255, 255, 0)"
        text-color="#A0A0A0"
        active-text-color="#51EE48"
      >
        <el-menu-item index="index"> 首页</el-menu-item>
        <el-menu-item index="warehouse">货仓</el-menu-item>
        <el-menu-item index="culture">文化墙</el-menu-item>
        <el-submenu style="text-align: center; float: right;" index="self">
          <template slot="title">{{ info.name ? `你好 ${info.name}` : '未登录' }}</template>
          <el-menu-item index="myorder" style="text-align: left"> <i class="el-icon-user"></i>个人中心 </el-menu-item>
          <el-menu-item index="manageShop" style="text-align: left" v-if="isAdmin">
            <i class="el-icon-s-grid"></i>商品管理
          </el-menu-item>
          <el-menu-item index="manageCulture" style="text-align: left" v-if="isAdmin">
            <i class="el-icon-picture-outline"></i>文化管理
          </el-menu-item>
          <el-menu-item index="login" style="text-align: left" @click.native="logOff">
            <i class="el-icon-top-left"></i> 退出
          </el-menu-item>
        </el-submenu>
      </el-menu>
    </el-header>

    <!--    内容区域-->
    <el-main style="margin: 0px; padding: 0px;">
      <router-view />
    </el-main>

    <!--      尾 -->
    <el-footer
      style="color: rgba(100, 100, 100, 0.5); text-align: center; line-height: 60px; width: 100%;
    background: rgba(53, 73, 94, 0.0); bottom: 0px;"
      >你是最棒的！ Little Sheep!!!</el-footer
    >
  </el-container>
</template>

<style scoped></style>

<script>
import {mapState, mapActions} from 'vuex'
export default {
  data() {
    return {
      activeIndex: 'home'
    }
  },
  computed: {
    ...mapState('sheep/user', ['info']),
    isAdmin() {
      return this.info.name === 'admin'
    }
  },
  mounted() {
    window.addEventListener('unload', this.load())
    var path = this.$route.path + ''
    this.activeIndex = path.substring(path.lastIndexOf('/') + 1)
  },
  methods: {
    ...mapActions('sheep/account', ['logout', 'load']),
    /**
     * @description 登出
     */
    logOff() {
      this.logout({
        confirm: true
      })
    },
    /**
     * @description 导航
     */
    handleSelect(key, keyPath) {
      this.activeIndex = key
      if (key === 'login') {
        this.$router.push('/' + key)
      } else {
        this.$router.push('/home/' + key)
      }
    }
  }
}
</script>
