<script>
	import { page } from '$app/stores';
	import { signOut } from '@auth/sveltekit/client';

	console.log($page.data.session);
</script>

<div class="flex h-screen flex-col justify-center items-center">
	<h1>A pior plataforma de streaming do mundo.</h1>

	<h2>Logado como: {$page.data.session?.user?.name}</h2>
	<button class="border-2 p-2 rounded-lg font-medium" on:click={() => signOut()}>Logout</button>
</div>
