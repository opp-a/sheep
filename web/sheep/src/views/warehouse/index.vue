<template>
  <div>
    <!--      购物车-->
    <div
      :class="shopDrawer"
      style="position: fixed; text-align: right; width:20%; height: 80%; z-index: 100; right: 0px; min-width: 400px;"
    >
      <el-row style="height: 100%;">
        <el-col :span="3" style="height: 100%;">
          <el-button
            style="margin: 0px; padding: 6px; font-size: x-large; height: 100%;"
            type="info"
            icon="el-icon-shopping-cart-full"
            @click="drawer = !drawer"
          ></el-button>
        </el-col>
        <el-col :span="21" style="height: 100%;">
          <div
            style="width: 100%; height: calc(100% - 20px); background-color: white; box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1); padding: 10px;"
          >
            <div
              style="width:100%; color: rgba(100, 100, 100, 0.5); font-family: Montserrat-Bold; text-align: center; font-size: x-large;padding-top: 20px; padding-bottom: 20px;"
            >
              购物车
            </div>
            <el-table :data="shoppingCar" style="width: 100%" max-height="600">
              <el-table-column fixed prop="name" label="商品" width="105px"></el-table-column>
              <el-table-column fixed label="价格" width="60px">
                <template slot-scope="scope"
                  ><p style="color:#f24a4a">{{ scope.row.price }}￥</p></template
                >
              </el-table-column>
              <el-table-column fixed label="数量" width="140">
                <template slot-scope="scope">
                  <el-input-number
                    v-model="scope.row.num"
                    @change="editShoppingCar(scope.$index, scope.row)"
                    :min="0"
                    :max="100"
                    size="mini"
                  ></el-input-number>
                </template>
              </el-table-column>
            </el-table>
            <div style="position: fixed; bottom: 15px; width: 75%;">
              <el-row>
                <el-col :span="12"
                  ><div style="font-size: x-large; color: #3c5977; font-family: Calibri; text-align: left"
                    >共计{{ totalPrice }}￥</div
                  ></el-col
                >
                <el-col :span="12"><el-button @click="pay()">去支付</el-button></el-col>
              </el-row>
            </div>
          </div>
        </el-col>
      </el-row>
    </div>
    <div :class="shopDrawer2" style="position: fixed; text-align: right; height: 80%; z-index: 1000; right: 0px;">
      <el-button
        style="margin: 0px; padding: 6px; font-size: x-large; height: 100%;"
        type="info"
        icon="el-icon-shopping-cart-full"
        @click="drawer = !drawer"
      ></el-button>
    </div>
    <!--      商品列表-->
    <div style="overflow:auto;">
      <el-row v-for="row in shopRows" :key="row">
        <el-col :span="4" v-for="(o, index) in 4" :key="o" :offset="index > 0 ? 2 : 1">
          <el-card :body-style="{padding: '0px'}">
            <el-image
              src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
              style="width: 100%; display: block; box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);"
            ></el-image>
            <div style="padding: 14px;">
              <el-row>
                <el-col :span="12"
                  ><span>{{ shops[(row - 1) * 4 + index].shopName }}</span></el-col
                >
                <el-col :span="12" style="text-align: right;"
                  ><span style="width:100%; text-align: right; font-size: x-large; color: #f02e29;"
                    >{{ shops[(row - 1) * 4 + index].price }}￥</span
                  >
                </el-col>
              </el-row>
              <div style="margin-top: 15px;">
                <el-row>
                  <el-col :span="12"
                    ><span style="color: #b2b1b3; line-height: 35px;">
                      存货量：{{ shops[(row - 1) * 4 + index].num }}件</span
                    >
                  </el-col>
                  <el-col :span="12" style="text-align: right;">
                    <el-button
                      type="primary"
                      icon="el-icon-shopping-cart-full"
                      style="background-color: #6b9bb3; border: #6b9bb3;"
                      @click="addShopToCar(shops[(row - 1) * 4 + index])"
                      >添加
                    </el-button>
                  </el-col>
                </el-row>
              </div>
            </div>
          </el-card>
          <div style="height: 40px;"></div>
        </el-col>
      </el-row>
      <p v-if="loading" style="text-align: center; color: #8c939d;">加载中...</p>
      <p v-if="noMore" style="text-align: center; color: #8c939d;">没有更多了</p>
    </div>
  </div>
</template>

<script>
import '@/plugins/js/animate.js'
import {GetShops, GetCar, DoOrder, SaveOrder} from '@/api/sys.warehourse'

export default {
  name: 'wareHouse',
  data() {
    return {
      loading: false,
      noMore: false,
      shops: [],
      drawer: false,
      shoppingCar: []
    }
  },
  computed: {
    shopRows: function() {
      const rowNumber = 4
      return Math.ceil(this.shops.length / rowNumber)
    },
    // 购物车动画
    shopDrawer() {
      if (this.drawer) {
        return ['animated', 'bounceOutRight']
      } else {
        return ['animated', 'bounceInRight']
      }
    },
    shopDrawer2() {
      if (this.drawer) {
        return ['animated', 'bounceInRight']
      } else {
        return ['animated', 'bounceOutRight']
      }
    },
    totalPrice() {
      let total = 0
      for (let index = 0; index < this.shoppingCar.length; index++) {
        total += this.shoppingCar[index].price * this.shoppingCar[index].num
      }
      return total
    }
  },
  mounted: function() {
    window.addEventListener('scroll', this.handleScroll, true)
  },
  methods: {
    // 无限加载
    handleScroll() {
      // 变量windowHeight是可视区的高度
      const scrollTop = document.documentElement.scrollTop || document.body.scrollTop
      // 变量scrollHeight是滚动条的总高度
      const windowHeight = document.documentElement.clientHeight || document.body.clientHeight
      const scrollHeight = document.documentElement.scrollHeight || document.body.scrollHeight
      if (scrollTop + windowHeight === scrollHeight) {
        if (this.loading === false && this.noMore === false) {
          this.load()
        }
      }
    },
    load: function() {
      this.loading = true
      const pageNumber = 12
      const page = this.shops.length / pageNumber + 1
      // 开始请求获取商品
      GetShops({page: page, pageNumber: pageNumber})
        .then(async res => {
          this.shops.push.apply(this.shops, res)
          if (res.length < pageNumber) {
            this.noMore = true
          }
          this.loading = false
        })
        .catch(() => {
          this.loading = false
        })
    },
    // 购物车
    getShoppingCar() {
      return new Promise((resolve, reject) => {
        GetCar()
          .then(async res => {
            this.shoppingCar.splice(0, this.shoppingCar.length)
            this.shoppingCar = res
            resolve()
          })
          .catch(err => {
            reject(err)
          })
      })
    },
    addShopToCar(shop) {
      let exist = false
      for (let index = 0; index < this.shoppingCar.length; index++) {
        if (shop.shopName === this.shoppingCar[index].name) {
          this.shoppingCar[index].num++
          exist = true
        }
      }
      if (!exist) {
        this.shoppingCar.push({name: shop.shopName, price: shop.price, num: 1})
      }
    },
    editShoppingCar(index, row) {
      if (row.num === 0) {
        setTimeout(() => {
          this.shoppingCar.splice(index, 1)
        }, 5)
      }
      // console.log(index, row)
    },
    pay() {
      this.$confirm('确定要支付吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'info'
      })
        .then(() => {
          // 支付购物车
          DoOrder(this.shoppingCar)
            .then(async res => {
              this.getShoppingCar()
              this.$message({
                type: 'success',
                message: '支付已提交!'
              })
            })
            .catch(err => {
              this.$message({type: 'error', message: err})
            })
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '已取消支付'
          })
        })
    }
  },
  created() {
    this.load()
    this.getShoppingCar()
  },
  beforeDestroy() {
    // 保存购物车的数据
    // return new Promise((resolve, reject) => {
    SaveOrder(this.shoppingCar)
    // .then(async res => {
    // resolve()
    // })
    // .catch(err => {
    // reject(err)
    // })
    // })
  }
}
</script>

<style scoped></style>
