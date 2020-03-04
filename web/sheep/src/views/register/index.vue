<template>
  <div style="margin: 20px; max-width: 500px;">
    <el-form ref="ruleForm" :model="ruleForm" :rules="rules" label-width="100px">
      <el-form-item prop="name" label="账户名">
        <el-input v-model="ruleForm.name" placeholder="请输入用户名称"></el-input>
      </el-form-item>
      <el-form-item prop="password" label="密码">
        <el-input v-model="ruleForm.password" placeholder="请输入用户密码"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm(ruleForm)">立即注册</el-button>
        <el-button>取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import md5 from 'js-md5'
import {AccountRegister} from '@/api/sys.register'
export default {
  name: 'register',
  data() {
    return {
      ruleForm: {
        name: '',
        password: ''
      },
      rules: {
        name: [
          {required: true, message: '请输入账户名称', trigger: 'blur'},
          {min: 3, max: 64, message: '长度在 3 到 64 个字符', trigger: 'blur'}
        ],
        password: [
          {required: true, message: '密码不能为空'},
          {min: 6, max: 16, message: '长度在 6 到 16 个字符', trigger: 'blur'}
        ]
      }
    }
  },
  methods: {
    submitForm(form) {
      // 获取历史订单
      AccountRegister({name: form.name, password: md5.hex(form.password)})
        .then(async res => {
          this.ruleForm.name = ''
          this.ruleForm.password = ''
          this.$message({message: '成功注册用户', type: 'success'})
        })
        .catch(err => {
          this.loading = false
          if (err.toString().indexOf('[ code: 401 ]') !== -1) {
            this.$router.replace('/login')
          } else if (err.toString().indexOf('[ code: 404 ]') !== -1) {
            this.$router.replace('/404')
          } else {
            this.$message.error('未知错误！！！')
          }
        })
    }
  }
}
</script>

<style scoped></style>
