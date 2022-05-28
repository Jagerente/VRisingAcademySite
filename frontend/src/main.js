import { createApp } from 'vue'
import App from './App'
import components from '@/components/UI'
import router from "@/router/router"
import store from '@/store'
import Popper from "vue3-popper";
import "bootstrap/dist/js/bootstrap.js"
import "@/assets/vrising.scss"
import "@/assets/styles/va_variables.scss"

const app = createApp(App)

components.forEach(component => {
    app.component(component.name, component)
})

app
    .use(router)
    .use(store)
    .use(Popper)
    .mount('#app')
