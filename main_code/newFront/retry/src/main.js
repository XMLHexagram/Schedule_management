import Vue from 'vue'
import App from './App.vue'
import {Row,Col} from "vant";

Vue.config.productionTip = false
Vue.use(Row).use(Col)

new Vue({
  render: h => h(App),
}).$mount('#app')
