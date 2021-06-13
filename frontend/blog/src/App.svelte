<script>
	import { onMount } from "svelte";
	import Home from "./Home.svelte";
	import Login from "./Login.svelte";
	import { Router, Link, Route } from "svelte-routing";

	export let url = "";

	const apiURL = "http://jpoly1219devbox.xyz:8090/posts";
	let data = [];

	onMount(async () => {
		const response = await fetch(apiURL);
		data = await response.json();
	});
</script>

<Router url="{url}">
	<Route path="/"><Home posts={data}/></Route>
	<Route path="/login"><Login/></Route>
</Router>

<style global>
	@tailwind base;
	@tailwind components;
	@tailwind utilities;
</style>