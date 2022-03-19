<template>
  <v-card>
    <v-row>
      <v-col cols="3">
        <v-card-title>レコード編集</v-card-title>
      </v-col>
      <v-col class="d-flex align-center">
        <v-btn>
          <v-icon @click="putItem()">mdi-database-arrow-left-outline</v-icon>
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
  mounted() {
    if (this.$route.query.id) {
      this.$refs.child.childGetItem(this.$route.query.id);
    }
  },
  methods: {
    putItem() {
      let param = new URLSearchParams();
      let data = this.$refs.child.childSendItem();
      param.append("id", this.$route.query.id);
      param.append("data", JSON.stringify(data));
      axios
        .put(process.env.VUE_APP_ENDPOINT + "/item", param)
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
