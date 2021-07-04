<script>
    import { accessToken, activePage, currentUser } from "./stores.js"

    function resize() {
        this.style.height = 'auto'
        this.style.height = this.scrollHeight + 'px'
    }

    let title = ""
    let author = $currentUser
    let content = ""
    async function submit() {
        let postDetails = {
            title: title,
            author: author,
            content: content
        }

        const options = {
            method: "POST",
            body: JSON.stringify(postDetails),
            headers: {
                "Authorization": "Bearer " + $accessToken,
                "Content-Type": "application/json"
            },
            credentials: "include"
        }
        const res = await fetch("http://jpoly1219devbox.xyz:8090/post", options)
        alert("post submitted!")
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
        <textarea bind:value={content} on:input={resize} placeholder="Tell your story..." class="place-self-stretch border-none p-2 font-medium text-gray-900"></textarea>
    </div>
    <div class="flex flex-row place-content-center my-5">
        <button on:click={submit} class="inline-flex items-center bg-blue-400 border rounded-lg text-base p-3 mr-5">
            <span class="text-white">Publish</span>
        </button>
        <button on:click={cancel} class="inline-flex items-center border border-gray-500 rounded-lg p-3">
            <span class="text-gray-900">Cancel</span>
        </button>
    </div>
</div>