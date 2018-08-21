<template>
    <v-app>
        <router-view/>
   </v-app>
</template>

<script>
import formData from '@/sampleData.json'
export default {
  name: 'App',
  data() {
      return {}
  },
  created() {
      this.$store.dispatch('initFormData')
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
