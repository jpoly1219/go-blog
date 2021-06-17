<script>
    import Navbar from "./Navbar.svelte";

    let email = "";
    let password = "";
    let result = null;

    async function checkLogin() {
        let loginDetails = {
            email: email,
            password: password
        }

        const options = {
            method: 'POST',
            header: {
                "Content-Type": "text/plain"
            },
            body: JSON.stringify(loginDetails)
        };
        const res = await fetch("http://jpoly1219devbox.xyz:8090/auth/login", options);
        const json = await res.json();
        result = JSON.stringify(json);
        alert(result);
    }
</script>

<Navbar on:pressed/>
<div class="container mx-auto items-center flex flex-col text-gray-900">
    <div class="bg-white p-16 rounded-lg border-gray-900 shadow-lg my-44">
        <h2 class="text-2xl font-medium text-center mb-20">Welcome back.</h2>
        <form action="/" class="flex flex-col px-30">
            <div class="flex flex-col items-stretch px-30">
                <input bind:value={email} type="text" placeholder="Email address" class="border-b-2 py-4">
                <input bind:value={password} type="text" placeholder="Password" class="border-b-2 py-4">
            </div>
            <div class="flex flex-col items-center my-10">
                <button type="button" on:click={checkLogin} class="border border-gray-500 rounded-lg p-3">
                    <span class="mx-3 my-2">Log in</span>
                </button>
            </div>
            <p class="text-sm font-extralight">
                Don't have an account yet? <a href="/" class="text-blue-400">Sign up</a> today.
            </p>
        </form>
    </div>
</div>