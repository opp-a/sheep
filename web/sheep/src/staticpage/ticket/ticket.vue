<template>
  <div style="min-width: 1800px;min-height: 1000px; background:#00628f;">
    <el-row>
      <el-col :span="24">
        <div style="font-size:x-large;color: rgb(233, 170, 34);">刮一刮</div>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c1" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c2" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c3" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c4" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c5" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c6" width="280" height="110"></canvas>
        </div>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c7" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c8" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c9" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c10" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c11" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c12" width="280" height="110"></canvas>
        </div>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c13" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c14" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c15" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c16" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c17" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c18" width="280" height="110"></canvas>
        </div>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c19" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c20" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c21" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c22" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c23" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c24" width="280" height="110"></canvas>
        </div>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c25" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c26" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c27" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c28" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c29" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c30" width="280" height="110"></canvas>
        </div>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c31" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c32" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c33" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c34" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c35" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c36" width="280" height="110"></canvas>
        </div>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c37" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c38" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c39" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c40" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c41" width="280" height="110"></canvas>
        </div>
      </el-col>
      <el-col :span="4">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c42" width="280" height="110"></canvas>
        </div>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="24">
        <div class="grid-content">
          <div class="result"></div>
          <canvas class="covercanvas" id="c43" width="280" height="110"></canvas>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import {mapState} from 'vuex'
export default {
  name: 'ticket',
  data() {
    return {
      ticketinfos: [],
      mousedown: false,
      device: true
    }
  },
  computed: {
    ...mapState('staticpage/ticket', ['info'])
  },
  methods: {
    initcanvas() {
      var canvas = document.getElementsByTagName('canvas')
      for (let i = 0; i < 43; i++) {
        var ctx = canvas[i].getContext('2d')
        if (i % 4 === 0) {
          ctx.fillStyle = '#1aa7df'
        } else if (i % 4 === 1) {
          ctx.fillStyle = '#ff0070'
        } else if (i % 4 === 2) {
          ctx.fillStyle = '#ff4000'
        } else if (i % 4 === 3) {
          ctx.fillStyle = '#00ff40'
        }
        ctx.textAlign = 'center'
        ctx.font = 'Bold 18px Arial'
        ctx.fillRect(0, 0, 280, 110)
        ctx.fillStyle = '#555'
        ctx.fillText('刮奖区', 140, 60)
        canvas[i].style.position = 'absolute'
        canvas[i].style.zIndex = '2'
        canvas[i].style.width = '280'
        canvas[i].style.height = '110'
        canvas[i].addEventListener('touchstart', this.eventDown)
        canvas[i].addEventListener('touchend', this.eventUp)
        canvas[i].addEventListener('mousedown', this.eventDown)
        canvas[i].addEventListener('mouseup', this.eventUp)
      }
    },
    eventDown(e) {
      e.preventDefault()
      this.mousedown = true
      e.currentTarget.addEventListener('touchmove', this.eventMove)
      e.currentTarget.addEventListener('mousemove', this.eventMove)
    },
    eventUp(e) {
      e.preventDefault()
      this.mousedown = false
      e.currentTarget.removeEventListener('touchmove', this.eventMove)
      e.currentTarget.removeEventListener('mousemove', this.eventMove)
      var ctx = e.currentTarget.getContext('2d')
      // 获取canvas全部的像素点，这里需要用到getImageData方法
      // getImageData(int x,int y,int width,int height)：该方法获取canvas上从(x,y)点开始，宽为width、高为height的图片区域的数据，该方法返回的是一个CanvasPixelArray对象，该对象具有width、height、data等属性。data属性为一个数组，该数组每4个元素对应一个像素点。
      const imageDate = ctx.getImageData(0, 0, e.currentTarget.width, e.currentTarget.height)
      // 记录整个canvas所有像素点
      const allPX = imageDate.width * imageDate.height
      // 记录刮开的像素点个数
      let iNum = 0
      // 通过for循环判断刮开了多少个像素点
      for (let i = 0; i < allPX; i++) {
        if (imageDate.data[i * 4 + 3] === 0) {
          iNum++
        }
      }
      // 当刮开的像素点数量大于等于总像素点数量的40%时隐藏canvas
      if (iNum >= allPX / 2) {
        e.currentTarget.style.display = 'none'
      }
    },
    eventMove(e) {
      e.preventDefault()
      if (this.mousedown) {
        // 获取鼠标或者手指距离屏幕的x坐标
        const x = this.device ? e.targetTouches[0].clientX : e.clientX
        // 获取鼠标或者手指距离屏幕的y坐标
        const y = this.device ? e.targetTouches[0].clientY : e.clientY
        // 获取canvas元素相对于视窗的位置
        const bbox = e.currentTarget.getBoundingClientRect()
        // 开始绘制涂抹动作的canvas
        var ctx = e.currentTarget.getContext('2d')
        ctx.beginPath()
        // canvas的globalCompositeOperation属性指的是canvas中的合成操作
        // 将该属性设置为"destination-out"会将在与canvas源不重叠的区域上保留目标，其他部分都变成透明
        ctx.globalCompositeOperation = 'destination-out'
        // 因为要模拟手机涂抹过程，所以涂抹的图形选择圆形，在这里绘制一个原型，半径为20px
        ctx.arc(x - bbox.left, y - bbox.top, 17, 0, Math.PI * 2, false)
        ctx.fill()
        ctx.closePath()
      }
    }
  },
  mounted: function() {
    var flag = []
    for (let i = 0; i < 43; i++) {
      flag[i] = false
    }
    var count = 0
    while (count < 43) {
      const randomindex = Math.floor(Math.random() * 43)
      if (flag[randomindex]) {
        continue
      } else {
        this.ticketinfos[count] = this.info[randomindex]
        flag[randomindex] = true
        count++
      }
    }
    this.device = /android|iphone|ipad|ipod|webos|iemobile|opear mini|linux/i.test(navigator.userAgent.toLowerCase()) // 判断是否是移动设备

    var objs = document.getElementsByClassName('result')
    objs.forEach((element, index, array) => {
      var customstyle = ''
      if (index % 5 === 0) {
        customstyle = '#00ff40'
      } else if (index % 5 === 1) {
        customstyle = '#ff4000'
      } else if (index % 5 === 2) {
        customstyle = '#1aa7df'
      } else if (index % 5 === 3) {
        customstyle = '#ff0070'
      } else if (index % 5 === 4) {
        customstyle = '#ad33ff'
      }
      element.style.color = customstyle
      element.innerText = this.ticketinfos[index]
    })
    this.initcanvas()
  }
}
</script>
<style scoped>
.el-row {
  padding-top: 10px;
  padding-bottom: 7px;
}
.el-col {
  border-radius: 4px;
  padding-left: 10px;
  padding-right: 10px;
}
.covercanvas {
  position: absolute;
  width: 280px;
  height: 110px;
  z-index: 2;
  margin-left: -140px;
}
.result {
  position: absolute;
  width: 280px;
  height: 110px;
  font-weight: 600;
  font-size: 20px;
  font-family: STXingkai;
  text-align: center;
  line-height: 110px;
  z-index: 1;
}
.grid-content {
  position: relative;
  color: antiquewhite;
  border-radius: 4px;
  height: 110px;
  width: 280px;
  box-shadow: 1px 1px 4px rgb(80, 167, 224);
  background-color: #005679;
}
</style>
