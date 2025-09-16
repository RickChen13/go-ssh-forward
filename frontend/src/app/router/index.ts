import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router';
import RootRouter from '@/app/router/node/root';

const routes: RouteRecordRaw[] = RootRouter; //.concat(IndexRouter,ArticleRouter)
const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes
});

export default router;
