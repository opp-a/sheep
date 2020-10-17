import Vue from 'vue'
import Vuex from 'vuex'

import sheep from './modules/sheep'
import staticpage from './modules/staticpage'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    sheep,
    staticpage
  }
})
