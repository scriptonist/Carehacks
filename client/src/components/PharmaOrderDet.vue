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

    <div class="container">
      <div class="row">
        <div class="col-xs-12 col-sm-12 col-md-12 col-lg-12">
          <ul class="card-holder">
                <li class="mad-card" v-for="med in medicines" v-on:click="orderDetailDisp">
                    <div class="mad-card-h1">{{ med.medName }}</div>
                    <div class="mad-card-h2">Quantity : - {{ med.medQty }}</div>
                    <div class="mad-card-h3">Price : - {{ med.medPrice }}</div>
                    <label class="m-label"><input type="checkbox" class="checkbox" id="check1" value="morning" />Morning</label>
                    <label class="m-label"><input type="checkbox" class="checkbox" id="check2" value="afternoon"/>Afternoon</label>
                    <label class="m-label"><input type="checkbox" class="checkbox" id="check3" value="evening"/>Evening</label>
                </li>
          </ul>
        </div>
      </div>
    </div>

    <div class="container">
      <div class="row">
        <div class="col-xs-12 col-sm-12 col-md-12 col-lg-12">
          <button class="m-btn mad-button-default" v-on:click="confirm">Confirmed</button>
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
      location: 'Trivandrum',
      medNo: 3,
      orderid: 1289367,
      status: 'Pending',
      medicines: [
        {
          medName: 'Adol 500mg',
          medQty: 3,
          medPrice: 29
        },
        {
          medName: 'Dolo 500mg',
          medQty: 3,
          medPrice: 29
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
  methods: {
    confirm: function () {
      this.$router.push({path: '/phome'})
    }
  },
  components: {
    'm-navbar': Navbar
  }
}
</script>

<style scoped>
.m-btn {
  margin-left: 80%;
}
.m-label {
  padding-top: 10px;
  padding-left: 20px;
  color: #ffffff;
  font-family: 'Quicksand';
}
.m-table {
  margin-left: 30%;
  margin-top: 5%;
}
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

body {
    font-family: 'Quicksand', sans-serif;
    overflow-x: hidden;
}

.container, .container-fluid{
    margin-left: -15px;
    margin-right: -15px;
}

.mad-font-weight-thick{
    font-weight: 700;
}

.mad-font-weight-medium{
    font-weight: 400;
}

.mad-font-weight-thin{
    font-weight: 100;
}


.white{
    color: white;
}

button{
    background-color: none;
    background: none;
}

.mad-button-default{
    background-color:#f74f61;
    padding-left: 25px;
    padding-right: 25px;
    height: 40px;
    border: none;
    border-radius: 4px;
    color: white;
    transition: all  0.35s ease-in;
    font-family: 'Quicksand', sans-serif;
    font-weight: 400;
    font-size: 1.5rem;
}

.mad-button-default:focus,
.mad-button-default:hover,
.mad-button-default:active{
    background-color:#795be2;
    box-shadow: 5px 5px 25px 1px rgba(0,0,0,0.2);
}

.mad-button-inverse-default{
    background-color:#795be2;
    padding-left: 25px;
    padding-right: 25px;
    height: 40px;
    border: none;
    border-radius: 4px;
    color: white;
    transition: all  0.35s ease-in;
    font-family: 'Quicksand', sans-serif;
    font-weight: 400;
    font-size: 1.5rem;
}

.mad-button-inverse-default:focus,
.mad-button-inverse-default:hover,
.mad-button-inverse-default:active{
    background-color:#f74f61;
    box-shadow: 5px 5px 25px 1px rgba(0,0,0,0.2);
}

.mad-button-border-default{
    border-color:white;
    border-width: 2.5px;
    padding-left: 25px;
    padding-right: 25px;
    height: 40px;
    border-radius: 25px;
    color: white;
    transition: all  0.35s ease-out;
    border-style: solid;
    background-color: none;    
    font-family: 'Quicksand', sans-serif;
    font-weight: 400;
    font-size: 1.5rem;
}

.mad-button-border-default:focus,
.mad-button-border-default:hover,
.mad-button-border-default:active{
    border-color:#f74f61;
    border-width: 2.5px;
    color: #f74f61;
    border-style: solid;
    background-color: none;
    border-radius: 18px;
}

.mad-button-border-purple{
    border-color:white;
    border-width: 2.5px;
    padding-left: 25px;
    padding-right: 25px;
    height: 40px;
    border-radius: 25px;
    color: white;
    transition: all  0.35s ease-out;
    border-style: solid;
    background-color: none;    
    font-family: 'Quicksand', sans-serif;
    font-weight: 400;
    font-size: 1.5rem;
}

.mad-button-border-purple:focus,
.mad-button-border-purple:hover,
.mad-button-border-purple:active{
    border-color:#795be2;
    border-width: 2.5px;
    color: #795be2;
    border-style: solid;
    background-color: none;
    border-radius: 18px;
}

/* Card Button */

.mad-button-card-default{
    border-color:white;
    border-width: 2.5px;
    padding-left: 25px;
    padding-right: 25px;
    height: 40px;
    border-radius: 25px;
    color: white;
    transition: all  0.35s ease-out;
    border-style: solid;
    background-color: none;    
    font-family: 'Quicksand', sans-serif;
    font-weight: 400;
    font-size: 1.5rem;
}

.mad-button-card-default:focus,
.mad-button-card-default:hover,
.mad-button-card-default:active{
    border-color:#f74f61;
    border-width: 2.5px;
    color: #f74f61;
    border-style: solid;
    background-color: none;
    border-radius: 18px;
}

.mad-button-card-full{
    background-color:white;
    padding-left: 25px;
    padding-right: 25px;
    height: 40px;
    border: none;
    border-radius: 4px;
    color: #f74f61;
    transition: all  0.35s ease-in;
    font-family: 'Quicksand', sans-serif;
    font-weight: 400;
    font-size: 1.5rem;
}

.mad-button-card-full:focus,
.mad-button-card-full:hover,
.mad-button-card-full:active{
    background-color:#f74f61;
    color: white;
    box-shadow: 5px 5px 25px 1px rgba(0,0,0,0.2);
}

.mad-button-card-full-inverse{
    background-color:#795be2;
    padding-left: 25px;
    padding-right: 25px;
    height: 40px;
    border: none;
    border-radius: 4px;
    color: white;
    transition: all  0.35s ease-in;
    font-family: 'Quicksand', sans-serif;
    font-weight: 400;
    font-size: 1.5rem;
}

.mad-button-card-full-inverse:focus,
.mad-button-card-full-inverse:hover,
.mad-button-card-full-inverse:active{
    background-color:#f74f61;
    box-shadow: 5px 5px 25px 1px rgba(0,0,0,0.2);
}

/* end of card button */

/* Input search Bar 
    it takes the complete width of the div */

input {
    font-family: 'Quicksand', sans-serif;
    font-weight: 400;
    font-size: 2.2rem;
    outline: none;
    border-style: solid;
    border-width: 0px 0px 2.5px 0px;
    border-color: #795be2;
    width: 100%;
    height: 40px;
    transition: all 0.3s ease-in;
    text-align: center;
    display: block;
    margin-left: auto;
    margin-right: auto;
    cursor: pointer;
}

input:active,
input:focus {
    border-color: #f74f61;
}


input~.glyphicon{
    position: absolute;
    right: 47.5%;
    font-size: 2.2rem;
    transition: all 0.3s ease-in;
    color: #795be2
}

input:focus~.glyphicon{
    right: 1%;
    color: #f74f61;
}

/* end of search bar */

/* Card */
.card-holder {
	background: white;
	display: flex;
    flex-wrap: wrap;  
}

.mad-card {
    flex-grow: 1;
    list-style-type: none;
    height: 265px; 
    width: 450px;
    min-width: 450px;
    max-width: 450px;
    margin: 50px;
    background: -o-linear-gradient(-45deg, #795be2, #724b9f);
    background: -moz-linear-gradient(-45deg, #795be2, #724b9f);
    background: linear-gradient(-45deg, #795be2, #724b9f);
    background: -webkit-linear-gradient(-45deg, #795be2, #724b9f);
    box-shadow: inset -1px -1px 20px rgba(1,1,1,0.2), 10px 10px 25px 1px rgba(0,0,0,0.3);
    cursor: pointer;
    transition: all 0.3s ease-in;
    border-radius: 25px;
    font-family: 'Quicksand', sans-serif;
}

.mad-card:hover{
    box-shadow: inset -1px -1px 20px rgba(1,1,1,0.2), 10px 10px 25px 1px rgba(0,0,0,0);
}

.mad-card .mad-card-h1{
    font-weight: 700;
    font-size: 3rem;
    color: white;
    padding-top: 20px;
    text-align: left;
    padding-left: 30px;
    padding-bottom: 10px;
    text-shadow: 5px 5px 15px rgba(1,1,1,0.5);
}

.mad-card .mad-card-h2{
    font-weight: 400;
    font-size: 2.5rem;
    color: white;
    text-align: left;
    padding-left: 30px;
    text-shadow: 5px 5px 15px rgba(1,1,1,0.5);
}

.mad-card .mad-card-h3{
    font-weight: 400;
    font-size: 2rem;
    color: white;
    text-align: left;
    padding-left: 30px;
    text-shadow: 5px 5px 15px rgba(1,1,1,0.5); 
}

.mad-card .mad-card-h4{
    font-weight: 400;
    font-size: 1.5rem;
    color: white;
    text-align: left;
    padding-left: 30px;
    text-shadow: 5px 5px 15px rgba(1,1,1,0.5);
}

.mad-card .mad-card-h5{
    font-weight: 100;
    font-size: 1rem;
    color: white;
    text-align: left;
    padding-left: 30px;
    text-shadow: 5px 5px 15px rgba(1,1,1,0.5);
}

.mad-card .mad-card-hr{
    margin-left: 30px;
    margin-top: -10px;
    width: 75%;
    border-bottom: 0.5px solid white;
    transition: all 0.3s ease-in;
}

.mad-card:hover .mad-card-hr{
    border-bottom: 0.5px solid #f74f61;
}

/* End of Card */

/* Calendar */

.mad-calendar{
    font-family: 'Quicksand', sans-serif;
    font-weight: 400;
    width: 60%;
    display: block;
    margin-left: auto;
    margin-right: auto;
}

/* End of Calendar */

/* Home page */

.mad-background-purple{
    background: -o-linear-gradient(-45deg, #795be2, #724b9f);
    background: -moz-linear-gradient(-45deg, #795be2, #724b9f);
    background: linear-gradient(-45deg, #795be2, #724b9f);
    background: -webkit-linear-gradient(-45deg, #795be2, #724b9f);
}

.mad-background-orange{
    background: -o-linear-gradient(-45deg, #f74f61, #ff570a);
    background: -moz-linear-gradient(-45deg, #f74f61, #ff570a);
    background: linear-gradient(-45deg, #f74f61, #ff570a);
    background: -webkit-linear-gradient(-45deg, #f74f61, #ff570a);
}

.mad-full-screen{
    height: 100vh;
}

.center{
    text-align: center;
}

.center-align{
    display: block;
    margin-left: auto;
    margin-right: auto;
}

h1{
    font-size: 3em;
}

.welcome{
    z-index: ;
}
</style>