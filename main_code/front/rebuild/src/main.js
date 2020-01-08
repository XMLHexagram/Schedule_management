import Vue from 'vue'
import App from './App.vue'
import router from './router'
import '@vant/touch-emulator';
import { Tab, Tabs,SwipeCell, Cell, CellGroup,Button,Popup,Overlay } from 'vant';


Vue.config.productionTip = false

Vue.use(Tab).use(Tabs);
Vue.use(SwipeCell);
Vue.use(Cell).use(CellGroup);
Vue.use(Button);
Vue.use(Popup);
Vue.use(Overlay);

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
