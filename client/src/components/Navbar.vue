<template>
  <div class="nav-div">
    <nav class="navbar-default navbar-toggleable-md navbar-light bg-faded">
      <a class="navbar-brand" href="#">Brand Company</a>
      <div class="collapse navbar-collapse" id="navbarNavDropdown">
        <ul class="navbar-nav">
          <li class="nav-item">
            <a class="nav-link" href="#">Home <span class="sr-only">(current)</span></a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="#">Features</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="#">Pricing</a>
          </li>
          <li class="nav-item dropdown">
            <a class="nav-link dropdown-toggle" href="http://example.com" id="navbarDropdownMenuLink" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
              Dropdown link <span class="glyphicon glyphicon-chevron-down"></span>
            </a>
            <div class="dropdown-menu" aria-labelledby="navbarDropdownMenuLink">
              <a class="dropdown-item" href="#">Action</a>
              <a class="dropdown-item" href="#">Another action</a>
              <a class="dropdown-item" href="#">Something else here</a>
            </div>
          </li>
        </ul>
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
.navbar-default {
  background-color: rgba(0,0,0,0.4);
  border-radius: 0px;
}

.navbar-default .navbar-brand {
  color: #795be2;
}
.navbar-default .navbar-brand:hover,
.navbar-default .navbar-brand:focus {
  color: #f74f61;
}
.navbar-default .navbar-text {
  color: #795be2;
}

.nav-item .active{
  background-color: rgba(0,0,0,0.4);
}

.navbar-default .dropdown-menu{
  background-color: rgba(0,0,0,0.4);
  border-radius: 0px 0px 5px 5px;
}

.navbar-default .dropdown-item{
  color: #795be2;
  display: block;
  height: 30px;
  padding-left: 8px;
  text-decoration: none;
}

.navbar-default .dropdown-menu .dropdown-item:hover,
.navbar-default .dropdown-menu .dropdown-item:active {
  color: #f74f61;
}

.navbar-default .navbar-nav > li  {
  color: #795be2;
  list-style-type: none;
  list-style: none;
  text-decoration-line: none;
  width: 150px;
  text-align: center;
  cursor: pointer;
}

.navbar-default .navbar-nav > li > a {
  color: #795be2;
  display: inline-block;
  text-decoration-line: none;
  width: 150px;
}
.navbar-default .navbar-nav > li > a:hover,
.navbar-default .navbar-nav > li > a:focus {
  background-color: rgba(0,0,0,0.4);
  color: #f74f61;
}
.navbar-default .navbar-nav > .active > a,
.navbar-default .navbar-nav > .active > a:hover,
.navbar-default .navbar-nav > .active > a:focus {
  color: #f74f61;
}
.navbar-default .navbar-nav > .open > a,
.navbar-default .navbar-nav > .open > a:hover,
.navbar-default .navbar-nav > .open > a:focus {
  color: #505050;
}
.navbar-default .navbar-toggle {
  border-color: none;
}
.navbar-default .navbar-toggle:hover,
.navbar-default .navbar-toggle:focus {
  background-color: #000;
}
.navbar-default .navbar-toggle .icon-bar {
  background-color: #000;
}
.navbar-default .navbar-collapse,
.navbar-default .navbar-form {
  border-color: #000;
}
.navbar-default .navbar-link {
  color: #ffffff;
}
.navbar-default .navbar-link:hover {
  color: #505050;
}

@media (max-width: 767px) {
  .navbar-default .navbar-nav .open .dropdown-menu > li > a {
    color: #ffffff;
  }
  .navbar-default .navbar-nav .open .dropdown-menu > li > a:hover,
  .navbar-default .navbar-nav .open .dropdown-menu > li > a:focus {
    color: #505050;
  }
  .navbar-default .navbar-nav .open .dropdown-menu > .active > a,
  .navbar-default .navbar-nav .open .dropdown-menu > .active > a:hover,
  .navbar-default .navbar-nav .open .dropdown-menu > .active > a:focus {
    color: #505050;
    background-color: #ffffff;
  }
}

</style>