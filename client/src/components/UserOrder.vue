<template>
  <div>
    <nav>
      <navbar v-bind:navlist="navlist"></navbar>
    </nav>

    <div class="center-align container" style="margin-top: 50px;">
        <div class="row">
            <div class="col-12 col-xs-12 col-md-12 col-lg-12">
              <form v-on:submit.prevent>
              <div>
                <input type="text" v-model="searchitem" v-on:keyup.enter="searchMed">
                <span class="glyphicon glyphicon-search"></span>  
              </div>
              </form>
            </div>
        </div>
    </div>

    <div class="center container-fluid" style="margin-top: 50px;">
      <div class="row">
        <div class="col-12 col-xs-12 col-md-12 col-lg-12">
          
            <ul>
              <h3><li style="padding-top:10px;" v-on:click='clicker(med.id)' v-for="med in medicines"> {{med.name}} </li></h3>
            </ul>
          
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import Navbar from '@/components/Navbar'
  // import UOrderService from '@/services/UOrderService'
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
        searchitem: [],
        medicines: null,
        store: []
      }
    },
    computed: {
      searchMed: async function () {
        if (this.searchitem === 'Dolo') {
          this.medicines = [
            {
              name: 'Dolo',
              id: 1
            }
          ]
        } else if (this.searchitem === 'Adol') {
          this.medicines = [
            {
              name: 'Adol',
              id: 2
            }
          ]
        }
      }
    },
    components: {
      'navbar': Navbar
    },
    methods: {
      clicker: function (id) {
        this.$store.dispatch('setAddMed', id)
      }
    },
    created: function () {
      this.$store.commit('setDestroyCart')
    }
  }
</script>

<style scoped>

li{
  list-style-type: none;
  cursor: pointer;
}
</style>