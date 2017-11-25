<!--<template>
    <div class="container-fluid" style="margin-top: 20px;">
        <div class="row">
            <div class="col-12 col-xs-12 col-sm-12 col-md-12 col-lg-12">
              <div class="mad-card">
                <div class="mad-card-h1">Sample Text<div class="mad-card-hr"></div></div>
                <br/>
                <div class="mad-card-h2">Sample Text</div>
                <div class="mad-card-h3">Sample Text</div>
                <div class="mad-card-h4">Sample Text</div>
                <div class="mad-card-h5">Sample Text</div>
              </div>
            </div>
            <div class="col-lg-4">
              <div class="mad-card">
                <div style="padding-top: 20px;">
                  <button class="mad-button-card-default">Submit</button> 
                </div>
                <div style="padding-top: 20px;">
                  <button class="mad-button-card-full">Submit</button> 
                </div>
                <div style="padding-top: 20px;">
                  <button class="mad-button-card-full-inverse">Submit</button> 
                </div>
              </div>
            </div>
            <div class="col-lg-4">
              <div class="mad-card"></div>
            </div>
        </div>
    </div>
  </div>-->


<template>
    <div class="container-fluid">
        <div class="col-12 col-xs-12 col-sm-12 col-md-12 col-lg-12">
            <ul class="card-holder">
                <li class="mad-card" v-for="dept in doctorsDetail">
                    <div class="mad-card-h1">{{ dept.name }}</div>
                    <div class="mad-card-hr"></div>
                    <div class="mad-card-h2">{{ dept.email }}</div>
                    <div class="mad-card-h3">{{ dept.number }}</div>
                </li>
            </ul>
        </div>
    </div>
</template>

<script>
export default {
  name: 'Card',
  data () {
    return {
      doctorList: null,
      doctorsDetail: [
        {
          name: 'Ondasd',
          number: '256464',
          email: 'loremipsum@lorem.com'
        },
        {
          name: 'sdawd',
          number: '256464',
          email: 'loremipsum@lorem.com'
        },
        {
          name: 'sdawd',
          number: '256464',
          email: 'loremipsum@lorem.com'
        },
        {
          name: 'sdawd',
          number: '256464',
          email: 'loremipsum@lorem.com'
        },
        {
          name: 'sdawd',
          number: '256464',
          email: 'loremipsum@lorem.com'
        }
      ]
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
