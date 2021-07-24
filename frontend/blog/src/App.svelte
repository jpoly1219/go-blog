<script>
	import Home from "./Home.svelte"
	import Login from "./Login.svelte"
	import Signup from "./Signup.svelte"
	import Notfound from "./Notfound.svelte"
	import Navbar from "./Navbar.svelte"
	import Singlepost from "./Singlepost.svelte"
	import Profile from "./Profile.svelte"
	import Editprofile from "./Editprofile.svelte"
	import Write from "./Write.svelte"
	import Editpost from "./Editpost.svelte"
	import { accessToken, activePage, authenticated, currentUser, expiration } from "./stores.js"
	import { afterUpdate, onMount } from "svelte";

	const pageMap = {
		home: Home,
		login: Login,
		signup: Signup,
		notfound: Notfound,
		singlepost: Singlepost,
		profile: Profile,
		editprofile: Editprofile,
		write: Write,
		editpost: Editpost
	}

	// onMount (so basically on reload or close/reopen tab)
	// try silent refresh: check for access token in stores.js
	// if exists and is valid, use that
	// if doesn't exist OR exists but invalid, use refresh token to get new refresh and access tokens
	// if refresh token is invalid, user has to manually log in

	async function refresh() {
		const options = {
			method: "POST",
			credentials: "include"
		}
		try {
			const res = await fetch("http://jpoly1219devbox.xyz:8090/auth/refresh", options)
			const json = await res.json()
			accessToken.set(json.accessToken)
			let payloadB64 = $accessToken.split(".")[1]
			expiration.set(JSON.parse(window.atob(payloadB64)).exp)
			authenticated.set(true)
		} catch(err) {
			alert(err)
		}
	}

	onMount(() => {
		let user = localStorage.getItem("user")
		let at = $accessToken
		console.log(user)
		if (user != null) {
			if (at == "") {
				console.log("refreshing")
				refresh()
			}
		}
	})

	function refreshTimer() {
		if ($expiration != "") {
			var i = Date.now()/1000;
			var timer = setInterval(() => {
				console.log($expiration)
				if (i >= Number($expiration)) {
					if (localStorage.getItem("user") != null) {
						refresh()
						clearInterval(timer)
					}
				}	
				console.log(Date.now()/1000)
				i++
			}, 1000)
		}
	}

	$: $accessToken, refreshTimer()
</script>

<main>
	<Navbar/>
	<svelte:component this={pageMap[$activePage]}/>
</main>

<style global>
	@tailwind base;
	@tailwind components;
	@tailwind utilities;
</style>