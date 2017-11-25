<template>
  <div>
    <nav>
      <navbar v-bind:navlist="navlist"></navbar>
    </nav>

    <div class="container-fluid">
      <div class="row">
        <div class="center-align col-12 col-xs-12 col-sm-12 col-md-12 col-lg-12">
            <ul class="card-holder">
                <li class="mad-card" v-for="pharma in pharmaDetail">
                    <div class="mad-card-h1">{{ pharma.name }}</div>
                    <div class="mad-card-hr"></div>
                    <div class="mad-card-h2">{{ pharma.dist }} km</div>
                    <!--<div class="mad-card-h3">{{ dept.number }}</div>-->
                </li>
            </ul>
        </div>
      </div>

    </div>


  </div>
</template>

<script>
  import Navbar from '@/components/Navbar'
  export default {
    name: 'UserCart',
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
        pharmaDetail: [
          {
            name: 'pharmacy Goodwill',
            dist: 25
          },
          {
            name: 'Lalaland',
            dist: 21
          },
          {
            name: 'Lalaland',
            dist: 21
          },
          {
            name: 'Lalaland',
            dist: 21
          }
        ],
        user: 'User',
        brandplaceholder: 'Brand Name'
      }
    },
    components: {
      'navbar': Navbar
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

</style>