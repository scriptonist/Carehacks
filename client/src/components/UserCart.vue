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