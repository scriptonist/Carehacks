<template>
  <div>

    <nav class="mad-nav">
      <div class="mad-nav-space">

        <div class="mad-brand">
          <a  href="#">Brand Company</a>
        </div>

        <div class="mad-nav-main-center">
          <ul>
            <li>
              <a>Order</a>
            </li>
            <li>
              <a>History</a>
            </li>
          </ul>
        </div>

      </div>
    </nav>

  </div>
</template>

<script>
  export default {
    name: 'Home'
  }
</script>

<style scoped>
.mad-nav{
  background-color: rgba(0,0,0,1);
  height: 50px;
  display: block;
  width: 100vw;
}

.mad-nav-space{
  padding-top: 10px;
  padding-left: 20px;
  padding-right: 20px;
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


</style>