<template>
  <div>
    <b-input-group v-if="edit === todo.id" class="mb-2">
      <b-form-input
        v-model="text"
        :ref="'todo' + todo.id"
        class="p-2"
        autofocus
        @blur="saveTodo()"
      ></b-form-input>
      <b-button variant="danger" @click="$emit('deleteTodo', todo.id)"
        ><font-awesome-icon icon="fa-solid fa-trash"
      /></b-button>
      <b-button variant="secondary" @click="$emit('cancelEdit')"
        ><font-awesome-icon icon="fa-solid fa-times"
      /></b-button>
    </b-input-group>
    <div v-else class="card mb-2">
      <div class="card-body d-flex p-2" @click.self="$emit('clicked')">
        <b-form-checkbox
          class="me-2 d-flex align-items-center"
          v-model="done"
          @change="$emit('doneTodo', todo.id, done)"
        ></b-form-checkbox>
        <p
          class="card-text d-flex align-items-center"
          @click.self="$emit('clicked')"
        >
          {{ todo.text }}
        </p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    todo: Object,
    edit: String,
  },
  data() {
    return {
      text: this.todo.text,
      done: this.todo.done,
      saveTimeout: null,
    };
  },
  beforeDestroy() {
    clearTimeout(this.saveTimeout);
  },
  methods: {
    saveTodo() {
      const value = this.$refs["todo" + this.todo.id].value;
      this.saveTimeout = setTimeout(() => {
        this.$emit("saveTodo", this.todo.id, value);
      }, 500);
    },
  },
};
</script>
