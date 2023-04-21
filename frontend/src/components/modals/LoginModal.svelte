<script lang="ts">
	import { applyAction, enhance } from '$app/forms';
	import { goto, invalidateAll } from '$app/navigation';
	import Modal from '../Modal.svelte';
	import FormInput from '../form/FormInput.svelte';

	export let showLoginModal = false;

	let loginFailureReason: string | null = null;
</script>

<Modal bind:showModal={showLoginModal}>
	<h1 class="font-bold text-2xl mb-3">Log in</h1>
	<form
		action="/"
		method="POST"
		use:enhance={() => {
			return async ({ result }) => {
				if (result.type === 'failure') {
					loginFailureReason = result.data?.message;
					return;
				}

				showLoginModal = false;
				await applyAction(result);
			};
		}}
	>
		<div class="flex flex-col mb-4">
			<label class="font-semibold text-sm" for="username">Username</label>
			<FormInput name="username" type="text" id="username" required />
		</div>

		<div class="flex flex-col mb-4">
			<label class="font-semibold text-sm" for="password">Password</label>
			<FormInput name="password" type="password" id="password" required />
		</div>

		<div class="flex justify-center items-center">
			<button
				class="bg-blue-600 text-white p-1 font-semibold text-sm w-96 rounded-[5px]"
				type="submit">Log in</button
			>
		</div>
	</form>

	{#if loginFailureReason}
		<p class="text-red-500 font-semibold text-sm pt-6">{loginFailureReason}</p>
	{/if}

	<a href="/" class="text-blue-600 font-semibold text-sm pt-6">Don't have an account? Sign up</a>
</Modal>
