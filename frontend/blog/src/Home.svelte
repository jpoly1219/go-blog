<script>
    import { onMount } from "svelte"
    import Post from "./Post.svelte"
    
    const apiURL = "http://jpoly1219devbox.xyz:8090/posts";
	let data = [];

	onMount(async () => {
		const res = await fetch(apiURL);
		data = await res.json();
	});

	function loadSinglePost() {
		
	}
</script>

{#each data as item}
    <Post post={item} on:click={loadSinglePost}/>
{/each}