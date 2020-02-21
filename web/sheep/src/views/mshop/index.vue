<template>
  <div style="padding: 10px 50px 50px 50px;">
    <el-collapse v-model="activeName" accordion>
      <el-collapse-item name="addOrEditShop" style="left: 50px; color:#17517493">
        <template slot="title">
          <span style="font-size:larger; color: rgb(59, 126, 228); margin: 20px;">{{ addOredit }}</span>
        </template>
        <el-card shadow="never" style="width: 100%; text-align: left;">
          <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
            <el-form-item v-show="false" label="商品ID" prop="id">
              <el-input v-model="ruleForm.id"></el-input>
            </el-form-item>
            <el-form-item label="商品名称" prop="name">
              <el-input v-model="ruleForm.name"></el-input>
            </el-form-item>
            <el-form-item label="商品封面" prop="delivery">
              <el-upload action="#" list-type="picture-card" :auto-upload="false">
                <i slot="default" class="el-icon-plus"></i>
                <div slot="file" slot-scope="{file}">
                  <img class="el-upload-list__item-thumbnail" :src="file.url" alt="" />
                  <span class="el-upload-list__item-actions">
                    <span class="el-upload-list__item-preview" @click="handlePictureCardPreview(file)">
                      <i class="el-icon-zoom-in"></i>
                    </span>
                    <span v-if="!disabled" class="el-upload-list__item-delete" @click="handleDownload(file)">
                      <i class="el-icon-download"></i>
                    </span>
                    <span v-if="!disabled" class="el-upload-list__item-delete" @click="handleRemove(file)">
                      <i class="el-icon-delete"></i>
                    </span>
                  </span>
                </div>
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
            <el-form-item label="新增商品数量" prop="num">
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
      <el-table-column prop="name" label="名称" width="230"> </el-table-column>
      <el-table-column prop="pricein" label="进货价(￥)" width="180"> </el-table-column>
      <el-table-column prop="priceout" label="出货价(￥)" width="180"> </el-table-column>
      <el-table-column prop="num" label="存货量" width="180"> </el-table-column>
      <el-table-column prop="describe" label="备注" width="350"> </el-table-column>
      <el-table-column prop="addtime" label="上新时间"> </el-table-column>
      <el-table-column fixed="right" label="操作" width="120">
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
      dialogImageUrl: '',
      dialogVisible: false,
      disabled: false,

      activeName: '',
      ruleForm: {
        id: '',
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
          {min: 3, max: 5, message: '长度在 3 到 32 个字符', trigger: 'blur'}
        ],
        icons: [{type: 'array', required: true, message: '请至少选择一个商品封面', trigger: 'change'}],
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
      handleRemove(file) {
        console.log(file)
      },
      handlePictureCardPreview(file) {
        this.dialogImageUrl = file.url
        this.dialogVisible = true
      },
      handleDownload(file) {
        console.log(file)
      },
      shops: [
        {
          id: 'xxx',
          name: '王小虎',
          pricein: 56,
          priceout: 80,
          num: 100,
          describe: '上海市普陀区金沙江路 1518 弄',
          addtime: '2016-05-02'
        },
        {
          name: '王小虎',
          pricein: 56,
          priceout: 80,
          num: 100,
          describe: '上海市普陀区金沙江路 1518 弄',
          addtime: '2016-05-02'
        },
        {
          name: '王小虎',
          pricein: 56,
          priceout: 80,
          num: 100,
          describe: '上海市普陀区金沙江路 1518 弄',
          addtime: '2016-05-02'
        },
        {
          name: '王小虎',
          pricein: 56,
          priceout: 80,
          num: 100,
          describe: '上海市普陀区金沙江路 1518 弄',
          addtime: '2016-05-02'
        }
      ],
      pageindex: 1,
      total: 0
    }
  },
  computed: {
    addOredit: function() {
      if (this.ruleForm.id === undefined || this.ruleForm.id === '') {
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
      this.ruleForm.id = rows[index].id
      this.ruleForm.name = rows[index].name
      this.ruleForm.icons = rows[index].icons
      this.ruleForm.pricein = rows[index].pricein
      this.ruleForm.priceout = rows[index].priceout
      this.ruleForm.num = rows[index].num
      this.ruleForm.desc = rows[index].desc
      this.activeName = 'addOrEditShop'
    },
    listshops() {
      GetMShops({page: this.pageindex, pagenumber: 20})
        .then(async res => {
          // this.shops.splice(0, this.shops.lenght)
          this.total = res.total
          this.shops = res
        })
        .catch(() => {})
    },
    handleDeleteRow(index, rows) {
      DeleteMShops({shopid: rows[index].id})
      rows.splice(index, 1)
    },
    getSummaries(param) {
      const {columns, data} = param
      const sums = []
      columns.forEach((column, index) => {
        if (index === 0) {
          sums[index] = '合计'
        } else if (index === 1 || index === 2) {
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
          alert('submit!')
          AddUpdateMShops(this.ruleForm)
          this.ruleForm.id = ''
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
    }
  },
  created() {
    this.listshops()
  }
}
</script>

<style scoped></style>
