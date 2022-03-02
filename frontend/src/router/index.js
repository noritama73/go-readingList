import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    component: Home,
  },
  {
    path: "/itemList",
    component: () => import("@/views/itemList.vue"),
  },
  {
    path: "/item",
    component: () => import("@/views/item.vue"),
  },
  {
    path: "/item/edit",
    component: () => import("@/views/itemEdit.vue"),
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
