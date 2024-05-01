<svelte:options immutable />

<script lang="ts">
	import { base } from '$app/paths';
	import LL from '$lib/i18n/i18n-svelte';
</script>

<svelte:head>
	<title>{$LL.maintenance.title()}</title>
</svelte:head>

<div class="flex flex-col flex-grow justify-end md:justify-start mt-10">
	<div class="sm:container mx-4 sm:mx-auto bg-white pt-4 px-4 pb-8 text-justify">
		<h4 class="uppercase mb-11">{$LL.maintenance.header()}</h4>
		<p>
			{@html $LL.maintenance.body({ link: `${base}/` })}
		</p>
	</div>
</div>
