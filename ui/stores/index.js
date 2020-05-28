import { writable } from 'svelte/store'

export const playerStore = writable({
    liquid: 0,
    generators: []
})

export const userStore = writable({
    id: 0,
    name: '',
    email: '',
    image_url: ''
})
