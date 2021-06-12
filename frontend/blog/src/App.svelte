<script>
	import { onMount } from "svelte";
	import { each } from "svelte/internal";
	import Navbar from "./Navbar.svelte";
	import Post from "./Post.svelte";

	const apiURL = "http://jpoly1219devbox.xyz:8090/posts";
	let data = [];

	onMount(async () => {
		const response = await fetch(apiURL);
		data = await response.json();
	});
</script>

<main>
	<Navbar/>
	{#each data as item}
		<Post post={item}/>
	{/each}
</main>

<style>
	main {
		display: grid;
		grid-template-columns: 15% 15% 40% 15% 15%;
	}
</style>