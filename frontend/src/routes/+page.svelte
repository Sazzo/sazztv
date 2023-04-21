<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<div class="flex h-screen flex-col justify-center items-center">
	<h1>The world worst streaming platform.</h1>
	{#if data.currentUser}
		<h1>Logged as: {data.currentUser.username} ({data.currentUser.id})</h1>
	{/if}
</div>
