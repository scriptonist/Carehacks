<template>
  <div>
    <nav>
      <m-navbar></m-navbar>
    </nav>
    <div class="continer">
      <div class="row">
        <div class="col-xs-12 col-sm-12 col-md-12 col-lg-12">
         <div class="m-card-depth"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Navbar from '@/components/Navbar'
export default {
  components: {
    'm-navbar': Navbar
  }
}
</script>

<style scoped>
.m-card-depth {
  background: #795be2;
  display:
}
</style>