<template>
  <div>
    <nav>
      <navbar v-bind:navlist="navlist"></navbar>
    </nav>

    <div class="container-fluid">
      <div class="col-12 col-xs-12 col-md-12 col-sm-12 col-lg-12">
        <h2 style="padding-top: 50px;" class="shadow center">Upload your prescription</h2>
      </div>
    </div>

    <div class="container-fluid">
      <div class="row">
        <div id="app" class="component">
          <div class="col-6 col-xs-6 col-sm-6 col-md-6 col-lg-6">
            <webcam ref="webcam"></webcam>
          </div>
          <div class="col-6 col-xs-6 col-sm-6 col-md-6 col-lg-6">
            <img class="image-div" :src="this.img" style="width:500px;height:435px;" />  
          </div>
        </div>
      </div>
    </div>

    <div class="container-fluid">
      <div class="row">
        <div class="col-12 col-xs-12 col-sm-12 col-md-12 col-lg-12">
          <div style="padding-bottom: 150px;" class="center">
            <button class="mad-button-default" type="button" @click="photo">{{msg}}</button>
            <button class="mad-button-default" v-if="imgpresent" v-on:click='clicker()'>Proceed</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import Webcam from 'vue-web-cam/src/webcam'
  import Navbar from '@/components/Navbar'
  export default {
    name: 'UserHome',
    data () {
      return {
        msg: 'Capture Prescription',
        navlist: [
        ],
        img: null,
        imgpresent: false
      }
    },
    methods: {
      photo: function () {
        this.msg = 'Recapture Prescription'
        this.img = this.$refs.webcam.capture()
        this.imgpresent = true
      },
      clicker: function () {
        this.$router.push({path: 'uconfirm'})
      }
    },
    components: {
      'navbar': Navbar,
      'webcam': Webcam
    },
    created: function () {
      console.log('created')
      this.$store.commit('setCart', false)
    },
    destroyed: function () {
      this.$store.commit('setCart', true)
    }
  }
</script>

<style scoped>
.component {
  display: block;
  margin-left: auto;
  margin-right: auto;
  text-align: center;
}

.image-div {
  padding-top: 65px;
}
</style>