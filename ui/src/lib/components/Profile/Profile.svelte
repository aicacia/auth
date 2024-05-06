<svelte:options immutable />

<script lang="ts">
	import type { User } from '$lib/openapi/auth';
	import ResetPassword from './ResetPassword.svelte';
	import Username from './Username.svelte';
	import { currentUser, updateCurrentUser } from '$lib/stores/user';
	import { currentUserApi } from '$lib/openapi';
	import { createNotification } from '$lib/stores/notifications';
	import { get } from 'svelte/store';
	import LL from '$lib/i18n/i18n-svelte';
	import { invalidateAll } from '$app/navigation';
	import Emails from './Emails.svelte';

	export let user: User;

	async function onUsernameUpdate(username: string) {
		await currentUserApi.updateUsername({ username });
		user.username = username;
		updateCurrentUser(user);
		createNotification(get(LL).profile.notification.usernameChangedSuccess(), 'success');
		await invalidateAll();
	}

	async function onEmailUpdate(_user: User) {
		await invalidateAll();
	}
</script>

<div class="flex flex-col justify-end md:justify-start px-4 mt-8">
	<div
		class="flex flex-col flex-shrink w-full max-w-lg mx-auto mb-4 bg-white dark:bg-gray-800 shadow p-4"
	>
		<div class="mb-2">
			<h3 class="mb-1">{$LL.profile.updateUsername()}</h3>
			<Username bind:user onUpdate={onUsernameUpdate} />
		</div>
	</div>
</div>
<div class="flex flex-col justify-end md:justify-start px-4 mt-8">
	<div
		class="flex flex-col flex-shrink w-full max-w-lg mx-auto mb-4 bg-white dark:bg-gray-800 shadow p-4"
	>
		<div class="mb-2">
			<h3 class="mb-1">{$LL.profile.updateEmails()}</h3>
			<Emails bind:user onUpdate={onEmailUpdate} />
		</div>
	</div>
</div>
{#if user.id === $currentUser?.id}
	<div class="flex flex-col justify-end md:justify-start px-4">
		<div
			class="flex flex-col flex-shrink w-full max-w-lg mx-auto mb-4 bg-white dark:bg-gray-800 shadow p-4"
		>
			<div class="mb-2">
				<h3 class="mb-1">{$LL.auth.resetPassword()}</h3>
				<ResetPassword />
			</div>
		</div>
	</div>
{/if}
