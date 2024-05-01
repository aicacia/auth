<svelte:options immutable />

<script lang="ts">
	import { base } from '$app/paths';
	import LL from '$lib/i18n/i18n-svelte';
</script>

<svelte:head>
	<title>{$LL.health.title()}</title>
</svelte:head>

<div class="flex flex-col flex-grow justify-end md:justify-start">
	<div
		class="flex flex-col flex-shrink md:w-80 w-full mx-auto my-10 bg-white dark:bg-gray-800 shadow p-4"
	>
		<h1 class="text-center">{$LL.health.header()}</h1>
		<p class="py-2 text-center">
			<a href={`${base}/signin`}>{$LL.auth.signIn()}</a>
		</p>
	</div>
</div>
