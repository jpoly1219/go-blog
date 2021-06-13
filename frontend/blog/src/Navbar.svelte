<script>
    import { Link } from "svelte-routing";
</script>

<header class="sticky top-0 bg-white text-gray-600 body-font z-10">
    <div class="container mx-auto flex flex-wrap p-5 flex-col md:flex-row items-center">
        <Link to="/" class="flex title-font font-medium items-center text-gray-900 mb-4 md:mb-0">
            <span class="ml-3 text-xl">Goblog</span>
        </Link>
        <nav class="md:ml-auto flex flex-wrap item-center text-base justify-center">
            <Link to="/login">
                <button class="mr-5 inline-flex items-center border border-gray-500 rounded-lg">
                    <span class="mx-3 my-2">Log in</span>
                </button>
            </Link>
            <button class="inline-flex items-center bg-blue-400 border border-blue-900 rounded-lg text-base">
                <Link to="/signup" class="mx-3 my-2 text-white">Sign up</Link>
            </button>
        </nav>
    </div>
</header>