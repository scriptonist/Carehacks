<template>
  <div class="hello">
    <div class="container-fluid">
        <div class="row">
            <div class="col-lg-12">
                <button class="mad-button-default">Submit</button>
            </div>
        </div>
    </div>
    <div class="container-fluid" style="margin-top: 20px;">
        <div class="row">
            <div class="col-lg-12">
                <button class="mad-button-inverse-default">Submit</button>
            </div>
        </div>
    </div>
    <div class="container-fluid" style="margin-top: 20px;">
        <div class="row">
            <div class="col-lg-12">
                <button class="mad-button-border-default">Submit</button>
            </div>
        </div>
    </div>
    <div class="container-fluid" style="margin-top: 20px;">
        <div class="row">
            <div class="col-lg-12">
              <div class="search">
                <input type="text">
                <span class="glyphicon glyphicon-search"></span>  
              </div>
            </div>
        </div>
    </div>

    <div class="container">
      <div class="row">
        <div class="col-12">
          <div style="margin-top: 50px;">
            <div class="mad-calendar">
              <v-date-picker
                is-inline
                is-expanded
                mode='single'
                :select-attribute='sat'
                :theme-styles='themeStyles'
                :click='clicker()'
                v-model='selectedDate'>
              </v-date-picker>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Vue from 'vue'
import VCalendar from 'v-calendar'
import 'v-calendar/lib/v-calendar.min.css'

// Access v-calendar, v-date-packer and v-popover components
Vue.use(VCalendar)
export default {
  name: 'HelloWorld',
  data () {
    return {
      date: null,
      day: null,
      month: null,
      year: null,
      selectedDate: new Date(),
      themeStyles: {
        wrapper: {
          background: 'linear-gradient(-45deg, #795be2, #724b9f)',
          borderRadius: '25px'
        },
        headerTitle: {
          fontSize: '1.4em',
          fontWeight: '100',
          color: 'white'
        },
        headerArrows: {
          fontSize: '1.4em',
          color: 'white'
        },
        dayContent: {
          fontSize: '1.4em',
          color: 'white',
          borderRadius: '10px',
          padding: '16px'
        },
        dayContentHover: {
          background: 'linear-gradient(-45deg, #f74f61, #ff570a)',
          borderRadius: '10px',
          padding: '16px'
        }
      },
      sat: {
        highlight: {
          backgroundColor: '#f74f61',
          borderWidth: '1px',
          borderColor: '#f74f61',
          fontSize: '1.4em',
          color: 'white',
          borderRadius: '10px',
          height: '50px'
        }
      }
    }
  },
  methods: {
    clicker: function () {
      this.date = this.selectedDate.getDate()
      this.day = this.selectedDate.getDay()
      this.month = this.selectedDate.getMonth()
      this.year = this.selectedDate.getFullYear()
      // Set functions
      this.$store.dispatch('setDate', this.date)
      this.$store.dispatch('setYear', this.year)
    }
  },
  computed: {
    picker: function () {
      return this.$store.getters.getYear
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
