<script>
    export let post;
</script>
<div class="container mx-auto px-96 pb-10 flex flex-col">
    <div class="border-b-2">
        <h2 class="text-2xl font-medium text-gray-900 mt-4 mb-4">
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