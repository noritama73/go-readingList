<template>
  <div>
    <v-card>
      <v-snackbar v-model="submitSnackbar" vertical top>
        {{ submitMessage }}
        <v-btn dark text color="indigo" @click="submitSnackbar = false">
          Close
        </v-btn>
      </v-snackbar>

      <v-card-text>
        <v-row>
          <v-col>
            <v-text-field
              v-model="submits.title"
              label="タイトル"
            ></v-text-field>
          </v-col>
        </v-row>
      </v-card-text>
    </v-card>
    <v-card>
      <v-card-text>
        <v-row>
          <v-col>
            <v-textarea
              v-model="submits.memo"
              label="メモ"
              auto-grow
            ></v-textarea>
          </v-col>
        </v-row>
      </v-card-text>
    </v-card>
    <v-card>
      <v-row justify="center">
        <v-col cols="9" class="mb-2">
          <v-text-field v-model="submits.url" label="関連リンク"></v-text-field>
        </v-col>
        <v-col cols="2" class="mb-2">
          <v-select :items="tags" v-model="submits.tag" label="タグ"></v-select>
        </v-col>
      </v-row>
    </v-card>
  </div>
</template>

<script>
import axios from "axios";
export default {
  props: {
    title: {
      type: String,
      required: true,
    },
  },
  data() {
    return {
      tags: [
        { text: "読了", value: "読了" },
        { text: "後で読む", value: "後で読む" },
        { text: "読みかけ", value: "読みかけ" },
      ],
      submits: {
        title: "",
        memo: "",
        url: "",
        tag: "",
      },

      submitSnackbar: false,
      submitMessage: "",
    };
  },
  methods: {
    childGetItem(id) {
      axios
        .get(process.env.VUE_APP_ENDPOINT + `/item?id=${id}`)
        .then((res) => {
          this.submits.title = res.data.Detail.Title;
          this.submits.memo = res.data.Detail.Memo;
          this.submits.url = res.data.Detail.URL;
          this.submits.tag = res.data.Detail.Tag;
        })
        .catch((err) => {
          this.openSnackbarError(err);
        });
    },
    childSendItem() {
      const submitData = {
        title: this.submits.title,
        memo: this.submits.memo,
        url: this.submits.url,
        tag: this.submits.tag,
      };
      return submitData;
    },
    openSnackbarMessage(message) {
      this.submitSnackbar = true;
      this.submitMessage = message;
    },
    openSnackbarError(error) {
      this.submitSnackbar = true;
      this.submitMessage =
        error.response.status +
        "(" +
        error.response.data.code +
        ")" +
        error.response.data.errors;
    },
  },
};
</script>

<style></style>
