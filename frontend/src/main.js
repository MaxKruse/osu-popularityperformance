import { createApp } from 'vue'
import App from './App.vue'
import { createRouter, createWebHistory } from "vue-router"

import { routes } from "./routes";

const router = createRouter({
    routes: routes,
    history: createWebHistory()
})

import { Quasar, Cookies, Notify } from 'quasar'
import quasarIconSet from 'quasar/icon-set/fontawesome-v6'

// Import icon libraries
import '@quasar/extras/fontawesome-v6/fontawesome-v6.css'

// A few examples for animations from Animate.css:
// import @quasar/extras/animate/fadeIn.css
// import @quasar/extras/animate/fadeOut.css

// Import Quasar css
import 'quasar/dist/quasar.css'

// Pinia Store
import { createPinia } from 'pinia'

createApp(App)
  .use(router)
  .use(createPinia())
  .use(Quasar, {
    plugins: {
      Cookies,
      Notify
    }, // import Quasar plugins and add here
    iconSet: quasarIconSet,
  })
  .mount('#app')
