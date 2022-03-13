<template>
  <v-data-table :headers="itemHeaders" :items="items">
    <template v-slot:[`item.detail`]="{ item }">
      <router-link :to="{ path: 'items', query: { id: item.ID } }">
        <v-icon> mdi-account-search </v-icon>
      </router-link>
    </template>
  </v-data-table>
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
        { test: "詳細", value: "detail", sortable: false },
      ],
    };
  },
  mounted() {
    axios.get(process.env.VUE_APP_ENDPOINT + "/itemList").then((res) => {
      this.items = res.data != null ? res.data : [];
    });
  },
  methods: {},
};
</script>

<style></style>
