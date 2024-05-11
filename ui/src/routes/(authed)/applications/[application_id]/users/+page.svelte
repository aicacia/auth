<svelte:options immutable />

<script lang="ts">
	import type { PageData } from './$types';
	import ArrowLeft from 'lucide-svelte/icons/arrow-left';
	import LL from '$lib/i18n/i18n-svelte';
	import { base } from '$app/paths';
	import Users from '$lib/components/Users/Users.svelte';

	export let data: PageData;

	$: application = data.application;
</script>

<svelte:head>
	<title>{application.description}: {$LL.users.title()}</title>
</svelte:head>

<div class="flex flex-col justify-end md:justify-start px-4">
	<div
		class="flex flex-col flex-shrink w-full max-w-6xl mx-auto mt-4 bg-white dark:bg-gray-800 shadow p-4"
	>
		<div class="flex flex-row justify-between">
			<div class="flex flex-row">
				<a class="btn icon primary me-3" href={`${base}/applications/${application.id}`}>
					<ArrowLeft />
				</a>
				<h4>{application.description} / {$LL.users.title()}</h4>
			</div>
		</div>
	</div>
</div>

<div class="flex flex-col flex-grow px-4">
	<div
		class="flex flex-col flex-grow w-full max-w-6xl mx-auto mt-4 bg-white dark:bg-gray-800 shadow p-4"
	>
		<Users bind:application />
	</div>
</div>
