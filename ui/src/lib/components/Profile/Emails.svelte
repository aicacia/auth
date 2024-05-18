<svelte:options immutable />

<script lang="ts">
	import { currentUserApi } from '$lib/openapi';
	import type { User, Email } from '$lib/openapi/auth';
	import EmailComponent from './Email.svelte';
	import NewEmail from './NewEmail.svelte';

	export let user: User;
	export let onUpdate: (user: User) => Promise<void>;

	let newEmail: Email | undefined;
	async function onCreateEmail(email: string) {
		newEmail = await currentUserApi.createEmail({ email });
		user = { ...user, emails: [...user.emails, newEmail] };
		onUpdate(user);
	}
</script>

<div class="flex flex-col">
	{#if user.email}
		<EmailComponent bind:user bind:email={user.email} primary {onUpdate} />
		{#if user.emails.length}
			<hr class="my-1" />
		{/if}
	{/if}
	{#each user.emails as email, index (email.id)}
		<EmailComponent
			bind:user
			bind:email
			sentEmailConfirmation={newEmail?.id === email.id}
			{onUpdate}
		/>
		{#if index < user.emails.length - 1}
			<hr class="my-1" />
		{/if}
	{/each}
</div>

<NewEmail onCreate={onCreateEmail} />
