<template>
  <v-card>
    <v-snackbar v-model="submitSnackbar" vertical top>
      {{ submitMessage }}
      <v-btn dark text color="indigo" @click="submitSnackbar = false">
        Close
      </v-btn>
    </v-snackbar>

    <Delete ref="child" />

    <v-data-table :headers="itemHeaders" :items="items">
      <template v-slot:[`item.detail`]="{ item }">
        <router-link :to="{ path: 'items', query: { id: item.ID } }">
          <v-icon> mdi-account-search </v-icon>
        </router-link>
      </template>
      <template v-slot:[`item.delete`]="{ item }">
        <v-btn @click="deleteItemCheck(item.ID)">削除</v-btn>
      </template>
    </v-data-table>
  </v-card>
</template>

<script>
import axios from "axios";
import Delete from "@/components/delete";
export default {
  data() {
    return {
      items: [],
      itemHeaders: [
        { text: "タイトル", value: "Title" },
        { text: "更新日時", value: "Updated_at" },
        { text: "タグ", value: "Tag" },
        { text: "詳細", value: "detail", sortable: false },
        { text: "削除", value: "delete", sortable: false },
      ],

      submitSnackbar: false,
      submitMessage: "",
      deleteID: "",
    };
  },
  components: {
    Delete,
  },
  mounted() {
    axios.get(process.env.VUE_APP_ENDPOINT + "/itemList").then((res) => {
      this.items = res.data != null ? res.data : [];
    });
  },
  methods: {
    openSnackbarMessage(message) {
      this.submitSnackbar = true;
      this.submitMessage = message;
    },
    openSnackbarError(error) {
      console.log(error);
      this.submitSnackbar = true;
      this.submitMessage =
        error.response.status +
        "(" +
        error.response.data.code +
        ")" +
        error.response.data.errors;
    },
    deleteItemCheck(id) {
      this.$refs.child.openDeleteSnackbarMessage(id);
    },
  },
};
</script>

<style></style>
