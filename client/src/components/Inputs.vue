<template>
  <div class="pol">
     <label>
        <input type="checkbox" checked="checked">
        <span class="checkmark"></span>
      </label>
  </div>
</template>

<script>
  export default {
  }
</script>

<style scoped>
 /* Customize the label (the container) */
 .pol {
   margin-left: 200px;
 }
 
/* Hide the browser's default checkbox */
input {
  position: absolute;
  opacity: 0;
}

/* Create a custom checkbox */
.checkmark {
  position: absolute;
  top: 0;
  left: 0;
  height: 25px;
  width: 25px;
  background-color: #eee;
}

/* On mouse-over, add a grey background color */
input ~ .checkmark {
  background-color: #795be2;
}

/* When the checkbox is checked, add a blue background */
input:checked ~ .checkmark {
  background-color: #f74f61;
}

/* Create the checkmark/indicator (hidden when not checked) */
.checkmark:after {
  content: "";
  position: absolute;
  display: none;
}

/* Show the checkmark when checked */
input:checked ~ .checkmark:after {
  display: block;
}

/* Style the checkmark/indicator */
.checkmark:after {
  left: 9px;
  top: 6px;
  width: 6px;
  height: 10px;
  border: solid white;
  border-width: 0 3px 3px 0;
  -webkit-transform: rotate(45deg);
  -ms-transform: rotate(45deg);
  transform: rotate(45deg);
} 


/************************************************************************************/

/* Hide the browser's default radio button */
input {
  position: absolute;
  opacity: 0;
}

/* Create a custom radio button */
.checkmark {
  position: absolute;
  top: 0;
  left: 0;
  height: 25px;
  width: 25px;
  background-color: #eee;
  border-radius: 50%;
}

/* On mouse-over, add a grey background color */
input ~ .checkmark {
  background-color: #795be2;
}

/* When the radio button is checked, add a blue background */
input:checked ~ .checkmark {
  background-color: #f74f61;
}

/* Create the indicator (the dot/circle - hidden when not checked) */
.checkmark:after {
  content: "";
  position: absolute;
  display: none;
}

/* Show the indicator (dot/circle) when checked */
input:checked ~ .checkmark:after {
  display: block;
}

/* Style the indicator (dot/circle) */
.checkmark:after {
  top: 7.6px;
  left: 7.6px;
  width: 10px;
  height: 10px;
  border-radius: 50%;
  background: white;
} 

</style>