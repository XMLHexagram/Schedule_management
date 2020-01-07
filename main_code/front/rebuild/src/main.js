import Vue from 'vue'
import App from './App.vue'
import router from './router'
import { Tab, Tabs,SwipeCell, Cell, CellGroup,Button } from 'vant';


Vue.config.productionTip = false

Vue.use(Tab).use(Tabs);
Vue.use(SwipeCell);
Vue.use(Cell).use(CellGroup);
Vue.use(Button);

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
