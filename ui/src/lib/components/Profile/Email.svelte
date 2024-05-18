<svelte:options immutable />

<script lang="ts">
	import Dropdown from '$lib/components/Dropdown.svelte';
	import type { Email, User } from '$lib/openapi/auth';
	import EllipsisVertical from 'lucide-svelte/icons/ellipsis-vertical';
	import CircleCheck from 'lucide-svelte/icons/circle-check';
	import Send from 'lucide-svelte/icons/send';
	import Mail from 'lucide-svelte/icons/mail';
	import Trash from 'lucide-svelte/icons/trash';
	import { handleError } from '$lib/errors';
	import { createNotification } from '$lib/stores/notifications';
	import Modal from '$lib/components/Modal.svelte';
	import { currentUserApi } from '$lib/openapi';
	import LL from '$lib/i18n/i18n-svelte';

	export let user: User;
	export let email: Email;
	export let primary = false;
	export let sentEmailConfirmation = false;
	export let onUpdate: (user: User) => Promise<void>;

	let open = false;

	async function onSetPrimaryInternal() {
		try {
			await currentUserApi.setPrimaryEmail(email.id);
			const newEmails = user.emails.slice();
			const index = newEmails.findIndex((e) => e.id === email.id);
			if (index !== -1) {
				newEmails.splice(index, 1);
			}
			if (user.email) {
				newEmails.push(user.email);
			}
			user = { ...user, email, emails: newEmails };
			await onUpdate(user);
			open = false;
		} catch (error) {
			await handleError(error);
		}
	}
	async function onDeleteEmail() {
		try {
			await currentUserApi.deleteEmail(email.id);
			const newEmails = user.emails.slice();
			const index = newEmails.findIndex((e) => e.id === email.id);
			if (index !== -1) {
				newEmails.splice(index, 1);
			}
			user = { ...user, emails: newEmails };
			await onUpdate(user);
			deleteEmailOpen = false;
		} catch (error) {
			await handleError(error);
		}
	}

	let emailConfirmation: string;
	async function onSendConfirmation() {
		try {
			await currentUserApi.sendConfirmationToEmail(email.id);
			open = false;
			sentEmailConfirmation = true;
			createNotification($LL.profile.notification.sentEmailConfirmation(), 'info');
		} catch (error) {
			await handleError(error);
		}
	}
	async function onConfirmEmail() {
		try {
			email = await currentUserApi.confirmEmail(email.id, {
				token: emailConfirmation
			});
			const newEmails = user.emails.slice();
			const index = newEmails.findIndex((e) => e.id === email.id);
			if (index !== -1) {
				newEmails[index] = email;
			}
			user = { ...user, emails: newEmails };
			sentEmailConfirmation = false;
			await onUpdate(user);
			createNotification($LL.profile.notification.emailConfirmed(), 'success');
		} catch (error) {
			await handleError(error);
		}
	}

	let deleteEmailOpen = false;
	function onDeleteEmailOpen() {
		deleteEmailOpen = true;
		open = false;
	}
</script>

<div class="flex flex-grow flex-row items-center justify-between">
	<div class="relative flex flex-grow">
		<input class="w-full" type="email" value={email.email} readonly />
		{#if email.confirmed}
			<span class="absolute right-0 top-0 me-1 mt-1 cursor-help text-green-600" title="Confirmed"
				><CircleCheck size={22} /></span
			>
		{/if}
	</div>
	{#if !primary || !email.confirmed}
		<div class="flex flex-shrink">
			<Dropdown bind:open>
				<EllipsisVertical slot="button" />
				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
				{#if !email.confirmed}
					<li
						class="flex cursor-pointer flex-row justify-between p-2 hover:bg-gray-200 dark:hover:bg-gray-600"
						on:click={onSendConfirmation}
					>
						<Mail /><span class="ms-4">{$LL.profile.emails.sendConfirmation()}</span>
					</li>
				{:else}
					<li
						class="flex cursor-pointer flex-row justify-between p-2 hover:bg-gray-200 dark:hover:bg-gray-600"
						on:click={onSetPrimaryInternal}
					>
						<Send /><span class="ms-4">{$LL.profile.emails.setAsPrimary()}</span>
					</li>
				{/if}
				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
				<li
					class="flex cursor-pointer flex-row justify-between p-2 hover:bg-gray-200 dark:hover:bg-gray-600"
					on:click={onDeleteEmailOpen}
				>
					<Trash /><span class="ms-4">{$LL.profile.emails.delete()}</span>
				</li>
			</Dropdown>
		</div>
	{/if}
</div>

<Modal bind:open={sentEmailConfirmation} backgroundClose={false}>
	<h4 slot="title">{$LL.profile.emails.checkYourEmail()}</h4>
	<form on:submit|preventDefault={onConfirmEmail}>
		<div class="flex flex-col">
			<input
				type="text"
				class="flex flex-grow"
				placeholder="Confirmation Token"
				bind:value={emailConfirmation}
			/>
		</div>
		<div class="mt-2 flex flex-row justify-end">
			<button class="btn primary" type="submit">{$LL.profile.emails.confirmCode()}</button>
		</div>
	</form>
</Modal>

<Modal bind:open={deleteEmailOpen}>
	<h4 slot="title">{$LL.profile.emails.deleteEmail(email.email)}</h4>
	<form on:submit|preventDefault={onDeleteEmail}>
		<div class="mt-2 flex flex-row justify-end">
			<button class="btn danger" type="submit">{$LL.profile.emails.delete()}</button>
		</div>
	</form>
</Modal>
