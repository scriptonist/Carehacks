<template>
  <div class="mad-background-orange mad-full-screen">
    <nav>
      <navbar v-bind:navlist="navlist"></navbar>
    </nav>
    <div class="container-fluid">
      <div class="col-12 col-xs-12 col-md-12 col-sm-12 col-lg-12">
        <h2 style="padding-top: 50px;" class="shadow center white">Success, Your QR Code is: </h2>
      </div>
    </div>

    <div class="container-fluid welcome">
      <div class="col-12 col-xs-12 col-md-12 col-sm-12 col-lg-12">
        <div class="center" style="padding-top: 80px;">
          <qrcode v-bind:val="q"></qrcode>
        </div>
      </div>
    </div>


  </div>
</template>

<script>
  import Navbar from '@/components/Navbar'
  import QRcode from '@/components/QRcode.vue'
  export default {
    name: 'UserHome',
    data () {
      return {
        q: 'awdefaefef',
        navlist: [
          {
            name: 'History',
            link: '/uhistory'
          },
          {
            name: 'Order',
            link: '/uorder'
          }
        ],
        user: 'User',
        brandplaceholder: 'Brand Name'
      }
    },
    components: {
      'navbar': Navbar,
      'qrcode': QRcode
    }
  }
</script>

<style scoped>
.shadow {
  text-shadow: 5px 5px 5px rgba(1,1,1,0.5);
}
</style>