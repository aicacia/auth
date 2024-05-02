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

<div class="flex flex-col justify-end md:justify-start px-4">
	<div
		class="flex flex-col flex-shrink w-full max-w-6xl mx-auto mt-4 bg-white dark:bg-gray-800 shadow p-4"
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
					class="btn primary flex flex-row me-2"
					href={`${base}/applications/${application.id}/users`}
				>
					<Users class="me-2" />
					{$LL.users.title()}
				</a>
				<a
					class="btn primary flex flex-row me-2"
					href={`${base}/applications/${application.id}/tenents`}
				>
					<KeySquare class="me-2" />
					{$LL.tenents.title()}
				</a>
				<a
					class="btn primary flex flex-row me-2"
					href={`${base}/applications/${application.id}/permissions`}
				>
					<Lock class="me-2" />
					{$LL.permissions.title()}
				</a>
			</div>
		</div>
	</div>
</div>

<Application bind:application />
