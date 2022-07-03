import { createApp } from 'vue';
import App from './App';
import components from '@/components/UI';
import customComponents from '@/components/custom';
import router from '@/router/router';
import store from '@/store';

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