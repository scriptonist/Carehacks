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
                    <div v-on:click="clicker(pharma.id)" class="button-div">
                      <button class="mad-button-border-default">Select</button>
                    </div>
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
            id: 121,
            name: 'pharmacy Goodwill',
            dist: 25
          },
          {
            id: 125,
            name: 'Lalaland',
            dist: 21
          },
          {
            id: 123,
            name: 'Lalaland',
            dist: 21
          },
          {
            id: 120,
            name: 'Lalaland',
            dist: 21
          }
        ],
        user: 'User',
        brandplaceholder: 'Brand Name'
      }
    },
    methods: {
      clicker: function (id) {
        console.log(id)
        this.$store.commit('setPharmaId', id)
        this.$router.push({path: 'uuploadprescription'})
      }
    },
    components: {
      'navbar': Navbar
    }
  }
</script>

<style scoped>
.button-div {
  float: right;
  padding-right: 25px;
  padding-top: 90px;
}
</style>