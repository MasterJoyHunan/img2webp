import Vue from 'vue'
import App from './App.vue'
import { Upload, Table, TableColumn, Button, Input, Message } from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';

Vue.use(Upload)
Vue.use(Table)
Vue.use(TableColumn)
Vue.use(Button)
Vue.use(Input)
Vue.prototype.$message = Message
Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')

