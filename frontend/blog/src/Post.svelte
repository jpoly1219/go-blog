<script>
    export let post;
</script>

<div>
    <h3>{post.title}</h3>
    <h4>{post.author}</h4>
    <p>{post.content}</p>
</div>

<style>
    div {
        border: solid 2px grey;
        border-radius: 30px;
        margin: 1em;
        padding: 10px;
		grid-column: 3 / 4;
    }
</style>