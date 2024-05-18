<svelte:options immutable />

<script lang="ts">
	import Dropdown from '$lib/components/Dropdown.svelte';
	import type { PhoneNumber, User } from '$lib/openapi/auth';
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
	import { formatPhoneNumber } from '$lib/util';

	export let user: User;
	export let phoneNumber: PhoneNumber;
	export let primary = false;
	export let sentPhoneNumberConfirmation = false;
	export let onUpdate: (user: User) => Promise<void>;

	$: formatedPhoneNumber = formatPhoneNumber(phoneNumber.phoneNumber);
	let open = false;

	async function onSetPrimaryInternal() {
		try {
			await currentUserApi.setPrimaryPhoneNumber(phoneNumber.id);
			const newPhoneNumbers = user.phoneNumbers.slice();
			const index = newPhoneNumbers.findIndex((e) => e.id === phoneNumber.id);
			if (index !== -1) {
				newPhoneNumbers.splice(index, 1);
			}
			if (user.phoneNumber) {
				newPhoneNumbers.push(user.phoneNumber);
			}
			user = { ...user, phoneNumber, phoneNumbers: newPhoneNumbers };
			await onUpdate(user);
			open = false;
		} catch (error) {
			await handleError(error);
		}
	}
	async function onDeletePhoneNumber() {
		try {
			await currentUserApi.deletePhoneNumber(phoneNumber.id);
			const newPhoneNumbers = user.phoneNumbers.slice();
			const index = newPhoneNumbers.findIndex((e) => e.id === phoneNumber.id);
			if (index !== -1) {
				newPhoneNumbers.splice(index, 1);
			}
			user = { ...user, phoneNumbers: newPhoneNumbers };
			await onUpdate(user);
			deletePhoneNumberOpen = false;
		} catch (error) {
			await handleError(error);
		}
	}

	let phoneNumberConfirmation: string;
	async function onSendConfirmation() {
		try {
			await currentUserApi.sendConfirmationToPhoneNumber(phoneNumber.id);
			open = false;
			sentPhoneNumberConfirmation = true;
			createNotification($LL.profile.notification.sentPhoneNumberConfirmation(), 'info');
		} catch (error) {
			await handleError(error);
		}
	}
	async function onConfirmPhoneNumber() {
		try {
			phoneNumber = await currentUserApi.confirmPhoneNumber(phoneNumber.id, {
				token: phoneNumberConfirmation
			});
			const newPhoneNumbers = user.phoneNumbers.slice();
			const index = newPhoneNumbers.findIndex((e) => e.id === phoneNumber.id);
			if (index !== -1) {
				newPhoneNumbers[index] = phoneNumber;
			}
			user = { ...user, phoneNumbers: newPhoneNumbers };
			await onUpdate(user);
			sentPhoneNumberConfirmation = false;
			createNotification($LL.profile.notification.phoneNumberConfirmed(), 'success');
		} catch (error) {
			await handleError(error);
		}
	}

	let deletePhoneNumberOpen = false;
	function onDeletePhoneNumberOpen() {
		deletePhoneNumberOpen = true;
		open = false;
	}
</script>

<div class="flex flex-grow flex-row items-center justify-between">
	<div class="relative flex flex-grow">
		<input class="w-full" type="phoneNumber" value={formatedPhoneNumber} readonly />
		{#if phoneNumber.confirmed}
			<span class="absolute right-0 top-0 me-1 mt-1 cursor-help text-green-600" title="Confirmed"
				><CircleCheck size={22} /></span
			>
		{/if}
	</div>
	{#if !primary || !phoneNumber.confirmed}
		<div class="flex flex-shrink">
			<Dropdown bind:open>
				<EllipsisVertical slot="button" />
				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
				{#if !phoneNumber.confirmed}
					<li
						class="flex cursor-pointer flex-row justify-between p-2 hover:bg-gray-200 dark:hover:bg-gray-600"
						on:click={onSendConfirmation}
					>
						<Mail /><span class="ms-4">{$LL.profile.phoneNumbers.sendConfirmation()}</span>
					</li>
				{:else}
					<li
						class="flex cursor-pointer flex-row justify-between p-2 hover:bg-gray-200 dark:hover:bg-gray-600"
						on:click={onSetPrimaryInternal}
					>
						<Send /><span class="ms-4">{$LL.profile.phoneNumbers.setAsPrimary()}</span>
					</li>
				{/if}
				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
				<li
					class="flex cursor-pointer flex-row justify-between p-2 hover:bg-gray-200 dark:hover:bg-gray-600"
					on:click={onDeletePhoneNumberOpen}
				>
					<Trash /><span class="ms-4">{$LL.profile.phoneNumbers.delete()}</span>
				</li>
			</Dropdown>
		</div>
	{/if}
</div>

<Modal bind:open={sentPhoneNumberConfirmation} backgroundClose={false}>
	<h4 slot="title">{$LL.profile.phoneNumbers.checkYourPhone()}</h4>
	<form on:submit|preventDefault={onConfirmPhoneNumber}>
		<div class="flex flex-col">
			<input
				type="text"
				class="flex flex-grow"
				placeholder="Confirmation Token"
				bind:value={phoneNumberConfirmation}
			/>
		</div>
		<div class="mt-2 flex flex-row justify-end">
			<button class="btn primary" type="submit">{$LL.profile.phoneNumbers.confirmCode()}</button>
		</div>
	</form>
</Modal>

<Modal bind:open={deletePhoneNumberOpen}>
	<h4 slot="title">{$LL.profile.phoneNumbers.deletePhoneNumber(formatedPhoneNumber)}</h4>
	<form on:submit|preventDefault={onDeletePhoneNumber}>
		<div class="mt-2 flex flex-row justify-end">
			<button class="btn danger" type="submit">{$LL.profile.phoneNumbers.delete()}</button>
		</div>
	</form>
</Modal>
