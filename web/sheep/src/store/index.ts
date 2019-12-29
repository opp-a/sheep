import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        userName: ''
    },
    mutations: {
        setUserName (state, payload) {
            state.userName = payload
        }
    },
    actions: {
        minusPriceAsync (context, payload) {
            setTimeout(() => {
                context.commit('setUserName', payload) // context??
            }, 2000)
        }
    },
    modules: {},
    getters: {
        getMyName: (state) => {
            return state.userName
        }
    }
})
