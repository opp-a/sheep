export default {
  namespaced: true,
  state: {
    // 彩票数据
    info: ['试卷一张', '试卷一张', '模拟卷一张', '模拟卷一张', '深蹲十个', '深蹲十个', '深蹲十个', '俯卧撑十个', '俯卧撑十个', '俯卧撑十个',
      '一道练习题', '一道练习题', '一道练习题', '一道练习题', '一道练习题', '一道练习题', '一道练习题', '一道练习题', '一道练习题', '大笑十秒钟（可指定一人代做）',
      '大白兔奶糖三颗', '大白兔奶糖三颗', '大白兔奶糖三颗', '大白兔奶糖三颗', '分享你学习的宝贵经验', '一根棒棒糖', '一根棒棒糖', '一根棒棒糖', '一杯暖暖的奶茶', '一杯暖暖的奶茶',
      '开心一下', '一个鬼脸', '一个鬼脸', '大笑十秒钟', '大笑十秒钟', '大笑十秒钟', '讲一个笑话', '说出你的梦想', '说出你的梦想', '啥也没有O(∩_∩)O',
      '深蹲十个（可指定一人代做）', '指定两人抽奖结果互换', '抢夺他人抽奖结果']
  },
  actions: {
    /**
     * @description 设置彩票数据
     * @param {Object} context
     * @param {*} info info
     */
    set({state, dispatch}, info) {
      return new Promise(async resolve => {
        // store 赋值
        state.info = info
        // 持久化
        await dispatch(
          'staticpage/db/set',
          {
            dbName: 'sys',
            path: 'tickets.info',
            value: info,
            user: false
          },
          {root: true}
        )
        // end
        resolve()
      })
    },
    /**
     * @description 从数据库取彩票数据
     * @param {Object} context
     */
    load({state, dispatch}) {
      return new Promise(async resolve => {
        // store 赋值
        state.info = await dispatch(
          'staticpage/db/get',
          {
            dbName: 'sys',
            path: 'tickets.info',
            defaultValue: state.info,
            user: false
          },
          {root: true}
        )
        // end
        resolve()
      })
    }
  }
}
