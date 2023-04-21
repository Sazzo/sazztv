<script lang="ts">
	export let name: string;
	export let type: 'text' | 'password';
	export let id: string;
	export let required: boolean;
</script>

<input class="bg-gray-200 rounded-[5px] w-96 text-sm p-1" {name} {type} {id} {required} />
