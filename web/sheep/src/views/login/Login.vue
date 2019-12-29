<template>
  <div class="container-login100">
    <div class="wrap-login100">
      <div class="login100-pic" data-tilt>
        <img src="../../assets/login/img-01.png" alt="IMG">
      </div>

      <div class="login100-form" id="loginForm">
        <span class="login100-form-title">会员登陆</span>

        <div class="wrap-input100">
          <input id="phone" class="input100" type="text" name="phone" placeholder="手机号">
          <span class="focus-input100"></span>
          <span class="symbol-input100"><i class="fa fa-mobile-phone" aria-hidden="true"></i></span>
        </div>

        <div class="wrap-input100">
          <input id="password" class="input100" type="password" name="pass" placeholder="密码">
          <span class="focus-input100"></span>
          <span class="symbol-input100"><i class="fa fa-lock" aria-hidden="true"></i></span>
        </div>

        <div class="container-login100-form-btn">
          <button class="login100-form-btn" v-on:click="submit">
            登陆
          </button>
        </div>

        <div class="text-center" style="padding-top: 12px; text-align: center">
          <a class="txt2" href="javascript:">
            忘记密码？
          </a>
        </div>

        <div class="text-center" style="padding-top: 136px; text-align: center">
          <a class="txt2" href="javascript:" target="_blank">
            还没有账号？立即注册
            <i class="fa fa-long-arrow-right m-l-5" aria-hidden="true"></i>
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import 'animate.css'
import 'font-awesome/css/font-awesome.min.css'
import { login } from '../../api/getData'

export default {
  name: 'Login',
  data () {
    return {}
  },
  methods: {
    submit: function () {
      let phone = document.getElementById('phone').value
      let password = document.getElementById('password').value
      console.log(phone)
      console.log(password)
      let res = login({ phone: phone, password: password })
      if (res.name === 'admin') {
        alert('login success!!!')
        // this.$store.commit('setUserName', res.name)
        localStorage.setItem('userName', res.name)
        this.$router.push('/')
      } else {
        alert('手机号请输入123')
      }
    }
  },
  beforeCreate () {
    localStorage.clear()
  }
}
</script>

<style scoped>
@import "./css/main.css";
</style>
