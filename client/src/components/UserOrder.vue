<template>
  <div>
    <nav>
      <navbar v-bind:navlist="navlist"></navbar>
    </nav>

    <div class="center-align container" style="margin-top: 50px;">
        <div class="row">
            <div class="col-12 col-xs-12 col-md-12 col-lg-12">
              <div >
                <input type="text">
                <span class="glyphicon glyphicon-search"></span>  
              </div>
            </div>
        </div>
    </div>

    <div class="center container-fluid" style="margin-top: 50px;">
      <div class="row">
        <div class="col-12 col-xs-12 col-md-12 col-lg-12">
          
            <ul>
              <h3 ><li style="padding-top:10px;" v-on:click='clicker(med.id)' v-for="med in medicines"> {{med.name}} </li></h3>
            </ul>
          
        </div>
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
            name: 'Go back To Home',
            link: '/uhome'
          }
        ],
        medicines: [
          {
            name: 'Panadol',
            id: 1
          },
          {
            name: 'Dolo',
            id: 2
          }
        ]
      }
    },
    components: {
      'navbar': Navbar
    },
    methods: {
      clicker: function (id) {
        this.$store.dispatch('setAddMed', id)
      }
    }
  }
</script>

<style scoped>

li{
  list-style-type: none;
  cursor: pointer;
}
</style>