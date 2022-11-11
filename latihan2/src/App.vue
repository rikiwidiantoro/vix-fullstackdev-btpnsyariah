<!-- eslint-disable no-unused-vars -->
<template>
  <img alt="Vue logo" src="./assets/logo.png">
  <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
  <Postingan />
</template>

<script>
// eslint-disable-next-line no-unused-vars
import Postingan from './components/Postingan.vue'
// import BuatPostingan from './components/BuatPostingan.vue'

export default {
  name: 'App',
  components: {
    // HelloWorld
    // eslint-disable-next-line vue/no-unused-components
    Postingan,
    // eslint-disable-next-line vue/no-unused-components
    // BuatPostingan
    // Home
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
