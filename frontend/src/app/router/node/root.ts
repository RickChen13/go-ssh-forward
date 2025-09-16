import { type RouteRecordRaw } from 'vue-router';

const root: Array<RouteRecordRaw> = [
    {
        path: '/',
        name: 'Index',
        component: () => import('@/app/views/index/Index.vue'),
        meta: {
            keepAlive: true
        }
    },
];

export default root;
