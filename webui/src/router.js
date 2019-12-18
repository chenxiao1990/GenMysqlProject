import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import genselectcode from './views/genselectcode.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/genselectcode',
      name: 'genselectcode',
      component: genselectcode
    }  
  ]
})
