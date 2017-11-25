import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    date: null,
    day: null,
    month: null,
    year: null
  },
  getters: {
    getDate: (state) => {
      return state.date
    },
    getDay: (state) => {
      return state.day
    },
    getMonth: (state) => {
      return state.month
    },
    getYear: (state) => {
      return state.year
    }
  },
  mutations: {
    setDate: (state, date) => {
      state.date = date
    },
    setDay: (state, day) => {
      state.day = day
    },
    setMonth: (state, month) => {
      state.month = month
    },
    setYear: (state, year) => {
      state.year = year
    }
  },
  actions: {
    setDate: (context, date) => {
      context.commit('setDate', date)
    },
    setDay: (context, day) => {
      context.commit('setDate', day)
    },
    setMonth: (context, month) => {
      context.commit('setDate', month)
    },
    setYear: (context, year) => {
      context.commit('setDate', year)
    }
  }
})
