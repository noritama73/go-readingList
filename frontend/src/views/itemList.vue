<template>
  <v-data-table :header="itemHeaders" :items="items"> </v-data-table>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      items: [],
      itemHeaders: [
        { text: "ID", value: "id" },
        { text: "タイトル", value: "title" },
        { text: "更新日時", value: "updated_at" },
        { text: "タグ", value: "tag" },
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
