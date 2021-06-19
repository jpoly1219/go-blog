<script>
	import Home from "./Home.svelte";
	import Login from "./Login.svelte";
	import Signup from "./Signup.svelte"
	import Notfound from "./Notfound.svelte"
	import Navbar from "./Navbar.svelte";
	import { activePage } from "./stores.js"

	const pageMap = {
		home: Home,
		login: Login,
		signup: Signup,
		notfound: Notfound
	};
</script>

<main>
	<Navbar/>
	<svelte:component this={pageMap[$activePage]}/>
</main>

<style global>
	@tailwind base;
	@tailwind components;
	@tailwind utilities;
</style>