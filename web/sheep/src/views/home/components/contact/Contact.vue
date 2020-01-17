<template>
  <div class="wrap-contact100">
    <el-form
      label-position="right"
      label-width="80px"
      :model="addData"
      :rules="addRules"
      ref="addForm"
      show-message
      inline-message
      status-icon
      style="width: calc(100% - 75%);background-color: white; padding: 170px 40px 180px 40px"
    >
      <span
        style=" width: 100%; display: block; font-family: Poppins-Regular; font-size: 30px;
    color: #333333; line-height: 1.2; text-align: center; padding-bottom: 48px;"
        >联系我们</span
      >
      <el-form-item label="名称" prop="name" required>
        <el-input v-model="addData.name" placeholder="请输入姓名"></el-input>
      </el-form-item>
      <el-form-item label="邮箱" prop="email" required>
        <el-input v-model="addData.email" placeholder="请输入电子邮箱"></el-input>
      </el-form-item>
      <el-form-item label="服务类型" prop="service">
        <el-select v-model="addData.service" placeholder="请选择" style="width: 100%;">
          <el-option label="业务咨询" value="业务咨询"></el-option>
          <el-option label="留言" value="留言"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="内容" prop="content" required>
        <el-input type="textarea" v-model="addData.content" placeholder="请输入留言" :rows="5" maxlength="300" show-word-limit></el-input>
      </el-form-item>
      <div style="margin: 120px;"></div>
      <el-form-item>
        <div style="text-align: center;">
          <el-button type="primary" @click="submit('addForm')" style="width: 80%; position: relative;">提交 </el-button>
        </div>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import 'font-awesome/css/font-awesome.min.css'
import {mapActions} from 'vuex'

export default {
  name: 'Contact',
  data: function() {
    var validateEmail = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('您的电子邮箱是？'))
      } else if (this.addData.email.trim().match(/^([a-zA-Z0-9_\-\.]+)@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.)|(([a-zA-Z0-9\-]+\.)+))([a-zA-Z]{1,5}|[0-9]{1,3})(\]?)$/) == null) {
        callback(new Error('请输入有效电子邮箱'))
      } else {
        callback()
      }
    }
    var validateName = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('您的尊称？'))
      } else {
        callback()
      }
    }
    var validateContent = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('您的问题是？'))
      } else {
        callback()
      }
    }
    return {
      addData: {
        name: '',
        email: '',
        service: '',
        content: ''
      },
      addRules: {
        name: [{validator: validateName, trigger: 'blur'}],
        email: [{validator: validateEmail, trigger: 'blur'}],
        content: [{validator: validateContent, trigger: 'blur'}]
      }
    }
  },
  methods: {
    ...mapActions('sheep/contact', ['newContact']),
    submit: function(formName) {
      this.$refs[formName].validate(valid => {
        if (valid) {
          this.newContact(this.addData).then(() => {
            // 重定向对象不存在则返回顶层路径
            // this.$router.replace(this.$route.query.redirect || '/')
            this.$refs[formName].resetFields()
          })
        }
      })
    }
  }
}
</script>

<style scoped>
.el-form-item {
  height: 60px;
}

.wrap-contact100 {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: stretch;
  flex-direction: row-reverse;
  background-image: url('./images/bg-01.jpg');
}
</style>
