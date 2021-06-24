<script>
    import { onMount } from "svelte"
    import Post from "./Post.svelte"
    
    const apiURL = "http://jpoly1219devbox.xyz:8090/posts"
	let postList = []

	onMount(async () => {
		const res = await fetch(apiURL)
		postList = await res.json()
	})
</script>

{#each postList as post}
    <Post post={post} on:message/>
{/each}