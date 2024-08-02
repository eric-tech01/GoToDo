import NProgress from 'nprogress'; // progress bar
import 'nprogress/nprogress.css';
import { createRouter, createWebHashHistory } from 'vue-router';
import HomeView from '../views/home/index.vue';

import SettingView from "../views/setting/index.vue";

NProgress.configure({
  easing: 'ease',  // 动画方式，和css动画属性一样（默认：ease）
  speed: 500,  // 递增进度条的速度，单位ms（默认： 200）}); // NProgress Configuration
})
const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: '/',
      name: 'Home',
      component: HomeView
    }, {
      path: '/setting',
      name: 'Setting',
      component: SettingView
    },
  ],
  scrollBehavior() {
    return { top: 0 };
  },
});




export default router