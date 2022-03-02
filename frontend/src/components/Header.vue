<template>
  <header>
    <v-navigation-drawer v-model="drawer" app clipped>
      <v-list dense>
        <v-list-item class="px-1">
          <v-list-item-title>ACTIONS</v-list-item-title>
        </v-list-item>

        <v-divider></v-divider>

        <v-list-item-group v-model="selectedItem" color="primary">
          <v-list-item
            v-for="(item, i) in HeaderItems"
            :key="i"
            :to="item.path"
          >
            <v-list-item-icon>
              <v-icon v-text="item.icon"></v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title v-text="item.text"></v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list-item-group>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar app clipped-left>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title>READINGLIST</v-toolbar-title>
    </v-app-bar>
  </header>
</template>

<script>
export default {
  data() {
    return {
      selectedItem: null,
      drawer: true,
      HeaderItems: [
        {
          text: "レコード一覧",
          icon: "mdi-format-list-bulleted",
          path: "/itemList",
        },
        {
          text: "レコード作成",
          icon: "mdi-database-arrow-left-outline",
          path: "/item",
        },
      ],
    };
  },
};
</script>

<style></style>
