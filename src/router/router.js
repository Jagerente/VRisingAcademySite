import Main from "@/pages/Main";
import ItemsPage from "@/pages/ItemsPage"

import { createRouter, createWebHistory } from "vue-router";

const routes = [
    {
        path: '/',
        component: Main
    },
    {
        path: '/items',
        component: ItemsPage
    },
    {
        path: '/map',
        component: ItemsPage
    },
    {
        path: '/guides',
        component: ItemsPage
    },
]

const router = createRouter({
    routes,
    history: createWebHistory()
});

export default router;