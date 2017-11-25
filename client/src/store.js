import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    date: null,
    day: null,
    month: null,
    year: null,
    username: null,
    usernumber: null,
    useremail: null,
    userlocation: null,
    usermedno: null,
    userstatus: null
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
    },
    getUserName: (state) => {
      return state.username
    },
    getUserEmail: (state) => {
      return state.useremail
    },
    getUserLocation: (state) => {
      return state.userlocation
    },
    getUserMedNo: (state) => {
      return state.usermedno
    },
    getUserStatus: (state) => {
      return state.userstatus
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
    },
    setUserName: (state, username) => {
      state.username = username
    },
    setUserEmail: (state, useremail) => {
      state.useremail = useremail
    },
    setUserLocation: (state, userlocation) => {
      state.userlocation = userlocation
    },
    setUserMedNo: (state, usermedno) => {
      state.usermedno = usermedno
    },
    setUserStatus: (state, userstatus) => {
      state.userstatus = userstatus
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
