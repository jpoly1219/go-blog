<script>
    import { isSinglePost } from "./stores"
    export let singlePost;
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
    <button type="button" on:click={() => isSinglePost.set(false)} class="border border-gray-500 rounded-lg p-3">
        <span class="mx-3 my-2">Go back to list</span>
    </button>
</div>