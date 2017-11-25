import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    date: null,
    day: null,
    month: null,
    year: null,
    meds: [],
    cart: true,
    username: null,
    usernumber: null,
    useremail: null,
    userlocation: null,
    usermedno: null,
    userstatus: null,
    pharmaId: null,
    usermedicines: null,
    orderdetail: null
  },
  getters: {
    getOrderDetail: (state) => {
      return state.orderdetail
    },
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
    getTotalCart: (state) => {
      return state.meds.length
    },
    getCart: (state) => {
      if (state.meds.length >= 1 && state.cart === true) {
        return true
      } else {
        return false
      }
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
    },
    getUserNumber: (state) => {
      return state.usernumber
    },
    getPharmaId: (state) => {
      return state.pharmaId
    },
    getUserMedicines: (state) => {
      return state.usermedicines
    }
  },
  mutations: {
    setOrderDetail: (state, orderdetail) => {
      state.orderdetail = orderdetail
    },
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
    setCart: (state, cart) => {
      state.cart = cart
    },
    setAddMed: (state, id) => {
      var number = 0
      var flag = false
      if (state.meds.length === 0) {
        let n = {'id': id, 'number': number}
        state.meds.push(n)
      } else {
        let n = {'id': id, 'number': number}
        for (var i = 0; i < state.meds.length; i++) {
          if (state.meds[i].id === id) {
            flag = true
          }
        }
        if (flag === true) {
          state.meds.pop(i)
        } else {
          state.meds.push(n)
        }
      }
      console.log(state.meds.length)
      // console.log(state.meds[0].id)
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
    },
    setUserNumber: (state, usernumber) => {
      state.usernumber = usernumber
    },
    setPharmaId: (state, pharmId) => {
      state.pharmId = pharmId
    },
    setUserMedicines: (state, usermedicines) => {
      state.usermedicines = usermedicines
    }
  },
  actions: {
    setDate: (context, date) => {
      console.log('sdads')
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
    },
    setAddMed: (context, id) => {
      /* for(var i = 0; i<context.meds.length; i++)
      {
      } */
      // console.log(context.meds.length)
      // context.commit('setAddMed', id)
      context.commit('setAddMed', id)
    }
  }
})
