import { writable } from "svelte/store"

export const authenticated = writable(false)
export const accessToken = writable("")
export const expiration = writable("")
export const currentUser = writable("")
export const activePage = writable("home")
export const postId = writable("")
export const viewUser = writable("")
export const postToEdit = writable({"id": "", "title": "", "author": "", "content": ""})