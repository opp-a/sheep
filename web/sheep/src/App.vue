<template>
  <router-view></router-view>
</template>

<script>
export default {
  name: 'app',
  updated () {
  let userName = localStorage.getItem('userName')
    console.log(userName)
    if (userName === null || userName === undefined || userName === '') {
      this.$router.push('/login')
    } else {
      this.$store.commit('setUserName', userName)
    }
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
  margin-top: 60px;
}
</style>
