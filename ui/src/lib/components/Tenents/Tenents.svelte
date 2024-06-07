<svelte:options immutable />

<script lang="ts">
	import LL from '$lib/i18n/i18n-svelte';
	import { tenentApi } from '$lib/openapi';
	import type { Application, Tenent } from '$lib/openapi/auth';
	import { EllipsisVertical, Pencil, Trash } from 'lucide-svelte';
	import { onMount } from 'svelte';
	import Dropdown from '$lib/components/Dropdown.svelte';
	import Modal from '../Modal.svelte';
	import { goto, invalidateAll } from '$app/navigation';
	import { handleError } from '$lib/errors';
	import TenentEditor, { type TenentEditorForm } from './TenentEditor.svelte';
	import TenentKey, { type TenentKeyForm } from './TenentKey.svelte';
	import { PUBLIC_TENENT_ID } from '$env/static/public';
	import { signOut } from '$lib/stores/user';

	export let application: Application;
	export let tenents: Tenent[] = [];

	onMount(async () => {
		const pagination = await tenentApi.tenents(application.id, 20, 0);
		tenents = pagination.items;
	});

	let editOpen = false;
	let editTenent: Tenent | undefined;
	function onOpenEdit(tenent: Tenent) {
		editOpen = true;
		editTenent = tenent;
	}
	async function onTenentUpdate(form: TenentEditorForm) {
		if (editTenent) {
			const tenent = await tenentApi.updateTenent(editTenent.applicationId, editTenent.id, form);
			tenentUpdate(tenent);
			await invalidateAll();
		}
	}
	function createGetTenentPrivateKey(tenent: Tenent) {
		return async () => {
			try {
				return await tenentApi.tenentPrivateKey(tenent.applicationId, tenent.id);
			} catch (error) {
				await handleError(error);
			}
		};
	}
	async function onTenentKeyUpdate(data: TenentKeyForm) {
		if (editTenent) {
			try {
				const tenent = await tenentApi.updateTenent(editTenent.applicationId, editTenent.id, data);
				if (tenent.clientId === PUBLIC_TENENT_ID) {
					signOut();
					await goto('/signin');
				} else {
					tenentUpdate(tenent);
					await invalidateAll();
				}
			} catch (error) {
				await handleError(error);
			}
		}
	}

	let deleteOpen = false;
	let deleteTenent: Tenent | undefined;
	function onOpenDelete(tenent: Tenent) {
		deleteOpen = true;
		deleteTenent = tenent;
	}
	let deleting = false;
	async function onDelete() {
		if (deleteTenent) {
			try {
				deleting = true;
				await tenentApi.deleteTenent(application.id, deleteTenent.id);
				tenentDelete(deleteTenent);
				await invalidateAll();
				deleteOpen = false;
				deleteTenent = undefined;
			} catch (error) {
				await handleError(error);
			} finally {
				deleting = false;
			}
		}
	}

	function tenentUpdate(tenent: Tenent) {
		const index = tenents.findIndex((u) => u.id === tenent.id);
		if (index !== -1) {
			const newTenents = tenents.slice();
			newTenents[index] = tenent;
			tenents = newTenents;
		}
	}
	function tenentDelete(tenent: Tenent) {
		const index = tenents.findIndex((u) => u.id === tenent.id);
		if (index !== -1) {
			const newTenents = tenents.slice();
			newTenents.splice(index, 1);
			tenents = newTenents;
		}
	}
</script>

<table class="w-full">
	<thead class="sticky top-0">
		<tr class="border-b border-gray-300 text-left shadow dark:border-gray-900">
			<th>{$LL.tenents.id()}</th>
			<th>{$LL.tenents.description()}</th>
			<th>{$LL.tenents.uri()}</th>
			<th>{$LL.common.updatedAt()}</th>
			<th>{$LL.common.createdAt()}</th>
			<th></th>
		</tr>
	</thead>
	<tbody class="h-full overflow-y-auto">
		{#each tenents as tenent (tenent.id)}
			<tr>
				<td>{tenent.id}</td>
				<td>{tenent.description}</td>
				<td>{tenent.uri}</td>
				<td>{tenent.updatedAt.toLocaleString()}</td>
				<td>{tenent.createdAt.toLocaleString()}</td>
				<td class="flex flex-row justify-end">
					<Dropdown>
						<EllipsisVertical slot="button" />
						<button
							class="flex cursor-pointer flex-row justify-between p-2 hover:bg-gray-200 dark:hover:bg-gray-600"
							on:click={() => onOpenEdit(tenent)}
						>
							<Pencil /><span class="ms-4">{$LL.tenents.edit.button()}</span>
						</button>
						<button
							class="flex cursor-pointer flex-row justify-between p-2 hover:bg-gray-200 dark:hover:bg-gray-600"
							on:click={() => onOpenDelete(tenent)}
						>
							<Trash /><span class="ms-4">{$LL.tenents.delete.button()}</span>
						</button>
					</Dropdown>
				</td>
			</tr>
		{/each}
	</tbody>
</table>

<Modal bind:open={editOpen}>
	<h4 slot="title">{editTenent?.description}</h4>
	{#if editTenent}
		<TenentEditor
			description={editTenent.description}
			uri={editTenent.uri}
			authorizationWebsite={editTenent.authorizationWebsite}
			registrationWebsite={editTenent.registrationWebsite}
			expiresInSeconds={editTenent.expiresInSeconds}
			refreshExpiresInSeconds={editTenent.refreshExpiresInSeconds}
			passwordResetExpiresInSeconds={editTenent.passwordResetExpiresInSeconds}
			onUpdate={onTenentUpdate}
		/>
		<h4>{$LL.tenents.key()}</h4>
		<TenentKey
			algorithm={editTenent.algorithm}
			publicKey={editTenent.publicKey}
			getPrivateKey={createGetTenentPrivateKey(editTenent)}
			onUpdate={onTenentKeyUpdate}
		/>
	{/if}
</Modal>

<Modal bind:open={deleteOpen}>
	<h4 slot="title">{$LL.tenents.delete.confirmTitle()}</h4>
	{#if deleteTenent}
		<p>{$LL.tenents.delete.confirmMessage(deleteTenent.description)}</p>
		<div class="flex flex-row justify-end">
			<button class="btn danger" on:click={onDelete} disabled={deleting}
				>{$LL.tenents.delete.confirm()}</button
			>
		</div>
	{/if}
</Modal>
