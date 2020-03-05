<template>
  <div style="padding: 10px 50px 50px 50px;">
    <el-collapse v-model="activeName" accordion>
      <el-collapse-item name="addOrEditShop" style="left: 50px; color:#17517493">
        <template slot="title">
          <span style="font-size:larger; color: rgb(59, 126, 228); margin: 20px;">{{ addOredit }}</span>
        </template>
        <el-card shadow="never" style="width: 100%; text-align: left;">
          <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
            <el-form-item v-show="true" label="商品ID" prop="shopid">
              <el-input v-model="ruleForm.shopid" :disabled="true"></el-input>
            </el-form-item>
            <el-form-item label="商品名称" prop="name">
              <el-input v-model="ruleForm.name"></el-input>
            </el-form-item>
            <el-form-item label="商品封面" prop="icons">
              <el-upload
                action=""
                :http-request="httpRequest"
                list-type="picture-card"
                :file-list="fileList"
                :limit="5"
                :on-preview="handlePictureCardPreview"
                :on-remove="handleRemove"
                :before-upload="handleBeforeUpload"
              >
                <i class="el-icon-plus"></i>
              </el-upload>
              <el-dialog :visible.sync="dialogVisible">
                <img width="100%" :src="dialogImageUrl" alt="" />
              </el-dialog>
            </el-form-item>
            <el-form-item label="进货价格" prop="pricein">
              <el-input
                type="text"
                v-model.number="ruleForm.pricein"
                autocomplete="off"
                placeholder="请输入进货价格"
                style="max-width: 200px;"
              ></el-input>
              <span style="font-size: larger;">￥</span>
            </el-form-item>
            <el-form-item label="出货价格" prop="priceout">
              <el-input
                type="text"
                v-model.number="ruleForm.priceout"
                autocomplete="off"
                placeholder="请输入出货价格"
                style="max-width: 200px;"
              ></el-input>
              <span style="font-size: larger;">￥</span>
            </el-form-item>
            <el-form-item label="商品数量" prop="num">
              <el-input-number v-model="ruleForm.num" :min="1" :max="10000"></el-input-number>
            </el-form-item>
            <el-form-item label="商品描述" prop="desc">
              <el-input type="textarea" v-model="ruleForm.desc" :autosize="{minRows: 4, maxRows: 6}"></el-input>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="submitForm('ruleForm')">确定提交</el-button>
              <el-button @click="resetForm('ruleForm')">重置</el-button>
            </el-form-item>
          </el-form>
        </el-card>
      </el-collapse-item>
    </el-collapse>
    <el-table :data="shops" :summary-method="getSummaries" border show-summary stripe style="width: 100%;">
      <el-table-column prop="shopid" label="商品ID" width="250"> </el-table-column>
      <el-table-column prop="name" label="名称"> </el-table-column>
      <el-table-column prop="pricein" label="进货价(￥)" width="110"> </el-table-column>
      <el-table-column prop="priceout" label="出货价(￥)" width="110"> </el-table-column>
      <el-table-column prop="num" label="存货量" width="90"> </el-table-column>
      <el-table-column prop="desc" label="备注"> </el-table-column>
      <el-table-column prop="addtime" label="上新时间"> </el-table-column>
      <el-table-column fixed="right" label="操作" width="140">
        <template slot-scope="scope">
          <el-button @click.native.prevent="handleEditShop(scope.$index, shops)" type="text" size="small">
            编辑
          </el-button>
          <el-button
            @click.native.prevent="handleDeleteRow(scope.$index, shops)"
            type="text"
            size="small"
            style="color: red;"
          >
            移除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      @prev-click="handleCurrentChange"
      @next-click="handleCurrentChange"
      @current-change="handleCurrentChange"
      :current-page="pageindex"
      :total="total"
      :page-size="20"
      :pager-count="7"
      layout="prev, pager, next"
      style="text-align: left; margin: 15px 15px 0px 0px;"
    >
    </el-pagination>
  </div>
</template>

<script>
import {GetMShops, DeleteMShops, AddUpdateMShops} from '@/api/sys.mshop'
export default {
  name: 'mShop',
  data() {
    return {
      // 图片预览
      dialogImageUrl: '',
      dialogVisible: false,
      fileList: [],

      // 表单
      activeName: '',
      ruleForm: {
        shopid: '',
        name: '',
        icons: [],
        pricein: 0,
        priceout: 0,
        num: 0,
        desc: ''
      },
      rules: {
        name: [
          {required: true, message: '请输入商品名称', trigger: 'blur'},
          {min: 3, max: 64, message: '长度在 3 到 64 个字符', trigger: 'blur'}
        ],
        pricein: [
          {required: true, message: '进货价格不能为空'},
          {type: 'number', message: '进货价格必须为数字值'}
        ],
        priceout: [
          {required: true, message: '出货价格不能为空'},
          {type: 'number', message: '出货价格必须为数字值'}
        ],
        num: [
          {required: true, message: '商品数量不能为空'},
          {type: 'number', message: '商品数量必须为数字值'}
        ],
        desc: [{required: true, message: '请填写备注', trigger: 'blur'}]
      },

      // 商品列表
      shops: [],
      pageindex: 1,
      total: 0
    }
  },
  computed: {
    addOredit: function() {
      if (this.ruleForm.shopid === undefined || this.ruleForm.shopid === '') {
        return '新增商品'
      } else {
        return '编辑商品'
      }
    }
  },
  methods: {
    // 分页
    handleCurrentChange: function(currentPage) {
      this.pageindex = currentPage
      this.listshops()
    },

    // 表格
    handleEditShop: function(index, rows) {
      this.ruleForm.shopid = rows[index].shopid
      this.ruleForm.name = rows[index].name
      this.ruleForm.icons = rows[index].icons
      for (const index in this.ruleForm.icons) {
        this.fileList.push({
          name: '',
          url: this.ruleForm.icons[index]
        })
      }
      this.ruleForm.pricein = rows[index].pricein
      this.ruleForm.priceout = rows[index].priceout
      this.ruleForm.num = rows[index].num
      this.ruleForm.desc = rows[index].desc
      this.activeName = 'addOrEditShop'
    },
    listshops() {
      GetMShops({page: this.pageindex, pagenumber: 20})
        .then(async res => {
          this.total = res.total
          this.shops = res.infos
          for (let i = 0; i < this.shops.length; i++) {
            this.shops[i].addtime = this.$moment(new Date(this.shops[i].addtime)).format('YYYY-MM-DD HH:mm:ss')
          }
        })
        .catch(err => {
          if (err.toString().indexOf('[ code: 401 ]') !== -1) {
            this.$router.replace('/login')
          } else if (err.toString().indexOf('[ code: 404 ]') !== -1) {
            this.$router.replace('/404')
          } else {
            this.$message.error('未知错误！！！')
          }
        })
    },
    handleDeleteRow(index, rows) {
      DeleteMShops({shopid: rows[index].shopid})
        .then(async res => {
          this.$message({message: '成功删除了一条商品记录', type: 'success'})
        })
        .catch(err => {
          if (err.toString().indexOf('[ code: 401 ]') !== -1) {
            this.$router.replace('/login')
          } else if (err.toString().indexOf('[ code: 404 ]') !== -1) {
            this.$router.replace('/404')
          } else {
            this.$message.error('未知错误！！！')
          }
        })
      rows.splice(index, 1)
    },
    getSummaries(param) {
      const {columns, data} = param
      const sums = []
      columns.forEach((column, index) => {
        if (index === 0) {
          sums[index] = '合计'
        } else if (index === 2 || index === 3) {
          const values = data.map(item => Number(item[column.property]))
          if (!values.every(value => isNaN(value))) {
            sums[index] = values.reduce((prev, curr, currindex) => {
              const value = Number(curr)
              if (!isNaN(value)) {
                return prev + curr * data[currindex].num
              } else {
                return prev
              }
            }, 0)
            sums[index] = sums[index] + ' 元'
          } else {
            sums[index] = ''
          }
        } else if (index === 3) {
          const values = data.map(item => Number(item[column.property]))
          if (!values.every(value => isNaN(value))) {
            sums[index] = values.reduce((prev, curr) => {
              const value = Number(curr)
              if (!isNaN(value)) {
                return prev + curr
              } else {
                return prev
              }
            }, 0)
            sums[index] = sums[index] + '件'
          } else {
            sums[index] = 'N/A'
          }
        } else {
          sums[index] = ''
        }
      })

      return sums
    },

    // 表单
    submitForm(formName) {
      this.$refs[formName].validate(valid => {
        if (valid) {
          this.ruleForm.icons = []
          for (const index in this.fileList) {
            this.ruleForm.icons.push(this.fileList[index].url)
          }
          AddUpdateMShops(this.ruleForm)
            .then(async res => {
              if (this.addOredit === '新增商品') {
                this.$message({message: '成功添加了一条商品记录', type: 'success'})
              } else {
                this.$message({message: '成功修改了一条商品记录', type: 'success'})
              }
              this.listshops()

              this.ruleForm.shopid = ''
              this.ruleForm.name = ''
              this.ruleForm.icons = []
              this.ruleForm.pricein = 0
              this.ruleForm.priceout = 0
              this.ruleForm.num = 1
              this.ruleForm.desc = ''
              this.fileList = []
              this.activeName = ''
            })
            .catch(err => {
              console.log(err.toString())
              if (err.toString().indexOf('[ code: 401 ]') !== -1) {
                this.$router.replace('/login')
              } else if (err.toString().indexOf('[ code: 404 ]') !== -1) {
                this.$router.replace('/404')
              } else {
                this.$message.error('未知错误！！！')
              }
            })
        } else {
          return false
        }
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
      this.fileList = []
      this.activeName = ''
    },
    // 表单图片
    handlePictureCardPreview(file) {
      this.dialogImageUrl = file.url
      this.dialogVisible = true
    },
    handleRemove(file, fileList) {
      for (const index in this.fileList) {
        if (this.fileList[index].url === file.url) {
          this.fileList.splice(index, 1)
          return
        }
      }
    },
    handleBeforeUpload(file) {
      // check file
      const isJPG = file.type === 'image/jpeg'
      const isLt2M = file.size / 1024 / 1024 < 2

      if (!isJPG) {
        this.$message.error('上传图片只能是 JPG 格式!')
      }
      if (!isLt2M) {
        this.$message.error('上传图片大小不能超过 2MB!')
      }
      if (isJPG && isLt2M) {
        return true
      } else {
        return false
      }
    },
    httpRequest(options) {
      var This = this
      const file = options.file
      if (file) {
        var reader = new FileReader()
        reader.readAsDataURL(file)
        reader.onload = function(e) {
          This.fileList.push({name: file.name, url: e.target.result})
        }
      }
    },
    // 备用 将图片base64数据转blob url
    dataURItoBlob(base64Data) {
      var byteString
      if (base64Data.split(',')[0].indexOf('base64') >= 0) byteString = atob(base64Data.split(',')[1])
      else byteString = unescape(base64Data.split(',')[1])
      var mimeString = base64Data
        .split(',')[0]
        .split(':')[1]
        .split(';')[0]
      var ia = new Uint8Array(byteString.length)
      for (var i = 0; i < byteString.length; i++) {
        ia[i] = byteString.charCodeAt(i)
      }
      return new Blob([ia], {
        type: mimeString
      })
    }
  },
  created() {
    this.listshops()
  }
}
</script>

<style scoped></style>
