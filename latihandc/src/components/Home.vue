<template>
  <h1>Dinamic Component</h1>
  <a href="#" @click="dilihat='Php'" class="btn btn-success mx-5">PHP</a>
  <a href="#" @click="dilihat='Js'" class="btn btn-success">JavaScript</a><br><br><br>
  <p>Pesan :</p>
  <component :is="dilihat"></component>
  <!-- <Php />
  <Js /> -->
</template>

<script>
  import Php from './Php.vue'
  import Js from './Js.vue'

export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: 'Home',
  components: {
    Php,
    Js
  },
  data() {
    return {
      dilihat: 'Php'
    }
  }
}
</script>

<style>
  h1 {
    color: red;
  }
</style>
