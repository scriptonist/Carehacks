<template>
  <div>
    <nav>
      <navbar></navbar>
    </nav>
  </div>
</template>

<script>
  import Navbar from '@/components/Navbar'

  export default {
    name: 'UserHome',
    components: {
      'navbar': Navbar
    }
  }
</script>

<style scoped>

</style>