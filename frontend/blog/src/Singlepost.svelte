<script>
    import { onMount } from "svelte"
    import { activePage, authenticated, postId, postToEdit } from "./stores"

    let singlePost = {
        title: "",
        author: "",
        content: ""
    }

    onMount(async () => {
        let id = $postId
        console.log(id)
        let apiURL = "http://jpoly1219devbox.xyz:8090/posts/" + id
        const res = await fetch(apiURL)
        singlePost = await res.json()
    })

    function handleClick() {
        postToEdit.set(singlePost)
        activePage.set("editpost")
    }
</script>

<div class="container mx-auto px-96 pb-10 flex flex-col">
    <div class="border-b-2">
        <h2 class="text-2xl font-medium text-gray-900 mt-4 mb-4">
            {singlePost.title}
        </h2>
        <h3 class="font-medium text-gray-900">
            {singlePost.author}
        </h3>
        <p class="leading-relaxed mb-8">
            {singlePost.content}
        </p>
    </div>
    {#if $authenticated == true}
        <button type="button" on:click={handleClick} class="bg-blue-400 rounded-lg p-3 flex-grow-0">
            <span class="mx-3 my-2 text-white">Edit post</span>
        </button>
    {/if}
    <button type="button" on:click={() => activePage.set("home")} class="border border-gray-500 rounded-lg p-3 flex-grow-0">
        <span class="mx-3 my-2">Go back to list</span>
    </button>
</div>