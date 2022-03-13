<template>
  <div>
    <v-card>
      <v-card-text>
        <v-row>
          <v-col> タイトル </v-col>
        </v-row>
      </v-card-text>
    </v-card>
    <v-card>
      <v-card-text>
        <v-row>
          <v-col> メモ </v-col>
        </v-row>
      </v-card-text>
    </v-card>
    <v-card>
      <v-row justify="center">
        <v-col cols="9" class="mb-2"> 関連リンク </v-col>
        <v-col cols="2" class="mb-2"> タグ </v-col>
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
    };
  },
  mounted() {
    if (this.$route.query.id) {
      axios
        .get(process.env.VUE_APP_ENDPOINT + `/item?id=${this.$route.query.id}`)
        .then((res) => {
          this.item = res.data;
        })
        .catch((err) => {
          console.log(err);
        });
    }
  },
};
</script>

<style></style>
