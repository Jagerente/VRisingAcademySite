import Main from "@/pages/Main";
import ItemsPage from "@/pages/ItemsPage"
import SpellsPage from "@/pages/SpellsPage"
import BloodTypesPage from "@/pages/BloodTypesPage"
import ContributePage from "@/pages/ContributePage"
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
        path: '/bloodtypes',
        component: BloodTypesPage
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
        path: '/contribute',
        component: ContributePage
    },
    {
        path: '/contribute/books',
        component: ContributePage,
        beforeEnter(to, from, next) {
            window.open("https://docs.google.com/forms/d/e/1FAIpQLSe4AtNpDE91s9xNi2JHRtKNle4WcotFB266wGjFRQ3RXP9aYQ/viewform?usp=sf_link", '_blank');
        }
    },
    {
        path: '/contribute/cosmetics',
        component: ContributePage,
        beforeEnter(to, from, next) {
            window.open("https://docs.google.com/forms/d/e/1FAIpQLSfR_8Dmx9dHCIieJDq2v41EIhVAS22cLxjyU8F6ZdFwJvifhg/viewform?usp=sf_link", '_blank');
        }
    },
    {
        path: '/contribute/servants/bloodquality',
        component: ContributePage,
        beforeEnter(to, from, next) {
            window.open("https://docs.google.com/forms/d/e/1FAIpQLSfKpLIMMY6ELxoYwnsEE3J5pDG2Y5PkAPNVYvbxo81HZdtk-g/viewform?usp=sf_link", '_blank');
        }
    },
    {
        path: '/credits/bobrokrot/discord',
        beforeEnter(to, from, next) {
            window.open("https://discordapp.com/users/142979124695138304", '_blank');
        }
    },
    {
        path: '/credits/jagerente/discord',
        beforeEnter(to, from, next) {
            window.open("https://discordapp.com/users/584324991471058946", '_blank');
        }
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