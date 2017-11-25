<template>
  <div>
    <nav>
      <ul class="m-nav-ul">
        <li class="m-nav-left-li m-opensans-300"><a href="">Logo</a></li>
        <li class="m-nav-left-li m-opensans-300"><a href="">Orders</a></li>
        <li class="m-nav-left-li m-opensans-300"><a href="">Store</a></li>
        <li class="m-nav-right-li m-opensans-300"><a href="">Hi, Pharmacy</a></li>
      </ul>
    </nav>
    <div class="continer">
      <div class="row">
        <div class="col-xs-12 col-sm-12 col-md-12 col-lg-12">
         <div class="m-card-depth"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
</script>

<style scoped>
.m-nav-ul {
  list-style-type: none;
  margin: 0;
  padding: 0;
  overflow: hidden;
  background-color: #000000;
  height: 80px;
  box-shadow: 0 8px 16px 0 rgba(0,0,0,0.2), 0 6px 20px 0 rgba(0,0,0,0.19);
}

.m-nav-left-li {
  float: left;
}

.m-nav-right-li {
  float: right;
}

.m-nav-left-li>a {
  display: block;
  color: white;
  text-align: center;
  padding: 30px 16px;
  text-decoration: none;
  border: none;
}

.m-nav-left-li>a {
  display: block;
  color: white;
  text-align: center;
  padding: 30px 16px;
  text-decoration: none;
  border: none;
}

.m-nav-right-li>a {
  display: block;
  color: white;
  text-align: center;
  padding: 30px 16px;
  text-decoration: none;
  border: none;
}

.m-nav-left-li>a:hover {
  background-color: #000000;
}

.m-nav-right-li>a:hover {
  background-color: #000000;
}

.m-opensans-300 {
  font-family: 'Open Sans';
  font-weight: 300;
  font-size: 1.2em;
}

.m-card-depth {
    
}
</style>