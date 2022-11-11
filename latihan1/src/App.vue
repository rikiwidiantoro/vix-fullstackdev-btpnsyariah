<!-- eslint-disable vue/require-v-for-key -->
<!-- <template>
  <img alt="Vue logo" src="./assets/logo.png">
  <HelloWorld msg="Welcome to Your Vue.js App"/>
</template> -->

<!-- reading 1 -->
<template>
  <input type="text" v-model="name" />
  <h4>Name : {{ name }}</h4>
</template>
<!-- reading 1 -->

<script>
// import HelloWorld from './components/HelloWorld.vue'

export default {
  name: 'App',

  // reading 1
  data() {
    return {
      name: ""
    }
  },
    // reading 1


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
