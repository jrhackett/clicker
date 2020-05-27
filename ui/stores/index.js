import { writable } from 'svelte/store'

export const playerStore = writable({
    liquid: 0,
    generators: []
})
