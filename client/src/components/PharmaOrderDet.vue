<template>
  <div>
    
    <nav>
      <m-navbar v-bind:navlist="navlist"></m-navbar>
    </nav>

  </div>
</template>

<script>
import Navbar from '@/components/Navbar'

export default {
  data () {
    return {
      userTitle: 'Hi, Pharmacy',
      doctorList: null,
      navlist: [
        {
          name: 'Store',
          link: '/pstore'
        },
        {
          name: 'Order',
          link: '/porder'
        }
      ],
      orderDetail: [
        {
          name: 'John Smith',
          number: '8086699507',
          email: 'johnsmith@gmail.com',
          location: 'Kottayam',
          medNo: 3,
          status: 'Pending'
        }
      ]
    }
  },
  components: {
    'm-navbar': Navbar
  }
}
</script>

<style scoped>
</style>