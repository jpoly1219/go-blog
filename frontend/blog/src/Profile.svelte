<script>
    import { onMount } from "svelte"
    import Post from "./Post.svelte"
    import { viewUser } from "./stores";

    let userId
	let userData = {
        name: "",
        username: "",
        email: ""
    }
    let postData = []

	onMount(async () => {
        if ($viewUser != "") {
            const res = await fetch("http://jpoly1219devbox.xyz:8090/"+$viewUser)
            userData = await res.json()
            userId = userData.id.toString()
            
            const res2 = await fetch("http://jpoly1219devbox.xyz:8090/"+userId+"/posts")
            postData = await res2.json()
            console.log(postData)
        }
	})
</script>

<div class="flex flex-row static">
    <div class="fixed w-1/3">
        <div class="absolute top-0 right-10 w-1/2 mx-auto bg-white rounded-lg overflow-hidden shadow">
            <div class="bg-cover h-40" style="background-image: url(https://free4kwallpapers.com/uploads/originals/2021/04/14/minimal-pacman-wallpaper.jpg)"></div>
            <div class="border-b px-4 pb-4">
                <div class="flex flex-col text-center">
                    <img class="place-self-center h-32 w-32 rounded-full border-white border-4 -mt-16" src="https://widgetwhats.com/app/uploads/2019/11/free-profile-photo-whatsapp-4-300x300.png" alt>
                    <div class="pt-2 pb-4">
                        <h3 class="font-bold text-2x1 mb-1">{userData.name}</h3>
                        <p class="text-gray-500">{userData.email}</p>
                        <p class="text-gray-500">{userData.username}</p>
                    </div>
                    <button type="button" class="place-self-center bg-blue-400 text-white text-base rounded-lg px-3 py-2">
                        <span>Follow</span>
                    </button>
                </div>
            </div>
        </div>
    </div>
    <div class="w-1/3"></div>
    <div class="container mx-auto w-1/3 flex flex-col">
        {#if postData != null}
            {#each postData as post}
                <Post post={post}/>
            {/each}
        {/if}
    </div>
    <div class="w-1/3"></div>
</div>