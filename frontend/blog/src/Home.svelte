<script>
    import { onMount } from "svelte"
	import { pagination } from "./stores.js"
    import Post from "./Post.svelte"
    
    const apiURL = "http://jpoly1219devbox.xyz:8090/posts"
	let postList = []

	onMount(async () => {
		const res = await fetch(apiURL)
		postList = await res.json()
		console.log(postList)
	})
</script>

<div class="container mx-auto w-1/3 flex flex-col">
	{#each postList as post}
    	<Post post={post} on:message/>
	{/each}
	{#if postList != []}
	<div class="text-center my-48">
		<span class="text-2xl font-medium text-gray-500">Loading new posts...</span>
	</div>
	{/if}
</div>

<!-- https://svelte.dev/repl/4b8ccdf1d01545baa0ab6a858bc05abb?version=3.32.1 -->
<!-- https://levelup.gitconnected.com/loading-more-results-on-scroll-with-svelte-js-restful-apis-svelte-infinite-scrolling-ad80a09b5e33 -->