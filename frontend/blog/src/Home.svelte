<script>
    import Navbar from "./Navbar.svelte"
    import Post from "./Post.svelte"
    export let posts;
</script>

<Navbar/>
{#each posts as item}
    <Post post={item}/>
{/each}