<template>
  <div class="mad-background-purple mad-full-screen">
    <nav>
      <navbar v-bind:navlist="navlist"></navbar>
    </nav>
    <div class="container-fluid">
      <div class="col-12 col-xs-12 col-md-12 col-sm-12 col-lg-12">
        <h2 style="padding-top: 50px;" class="shadow center white">Hello, {{user}} </h2>
      </div>
    </div>

    <div class="container-fluid welcome">
      <div class="col-12 col-xs-12 col-md-12 col-sm-12 col-lg-12">
        <h1 style="margin-top: 25vh;" class="shadow center white">Welcome to {{brandplaceholder}}</h1>
      </div>
    </div>


  </div>
</template>

<script>
  import Navbar from '@/components/Navbar'

  export default {
    name: 'UserHome',
    data () {
      return {
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
      'navbar': Navbar
    }
  }
</script>

<style scoped>
.shadow {
  text-shadow: 5px 5px 15px rgba(1,1,1,0.5);
}
</style>