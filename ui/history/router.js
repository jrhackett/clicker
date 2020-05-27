import Home from 'components/home'
import { writable } from 'svelte/store'

const router = {
    '/': Home
}

export default router
export const curRoute = writable('/')
