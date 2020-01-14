<template>
  <div>
    <!--        品牌宣传-->
    <div style="text-align: center; width: 100%; margin: 0px; padding: 0px;">
      <video loop muted data-autoplay style="min-width: 100%;">
        <source src="../../assets/home/brand/hello.mp4" type="video/mp4" />
      </video>
    </div>
    <!--    产品宣传展示-->
    <div style="margin:0px; padding:0px; min-width: 100%;">
      <div v-show="showShops" style="position: relative;">
        <el-row :gutter="20" :style="{position: 'absolute', top: '10%', width: '100%'}">
          <el-col v-for="(shop, index) in shops" :key="shop.id" :span="8">
            <div v-if="index <= 2" :class="getRandomAnimate()" style="width: 100%; text-align: center">
              <img style="max-height: 300px; max-width: 500px" :src="getShopSrc(shop)" />
            </div>
          </el-col>
        </el-row>
      </div>
      <div v-show="showShops" style="position: relative;">
        <el-row :gutter="20" :style="{position: 'absolute', top: '50%', width: '100%'}">
          <el-col v-for="(shop, index) in shops" :key="shop.id" :span="8">
            <div v-if="index > 2" :class="getRandomAnimate()" style="width: 100%; text-align: center">
              <img style="max-height: 300px; max-width: 500px" :src="getShopSrc(shop)" />
            </div>
          </el-col>
        </el-row>
      </div>
      <div>
        <img src="../../assets/warehouse/img1.jpeg" style="min-width:100%; margin: 0px; padding:0px;" />
      </div>
    </div>
    <!--    公司简介-->
    <div>
      <div v-show="showBrief" style="position: relative;">
        <div class="animated lightSpeedIn slower" :style="{position: 'absolute', top: '100px', left: '100px'}">
          <img style="max-height: 500px; max-width: 900px" :src="require('../../assets/warehouse/layer_0020.png')" />
        </div>
        <div
          class="animated fadeInUp slower delay-1s"
          :style="{position: 'absolute', top: '400px', left: '900px', width: '600px'}"
        >
          <p
            style="font-size: x-large; font-family: 宋体; color: rgba(176,161,149,0.81); background-color: rgba(68,92,116,0.45); border-radius: 25px; padding: 30px; box-shadow: 10px 10px 5px rgba(48,63,86,0.75)"
          >
            华为创立于1987年，是全球领先的ICT（信息与通信）基础设施和智能终端提供商，我们致力于把数字世界带入每个人、每个家庭、每个组织，构建万物互联的智能世界。目前华为有19.4万员工，业务遍及170多个国家和地区，服务30多亿人口。
            我们在通信网络、IT、智能终端和云服务等领域为客户提供有竞争力、安全可信赖的产品、解决方案与服务，与生态伙伴开放合作，持续为客户创造价值，释放个人潜能，丰富家庭生活，激发组织创新。华为坚持围绕客户需求持续创新，加大基础研究投入，厚积薄发，推动世界进步。
          </p>
        </div>
      </div>
      <div style="width: 100%; height: 100%; text-align: center;">
        <img style="height: 100%; text-align: center" src="../../assets/home/brief.jpg" />
      </div>
    </div>
    <!--    企业文化照片墙-->
    <div>
      <div class="cultureBackground">
        <grid-layout
          style="margin-top: 65px; margin-bottom: 65px; margin-left: 160px; margin-right: 160px; width: 1580px"
          v-bind="layout"
        >
          <grid-item v-for="(gItem, index) in layout.layout" :key="index" v-bind="gItem">
            <template
              v-if="gItem.i === '0' || gItem.i === '2' || gItem.i === '6' || gItem.i === '8'"
              style="width: 100%; height: 100%;"
            >
              <grid-layout style="width: 100%; height: 100%;" v-bind="layoutItems">
                <grid-item v-for="(item, index) in layoutItems.layout" :key="index" v-bind="item">
                  <div style="text-align: center;">
                    <img
                      style="height: 111px;"
                      :src="require('../../assets/home/culture/' + gItem.title + item.i + '.jpg')"
                    />
                  </div>
                </grid-item>
              </grid-layout>
            </template>
            <template v-else-if="gItem.i === '4'">
              <grid-layout style="width: 100%; height: 100%;" v-bind="layoutTitles">
                <grid-item v-for="(item, index) in layoutTitles.layout" :key="index" v-bind="item">
                  <div style="text-align: center;">
                    <img style="height: 177px;" :src="item.src" />
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
    <contact></contact>
  </div>
</template>

<style scoped>
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
  background: url(../../assets/home/culture/background.jpg);
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
import '@/plugins/js/animate.js'
import VueGridLayout from 'vue-grid-layout'
import Contact from '../../components/contact/Contact.vue'

export default {
  data() {
    return {
      // 货仓 数据
      shops: [
        {id: '0000', src: 'layer_0000.png'},
        {id: '0001', src: 'layer_0001.png'},
        {id: '0002', src: 'layer_0002.png'},
        {id: '0006', src: 'layer_0006.png'},
        {id: '00022', src: 'layer_0022.png'},
        {id: '00023', src: 'layer_0023.png'}
      ], // get from api
      animates: [
        'bounceIn',
        'bounceInDown',
        'bounceInLeft',
        'bounceInRight',
        'bounceInUp',
        'fadeIn',
        'fadeInDown',
        'fadeInDownBig',
        'fadeInLeft',
        'fadeInLeftBig',
        'fadeInRight',
        'fadeInRightBig',
        'fadeInUp',
        'fadeInUpBig',
        'flipInX',
        'flipInY',
        'lightSpeedIn',
        'rotateIn',
        'rotateInDownLeft',
        'rotateInDownRight',
        'rotateInUpLeft',
        'rotateInUpRight',
        'hinge',
        'jackInTheBox',
        'rollIn',
        'zoomIn',
        'zoomInDown',
        'zoomInLeft',
        'zoomInRight',
        'zoomInUp',
        'slideInDown',
        'slideInLeft',
        'slideInRight',
        'slideInUp'
      ],
      showShops: false,
      // 文化墙
      layout: {
        layout: [
          {x: 0, y: 0, w: 4, h: 6, i: '0', title: 'attitude'}, // 态度
          {x: 4, y: 0, w: 6, h: 4, i: '1'}, // 留白
          {x: 10, y: 0, w: 4, h: 6, i: '2', title: 'quality'}, // 品质
          {x: 0, y: 4, w: 4, h: 2, i: '3'}, // 留白
          {x: 4, y: 4, w: 6, h: 6, i: '4'}, // 观点
          {x: 10, y: 4, w: 4, h: 2, i: '5'}, // 留白
          {x: 0, y: 10, w: 4, h: 6, i: '6', title: 'details'}, // 细心
          {x: 4, y: 10, w: 6, h: 4, i: '7'}, // 留白
          {x: 10, y: 10, w: 4, h: 6, i: '8', title: 'responsibility'} // 责任
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
          {x: 0, y: 0, w: 1, h: 2, i: '1'},
          {x: 1, y: 0, w: 1, h: 2, i: '2'},
          {x: 0, y: 1, w: 1, h: 2, i: '3'},
          {x: 1, y: 1, w: 1, h: 2, i: '4'}
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
          {x: 0, y: 0, w: 1, h: 1, i: '0', src: require('../../assets/home/culture/culture1.jpg')},
          {x: 1, y: 0, w: 1, h: 1, i: '1', src: require('../../assets/home/culture/culture2.jpg')},
          {x: 0, y: 1, w: 1, h: 1, i: '2', src: require('../../assets/home/culture/culture3.jpg')},
          {x: 1, y: 1, w: 1, h: 1, i: '3', src: require('../../assets/home/culture/culture4.jpg')}
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
  methods: {
    // element 导航
    handleSelect(key, keyPath) {
      this.activeIndex = key
    },
    // 自定义
    getShopSrc: function(shop) {
      return require('../../assets/warehouse/' + shop.src)
    },
    getRandomAnimate: function() {
      const randomCss = ['animated', 'delay-1s']
      randomCss.push(this.animates[Math.floor(Math.random() * parseInt(this.animates.length))])
      return randomCss
    }
  }
}
</script>
