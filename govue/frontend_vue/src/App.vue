<template>
  <div id="app">
    <IndexContainer></IndexContainer>
     <!-- <router-view></router-view> -->
  </div>
</template>

<script>
import IndexContainer from './components/IndexContainer'

export default {
  name: 'App',
  components: {
    IndexContainer,
  }
}
</script>

<style>
*{
  margin: 0;
  padding: 0;
}

#app {
  min-width: 100vw;
  min-height: 100vh;
  display:flex;
  /* background-color: #2f363e; */
      background: linear-gradient(to right, #4158d0, rgb(201, 180, 129), #c850c0);

}
</style>
