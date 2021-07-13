<script>
import { onMount } from "svelte";

    import { accessToken, activePage, currentUser, postToEdit } from "./stores.js"

    function resize() {
        this.style.height = 'auto'
        this.style.height = this.scrollHeight + 'px'
    }

    let contentDiv
    onMount(() => {
        contentDiv.style.height = 'auto'
        contentDiv.style.height = contentDiv.scrollHeight + 'px'
    })

    let title = $postToEdit.title
    let author = $currentUser
    let content = $postToEdit.content
    async function submit() {
        let postDetails = {
            title: title,
            author: author,
            content: content
        }

        const options = {
            method: "PUT",
            body: JSON.stringify(postDetails),
            headers: {
                "Authorization": "Bearer " + $accessToken,
                "Content-Type": "application/json"
            },
            credentials: "include"
        }
        const res = await fetch("http://jpoly1219devbox.xyz:8090/posts/" + $postToEdit.id, options)
        alert("post editted!")
        activePage.set("profile")
    }

    function cancel() {
        alert("canceled")
        activePage.set("home")
    }
</script>

<div class="container mx-auto px-96 pb-10 flex flex-col">
    <div class="flex flex-col my-5">
        <input bind:value={title} type="text" placeholder="Title" class="place-self-stretch border-none p-2 text-2xl font-medium mb-5 text-gray-900">
        <textarea bind:value={content} bind:this={contentDiv} on:input={resize} placeholder="Tell your story..." class="place-self-stretch border-none p-2 font-medium text-gray-900"></textarea>
    </div>
    <div class="flex flex-row place-content-center my-5">
        <button on:click={submit} class="inline-flex items-center bg-blue-400 border rounded-lg text-base p-3 mr-5">
            <span class="text-white">Edit</span>
        </button>
        <button on:click={cancel} class="inline-flex items-center border border-gray-500 rounded-lg p-3">
            <span class="text-gray-900">Cancel</span>
        </button>
    </div>
</div>