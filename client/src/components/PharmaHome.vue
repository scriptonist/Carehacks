<template>
  <div>

    <nav>
      <m-navbar v-bind:navlist="navlist">></m-navbar>
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