import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/home",
    name: "Home",
    component: Home,
  },
  {
    path: "/about",
    name: "About",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/About.vue"),
  },
  {
    path: "/",
    name:"signup",
    component: () => import ("../views/home/index.vue")
  },
  {
    path:"/login",
    name:"login",
    component: ()=>import("../views/login/index.vue")
  },
  {
    path:"/dash",
    name:"dash",
    component: () => import("../views/dashboard/index.vue")
  },
  {
    path:"/search",
    name:"search",
    component: ()=> import ("../views/searchbus/index.vue")
  },
  {
    path:"/seatbook",
    name:"seatbook",
    component: () => import("../views/seatbook/index.vue")
  },
  {
    path:"/booking",
    name:"booking",
    component: () => import ("../views/booking/index.vue")
  },
  {
    path:"/allbook",
    name:"allbook",
    component: ()=> import ("../views/allbooks/index.vue")
  },
  {
    path:"/gpay",
    name:"gpay",
    component: ()=> import("../components/allbook/component/gpay.vue")
  },
  {
    path:"/payoption",
    name:"option",
    component: ()=>import("../components/allbook/component/optionpay.vue")
  },
  {
    path:"/otp",
    name:"otp",
    component: ()=> import ("../components/allbook/component/otp.vue")
  },
  {
    path:"/phone",
    name:"phone",
    component: () => import("../components/allbook/component/phoneno.vue")
  },
  {
    path:"/ticket",
    name:"ticket",
    component: () => import("../components/allbook/component/ticket.vue")
  },
  {
    path:"/setting",
    name:"setting",
    component: ()=> import("../components/allbook/component/contact.vue")
  },
  {
    path:"/allticket",
    name:"allticket",
    component: ()=> import ("../components/allbook/component/allticket.vue")
  }
];

const router = new VueRouter({
  routes,
});

export default router;
