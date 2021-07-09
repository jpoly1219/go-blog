<script>
    import { onDestroy, onMount } from "svelte"
    import Post from "./Post.svelte"
    
	let pageNumber = 0
	let postList = []
	let hasMore = true
	let loadingDiv

	const fetchPosts = async (page) => {
		let apiURL = "http://jpoly1219devbox.xyz:8090/posts/batch/" + page
		let res = await fetch(apiURL)
		let post = await res.json()

		if (post.length > 0) {
			postList = [...postList, ...post]
			console.log(postList)
		}
		else {
			hasMore = false
		}
	}

	function handleIntersect() {
		pageNumber++
		if (hasMore) {
			fetchPosts(pageNumber)
		}
	}

	let observer = new IntersectionObserver((entries) => {
		entries.forEach((entry) => {
			if (entry.isIntersecting) {
				handleIntersect()
			}
		})
	})

	onMount(async () => {
		observer.observe(loadingDiv)
		fetchPosts(pageNumber)
	})
	onDestroy(() => {
		observer.disconnect()
	})
</script>

<div class="container mx-auto w-1/3 flex flex-col">
	{#if postList}
		{#each postList as post}
			<Post post={post} on:message/>
		{/each}
	{/if}
	{#if hasMore}
		<div bind:this={loadingDiv} class="text-center my-48">
			<span class="text-2xl font-medium text-gray-500">Loading new posts...</span>
		</div>
	{/if}
</div>

<!-- https://svelte.dev/repl/4b8ccdf1d01545baa0ab6a858bc05abb?version=3.32.1 -->
<!-- https://levelup.gitconnected.com/loading-more-results-on-scroll-with-svelte-js-restful-apis-svelte-infinite-scrolling-ad80a09b5e33 -->