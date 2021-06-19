<script>
    import { authenticated, accessToken, activePage } from "./stores.js"

    let email = "";
    let password = "";
    // let result = null;

    async function submit() {
        let loginDetails = {
            email: email,
            password: password
        }

        const options = {
            method: "POST",
            body: JSON.stringify(loginDetails),
            headers: {
                "Content-Type": "application/json"
            }
        };
        const res = await fetch("http://jpoly1219devbox.xyz:8090/auth/login", options);
        const json = await res.json();
        // result = JSON.stringify(json);
        accessToken.set(json.accessToken);

        if ($accessToken != undefined && $accessToken != "") {
            authenticated.set(true);
            activePage.set("home");
        }
        else {
            alert("login failed.")
        }
    }
</script>

<div class="container mx-auto items-center flex flex-col text-gray-900">
    <div class="bg-white p-16 rounded-lg border-gray-900 shadow-lg my-44">
        <h2 class="text-2xl font-medium text-center mb-20">Welcome back.</h2>
        <form on:submit|preventDefault={submit} class="flex flex-col px-30">
            <div class="flex flex-col items-stretch px-30">
                <input bind:value={email} type="text" placeholder="Email address" class="border-b-2 py-4">
                <input bind:value={password} type="text" placeholder="Password" class="border-b-2 py-4">
            </div>
            <div class="flex flex-col items-center my-10">
                <button type="submit" class="border border-gray-500 rounded-lg p-3">
                    <span class="mx-3 my-2">Log in</span>
                </button>
            </div>
            <p class="text-sm font-extralight">
                Don't have an account yet? <span on:click={() => activePage.set("signup")} class="text-blue-400 cursor-pointer">Sign up</span> today.
            </p>
        </form>
    </div>
</div>