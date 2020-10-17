<template>
  <div class="contain" ref="wrapper">
    <!--        品牌宣传-->
    <div style="text-align: center; background-color: black;">
      <video loop muted autoplay src="./videos/hello.mp4" height="100%" width="100%"></video>
    </div>
    <!--    产品宣传展示-->
    <div id="shops" class="shopsBackground">
      <grid-layout v-bind="layoutShops" v-show="showShops" style="height: 100%;">
        <grid-item v-for="(item, index) in layoutShops.layout" :key="index" v-bind="item" style="display: flex;justify-content: center;align-items: center;flex-direction: column;">
          <div class="animated fadeInLeft slower">
            <img :src="require('./images/' + item.src)" style=" max-width: 200px;" />
          </div>
        </grid-item>
      </grid-layout>
    </div>
    <!--    公司简介-->
    <div id="brief" class="briefBackground">
      <div v-show="showBrief" style="position: relative;">
        <div style="position: absolute; position: relative;">
          <div class="animated lightSpeedIn slower" style="position: absolute; top: 60px; left: 8%">
            <img style="max-height: 500px; max-width: 900px" :src="require('./images/layer_0000.png')" />
          </div>
          <div class="animated fadeInUp slower delay-1s" style="position: absolute; top: 300px; left: 50%; width: 33%; min-width: 400px;">
            <p
              style="font-size: x-large; font-family: 宋体; color: rgba(176,161,149,0.81); background-color: rgba(68,92,116,0.45); border-radius: 25px; padding: 30px; box-shadow: 10px 10px 5px rgba(48,63,86,0.75)"
            >
              华为创立于1987年，是全球领先的ICT（信息与通信）基础设施和智能终端提供商，我们致力于把数字世界带入每个人、每个家庭、每个组织，构建万物互联的智能世界。目前华为有19.4万员工，业务遍及170多个国家和地区，服务30多亿人口。
              我们在通信网络、IT、智能终端和云服务等领域为客户提供有竞争力、安全可信赖的产品、解决方案与服务，与生态伙伴开放合作，持续为客户创造价值，释放个人潜能，丰富家庭生活，激发组织创新。华为坚持围绕客户需求持续创新，加大基础研究投入，厚积薄发，推动世界进步。
            </p>
          </div>
        </div>
      </div>
    </div>
    <!--    企业文化照片墙-->
    <div>
      <div id="culture" class="bg bg-blur"></div>
      <grid-layout v-bind="layoutCuture">
        <grid-item v-for="(item, index) in layoutCuture.layout" :key="index" v-bind="item">
          <div style="display: flex;justify-content: center;align-items: center;flex-direction: column; height: 100%; width: 100%;">
            <img :class="item.myclass" :src="require('./images/culture/' + item.src)" style="max-width: 100%; max-height: 100%;" />
          </div>
        </grid-item>
      </grid-layout>
    </div>
    <!--    联系方式-->
    <contact></contact>
  </div>
</template>

<style scoped>
/*隐藏滚动条*/
.contain::-webkit-scrollbar {
  width: 0 !important;
  overflow: hidden;
  overflow-x: inherit;
}

/*商品背景*/
.shopsBackground {
  /*background*/
  min-width: 100%;
  background: url(./images/bg1.jpeg);
  background-repeat: no-repeat;
  background-position-x: center;
  background-position-y: center;
}

/*简介背景*/
.briefBackground {
  /*background*/
  min-width: 100%;
  height: 100%;
  background: url(./images/bg2.jpg);
  background-repeat: no-repeat;
  background-size: cover;
  background-position-x: center;
  background-position-y: center;
}

/*文化背景*/
.bg {
  background: url(./images/bg3.jpg);
  /* height:600px; */
  text-align: center;
  /* line-height: 600px; */
}
.bg-blur {
  float: left;
  width: 100%;
  background-repeat: no-repeat;
  background-position: center;
  background-size: cover;
  -webkit-filter: blur(15px);
  -moz-filter: blur(15px);
  -o-filter: blur(15px);
  -ms-filter: blur(15px);
  filter: blur(15px);
}
</style>

<script>
import 'font-awesome/css/font-awesome.min.css'
import VueGridLayout from 'vue-grid-layout'
import '@/plugins/js/animate.js'
import scroll from '@/components/mixins/normal'
import Contact from './components/contact/Contact.vue'

export default {
  mixins: [scroll],
  data() {
    return {
      // 货仓 数据
      layoutShops: {
        layout: [
          {x: 1, y: 0, w: 2, h: 1, i: '0', src: 'layer_0000.png'},
          {x: 5, y: 0, w: 2, h: 1, i: '1', src: 'layer_0001.png'},
          {x: 9, y: 0, w: 2, h: 1, i: '2', src: 'layer_0002.png'},
          {x: 1, y: 1, w: 2, h: 1, i: '3', src: 'layer_0006.png'},
          {x: 5, y: 1, w: 2, h: 1, i: '4', src: 'layer_0022.png'},
          {x: 9, y: 1, w: 2, h: 1, i: '5', src: 'layer_0023.png'}
        ],
        rowHeight: (document.documentElement.clientHeight - 20) / 2,
        isDraggable: false,
        isResizable: false,
        autoSize: false
      },
      showShops: false,
      // 文化墙
      layoutCuture: {
        layout: [
          {x: 0, y: 0, w: 3, h: 1, i: '0', src: 'culture1.jpg', myclass: ['animated', 'bounceInDown']},
          {x: 3, y: 0, w: 2, h: 1, i: '1', src: 'attitude1.jpg', myclass: ['animated', 'bounceInDown']},
          {x: 5, y: 0, w: 2, h: 1, i: '2', src: 'attitude2.jpg', myclass: ['animated', 'bounceInDown']},
          {x: 7, y: 0, w: 2, h: 1, i: '3', src: 'attitude3.jpg', myclass: ['animated', 'bounceInDown']},
          {x: 9, y: 0, w: 2, h: 1, i: '4', src: 'attitude4.jpg', myclass: ['animated', 'bounceInDown']},
          {x: 0, y: 1, w: 3, h: 1, i: '5', src: 'culture2.jpg', myclass: ['animated', 'bounceInDown']},
          {x: 3, y: 1, w: 2, h: 1, i: '6', src: 'quality1.jpg', myclass: ['animated', 'bounceInDown']},
          {x: 5, y: 1, w: 2, h: 1, i: '7', src: 'quality2.jpg', myclass: ['animated', 'bounceInDown']},
          {x: 7, y: 1, w: 2, h: 1, i: '8', src: 'quality3.jpg', myclass: ['animated', 'bounceInDown']},
          {x: 9, y: 1, w: 2, h: 1, i: '9', src: 'quality4.jpg', myclass: ['animated', 'bounceInDown']},
          {x: 0, y: 2, w: 3, h: 1, i: '10', src: 'culture3.jpg', myclass: ['animated', 'bounceInDown']},
          {x: 3, y: 2, w: 2, h: 1, i: '11', src: 'details1.jpg', myclass: ['animated', 'bounceInDown']},
          {x: 5, y: 2, w: 2, h: 1, i: '12', src: 'details2.jpg', myclass: ['animated', 'bounceInDown']},
          {x: 7, y: 2, w: 2, h: 1, i: '13', src: 'details3.jpg', myclass: ['animated', 'bounceInDown']},
          {x: 9, y: 2, w: 2, h: 1, i: '14', src: 'details4.jpg', myclass: ['animated', 'bounceInDown']},
          {x: 0, y: 3, w: 3, h: 1, i: '15', src: 'culture4.jpg', myclass: ['animated', 'bounceInDown']},
          {x: 3, y: 3, w: 2, h: 1, i: '16', src: 'responsibility1.jpg', myclass: ['animated', 'bounceInDown']},
          {x: 5, y: 3, w: 2, h: 1, i: '17', src: 'responsibility2.jpg', myclass: ['animated', 'bounceInDown']},
          {x: 7, y: 3, w: 2, h: 1, i: '18', src: 'responsibility3.jpg', myclass: ['animated', 'bounceInDown']},
          {x: 9, y: 3, w: 2, h: 1, i: '19', src: 'responsibility4.jpg', myclass: ['animated', 'bounceInDown']}
        ],
        colNum: 11,
        rowHeight: (document.documentElement.clientHeight - 50) / 4,
        isDraggable: false,
        isResizable: false,
        autoSize: false
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
  mounted() {
    // 增加滚动事件监听
    this.addScrollListener()
    window.addEventListener('scroll', this.myHandleScroll, true) // 监听（绑定）滚轮滚动事件

    // 设置元素屏高
    const height = document.documentElement.clientHeight || document.body.clientHeight
    document.getElementById('shops').style.height = height + 'px'
    document.getElementById('culture').style.height = height + 'px'
    document.getElementById('brief').style.height = height + 'px'
  },

  methods: {
    myHandleScroll: function(event) {
      const objShops = document.getElementById('shops')
      if (window.scrollY > objShops.offsetTop - 200 && window.scrollY < objShops.offsetTop + objShops.offsetHeight && this.showShops === false) {
        this.showShops = true
      } else if ((window.scrollY > objShops.offsetTop + objShops.offsetHeight || window.scrollY < objShops.offsetTop - objShops.offsetHeight) && this.showShops === true) {
        this.showShops = false
      }
      const objBrief = document.getElementById('brief')
      if (window.scrollY > objBrief.offsetTop - 200 && window.scrollY < objBrief.offsetTop + objBrief.offsetHeight && this.showBrief === false) {
        this.showBrief = true
      } else if ((window.scrollY > objBrief.offsetTop + objBrief.offsetHeight || window.scrollY < objBrief.offsetTop - objBrief.offsetHeight) && this.showBrief === true) {
        this.showBrief = false
      }
    }
  },
  beforeDestroy() {
    // 移除滚动事件监听
    this.removeScrollListener()
  },
  destroyed: function() {
    window.removeEventListener('scroll', this.myHandleScroll) //  离开页面清除（移除）滚轮滚动事件
  }
}
</script>
