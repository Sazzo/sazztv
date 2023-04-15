<script>
	import Logo from './Logo.svelte';
	import { signIn, signOut } from '@auth/sveltekit/client';
	import { page } from '$app/stores';
</script>

<nav class="flex h-14 items-center bg-white border-b-2">
	<div class="pl-4">
		<Logo />
	</div>

	<div class="pl-5 font-medium">
		<a href="/">Home</a>
	</div>
	<div class="ml-auto pr-4">
		{#if !$page.data.session}
			<button class="border-2 p-2 rounded-lg text-sm font-medium" on:click={() => signIn()}
				>Login</button
			>
		{:else}
			<div class="flex flex-row items-center">
				<img
					src={$page.data.session.user?.image}
					alt={`Avatar de ${$page.data.session.user?.name}`}
					class="w-8 h-8 rounded-full"
				/>
			</div>
		{/if}
	</div>
</nav>
