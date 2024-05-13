<svelte:options immutable />

<script lang="ts">
	import { base } from '$app/paths';
	import Users from 'lucide-svelte/icons/users';
	import Lock from 'lucide-svelte/icons/lock';
	import ArrowLeft from 'lucide-svelte/icons/arrow-left';
	import KeySquare from 'lucide-svelte/icons/key-square';
	import Application from '$lib/components/Application/Application.svelte';
	import type { PageData } from './$types';
	import LL from '$lib/i18n/i18n-svelte';

	export let data: PageData;

	$: application = data.application;
</script>

<svelte:head>
	<title>{application.description}</title>
</svelte:head>

<div class="flex flex-col justify-end px-4 md:justify-start">
	<div
		class="mx-auto mt-4 flex w-full max-w-6xl flex-shrink flex-col bg-white p-4 shadow dark:bg-gray-800"
	>
		<div class="flex flex-row justify-between">
			<div class="flex flex-row">
				<a class="btn icon primary me-3" href={`${base}/applications`}>
					<ArrowLeft />
				</a>
				<h4>{application.description}</h4>
			</div>
			<div class="flex flex-row">
				<a
					class="btn primary me-2 flex flex-row"
					href={`${base}/applications/${application.id}/users`}
				>
					<Users class="me-2" />
					{$LL.users.title()}
				</a>
				<a
					class="btn primary me-2 flex flex-row"
					href={`${base}/applications/${application.id}/tenents`}
				>
					<KeySquare class="me-2" />
					{$LL.tenents.title()}
				</a>
				<a
					class="btn primary me-2 flex flex-row"
					href={`${base}/applications/${application.id}/permissions`}
				>
					<Lock class="me-2" />
					{$LL.permissions.title()}
				</a>
			</div>
		</div>
	</div>
</div>

<div class="flex flex-grow flex-col px-4">
	<Application bind:application />
</div>
