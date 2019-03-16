import Vue from "vue";
import Router from "vue-router";
import VueChartkick from "vue-chartkick";
import Chart from "chart.js";
import Devices from "./views/Devices.vue";
import RegisterDevice from "./views/RegisterDevice.vue";
import UnRegisterDevice from "./views/UnRegisterDevice.vue";
import Shadow from "./views/Shadow.vue";
import DeviceManagement from "./views/DeviceManagement.vue";
import LoginUser from "./views/LoginUser.vue";
import RegisterUser from "./views/RegisterUser.vue";
import LogoutUser from "./views/LogoutUser.vue";
import AccountManagement from "./views/AccountManagement.vue";
import HomeDashboard from "./views/HomeDashboard.vue";

Vue.use(Router, VueChartkick, { adapter: Chart });

const router = new Router({
  mode: "history",
  base: process.env.BASE_URL,
  routes: [
    {
      path: `/devices`,
      name: "Devices",
      component: Devices
    },
    {
      path: "/devices/manage",
      name: "Device Management",
      component: DeviceManagement
    },
    {
      path: "/devices/:id",
      name: "Device Shadow",
      component: Shadow
    },
    {
      path: "/devices/register",
      name: "Register Device",
      component: RegisterDevice
    },
    {
      path: "/namespaces/:namespace/devices/:id/delete",
      name: "Delete Device",
      component: UnRegisterDevice
    },
    {
      path: "/user/register",
      name: "Register new User",
      component: RegisterUser
    },
    {
      path: "/user/login",
      name: "Login",
      component: LoginUser
    },
    {
      path: "/user/logout",
      name: "Logout",
      component: LogoutUser
    },
    {
      path: "/accounts",
      name: "Account Management",
      component: AccountManagement
    },
    {
      path: "/dashboard",
      name: "Dashboard",
      component: HomeDashboard
    },
    {
      path: "*",
      redirect: "/namespaces/:namespace/devices"
    }
  ]
});

router.beforeEach((to, from, next) => {
  if (localStorage.token) {
    next();
  } else if (to.path === "/user/login" || to.path === "/user/logout") {
    next();
  } else {
    localStorage.loginError = true;
    next("/user/login");
  }
});

export default router;
