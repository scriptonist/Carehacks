<template>
  <div>
    <nav>
      <m-navbar></m-navbar>
    </nav>
    <div class="continer">
      <div class="row">
        <div class="col-xs-12 col-sm-12 col-md-12 col-lg-12">
          <h1 class="m-user-title m-opensans-300">{{userTitle}}</h1>
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
      userTitle: 'Hi, Pharmacy'
    }
  },
  components: {
    'm-navbar': Navbar
  }
}
</script>

<style scoped>
.m-user-title {
  color: #000000;
  text-align: center;
  margin-top: 3%;
}

.m-opensans-300 {
  font-family: 'Open Sans';
  font-weight: 300;
}
</style>