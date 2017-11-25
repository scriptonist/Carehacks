import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import Card from '@/components/Card'
import Navbar from '@/components/Navbar'
import Input from '@/components/Inputs'
import POrder from '@/components/PharmaOrders'
import UserHome from '@/components/UserHome'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'HelloWorld',
      component: HelloWorld
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
      path: '/porder',
      name: 'POrder',
      component: POrder
    },
    {
      path: '/uhome',
      name: 'UserHome',
      component: UserHome
    }
  ]
})
