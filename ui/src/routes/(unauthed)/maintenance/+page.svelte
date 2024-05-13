<svelte:options immutable />

<script lang="ts">
	import { base } from '$app/paths';
	import LL from '$lib/i18n/i18n-svelte';
</script>

<svelte:head>
	<title>{$LL.maintenance.title()}</title>
</svelte:head>

<div class="mt-10 flex flex-grow flex-col justify-end md:justify-start">
	<div class="mx-4 bg-white px-4 pb-8 pt-4 text-justify sm:container sm:mx-auto">
		<h4 class="mb-11 uppercase">{$LL.maintenance.header()}</h4>
		<p>
			{@html $LL.maintenance.body({ link: `${base}/` })}
		</p>
	</div>
</div>
