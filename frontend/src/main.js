import { createApp } from 'vue'
import App from './App.vue'
import mdiVue from 'mdi-vue/v3'
import * as mdijs from '@mdi/js'
import router from './router'
import i18n from './i18n'

import './output.css';
createApp(App)
  .use(i18n)
  .use(router)
  .use(mdiVue, {
    icons: mdijs
  })
  .mount('#app')
