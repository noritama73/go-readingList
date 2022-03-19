<template>
  <v-card>
    <v-row>
      <v-col cols="3">
        <v-card-title>レコード作成</v-card-title>
      </v-col>
      <v-col class="d-flex align-center">
        <v-btn>
          <v-icon @click="postItem()">mdi-database-arrow-left-outline</v-icon>
        </v-btn>
      </v-col>
    </v-row>
    <Form title="レコード作成" ref="child" />
  </v-card>
</template>

<script>
import axios from "axios";
import Form from "@/components/form";
export default {
  components: {
    Form,
  },
  methods: {
    postItem() {
      let param = new URLSearchParams();
      let data = this.$refs.child.childPostItem();
      param.append("data", JSON.stringify(data));
      axios
        .post(process.env.VUE_APP_ENDPOINT + "/item", param)
        .then((res) => {
          this.$refs.child.openSnackbarMessage(res.data);
        })
        .catch((err) => {
          this.$refs.child.openSnackbarError(err);
        });
    },
  },
};
</script>

<style></style>
