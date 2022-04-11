import Vue from 'vue';
import App from './App.vue';
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import NProgress from 'nprogress';
import 'nprogress/nprogress.css';


Vue.config.productionTip = false;

Vue.use(ElementUI);
Vue.use(NProgress);



new Vue({
  render: h => h(App),
}).$mount('#app')
