<svelte:options immutable />

<script lang="ts">
	import { currentUserApi } from '$lib/openapi';
	import type { User, PhoneNumber } from '$lib/openapi/auth';
	import PhoneNumberComponent from './PhoneNumber.svelte';
	import NewPhoneNumber from './NewPhoneNumber.svelte';

	export let user: User;
	export let onUpdate: (user: User) => Promise<void>;

	let newPhoneNumber: PhoneNumber | undefined;
	async function onCreatePhoneNumber(phoneNumber: string) {
		newPhoneNumber = await currentUserApi.createPhoneNumber({
			phoneNumber
		});
		user = { ...user, phoneNumbers: [...user.phoneNumbers, newPhoneNumber] };
		onUpdate(user);
	}
</script>

<div class="flex flex-col">
	{#if user.phoneNumber}
		<PhoneNumberComponent bind:user bind:phoneNumber={user.phoneNumber} primary {onUpdate} />
		{#if user.phoneNumbers.length}
			<hr class="my-1" />
		{/if}
	{/if}
	{#each user.phoneNumbers as phoneNumber, index (phoneNumber.id)}
		<PhoneNumberComponent
			bind:user
			bind:phoneNumber
			sentPhoneNumberConfirmation={newPhoneNumber?.id === phoneNumber.id}
			{onUpdate}
		/>
		{#if index < user.phoneNumbers.length - 1}
			<hr class="my-1" />
		{/if}
	{/each}
</div>

<NewPhoneNumber onCreate={onCreatePhoneNumber} />
