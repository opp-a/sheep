<template>
  <div class="block">
    <el-timeline>
      <el-timeline-item
        v-for="(activity, index) in activities"
        :key="index"
        icon="el-icon-more"
        type="primary"
        color="#0bbd87"
        size="large"
        :timestamp="formattime(activity.addtime)"
        placement="top"
      >
        <el-card>
          <el-table :data="activity.myshops" style="width: 100%">
            <el-table-column prop="name" label="商品" width="180"> </el-table-column>
            <el-table-column prop="priceout" label="价格" width="180"> </el-table-column>
            <el-table-column prop="num" label="数量"> </el-table-column>
          </el-table>
          <p style="font-size: 16px;">总价 {{ activity.pricetotal }}</p></el-card
        >
      </el-timeline-item>
    </el-timeline>
    <p v-if="loading" style="text-align: center; color: #8c939d;">加载中...</p>
    <p v-if="noMore" style="text-align: center; color: #8c939d;">没有更多了</p>
  </div>
</template>

<script>
import {GetMyOrders} from '@/api/sys.myorder'
export default {
  name: 'myorder',
  data() {
    return {
      loading: false,
      noMore: false,
      activities: []
    }
  },
  methods: {
    formattime(paramstime) {
      return this.$moment(new Date(paramstime)).format('YYYY-MM-DD HH:mm:ss')
    },
    loadOrders() {
      this.loading = true
      const pageNumber = 12
      const page = Math.floor(this.activities.length / pageNumber) + 1

      // 获取历史订单
      GetMyOrders({page: page, pageNumber: pageNumber})
        .then(async res => {
          this.activities.push.apply(this.activities, res.infos)
          if (res.length < pageNumber) {
            this.noMore = true
          }
          this.loading = false
        })
        .catch(() => {
          this.loading = false
        })
    },
    // 无限加载
    handleScroll() {
      // 变量windowHeight是可视区的高度
      const scrollTop = document.documentElement.scrollTop || document.body.scrollTop
      // 变量scrollHeight是滚动条的总高度
      const windowHeight = document.documentElement.clientHeight || document.body.clientHeight
      const scrollHeight = document.documentElement.scrollHeight || document.body.scrollHeight
      if (scrollTop + windowHeight === scrollHeight) {
        if (this.loading === false && this.noMore === false) {
          this.loadOrders()
        }
      }
    }
  },
  mounted: function() {
    window.addEventListener('scroll', this.handleScroll, true)
  },
  created() {
    this.loadOrders()
  }
}
</script>

<style scoped></style>
