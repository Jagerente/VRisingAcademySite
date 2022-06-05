import Main from "@/pages/Main";
import ItemsPage from "@/pages/ItemsPage"
import SpellsPage from "@/pages/SpellsPage"
import NotFoundPage from "@/pages/NotFoundPage"

import { createRouter, createWebHistory } from "vue-router";

const routes = [
    {
        path: '/:pathMatch(.*)*',
        redirect: '/404'
    },
    {
        path: '/',
        component: Main
    },
    {
        path: '/items',
        component: ItemsPage
    },
    {
        path: '/spells',
        component: SpellsPage
    },
    {
        path: '/map',
        component: ItemsPage
    },
    {
        path: '/guides',
        component: ItemsPage
    },
    {
        path: '/404',
        name: '404',
        component: NotFoundPage
    },
]

const router = createRouter({
    routes,
    history: createWebHistory()
});

export default router;