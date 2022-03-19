<template>
  <v-snackbar v-model="deleteSnackbar" vertical top>
    {{ deleteMessage }}
    <v-btn dark text color="indigo" @click="deleteItem()"> Delete </v-btn>
  </v-snackbar>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      deleteSnackbar: false,
      deleteMessage: "",
      deleteID: "",
    };
  },
  methods: {
    deleteItem() {
      this.deleteSnackbar = false;
      axios
        .delete(process.env.VUE_APP_ENDPOINT + `/item?id=${this.deleteID}`)
        .then(() => {
          this.$router.go({
            path: this.$router.currentRoute.path,
            force: true,
          });
        })
        .catch((err) => {
          this.$emit("openSnackbarError", err);
        });
    },
    openDeleteSnackbarMessage(id) {
      this.deleteSnackbar = true;
      this.deleteID = id;
      this.deleteMessage = this.deleteID + " のデータを削除しますか？";
    },
  },
};
</script>

<style></style>
