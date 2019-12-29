<template>
  <el-container>
    <!--    内容区域-->
    <div>
      <full-page ref="fullpage" :options="options">
        <!--        品牌宣传-->
        <div class="section" style="text-align: center">
          <video loop muted data-autoplay class="home-video">
            <source src="../assets/home/brand/hello.mp4" type="video/mp4">
          </video>
        </div>
        <!--    产品宣传展示-->
        <div class="section">
          <div class="slide" data-anchor="slide0">
            <div style="position: relative; top: 26em;  z-index:99;">
              <div @click="customMoveLeft()" style="left: 10px; position: absolute;">
                <i class=" fa fa-caret-left" aria-hidden="true"
                   style="color: rgba(255, 255, 255, 0.1);font-size: 11em;"></i>
              </div>
              <div @click="customMoveRight()" style="right: 10px; position: absolute;">
                <i class=" fa fa-caret-right" aria-hidden="true"
                   style="color: rgba(255, 255, 255, 0.1);font-size: 11em;"></i>
              </div>
            </div>
            <div v-show="showShops" style="position: relative; top: 2em;">
              <el-row :gutter="20" :style="{position: 'absolute', top: '10%', width: '100%'}">
                <el-col v-for="(shop, index) in shops" :key="shop.id" :span="8">
                  <div v-if="index<=2" :class="getRandomAnimate()" style="width: 100%; text-align: center"><img
                    style="max-height: 300px; max-width: 500px" :src='getShopSrc(shop)'></div>
                </el-col>
              </el-row>
            </div>
            <div v-show="showShops" style="position: relative; top: 30em;">
              <el-row :gutter="20" :style="{position: 'absolute', top: '50%', width: '100%', }">
                <el-col v-for="(shop, index) in shops" :key="shop.id" :span="8">
                  <div v-if="index>2" :class="getRandomAnimate()" style="width: 100%; text-align: center"><img
                    style="max-height: 300px; max-width: 500px" :src='getShopSrc(shop)'></div>
                </el-col>
              </el-row>
            </div>
            <img src="../assets/warehouse/img1.jpeg">
          </div>
          <div class="slide" data-anchor="slide1" style="position: relative">
            <div style="position: relative; top: 26em;">
              <div @click="customMoveLeft()" style="left: 10px; position: absolute;">
                <i class=" fa fa-caret-left" aria-hidden="true"
                   style="color: rgba(255, 255, 255, 0.1);font-size: 11em;"></i>
              </div>
              <div @click="customMoveRight()" style="right: 10px; position: absolute;">
                <i class=" fa fa-caret-right" aria-hidden="true"
                   style="color: rgba(255, 255, 255, 0.1);font-size: 11em;"></i>
              </div>
            </div>
            <div v-show="showShops" style="position: relative;">
              <div :class="getRandomAnimate()"
                   v-for="shop in shops" :key="shop.id"
                   :style="{position: 'absolute', top: getScreenHeigh() + 'px', left: getScreenWidth() + 'px'}">
                <img style="max-height: 300px; max-width: 500px" :src='getShopSrc(shop)'>
              </div>
            </div>
            <img src="../assets/warehouse/img2.jpeg">
          </div>
          <div class="slide" data-anchor="slide2" style="position: relative">
            <div style="position: relative; top: 26em;">
              <div @click="customMoveLeft()" style="left: 10px; position: absolute;">
                <i class=" fa fa-caret-left" aria-hidden="true"
                   style="color: rgba(255, 255, 255, 0.1);font-size: 11em;"></i>
              </div>
              <div @click="customMoveRight()" style="right: 10px; position: absolute;">
                <i class=" fa fa-caret-right" aria-hidden="true"
                   style="color: rgba(255, 255, 255, 0.1);font-size: 11em;"></i>
              </div>
            </div>
            <div v-show="showShops" style="position: relative;">
              <div :class="getRandomAnimate()"
                   v-for="shop in shops" :key="shop.id"
                   :style="{position: 'absolute', top: getScreenHeigh() + 'px', left: getScreenWidth() + 'px'}">
                <img style="max-height: 300px; max-width: 500px" :src='getShopSrc(shop)'>
              </div>
            </div>
            <img src="../assets/warehouse/img3.jpeg">
          </div>
        </div>
        <!--    公司简介-->
        <div class="section">
          <div v-show="showBrief" style="position: relative;">
            <div class="animated lightSpeedIn slower" :style="{position: 'absolute', top: '100px', left: '100px'}">
              <img style="max-height: 500px; max-width: 900px" :src="require('../assets/warehouse/layer_0020.png')">
            </div>
            <div class="animated fadeInUp slower delay-1s"
                 :style="{position: 'absolute', top: '400px', left: '900px', width: '600px'}">
              <p
                style="font-size: x-large; font-family: 宋体; color: rgba(176,161,149,0.81); background-color: rgba(68,92,116,0.45); border-radius: 25px; padding: 30px; box-shadow: 10px 10px 5px rgba(48,63,86,0.75)">
                华为创立于1987年，是全球领先的ICT（信息与通信）基础设施和智能终端提供商，我们致力于把数字世界带入每个人、每个家庭、每个组织，构建万物互联的智能世界。目前华为有19.4万员工，业务遍及170多个国家和地区，服务30多亿人口。
                我们在通信网络、IT、智能终端和云服务等领域为客户提供有竞争力、安全可信赖的产品、解决方案与服务，与生态伙伴开放合作，持续为客户创造价值，释放个人潜能，丰富家庭生活，激发组织创新。华为坚持围绕客户需求持续创新，加大基础研究投入，厚积薄发，推动世界进步。
              </p>
            </div>
          </div>
          <div style="width: 100%; height: 100%; text-align: center;">
            <img style="height: 100%; text-align: center" src="../assets/home/brief.jpg">
          </div>
        </div>
        <!--    企业文化照片墙-->
        <div class="section">
          <div class="cultureBackground">
            <grid-layout
              style="margin-top: 65px; margin-bottom: 65px; margin-left: 160px; margin-right: 160px; width: 1580px"
              v-bind="layout">
              <grid-item
                v-for="(gItem, index) in layout.layout"
                :key="index"
                v-bind="gItem">
                <template v-if="gItem.i === '0' || gItem.i === '2' || gItem.i === '6'|| gItem.i === '8'"
                          style="width: 100%; height: 100%;">
                  <grid-layout style="width: 100%; height: 100%;"
                               v-bind="layoutItems">
                    <grid-item
                      v-for="(item, index) in layoutItems.layout"
                      :key="index"
                      v-bind="item">
                      <div style="text-align: center;">
                        <img style="height: 111px;" :src="require('../assets/home/culture/'+gItem.title+item.i+'.jpg')">
                      </div>
                    </grid-item>
                  </grid-layout>
                </template>
                <template v-else-if="gItem.i === '4'">
                  <grid-layout style="width: 100%; height: 100%;"
                               v-bind="layoutTitles">
                    <grid-item
                      v-for="(item, index) in layoutTitles.layout"
                      :key="index"
                      v-bind="item">
                      <div style="text-align: center;">
                        <img style="height: 177px;" :src="item.src">
                      </div>
                    </grid-item>
                  </grid-layout>
                </template>
                <template v-else>
                  <div style="height: 100%; width: 100%;"></div>
                </template>
              </grid-item>
            </grid-layout>
          </div>
        </div>
        <!--    联系方式-->
        <div class="section">
          <contact></contact>
        </div>
      </full-page>
    </div>

    <!--    头 -->
    <div class="header">
      <el-header>
        <el-menu default-active="home" mode="horizontal" @select="handleSelect" router
                 style="border-bottom: none;" background-color="rgba(255, 255, 255, 0)" text-color="#A0A0A0"
                 active-text-color="#51EE48">
          <el-menu-item index="home">首页</el-menu-item>
          <el-menu-item index="warehouse">货仓</el-menu-item>
          <el-menu-item index="culture"> 文化墙</el-menu-item>
          <el-submenu style="text-align: center; float: right;" index="self">
            <template slot="title">{{this.$store.state.userName}}</template>
            <el-menu-item index="self" style="text-align: left"><i class="el-icon-user"></i>个人中心</el-menu-item>
            <el-menu-item index="login" style="text-align: left"><i class="el-icon-top-left"></i> 退出
            </el-menu-item>
          </el-submenu>
        </el-menu>
      </el-header>
    </div>

    <!--      尾 -->
    <el-footer style=" color: rgba(100, 100, 100, 0.5); text-align: center; line-height: 60px; width: 100%;
    background: rgba(53, 73, 94, 0.0); position: fixed; bottom: 0px;">你是最棒的！ Little Sheep!!!</el-footer>
  </el-container>
</template>

<style scoped>
/* 导航头尾区域 */
.header {
  color: #FFFFFF;
  text-align: center;
  line-height: 60px;
  width: 100%;
  background: rgba(255, 255, 255, 0);
  position: fixed;
  top: 0px;
}

/* 首页宣传视频样式*/
.home-video {
  z-index: 100;
  top: 50%;
  left: 50%;
  /*min-width: 100%;*/
  min-height: 100%;
  object-fit: fill; /*这里是关键*/
  width: auto;
  height: auto;
}

/*模糊背景*/
.cultureBackground {
  /*background*/
  width: 100%;
  height: 100%;
  background: hsla(0, 0%, 60%, 0.5);
  background-size: cover;
  overflow: hidden;
  position: relative;
  margin: 0px;
  padding: 0px;
}

.cultureBackground::before {
  position: absolute;
  background: url(../assets/home/culture/background.jpg);
  background-size: cover;
  width: 100%;
  height: 100%;
  content: '';
  filter: blur(12px);
  z-index: 0;
  margin: 0px;
  padding: 0px;
}
</style>

<script>
import 'font-awesome/css/font-awesome.min.css'
import '@/plugins/js/fullpage.js'
import '@/plugins/js/animate.js'
import VueGridLayout from 'vue-grid-layout'
import Contact from '../components/contact/Contact.vue'
import { getUserCount } from '../api/getData'

export default {
  data () {
    return {
      // fullpage params
      options: {
        afterLoad: this.afterLoad,
        anchors: ['brand', 'warehouse', 'brief', 'culture', 'contact'],
        sectionsColor: ['#000000', '#000000', '#303c52', '#000000', '#000000'],
        slidesNavigation: true,
        slidesNavPosition: 'bottom',
        controlArrows: false,
        recordHistory: false,
        parallax: false,
        lockAnchors: true,
        resetSliders: true,
        parallaxOptions: { type: 'reveal', percentage: 62, property: 'translate' }
      },
      // 货仓 数据
      shops: [
        { id: '0000', src: 'layer_0000.png' },
        { id: '0001', src: 'layer_0001.png' },
        { id: '0002', src: 'layer_0002.png' },
        { id: '0006', src: 'layer_0006.png' },
        { id: '00022', src: 'layer_0022.png' },
        { id: '00023', src: 'layer_0023.png' }
      ], // get from api
      animates: [
        'bounceIn', 'bounceInDown', 'bounceInLeft', 'bounceInRight', 'bounceInUp',
        'fadeIn', 'fadeInDown', 'fadeInDownBig', 'fadeInLeft', 'fadeInLeftBig', 'fadeInRight', 'fadeInRightBig', 'fadeInUp', 'fadeInUpBig',
        'flipInX', 'flipInY', 'lightSpeedIn', 'rotateIn', 'rotateInDownLeft', 'rotateInDownRight', 'rotateInUpLeft', 'rotateInUpRight',
        'hinge', 'jackInTheBox', 'rollIn', 'zoomIn', 'zoomInDown', 'zoomInLeft', 'zoomInRight', 'zoomInUp',
        'slideInDown', 'slideInLeft', 'slideInRight', 'slideInUp'
      ],
      showShops: false,
      // 文化墙
      layout: {
        layout: [
          { 'x': 0, 'y': 0, 'w': 4, 'h': 6, 'i': '0', 'title': 'attitude' }, // 态度
          { 'x': 4, 'y': 0, 'w': 6, 'h': 4, 'i': '1' }, // 留白
          { 'x': 10, 'y': 0, 'w': 4, 'h': 6, 'i': '2', 'title': 'quality' }, // 品质
          { 'x': 0, 'y': 4, 'w': 4, 'h': 2, 'i': '3' }, // 留白
          { 'x': 4, 'y': 4, 'w': 6, 'h': 6, 'i': '4' }, // 观点
          { 'x': 10, 'y': 4, 'w': 4, 'h': 2, 'i': '5' }, // 留白
          { 'x': 0, 'y': 10, 'w': 4, 'h': 6, 'i': '6', 'title': 'details' }, // 细心
          { 'x': 4, 'y': 10, 'w': 6, 'h': 4, 'i': '7' }, // 留白
          { 'x': 10, 'y': 10, 'w': 4, 'h': 6, 'i': '8', 'title': 'responsibility' } // 责任
        ],
        colNum: 14,
        maxRows: 14,
        rowHeight: 60,
        isDraggable: false,
        isResizable: false,
        isMirrored: false,
        autoSize: false,
        verticalCompact: true,
        margin: [0, 0],
        useCssTransforms: true
      },
      layoutItems: {
        layout: [
          { 'x': 0, 'y': 0, 'w': 1, 'h': 2, 'i': '1' },
          { 'x': 1, 'y': 0, 'w': 1, 'h': 2, 'i': '2' },
          { 'x': 0, 'y': 1, 'w': 1, 'h': 2, 'i': '3' },
          { 'x': 1, 'y': 1, 'w': 1, 'h': 2, 'i': '4' }
        ],
        colNum: 2,
        maxRows: 2,
        rowHeight: 111,
        isDraggable: false,
        isResizable: false,
        isMirrored: false,
        autoSize: false,
        verticalCompact: false,
        margin: [6, 6],
        useCssTransforms: true
      },
      layoutTitles: {
        layout: [
          { 'x': 0, 'y': 0, 'w': 1, 'h': 1, 'i': '0', src: require('../assets/home/culture/culture1.jpg') },
          { 'x': 1, 'y': 0, 'w': 1, 'h': 1, 'i': '1', src: require('../assets/home/culture/culture2.jpg') },
          { 'x': 0, 'y': 1, 'w': 1, 'h': 1, 'i': '2', src: require('../assets/home/culture/culture3.jpg') },
          { 'x': 1, 'y': 1, 'w': 1, 'h': 1, 'i': '3', src: require('../assets/home/culture/culture4.jpg') }
        ],
        colNum: 2,
        maxRows: 2,
        rowHeight: 177,
        isDraggable: false,
        isResizable: false,
        isMirrored: false,
        autoSize: false,
        verticalCompact: false,
        margin: [2, 2],
        useCssTransforms: true
      },
      // 简介
      showBrief: false
    }
  },
  components: {
    GridLayout: VueGridLayout.GridLayout,
    GridItem: VueGridLayout.GridItem,
    Contact: Contact
  },
  computed: {
    screenWidth: function () {
      return (document.documentElement.clientWidth || document.body.clientWidth)
    },
    screenHeight: function () {
      return (document.documentElement.clientHeight || document.body.clientHeight)
    }
  },
  methods: {
    // element 导航
    handleSelect (key, keyPath) {
      this.activeIndex = key
    },
    // 自定义
    getShopSrc: function (shop) {
      return require('../assets/warehouse/' + shop.src)
    },
    getScreenWidth: function () {
      return Math.floor(Math.random() * ((document.documentElement.clientWidth || document.body.clientWidth) - 500))
    },
    getScreenHeigh: function () {
      return Math.floor(Math.random() * ((document.documentElement.clientHeight || document.body.clientHeight) - 300))
    },
    getRandomAnimate: function () {
      let randomCss = ['animated', 'delay-1s']
      randomCss.push(this.animates[Math.floor(Math.random() * parseInt(this.animates.length))])
      return randomCss
    },
    customMoveLeft: function () {
      this.$refs.fullpage.api.moveSlideLeft()
      this.showShops = false
      this.showShops = true
    },
    customMoveRight: function () {
      this.$refs.fullpage.api.moveSlideRight()
      this.showShops = false
      this.showShops = true
    },
    // fullpage 回调
    afterLoad: function (origin, destination, direction) {
      if (destination.anchor === 'warehouse' && this.showShops === false) {
        this.showShops = true
      } else if (destination.anchor !== 'warehouse' && this.showShops === true) {
        this.showShops = false
      }

      if (destination.anchor === 'brief' && this.showBrief === false) {
        this.showBrief = true
      } else if (destination.anchor !== 'brief' && this.showBrief === true) {
        this.showBrief = false
      }
    },
    // api for demo
    async initData () {
      try {
        const countData = await getUserCount()
        if (countData.status === 1) {
          this.count = countData.count
        } else {
          throw new Error('获取数据失败')
        }
        this.getUsers()
      } catch (err) {
        console.log('获取数据失败', err)
      }
    },
    async getUsers () {
      const Users = await getUserCount(para).then((res) => {
        this.addLoading = false
        this.$message({
          message: '提交成功',
          type: 'success'
        })
        this.$refs['addForm'].resetFields()
        this.addFormVisible = false
        this.getUsers()
      }).catch(() => {

      })

      // getUserList({ offset: this.offset, limit: this.limit })
      // this.tableData = []
      // Users.forEach(item => {
      //   const tableData = {}
      //   tableData.username = item.username
      //   tableData.registe_time = item.registe_time
      //   tableData.city = item.city
      //   this.tableData.push(tableData)
      // })
    }
  }
}

</script>
