<template>
  <div id="app">
    <nav-bar></nav-bar>
    <todo-list></todo-list>
  </div>
</template>

<script>
import NavBar from "@/components/NavBar.vue"
import TodoList from "@/components/TodoList.vue"

export default {
  name: 'App',
  components: {
    NavBar,
    TodoList
  }
}
</script>

<style scoped>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>