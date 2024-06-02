<svelte:options immutable />

<script lang="ts" context="module">
	import { writable } from 'svelte/store';

	let search = writable('');
</script>

<script lang="ts">
	import { base } from '$app/paths';
	import { fuzzyEquals } from '@aicacia/string-fuzzy_equals';
	import type { PageData } from './$types';
	import Plus from 'lucide-svelte/icons/plus';
	import Users from 'lucide-svelte/icons/users';
	import Lock from 'lucide-svelte/icons/lock';
	import Settings from 'lucide-svelte/icons/settings';
	import KeySquare from 'lucide-svelte/icons/key-square';
	import EllipsisVertical from 'lucide-svelte/icons/ellipsis-vertical';
	import Modal from '$lib/components/Modal.svelte';
	import ApplicationEditor from '$lib/components/Application/ApplicationEditor.svelte';
	import type { Application } from '$lib/openapi/auth';
	import LL from '$lib/i18n/i18n-svelte';
	import Dropdown from '$lib/components/Dropdown.svelte';

	export let data: PageData;

	let hasMore = data.pagination.hasMore;
	$: applications = data.pagination.items;

	let filtered = new Set<number>();
	$: filtered = $search.length
		? applications.reduce((acc, a) => {
				if (!fuzzyEquals($search, a.description)) {
					acc.add(a.id);
				}
				return acc;
			}, new Set<number>())
		: new Set<number>();

	let addOpen = false;
	function onAddOpen() {
		addOpen = true;
	}

	function onCreateAppliction(application: Application) {
		applications = [application, ...applications];
		addOpen = false;
	}
</script>

<svelte:head>
	<title>{$LL.applications.title()}</title>
</svelte:head>

<div class="container mx-auto my-4">
	<div class="bg-white p-4 shadow dark:bg-gray-800">
		<div class="flex flex-grow flex-row justify-between">
			<input type="text" placeholder={$LL.applications.filter()} bind:value={$search} />
			<button class="btn primary icon" on:click={onAddOpen}><Plus /></button>
		</div>
		<hr class="mb-3 mt-1" />
		<div class="grid grid-cols-1 gap-2 sm:grid-cols-3 lg:grid-cols-6">
			{#each applications as application (application.id)}
				{@const hidden = filtered.has(application.id)}
				<div
					class="flex flex-col border border-gray-300 bg-white p-2 shadow hover:shadow-lg dark:border-gray-900 dark:bg-gray-700"
					class:hidden
				>
					<div class="flex flex-row justify-between">
						<h4 class="overflow-hidden text-ellipsis text-nowrap">
							{application.description}
						</h4>
						<Dropdown>
							<EllipsisVertical slot="button" />
							<a
								href={`${base}/applications/${application.id}`}
								class="default flex cursor-pointer flex-row justify-between p-2 hover:bg-gray-200 dark:hover:bg-gray-600"
							>
								<Settings /><span class="ms-4">{$LL.application.title()}</span>
							</a>
							<a
								href={`${base}/applications/${application.id}/users`}
								class="default flex cursor-pointer flex-row justify-between p-2 hover:bg-gray-200 dark:hover:bg-gray-600"
							>
								<Users /><span class="ms-4">{$LL.users.title()}</span>
							</a>
							<a
								href={`${base}/applications/${application.id}/tenents`}
								class="default flex cursor-pointer flex-row justify-between p-2 hover:bg-gray-200 dark:hover:bg-gray-600"
							>
								<KeySquare /><span class="ms-4">{$LL.tenents.title()}</span>
							</a>
							<a
								href={`${base}/applications/${application.id}/rbac`}
								class="default flex cursor-pointer flex-row justify-between p-2 hover:bg-gray-200 dark:hover:bg-gray-600"
							>
								<Lock /><span class="ms-4">{$LL.rbac.title()}</span>
							</a>
						</Dropdown>
					</div>
				</div>
			{/each}
		</div>
	</div>
</div>

<Modal bind:open={addOpen}>
	<h4 slot="title">Create Application</h4>
	<ApplicationEditor onDone={onCreateAppliction} />
</Modal>
