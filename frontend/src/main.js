import { createApp } from 'vue';
import App from './App';
import components from '@/components/UI';
import customComponents from '@/components/custom';
import router from '@/router/router';
import store from '@/store';

// import "bootstrap/dist/js/bootstrap.js";
// import "@/assets/styles/va_variables.scss";
// import "@/assets/styles/va_styles.scss";
// import "@/assets/styles/utility/va_slider.scss";

const app = createApp(App);

components.forEach(component => {
    app.component(component.name, component)
});

customComponents.forEach(component => {
    app.use(component)
});

app
    .use(router)
    .use(store)
    .mount('#app');
