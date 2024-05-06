<svelte:options immutable />

<script lang="ts">
	import { userApi } from '$lib/openapi';
	import type { User, Email } from '$lib/openapi/auth';
	import EmailComponent from './Email.svelte';
	import NewEmail from './NewEmail.svelte';

	export let user: User;
	export let onUpdate: (user: User) => Promise<void>;

	let newEmail: Email | undefined;
	async function onCreateEmail(email: string) {
		newEmail = await userApi.createEmail(user.applicationId, user.id as number, { email });
		user = { ...user, emails: [...user.emails, newEmail] };
		onUpdate(user);
	}
</script>

<div class="flex flex-col">
	{#if user.email}
		<EmailComponent bind:user bind:email={user.email} primary />
		{#if user.emails.length}
			<hr class="my-1" />
		{/if}
	{/if}
	{#each user.emails as email, index (email.id)}
		<EmailComponent bind:user bind:email sentEmailConfirmation={newEmail?.id === email.id} />
		{#if index < user.emails.length - 1}
			<hr class="my-1" />
		{/if}
	{/each}
</div>

<NewEmail onCreate={onCreateEmail} />
