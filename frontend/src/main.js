import sha256 from 'js-sha256'
import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementUI from 'element-ui';
import axios from 'axios'
import 'element-ui/lib/theme-chalk/index.css';


import "video.js/dist/video-js.css";
// import "vue-video-player/src/custom-theme.css";
import videoPlayer from "vue-video-player";
import "videojs-contrib-hls";
// import 'video.js'
// import "videojs-flash";
// import Vue from "vue";
Vue.use(videoPlayer);


Vue.use(ElementUI);
Vue.config.productionTip = false;
Vue.prototype.$http = axios;        // 把axios赋给全局变量http
Vue.prototype.$encrypt = sha256;
// axios.defaults.baseURL = 'http://localhost:8000'    // 全局地址
axios.defaults.baseURL = '/api'    // 全局地址

axios.interceptors.request.use(
  config => {
    if (localStorage.getItem('Authorization')) {
      config.headers.Authorization = localStorage.getItem('Authorization');
    }
    return config;
  },
  error => {
    return Promise.reject(error);
  }
)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
