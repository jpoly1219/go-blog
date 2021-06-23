<script>
	import Home from "./Home.svelte";
	import Login from "./Login.svelte";
	import Signup from "./Signup.svelte";
	import Notfound from "./Notfound.svelte";
	import Navbar from "./Navbar.svelte";
	import Singlepost from "./Singlepost.svelte";
	import { accessToken, activePage, expiration, isSinglePost } from "./stores.js";

	const pageMap = {
		home: Home,
		login: Login,
		signup: Signup,
		notfound: Notfound
	};

	let singlePost;
	function loadSinglePost(event) {
		singlePost = event.detail.post
		console.log(singlePost)
		isSinglePost.set(true)
	}

	function refreshTimer() {
		if ($expiration != "") {
			var i = Date.now()/1000;
			var timer = setInterval(() => {
				if (i >= Number($expiration)) {
					(async () => {
						console.log("running async")
						const options = {
							method: "POST",
							credentials: "include"
						};
						const res = await fetch("http://jpoly1219devbox.xyz:8090/auth/refresh", options);
						const json = await res.json();
						accessToken.set(json.accessToken);
						console.log($accessToken);
						let payloadB64 = $accessToken.split(".")[1];
						expiration.set(JSON.parse(window.atob(payloadB64)).exp);
					})();
					clearInterval(timer);
				}	
				console.log(Date.now()/1000);
				i++;
			}, 1000)
		}
	};

	$: $accessToken, refreshTimer();
</script>

<main>
	<Navbar/>
	{#if $isSinglePost == true}
		<Singlepost singlePost={singlePost}/>
	{:else}
		<svelte:component this={pageMap[$activePage]} on:message={loadSinglePost}/>
	{/if}
</main>

<style global>
	@tailwind base;
	@tailwind components;
	@tailwind utilities;
</style>