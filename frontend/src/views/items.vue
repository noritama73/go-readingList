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
            <v-subheader>タイトル</v-subheader>
            <v-text-field v-model="item.Title" readonly></v-text-field>
          </v-col>
        </v-row>
      </v-card-text>
    </v-card>
    <v-card>
      <v-card-text>
        <v-row>
          <v-col>
            <v-subheader>メモ</v-subheader>
            <v-textarea v-model="item.Memo" readonly></v-textarea>
          </v-col>
        </v-row>
      </v-card-text>
    </v-card>
    <v-card>
      <v-row justify="center">
        <v-col cols="9" class="mb-2">
          <v-subheader>関連リンク</v-subheader>
          <v-btn :href="item.URL" color="link">
            <span>{{ item.URL }}</span>
          </v-btn>
        </v-col>
        <v-col cols="2" class="mb-2">
          <v-subheader>タグ</v-subheader>
          <v-text-field v-model="item.Tag" readonly></v-text-field>
        </v-col>
      </v-row>
    </v-card>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      item: null,

      submitSnackbar: false,
      submitMessage: "",
    };
  },
  mounted() {
    if (this.$route.query.id) {
      axios
        .get(process.env.VUE_APP_ENDPOINT + `/item?id=${this.$route.query.id}`)
        .then((res) => {
          this.item = res.data.Detail;
          this.openSnackbarMessage(res);
          console.log(res);
        })
        .catch((err) => {
          this.openSnackbarError(err);
        });
    }
  },
  methods: {
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
