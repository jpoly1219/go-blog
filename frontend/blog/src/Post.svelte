<script>
    import { createEventDispatcher } from "svelte";
    export let post;

    const dispatch = createEventDispatcher();
    let singlePost;
    
    async function loadSinglePost() {
        let id = post.id;
        let apiURL = "http://jpoly1219devbox.xyz:8090/posts/" + id;
        const res = await fetch(apiURL)
        singlePost = await res.json()
        
        dispatch("message", {
            post: singlePost
        });
    }
</script>

<div class="container mx-auto px-96 pb-10 flex flex-col">
    <div class="border-b-2">
        <h2 on:click={loadSinglePost} class="text-2xl font-medium text-gray-900 mt-4 mb-4 cursor-pointer">
            {post.title}
        </h2>
        <h3 class="font-medium text-gray-900">
            {post.author}
        </h3>
        <p class="leading-relaxed mb-8">
            {post.content}
        </p>
    </div>
</div>