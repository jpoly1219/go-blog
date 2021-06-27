<script>
    import { activePage, postId } from "./stores"
    export let post
    
    function setPostId() {
        postId.set(post.id)
        activePage.set("singlepost")
    }
</script>

<div class="border-b-2">
    <h2 on:click={setPostId} class="text-2xl font-medium text-gray-900 mt-4 mb-4 cursor-pointer">
        {post.title}
    </h2>
    <h3 class="font-medium text-gray-900">
        {post.author}
    </h3>
    <p class="leading-relaxed mb-8">
        {post.content}
    </p>
</div>