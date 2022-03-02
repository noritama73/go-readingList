<template>
  <div>
    <v-card>
      <v-card-title>タイトル</v-card-title>
      <v-card-text>
        <v-row>
          <v-col>
            <v-text-field></v-text-field>
          </v-col>
        </v-row>
      </v-card-text>
    </v-card>
    <v-card>
      <v-card-title>メモ</v-card-title>
      <v-card-text>
        <v-row>
          <v-col>
            <v-textarea auto-grow></v-textarea>
          </v-col>
        </v-row>
      </v-card-text>
    </v-card>
    <v-card>
      <v-row justify="center">
        <v-col cols="9" class="mb-2">
          <v-card-title>関連リンク</v-card-title>
          <v-text-field></v-text-field>
        </v-col>
        <v-col cols="2" class="mb-2">
          <v-card-title>タグ</v-card-title>
          <v-select :items="tags"></v-select>
        </v-col>
      </v-row>
    </v-card>
  </div>
</template>

<script>
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
        { text: "読了", value: "finish" },
        { text: "後で読む", value: "later" },
        { text: "読みかけ", value: "ontheway" },
      ],
      submits: {
        title: "",
        memo: "",
        url: "",
        tag: "",
      },
    };
  },
  methods: {
    childPostItem() {
      const submitData = {
        title: this.submits.title,
        memo: this.submits.memo,
        url: this.sibmit.url,
        tag: this.submits.tag,
      };
      this.$emit("postItem", submitData);
    },
  },
};
</script>

<style></style>
