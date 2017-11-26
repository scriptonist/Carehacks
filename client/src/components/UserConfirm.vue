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

    <div class="container-fluid">
      <div class="row">
        <div style="padding-top: 50px;"class="col-12 col-xs-12 col-md-12 col-sm-12 col-lg-12 center">
          <div v-on:click='clicker()'>
            <button type="button" class="mad-button-border-purple">Proceed</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import Navbar from '@/components/NavbarT'
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
        brandplaceholder: 'Brand Name',
        lat: null,
        long: null
      }
    },
    components: {
      'navbar': Navbar,
      'qrcode': QRcode
    },
    created: function () {
      this.$store.commit('setDestroyCart')
      if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(this.savePosition)
      } else {
        this.msg1 = 'Geolocation is not supported by this browser.'
      }
    },
    methods: {
      savePosition: function (position) {
        this.lat = position.coords.latitude
        this.long = position.coords.longitude
      },
      clicker: function () {
        if (this.lat != null) {
          this.$router.push({path: 'uhome'})
        } else {
          alert('Click Allow Access Location')
        }
      }
    }
  }
</script>

<style scoped>
.shadow {
  text-shadow: 5px 5px 5px rgba(1,1,1,0.5);
}
</style>