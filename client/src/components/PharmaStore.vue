<template>
  <div>

    <nav>
      <m-navbar v-bind:navlist="navlist">></m-navbar>
    </nav>

    <div class="container">
      <div class="row">
        <div class="col-xs-12 col-sm-12 col-md-12 col-lg-12">
          
        </div>
      </div>
    </div>
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