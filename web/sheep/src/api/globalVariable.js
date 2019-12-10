// 屏幕宽
const gScreenWidth = '1920px'
// 屏幕高
const gScreenHeight = '1080px'

// 用export default 暴露出去,供其他vue文件使用
export default {
  gScreenWidth: '1920px',
  gScreenHeight: '1080px',

  // 设置属性方法
  setgScreenWidth (width) {
    this.gScreenWidth = width
  },
  setgScreenHeight (height) {
    this.gScreenHeight = height
  }
}
