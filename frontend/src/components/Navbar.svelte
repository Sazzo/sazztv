<script>
	import { page } from '$app/stores';
	import LoginModal from './modals/LoginModal.svelte';
	import Logo from './Logo.svelte';

	let showLoginModal = false;
</script>

<nav class="flex h-14 items-center bg-white border-b-2">
	<div class="pl-4">
		<Logo />
	</div>

	<div class="pl-5 font-medium">
		<a href="/">Home</a>
	</div>

	{#if !$page.data.currentUser}
		<div class="ml-auto pr-4">
			<button
				class="border-2 p-2 rounded-lg text-sm font-medium"
				on:click={() => (showLoginModal = true)}>Log in</button
			>
		</div>
	{:else}
		<div class="ml-auto pr-4">
			<p>{$page.data.currentUser.username}</p>
		</div>
	{/if}
</nav>

<LoginModal bind:showLoginModal />
