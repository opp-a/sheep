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
        :timestamp="activity.timestamp"
        placement="top"
      >
        <el-card> {{ activity.content }}</el-card>
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
      activities: [
        {
          content: '支持使用图标',
          timestamp: '2018-04-12 20:46'
        },
        {
          content: '支持自定义颜色',
          timestamp: '2018-04-03 20:46'
        },
        {
          content: '支持自定义尺寸',
          timestamp: '2018-04-03 20:46'
        },
        {
          content: '默认样式的节点',
          timestamp: '2018-04-03 20:46'
        }
      ]
    }
  },
  methods: {
    loadOrders() {
      this.loading = true
      const pageNumber = 12
      const page = this.activities.length / pageNumber + 1

      // 获取历史订单
      GetMyOrders({page: page, pageNumber: pageNumber})
        .then(async res => {
          this.activities.push.apply(this.activities, res)
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
