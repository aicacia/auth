<svelte:options immutable />

<script lang="ts">
	import type { PageData } from './$types';
	import ArrowLeft from 'lucide-svelte/icons/arrow-left';
	import Plus from 'lucide-svelte/icons/plus';
	import LL from '$lib/i18n/i18n-svelte';
	import { base } from '$app/paths';
	import Tenents from '$lib/components/Tenents/Tenents.svelte';
	import type { Tenent } from '$lib/openapi/auth';
	import Modal from '$lib/components/Modal.svelte';
	import { tenentApi } from '$lib/openapi';
	import TenentEditor, { type TenentEditorForm } from '$lib/components/Tenents/TenentEditor.svelte';

	export let data: PageData;

	$: application = data.application;
	let tenents: Tenent[] = [];

	let newTenentOpen = false;
	function onNewTenentOpen() {
		newTenentOpen = true;
	}
	async function onNewTenent(form: TenentEditorForm) {
		const tenent = await tenentApi.createTenent(application.id, form);
		tenents = [tenent, ...tenents];
		newTenentOpen = false;
	}
</script>

<svelte:head>
	<title>{application.description}: {$LL.tenents.title()}</title>
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
				<h4>{application.description} / {$LL.tenents.title()}</h4>
			</div>
			<div>
				<button class="btn icon primary" on:click={onNewTenentOpen}>
					<Plus />
				</button>
			</div>
		</div>
	</div>
</div>

<div class="flex flex-grow flex-col px-4 pb-16">
	<div
		class="mx-auto mt-4 flex w-full max-w-6xl flex-grow flex-col bg-white p-4 shadow dark:bg-gray-800"
	>
		<Tenents bind:application bind:tenents />
	</div>
</div>

<Modal bind:open={newTenentOpen}>
	<h4 slot="title">{$LL.tenents.newTenent.title()}</h4>
	<TenentEditor onUpdate={onNewTenent} />
</Modal>
