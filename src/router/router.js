import Main from "@/pages/Main";
import BestiaryPage from "@/pages/BestiaryPage"

import { createRouter, createWebHistory } from "vue-router";

const routes = [
    {
        path: '/',
        component: Main
    },
    {
        path: '/bestiary',
        component: BestiaryPage
    },
    {
        path: '/map',
        component: BestiaryPage
    },
    {
        path: '/guides',
        component: BestiaryPage
    },
]

const router = createRouter({
    routes,
    history: createWebHistory()
});

export default router;