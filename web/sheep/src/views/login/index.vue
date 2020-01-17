<template>
  <div class="container-login100">
    <div class="wrap-login100">
      <div class="login100-pic" data-tilt>
        <img src="./images/img-01.png" alt="IMG" />
      </div>

      <div class="login100-form" id="loginForm">
        <span class="login100-form-title">会员登陆</span>

        <div class="wrap-input100">
          <input id="phone" class="input100" type="text" name="phone" placeholder="手机号" />
          <span class="focus-input100"></span>
          <span class="symbol-input100">
            <i class="fa fa-mobile-phone" aria-hidden="true"></i>
          </span>
        </div>

        <div class="wrap-input100">
          <input id="password" class="input100" type="password" name="pass" placeholder="密码" />
          <span class="focus-input100"></span>
          <span class="symbol-input100">
            <i class="fa fa-lock" aria-hidden="true"></i>
          </span>
        </div>

        <div class="container-login100-form-btn">
          <button class="login100-form-btn" v-on:click="submit">登陆</button>
        </div>

        <div class="text-center" style="padding-top: 12px; text-align: center">
          <a class="txt2" href="javascript:">忘记密码？</a>
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
import {mapActions} from 'vuex'

export default {
  name: 'Login',
  data() {
    return {}
  },
  methods: {
    ...mapActions('sheep/account', ['login']),
    submit: function() {
      console.log(process.env.VUE_APP_API)
      const phone = document.getElementById('phone').value
      const password = document.getElementById('password').value
      this.login({
        username: phone,
        password: password
      })
        .then(() => {
          // 重定向对象不存在则返回顶层路径
          this.$router.replace(this.$route.query.redirect || '/')
        })
        .catch(() => {
          // 错误情况
        })
    }
  }
}
</script>

<style scoped>
a:hover {
  text-decoration: none;
  color: #57b846;
}

/*---------------------------------------------*/
input {
  outline: none;
  border: none;
}

/*---------------------------------------------*/
button {
  outline: none !important;
  border: none;
  background: transparent;
}

button:hover {
  cursor: pointer;
}

/*//////////////////////////////////////////////////////////////////
[ Utility ]*/
.txt1 {
  font-family: Poppins-Regular;
  font-size: 13px;
  line-height: 1.5;
  color: #999999;
}

.txt2 {
  font-family: Poppins-Regular;
  font-size: 13px;
  line-height: 1.5;
  color: #666666;
}

.container-login100 {
  position: fixed;
  box-sizing: border-box;
  width: 100%;
  min-height: 100%;
  background-color: #255873;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  align-items: center;
  box-shadow: 2px 2px 0px 10px rgba(37, 88, 115, 1);
}

.wrap-login100 {
  height: 500px;
  width: 850px;
  background: #efefef;
  border-radius: 10px;
  overflow: hidden;

  display: -webkit-box;
  display: -webkit-flex;
  display: -moz-box;
  display: -ms-flexbox;
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
  padding: 177px 130px 33px 95px;
}

/*------------------------------------------------------------------
[  ]*/
.login100-form {
  width: 290px;
}

.login100-form-title {
  font-family: Poppins-Bold;
  font-size: 24px;
  color: #555555;
  line-height: 1.2;
  text-align: center;

  width: 100%;
  display: block;
  padding-bottom: 54px;
}

/*---------------------------------------------*/
.wrap-input100 {
  box-sizing: border-box;
  position: relative;
  width: 100%;
  z-index: 1;
  margin-bottom: 10px;
}

.input100 {
  margin: 0px;
  box-sizing: border-box;
  font-family: Poppins-Medium;
  font-size: 15px;
  line-height: 1.5;
  color: #444444;

  display: block;
  width: 100%;
  background: #f6f6f6;
  height: 50px;
  border-radius: 25px;
  padding: 0 30px 0 68px;
}

.symbol-input100 {
  margin: 0px;
  box-sizing: border-box;
  font-size: 15px;

  display: -webkit-box;
  display: -webkit-flex;
  display: -moz-box;
  display: -ms-flexbox;
  display: flex;
  align-items: center;
  position: absolute;
  border-radius: 25px;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 100%;
  padding-left: 35px;
  pointer-events: none;
  color: #dedede;
  box-shadow: 0px 0px 9px 0.6px;

  -webkit-transition: all 0.4s;
  -o-transition: all 0.4s;
  -moz-transition: all 0.4s;
  transition: all 0.4s;
}

.input100:focus + .focus-input100 + .symbol-input100 {
  color: #b0b0b0;
  padding-left: 28px;
}

/*------------------------------------------------------------------
[ Button ]*/
.container-login100-form-btn {
  width: 100%;
  display: -webkit-box;
  display: -webkit-flex;
  display: -moz-box;
  display: -ms-flexbox;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  padding-top: 20px;
}

.login100-form-btn {
  font-family: Montserrat-Bold;
  font-size: 15px;
  line-height: 1.5;
  color: #fff;
  text-transform: uppercase;

  width: 100%;
  height: 50px;
  border-radius: 25px;
  background: #346d86;
  display: -webkit-box;
  display: -webkit-flex;
  display: -moz-box;
  display: -ms-flexbox;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 0 25px;

  -webkit-transition: all 0.4s;
  -o-transition: all 0.4s;
  -moz-transition: all 0.4s;
  transition: all 0.4s;
}

.login100-form-btn:hover {
  background: #3c5977;
}
</style>
