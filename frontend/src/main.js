import { createApp } from 'vue'
import App from './App'
import components from '@/components/UI'
import router from "@/router/router"
import directives from '@/directives'
import store from '@/store'
// import "bootstrap/dist/css/bootstrap.css"
import "bootstrap/dist/js/bootstrap.js"
import "@/assets/vrising.scss"
import PerfectScrollbar from 'vue3-perfect-scrollbar';
import 'vue3-perfect-scrollbar/dist/vue3-perfect-scrollbar.css'

const app = createApp(App)

components.forEach(component => {
    app.component(component.name, component)
})

directives.forEach(directive => {
    app.directive(directive.name, directive)
})

app
    .use(router)
    .use(store)
    .use(PerfectScrollbar)
    .mount('#app')
