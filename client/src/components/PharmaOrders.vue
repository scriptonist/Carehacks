<template>
  <div>

    <nav>
      <m-navbar v-bind:navlist="navlist">></m-navbar>
    </nav>

    <div class="continer">
      <div class="row">
        <div class="col-xs-12 col-sm-12 col-md-12 col-lg-12">
          <h1 class="m-user-title m-opensans-300">{{userTitle}}</h1>
        </div>
      </div>
    </div>

    <div class="container">
      <div class="row">
        <div class="col-xs-12 col-sm-12 col-md-12 col-lg-12">
          <ul class="card-holder">
                <li class="mad-card" v-for="order in orderDetail">
                    <div class="mad-card-h1">{{ order.name }}</div>
                    <div class="mad-card-hr"></div>
                    <div class="mad-card-h2">{{ order.email }}</div>
                    <div class="mad-card-h3">{{ order.number }}</div>
                    <div class="mad-card-h4">Location : - {{ order.location }}</div>
                    <div class="mad-card-h4">Medicines : - {{ order.medNo }}</div>
                    <div class="mad-card-h4">Status : - {{ order.status }}</div>
                </li>
          </ul>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
import Navbar from '@/components/Navbar'
// import POrderService from '@/services/POrderService'
export default {
  data () {
    return {
      userTitle: 'Hi, Pharmacy',
      doctorList: null,
      navlist: [
        {
          name: 'Store',
          link: '/pstore'
        },
        {
          name: 'Order',
          link: '/porder'
        }
      ],
      orderDetail: [
        {
          name: 'John Smith',
          number: '8086699507',
          email: 'johnsmith@gmail.com',
          location: 'Kottayam',
          medNo: 3,
          status: 'Pending'
        }
      ]
    }
  },
  computed: {
    orderDetails: () => {
      return this.$store.getters.getOrderDetail
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
  font-family: 'Quicksand', sans-serif;
  font-weight: 400;
}
</style>