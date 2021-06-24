import { writable } from "svelte/store"

export const authenticated = writable(false)
export const accessToken = writable("")
export const expiration = writable("")
export const activePage = writable("home")
export const postId = writable("")