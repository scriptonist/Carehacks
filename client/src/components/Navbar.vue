<template>
  <div>

    <nav class="mad-nav">
      <div class="mad-nav-space">

        <div class="mad-brand">
          <a  href="#">Brand Company</a>
        </div>

        <div class="mad-nav-main-center">
          <ul>
            <li v-for="links in navlist">
              <router-link v-bind:to="links.link"><a>{{links.name}}</a v-link="'\links.link'"></router-link>
            </li>
          </ul>
          
        </div>

        <div class="button-proceed">
           <button v-on:click="clicker()" v-if="cart" class="mad-button-default">Cart - {{total}} </button>
        </div>

      </div>
    </nav>

  </div>
</template>

<script>
  export default {
    props: ['navlist'],
    name: 'Home',
    computed: {
      total: function () {
        return this.$store.getters.getTotalCart
      },
      cart: function () {
        return this.$store.getters.getCart
      }
    },
    methods: {
      clicker: function () {
        console.log('asda')
        if (this.total > 0) {
          this.$router.push({path: '/ucart'})
        }
      }
    }
  }
</script>

<style scoped>
.mad-nav{
  background-color: rgba(0,0,0,1);
  height: 70px;
  display: block;
  width: 100vw;
  box-shadow: 0 10px 20px 0 rgba(0,0,0,0.2);
}

.mad-nav-space{
  padding-top: 20px;
  padding-left: 20px;
  padding-right: 40px;
  display: block;
  height: 50px;
}

.mad-brand{
  display: inline-block;
}

.mad-nav-space .mad-brand a {
  font-size: 20px;
  text-decoration: none;
  color: #795be2;
  -webkit-transition: all 0.5s ease-in;
  -moz-transition: all 0.5s ease-in;
  transition: all 0.5s ease-in;
}

.mad-nav-space .mad-brand a:hover,
.mad-nav-space .mad-brand a:active{
  color: #f74f61;
}

.mad-nav-space .mad-nav-main-center{
  float: right;
  padding-top: 3px;
}

.mad-nav-main-center ul li{
  cursor: pointer;
  list-style-type: none;
  float: right;
  padding-left: 25px;
}

.mad-nav-main-center ul li a{
  font-size: 17px;
  padding-top: 5px;
  text-decoration: none;
  color: #795be2;
  -webkit-transition: all 0.5s ease-in;
  -moz-transition: all 0.5s ease-in;
  transition: all 0.5s ease-in;
}

.mad-nav-main-center ul li a:hover,
.mad-nav-main-center ul li a:active{
  color: #f74f61;
}

.button-proceed{
  float: right;
  margin-top: -5px;
}
</style>