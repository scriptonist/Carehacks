<template>
  <div>
    
    <nav>
      <m-navbar v-bind:navlist="navlist"></m-navbar>
    </nav>

    <div class="continer">
      <div class="row">
        <div class="col-xs-12 col-sm-12 col-md-12 col-lg-12">
          <h1 class="m-user-title quicksand">OrderId : - {{orderid}}</h1>
        </div>
      </div>
    </div>

    <div class="container">
      <div class="row">
        <div class="col-12 col-xs-12 col-sm-12 col-md-12 col-lg-12">
          <div class="info-box quicksand">
            <p class="info-box-font">Name - {{ name }}</p>
            <p class="info-box-font">Number - {{ number }}</p>
            <p class="info-box-font">Location - {{ location }}</p>
            <p class="info-box-font">Status - {{ status }}</p>
          </div>
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
      userTitle: 'Hi, Pharmacy',
      doctorList: null,
    //   @TODO: After Connecting Api
    //   orderid: this.store.getters.getOrderDetail.orderid,
    //   name: this.$store.getters.getOrderDetail.name,
    //   number: this.$store.getters.getOrderDetail.number,
    //   email: this.$store.getters.getOrderDetail.email,
    //   medicines: this.$store.getters.getOrderDetail.medicines,
      name: 'John Smith',
      number: '8086699507',
      email: 'johnsmith@gmail.com',
      location: 'Kottayam',
      medNo: 3,
      orderid: 1289367,
      status: 'Pending',
      medicines: [
        {
          medName: 'Adol 500mg',
          medQty: 3,
          price: 29
        },
        {
          medName: 'Adol 500mg',
          medQty: 3,
          price: 29
        }
      ],
      navlist: [
        {
          name: 'Store',
          link: '/pstore'
        },
        {
          name: 'Order',
          link: '/porder'
        }
      ]
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
  font-size: 2em;
  margin-top: 3%;
}

.info-box{
    padding-top: 10%;
    margin-left: 5%;
}
.info-box-font{
    font-size: 2em;
}
.bp{
    display: inline-block;
    position: absolute;
    right: 40px;
    top: 45px;
}
.my-btn{
    background-color: #ff0000;
    border: 2px solid #ff0000;
    border-radius: 12px;
    padding: 10px;
    padding-left: 35px;
    padding-right: 35px;
    box-shadow: 5px 5px 20px rgba(0,0,0,0.4);
    transition: all 0.5s ease;
    color: white;
    position: absolute;
    margin-top: 60px;
    right:0;
}
.box{
    position: absolute;
    right:0;
    margin: 20px;
}
.my-btn:hover{
    box-shadow: 5px 5px 20px rgba(0,0,0,0);
}
.cancel{
    display: inline-block;
    margin-top: -5px;
    color: white;
    transition: color 0.2s ease;
    text-decoration: none;
    font-size: 0.6em;
}
a:hover{
    color: #ffcc00;
}
.order-user-name{
    display: inline-block;
    position: relative;
    font-size: 2.1em;
    top: 25px;
    left: 25px;
    font-weight: 700;
}
.order-user-number{
    display: inline-block;
    position: relative;
    top: 12px;
    left: 25px;
    font-size: 1.2em;
    font-weight: 300;
}
.primary-color{
    color: #3b2577;
}
.quicksand {
    font-family: 'Quicksand', sans-serif;
}
</style>