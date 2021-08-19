import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementUI from 'element-ui';
import axios from 'axios'
import 'element-ui/lib/theme-chalk/index.css';

Vue.use(ElementUI);
Vue.config.productionTip = false; 
Vue.prototype.$http = axios;        // 把axios赋给全局变量http
// axios.defaults.baseURL = 'http://localhost:8000'    // 全局地址
axios.defaults.baseURL = '/api'    // 全局地址

new Vue({
  router, 
  store,
  render: h => h(App)
}).$mount('#app')
