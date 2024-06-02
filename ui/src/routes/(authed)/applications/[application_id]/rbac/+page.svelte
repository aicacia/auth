<svelte:options immutable />

<script lang="ts">
	import type { PageData } from './$types';
	import ArrowLeft from 'lucide-svelte/icons/arrow-left';
	import LL from '$lib/i18n/i18n-svelte';
	import { base } from '$app/paths';

	export let data: PageData;

	$: application = data.application;
</script>

<svelte:head>
	<title>{application.description}: {$LL.rbac.title()}</title>
</svelte:head>

<div class="flex flex-col justify-end px-4 md:justify-start">
	<div
		class="mx-auto mt-4 flex w-full max-w-6xl flex-shrink flex-col bg-white p-4 shadow dark:bg-gray-800"
	>
		<div class="flex flex-row justify-between">
			<div class="flex flex-row">
				<a class="btn icon primary me-3" href={`${base}/applications/${application.id}`}>
					<ArrowLeft />
				</a>
				<h4>{application.description} / {$LL.rbac.title()}</h4>
			</div>
		</div>
	</div>
</div>

<div class="flex flex-grow flex-col px-4 pb-16">
	<div
		class="mx-auto mt-4 flex w-full max-w-6xl flex-grow flex-col bg-white p-4 shadow dark:bg-gray-800"
	></div>
</div>
