import { createApp } from 'vue'
import App from './App.vue'
import { createRouter, createWebHistory } from "vue-router"
import VueCookies from "vue-cookies"

import { routes } from "./routes";

const router = createRouter({
    routes: routes,
    history: createWebHistory()
})

createApp(App)
    .use(router)
    .use(VueCookies)
    .mount('#app')
