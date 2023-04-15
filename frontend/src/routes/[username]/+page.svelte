<script lang="ts">
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	import Hls from 'hls.js';
	import { invalidateAll } from '$app/navigation';

	export let data: PageData;

	let videoElement: HTMLVideoElement;
	let hls: Hls;

	if (data.user.stream) {
		// todo: hardcoded streaming server url
		const streamPlaybackUrl = `http://localhost/live/${data.user.username}/index.m3u8`;

		let noUpdateTimeout: number;

		onMount(() => {
			hls = new Hls();

			videoElement.addEventListener('playing', () => {
				console.log('stream is playing');

				clearTimeout(noUpdateTimeout);
			});

			videoElement.addEventListener('waiting', async () => {
				console.log('stream is waiting');
				await invalidateAll(); // invalidate the user data to check if the stream is still online

				if (data.user.stream) {
					console.log("stream seems to be online, but it's missing updates");
					noUpdateTimeout = setTimeout(async () => {
						console.log('stream seems to ended (forcefully)');
					}, 18000);
				}

				console.log('stream seems to ended');
			});

			hls.loadSource(streamPlaybackUrl);
			hls.attachMedia(videoElement);
		});
	}
</script>

<div class="h-screen bg-gray-100">
	{#if !data.user.stream}
		<div class="flex flex-col h-full justify-center items-center">
			<h1 class="text-2xl">{'):'}</h1>
			<h2 class="text-xl">O usuário está offline/não está transmitindo!</h2>
			<h3 class="text-lg">Última vez online: {data.user.last_stream_at}</h3>
		</div>
	{:else}
		<div class="flex flex-row h-full">
			<div class="flex flex-col flex-1">
				<video class="max-h-[510px] bg-black" bind:this={videoElement} autoplay muted>
					<track kind="captions" />
				</video>

				<div class="flex pt-5 pl-6">
					<img
						src="http://placekitten.com/80/80"
						alt="Avatar"
						class="w-20 h-20 bg-slate-300 rounded-lg"
					/>
					<div class="flex flex-col pl-5">
						<h1 class="text-2xl font-bold">
							{data.user.username} <span class="text-xs font-normal">({data.user.id})</span>
						</h1>
						<h2>Titulo (que ainda não foi implementado pois estou ficando maluco)</h2>
					</div>
				</div>
			</div>

			<div class="flex flex-col w-[400px] bg-white border-l-2">
				<div class="p-3 text-center border-b-2">CHAT DA LIVE</div>

				<div class="h-full" />
				<div class="pb-3 flex justify-center items-center">
					<input
						type="text"
						class="bg-gray-100 w-96 rounded-lg p-3"
						placeholder="Digite uma mensagem"
					/>
				</div>
			</div>
		</div>
	{/if}
</div>
