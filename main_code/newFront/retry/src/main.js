import Vue from 'vue'
import App from './App.vue'
import {Row,Col,Field,DatetimePicker} from "vant";

Vue.config.productionTip = false
Vue.use(Row).use(Col);
Vue.use(Field);
Vue.use(DatetimePicker);

new Vue({
  render: h => h(App),
}).$mount('#app')
