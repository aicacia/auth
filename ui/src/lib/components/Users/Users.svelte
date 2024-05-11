<svelte:options immutable />

<script lang="ts">
	import LL from '$lib/i18n/i18n-svelte';
	import { userApi } from '$lib/openapi';
	import type { Application, User } from '$lib/openapi/auth';
	import { EllipsisVertical, Pencil, Trash } from 'lucide-svelte';
	import { onMount } from 'svelte';
	import Dropdown from '$lib/components/Dropdown.svelte';
	import Modal from '../Modal.svelte';
	import { invalidateAll } from '$app/navigation';
	import { handleError } from '$lib/errors';
	import UserEditor from './UserEditor.svelte';

	export let application: Application;

	let users: User[] = [];
	onMount(async () => {
		const pagination = await userApi.users(application.id, 20, 0);
		users = pagination.items;
	});

	let editOpen = false;
	let editUser: User | undefined;
	function onOpenEdit(user: User) {
		editOpen = true;
		editUser = user;
	}
	let editing = false;
	async function onUserUpdate(user: User) {
		try {
			editing = true;
			const index = users.findIndex((u) => u.id === user.id);
			if (index !== -1) {
				const newUsers = users.slice();
				newUsers[index] = user;
				users = newUsers;
				editUser = user;
			}
		} finally {
			editing = false;
		}
	}

	let deleteOpen = false;
	let deleteUser: User | undefined;
	function onOpenDelete(user: User) {
		deleteOpen = true;
		deleteUser = user;
	}
	let deleting = false;
	async function onDelete() {
		if (deleteUser) {
			try {
				deleting = true;
				await userApi.deleteUserById(application.id, deleteUser.id);
				await invalidateAll();
				deleteOpen = false;
				deleteUser = undefined;
			} catch (error) {
				await handleError(error);
			} finally {
				deleting = false;
			}
		}
	}
</script>

<table class="w-full">
	<thead class="sticky top-0">
		<tr class="text-left border-b border-gray-300 dark:border-gray-900 shadow">
			<th>{$LL.users.id()}</th>
			<th>{$LL.users.username()}</th>
			<th>{$LL.users.email()}</th>
			<th>{$LL.users.phoneNumber()}</th>
			<th>{$LL.common.updatedAt()}</th>
			<th>{$LL.common.createdAt()}</th>
			<th></th>
		</tr>
	</thead>
	<tbody class="h-full overflow-y-auto">
		{#each users as user (user.id)}
			<tr>
				<td>{user.id}</td>
				<td>{user.username}</td>
				<td>{user.email?.email || ''}</td>
				<td>{user.phoneNumber?.phoneNumber || ''}</td>
				<td>{user.updatedAt.toLocaleString()}</td>
				<td>{user.createdAt.toLocaleString()}</td>
				<td class="flex flex-row justify-end">
					<Dropdown>
						<EllipsisVertical slot="button" />
						<button
							class="flex flex-row justify-between p-2 cursor-pointer hover:bg-gray-200 dark:hover:bg-gray-600"
							on:click={() => onOpenEdit(user)}
						>
							<Pencil /><span class="ms-4">{$LL.users.edit.button()}</span>
						</button>
						<button
							class="flex flex-row justify-between p-2 cursor-pointer hover:bg-gray-200 dark:hover:bg-gray-600"
							on:click={() => onOpenDelete(user)}
						>
							<Trash /><span class="ms-4">{$LL.users.delete.button()}</span>
						</button>
					</Dropdown>
				</td>
			</tr>
		{/each}
	</tbody>
</table>

<Modal bind:open={editOpen}>
	{#if editUser}
		<UserEditor bind:user={editUser} onUpdate={onUserUpdate} />
	{/if}
</Modal>

<Modal bind:open={deleteOpen}>
	<h4 slot="title">{$LL.users.delete.confirmTitle()}</h4>
	{#if deleteUser}
		<p>{$LL.users.delete.confirmMessage()}</p>
		<div class="flex flex-row justify-end">
			<button class="btn danger" on:click={onDelete} disabled={deleting}
				>{$LL.users.delete.confirm()}</button
			>
		</div>
	{/if}
</Modal>
