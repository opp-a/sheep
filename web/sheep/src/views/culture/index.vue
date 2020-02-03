<template>
  <div>
    <div style="height: 20px; width: 100%;"></div>
    <el-timeline>
      <el-timeline-item
        v-for="itemArray in culturePictures"
        :key="itemArray.time"
        :timestamp="itemArray.time"
        placement="top"
      >
        <div
          style="display: flex; flex-direction: row; flex-wrap: wrap; justify-content: flex-start; align-items: flex-start; align-content: flex-start"
        >
          <div v-for="item in itemArray.pictures" :key="item.time" style="flex: none; margin: 10px;">
            <img
              :src="item.src"
              style="margin: 0px; padding: 0px; height: 100%; width: 100%; box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1)"
            />
            <div style="width: 100%; text-align: center">{{ item.name }}</div>
          </div>
        </div>
      </el-timeline-item>
    </el-timeline>
    <p v-if="loading" style="text-align: center; color: #8c939d;">加载中...</p>
    <p v-if="noMore" style="text-align: center; color: #8c939d;">没有更多了</p>
  </div>
</template>

<script>
import {GetCulturePictures} from '@/api/sys.culture'
export default {
  name: 'Culture',
  data() {
    return {
      loading: false,
      noMore: false,
      culturePictures: [
        {
          time: '2019-12-19',
          pictures: [
            {
              name: 'picture-1',
              src: 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png',
              time: '2019-12-19'
            },
            {
              name: 'picture-2',
              src: 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png',
              price: '2019-12-19'
            },
            {
              name: 'picture-3',
              src: 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png',
              price: '2019-12-19'
            }
          ]
        },
        {
          time: '2019-12-17',
          pictures: [
            {
              name: 'picture-4',
              src: 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png',
              price: '2019-12-17'
            },
            {
              name: 'picture-5',
              src: 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png',
              price: '2019-12-17'
            }
          ]
        },
        {
          time: '2019-12-16',
          pictures: [
            {
              name: 'picture-6',
              src: 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png',
              price: '2019-12-16'
            },
            {
              name: 'picture-7',
              src: 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png',
              price: '2019-12-16'
            },
            {
              name: 'picture-8',
              src: 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png',
              price: '2019-12-16'
            },
            {
              name: 'picture-9',
              src: 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png',
              price: '2019-12-16'
            },
            {
              name: 'picture-10',
              src: 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png',
              price: '2019-12-16'
            },
            {
              name: 'picture-11',
              src: 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png',
              price: '2019-12-16'
            }
          ]
        }
      ]
    }
  },
  methods: {
    getPictures() {
      this.loading = true
      const pageNumber = 12
      const page = this.culturePictures.length / pageNumber + 1

      // 获取文化墙的图片
      GetCulturePictures({page: page, pageNumber: pageNumber})
        .then(async res => {
          this.culturePictures.push.apply(this.culturePictures, res)
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
          this.getPictures()
        }
      }
    }
  },
  mounted: function() {
    window.addEventListener('scroll', this.handleScroll, true)
  },
  created() {
    this.getPictures()
  }
}
</script>

<style scoped></style>
