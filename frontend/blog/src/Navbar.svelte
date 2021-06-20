<script>
import { beforeUpdate } from "svelte";

    import { authenticated, accessToken, activePage } from "./stores.js"

    function logout() {
        authenticated.set(false);
        accessToken.set("");
        activePage.set("home");
    }

    let username = "";
    beforeUpdate(() => {
        if ($authenticated == true) {
            let payloadB64 = $accessToken.split(".")[1]
            username = JSON.parse(window.atob(payloadB64)).user_name
            console.log(JSON.parse(window.atob(payloadB64)).exp)
            setTimeout(() => console.log(JSON.parse(window.atob(payloadB64)).exp), 2000)
        }
    });
</script>

<header class="sticky top-0 bg-white text-gray-600 body-font z-10">
    <div class="container mx-auto flex flex-wrap p-5 flex-col md:flex-row items-center">
        <div on:click={() => activePage.set("home")} class="flex title-font font-medium items-center text-gray-900 mb-4 md:mb-0 cursor-pointer">
            <span class="ml-3 text-xl">Goblog</span>
        </div>
        <nav class="md:ml-auto flex flex-wrap item-center text-base justify-center">
            {#if $authenticated == false}
                <button on:click={() => activePage.set("login")} class="mr-5 inline-flex items-center border border-gray-500 rounded-lg">
                    <span class="mx-3 my-2">Log in</span>
                </button>
                <button on:click={() => activePage.set("signup")} class="inline-flex items-center bg-blue-400 border border-blue-900 rounded-lg text-base">
                    <span class="mx-3 my-2 text-white">Sign up</span>
                </button>
            {:else}
                <span>{username}</span>
                <button on:click={logout} class="inline-flex items-center bg-blue-400 border border-blue-900 rounded-lg text-base">
                    <span class="mx-3 my-2 text-white">Log out</span>
                </button>
            {/if}
        </nav>
    </div>
</header>