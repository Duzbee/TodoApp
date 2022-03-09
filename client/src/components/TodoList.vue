<template>
  <div class="container">
    <div v-if="loading" class="w-100 text-center my-3">
      <font-awesome-icon icon="fa-solid fa-spinner" spin />
    </div>
    <div v-else-if="todos.length > 0" class="todo-list pt-5 pb-2">
      <todo-item
        v-for="todo in todos"
        :key="todo.id"
        :todo="todo"
        :edit="edit"
        @clicked="edit = todo.id"
        @cancelEdit="edit = null"
        @deleteTodo="deleteTodo"
        @saveTodo="saveTodo"
        @doneTodo="doneTodo"
      ></todo-item>
    </div>
    <b-card v-else class="my-3" bg-variant="secondary" text-variant="white">
      <b-card-text class="text-center"> The todo list is empty. </b-card-text>
    </b-card>
    <b-input-group prepend="Todo">
      <b-form-input
        v-model="text"
        v-on:keyup.enter="createTodo()"
        placeholder="Type here..."
      ></b-form-input>
      <b-button variant="success" @click="createTodo()">Add</b-button>
    </b-input-group>
  </div>
</template>

<script>
import { API_URL } from "@/constants";
import axios from "axios";
import TodoItem from "./TodoItem.vue";

export default {
  components: {
    TodoItem,
  },
  data() {
    return {
      text: null,
      todos: [],
      loading: true,
      edit: "",
    };
  },
  mounted() {
    this.getTodos();
  },
  methods: {
    getTodos() {
      this.loading = true;
      axios.get(`${API_URL}/todos`).then((res) => {
        if (res.status === 200 && res.data.success === true) {
          this.todos = res.data.todos;
          this.loading = false;
        }
      });
    },
    createTodo() {
      axios
        .post(`${API_URL}/create-todo`, {
          text: this.text,
        })
        .then((res) => {
          if (res.status === 200 && res.data.success === true) {
            this.text = null;
            this.todos.push(res.data.todo);
          }
        });
    },
    deleteTodo(id) {
      axios.delete(`${API_URL}/delete-todo/${id}`).then((res) => {
        if (res.status === 200 && res.data.success === true) {
          for (let i in this.todos) {
            if (this.todos[i].id == id) {
              this.todos.splice(i, 1);
              break;
            }
          }
        }
      });
    },
    doneTodo(id, done) {
      axios
        .patch(`${API_URL}/done-todo/${id}`, {
          done: done,
        })
        .then((res) => {
          if (res.status === 200 && res.data.success === true) {
            this.edit = null;
            for (let i in this.todos) {
              if (this.todos[i].id == id) {
                this.todos[i].done = done;
                break;
              }
            }
          }
        });
    },
    saveTodo(id, text) {
      axios
        .patch(`${API_URL}/edit-todo/${id}`, {
          text: text,
        })
        .then((res) => {
          if (res.status === 200 && res.data.success === true) {
            if (this.edit == id) this.edit = null;
            for (let i in this.todos) {
              if (this.todos[i].id == id) {
                this.todos[i].text = text;
                break;
              }
            }
          }
        });
    },
  },
};
</script>

<style>
.todo-list {
  max-height: 75vh;
  overflow-y: scroll;
}
</style>
