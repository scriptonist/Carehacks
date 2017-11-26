import Vue from 'vue'
import Router from 'vue-router'
import Card from '@/components/Card'
import Navbar from '@/components/Navbar'
import Input from '@/components/Inputs'
import POrder from '@/components/PharmaOrders'
import POrderDet from '@/components/PharmaOrderDet'
import PStore from '@/components/PharmaStore'
import PHome from '@/components/PharmaHome'
import UserHome from '@/components/UserHome'
import UserOrder from '@/components/UserOrder'
import UserCart from '@/components/UserCart'
import Upload from '@/components/Upload'
import Confirm from '@/components/UserConfirm'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'UserHome',
      component: UserHome
    },
    {
      path: '/card',
      name: 'Card',
      component: Card
    },
    {
      path: '/navbar',
      name: 'Navbar',
      component: Navbar
    },
    {
      path: '/input',
      name: 'Input',
      component: Input
    },
    {
      path: '/phome',
      name: 'PHome',
      component: PHome
    },
    {
      path: '/porder',
      name: 'POrder',
      component: POrder
    },
    {
      path: '/porder/details',
      name: 'POrderDet',
      component: POrderDet
    },
    {
      path: '/pstore',
      name: 'PStore',
      component: PStore
    },
    {
      path: '/uhome',
      name: 'UserHome',
      component: UserHome
    },
    {
      path: '/uorder',
      name: 'UserOrder',
      component: UserOrder
    },
    {
      path: '/ucart',
      name: 'UserCart',
      component: UserCart
    },
    {
      path: '/uuploadprescription',
      name: 'Upload',
      component: Upload
    },
    {
      path: '/uconfirm',
      name: 'Confirm',
      component: Confirm
    }
  ]
})
