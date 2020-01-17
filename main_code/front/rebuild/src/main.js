import Vue from 'vue'
import App from './App.vue'
import router from './router'
import '@vant/touch-emulator';
import { Tab, Tabs,SwipeCell, Cell, CellGroup,Button,Overlay,Icon,Sticky,Popup,Field,Row,Col,DatetimePicker } from 'vant';
import 'bootstrap/dist/css/bootstrap.min.css'
// import 'bootstrap/dist/js/bootstrap.min'

Vue.config.productionTip = false

Vue.use(Tab).use(Tabs);
Vue.use(SwipeCell);
Vue.use(Cell).use(CellGroup);
Vue.use(Button);
Vue.use(Overlay);
Vue.use(Icon);
Vue.use(Sticky);
Vue.use(Popup);
Vue.use(Field);
Vue.use(Row).use(Col);
Vue.use(DatetimePicker);

const port=":12210";
const URL="http://www.sweetbeecr.com";
export const baseURL=URL+port;

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
