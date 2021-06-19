import { writable } from "svelte/store"

export const authenticated = writable(false);
export const accessToken = writable("");
export let activePage = writable("home");