<svelte:options immutable />

<script lang="ts">
	import Header from '$lib/components/Header.svelte';
	import { page } from '$app/stores';
</script>

<div class="flex flex-col w-full h-full">
	<Header />
	<div class="flex flex-grow flex-col w-full mx-4 sm:container sm:mx-auto mt-4 mb-16">
		<h3>{$page.status}</h3>
		{#if $page.error?.message}
			<p>{$page.error.message}</p>
		{/if}
	</div>
</div>
