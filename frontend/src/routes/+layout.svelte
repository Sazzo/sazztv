<script>
	import '../app.css';

	import '@fontsource/overpass/200.css';
	import '@fontsource/overpass/300.css';
	import '@fontsource/overpass/400.css';
	import '@fontsource/overpass/500.css';
	import '@fontsource/overpass/600.css';
	import '@fontsource/overpass/700.css';
	import '@fontsource/overpass/800.css';
	import '@fontsource/overpass/900.css';
	import Navbar from '../components/Navbar.svelte';
</script>

<div class="flex flex-col h-screen">
	<Navbar />

	<slot />
</div>
