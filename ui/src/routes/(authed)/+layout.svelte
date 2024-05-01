<svelte:options immutable />

<script lang="ts">
	import Header from '$lib/components/Header.svelte';
</script>

<div class="flex flex-col w-full h-full">
	<Header />
	<div class="flex flex-col flex-grow overflow-x-hidden overflow-y-auto">
		<slot />
	</div>
</div>
