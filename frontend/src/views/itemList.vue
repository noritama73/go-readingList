<template>
  <v-data-table :headers="itemHeaders" :items="items"> </v-data-table>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      items: [],
      itemHeaders: [
        { text: "ID", value: "ID" },
        { text: "タイトル", value: "Title" },
        { text: "更新日時", value: "Updated_at" },
        { text: "タグ", value: "Tag" },
      ],
    };
  },
  mounted() {
    axios.get(process.env.VUE_APP_ENDPOINT + "/itemList").then((res) => {
      this.items = res.data != null ? res.data : [];
    });
  },
};
</script>

<style></style>
