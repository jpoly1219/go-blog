<script>
	import { onMount } from "svelte";
	import Home from "./Home.svelte";
	import Login from "./Login.svelte";
	import Signup from "./Signup.svelte"
	import Notfound from "./Notfound.svelte"

	const apiURL = "http://jpoly1219devbox.xyz:8090/posts";
	let data = [];

	onMount(async () => {
		const response = await fetch(apiURL);
		data = await response.json();
	});

	let currentPage = "home";
	function handlePressed(event) {
		currentPage = event.detail;
	}
</script>

<main>
	{#if currentPage == "home"}
		<Home on:pressed={handlePressed} posts={data}/>
	{:else if currentPage == "login"}
		<Login on:pressed={handlePressed}/>
	{:else if currentPage == "signup"}
		<Signup on:pressed={handlePressed}/>
	{:else}
		<Notfound/>
	{/if}
</main>

<style global>
	@tailwind base;
	@tailwind components;
	@tailwind utilities;
</style>