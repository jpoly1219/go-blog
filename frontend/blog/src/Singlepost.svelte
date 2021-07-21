<script>
    import { onMount } from "svelte"
    import { accessToken, activePage, authenticated, postId, postToEdit } from "./stores"

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

    function handleEdit() {
        postToEdit.set(singlePost)
        activePage.set("editpost")
    }

    let clickedDelete = false;
    function handleDelete() {
        if (!clickedDelete) {
            clickedDelete = true;
        }
    }
    async function handleDeleteConfirm() {
        const options = {
            method: "DELETE",
            headers: {
                "Authorization": "Bearer " + $accessToken
            },
            credentials: "include"
        }
        const res = await fetch("http://jpoly1219devbox.xyz:8090/posts/"+$postId, options)

        alert("post deleted!")
        clickedDelete = false;
        activePage.set("profile")
    }
    function handleDeleteCancel() {
        clickedDelete = false;
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
        <button type="button" on:click={handleEdit} class="bg-blue-400 rounded-lg p-3 flex-grow-0">
            <span class="mx-3 my-2 text-white">Edit post</span>
        </button>
        <button type="button" on:click={handleDelete} class="bg-red-400 rounded-lg p-3 flex-grow-0">
            <span class="mx-3 my-2 text-white">Delete post</span>
        </button>
        {#if clickedDelete}
        <div class="px-4 py-3 leading-normal text-white bg-red-700 rounded-lg" role="alert">
            <p>Are you sure you want to delete this post?</p>
            <div class="flex flex-row">
                <span on:click={handleDeleteConfirm} class="mx-2 my-1">Yes</span>
                <span on:click={handleDeleteCancel} class="mx-2 my-1">No</span>
            </div>
        </div>
        {/if}
    {/if}
    <button type="button" on:click={() => activePage.set("home")} class="border border-gray-500 rounded-lg p-3 flex-grow-0">
        <span class="mx-3 my-2">Go back to list</span>
    </button>
</div>