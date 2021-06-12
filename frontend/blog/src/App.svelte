<script>
	import { onMount } from "svelte";
	import { each } from "svelte/internal";

	const apiURL = "http://jpoly1219devbox.xyz:8090/posts";
	let data = [];

	onMount(async () => {
		const response = await fetch(apiURL);
		data = await response.json();
	});
</script>

<main>
	{#each data as item}
		<div>
			<p>{item.title}</p>
			<p>{item.author}</p>
			<p>{item.content}</p>
		</div>
	{/each}
</main>

<style>
	div {
		border: solid 2px grey;
		border-radius: 30px;
		margin: 1em;
		padding: 10px;
	}
</style>