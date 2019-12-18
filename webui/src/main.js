import Vue from 'vue'
import App from './App.vue'
import router from './router'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import axios from 'axios'

Vue.config.productionTip = false

Vue.use(ElementUI);

//这样各个组件就可以使用  this.$axios 来使用了，不必每次都引入axios
Vue.prototype.$axios = axios

//配置本地访问的baseurl
if (process.env.NODE_ENV == "development") {
 
  axios.defaults.baseURL = "http://127.0.0.1:8008"
 
}
new Vue({
  router,
  render: h => h(App)
}).$mount('#app')




