<script>
    import Navbar from "./Navbar.svelte";

</script>

<Navbar on:pressed/>
<div class="container mx-auto items-center flex flex-col text-gray-900">
    <div class="bg-white p-16 rounded-lg border-gray-900 shadow-lg my-20">
        <h2 class="text-2xl font-medium text-center mb-20">Join Goblog.</h2>
        <form action="/" class="flex flex-col">
            <div class="flex flex-col items-stretch">
                <input type="text" placeholder="Name" class="border-b-2 py-4">
                <input type="text" placeholder="Email address" class="border-b-2 py-4">
                <input type="text" placeholder="Username" class="border-b-2 py-4">
                <input type="text" placeholder="Password" class="border-b-2 py-4">
            </div>
            <div class="flex flex-col items-center my-10">
                <button type="submit" class="border bg-blue-400 border-blue-900 rounded-lg p-3">
                    <span class="mx-3 my-2 text-white">Sign up</span>
                </button>
            </div>
            <p class="text-sm font-extralight">
                Already have an account? <a href="/" class="text-blue-400">Log in</a>.
            </p>
        </form>
    </div>
</div>