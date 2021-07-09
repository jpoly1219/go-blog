<script>
    import { activePage, postId, viewUser } from "./stores"
    import { fade } from "svelte/transition"
    export let post
    
    function setPostId() {
        postId.set(post.id)
        activePage.set("singlepost")
    }

    function goToProfile() {
        viewUser.set(post.author)
        activePage.set("profile")
    }

    let originalContent = post.content
    let wordArray = originalContent.split(" ")

    function formatContent(wordArray) {
        wordArray = wordArray.slice(0, 50)
        const formattedContent = wordArray.join(" ").concat("...")
        return formattedContent
    }

    let newContent
    if (wordArray.length > 50) {
        newContent = formatContent(wordArray)
    }
    else {
        newContent = post.content
    }

</script>

<div in:fade class="border-b-2">
    <h2 on:click={setPostId} class="text-2xl font-medium text-gray-900 mt-4 mb-4 cursor-pointer">
        {post.title}
    </h2>
    <h3 on:click={goToProfile} class="font-medium text-gray-900 cursor-pointer">
        {post.author}
    </h3>
    <p class="leading-relaxed mb-8">
        {newContent}
    </p>
</div>