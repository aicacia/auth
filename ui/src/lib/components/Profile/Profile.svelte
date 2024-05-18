<svelte:options immutable />

<script lang="ts">
	import type { User, UserInfo } from '$lib/openapi/auth';
	import ResetPassword from './ResetPassword.svelte';
	import Username from './Username.svelte';
	import { updateCurrentUser } from '$lib/stores/user';
	import { currentUserApi } from '$lib/openapi';
	import { createNotification } from '$lib/stores/notifications';
	import { get } from 'svelte/store';
	import LL from '$lib/i18n/i18n-svelte';
	import { invalidateAll } from '$app/navigation';
	import Emails from './Emails.svelte';
	import UserInfoComponent from './UserInfo.svelte';
	import { onMount } from 'svelte';
	import Spinner from '../Spinner.svelte';
	import PhoneNumbers from './PhoneNumbers.svelte';

	export let user: User;

	async function onUsernameUpdate(username: string) {
		await currentUserApi.updateUsername({ username });
		user.username = username;
		updateCurrentUser(user);
		createNotification(get(LL).profile.notification.usernameChangedSuccess(), 'success');
		await invalidateAll();
	}

	async function onUserInfoUpdate(data: UserInfo) {
		const userInfo = await currentUserApi.updateCurrentUserInfo(data);
		createNotification(get(LL).profile.notification.userInfoChangedSuccess(), 'success');
		return userInfo;
	}

	async function onEmailUpdate(user: User) {
		updateCurrentUser(user);
		await invalidateAll();
	}

	async function onPhoneNumberUpdate(user: User) {
		updateCurrentUser(user);
		await invalidateAll();
	}

	let loading = true;
	let userInfo: UserInfo;
	onMount(async () => {
		try {
			userInfo = await currentUserApi.currentUserInfo();
		} finally {
			loading = false;
		}
	});
</script>

<div class="mt-8 flex flex-col justify-end px-4 md:justify-start">
	<div
		class="mx-auto mb-4 flex w-full max-w-lg flex-shrink flex-col bg-white p-4 shadow dark:bg-gray-800"
	>
		<div class="mb-2">
			<h3 class="mb-1">{$LL.profile.updateUsername()}</h3>
			<Username bind:user onUpdate={onUsernameUpdate} />
		</div>
	</div>
</div>
<div class="flex flex-col justify-end px-4 md:justify-start">
	<div
		class="mx-auto mb-4 flex w-full max-w-lg flex-shrink flex-col bg-white p-4 shadow dark:bg-gray-800"
	>
		<div class="mb-2">
			<h3 class="mb-1">{$LL.profile.updateEmails()}</h3>
			<Emails bind:user onUpdate={onEmailUpdate} />
		</div>
	</div>
</div>
<div class="flex flex-col justify-end px-4 md:justify-start">
	<div
		class="mx-auto mb-4 flex w-full max-w-lg flex-shrink flex-col bg-white p-4 shadow dark:bg-gray-800"
	>
		<div class="mb-2">
			<h3 class="mb-1">{$LL.profile.updatePhoneNumbers()}</h3>
			<PhoneNumbers bind:user onUpdate={onPhoneNumberUpdate} />
		</div>
	</div>
</div>
<div class="flex flex-col justify-end px-4 md:justify-start">
	<div
		class="mx-auto mb-4 flex w-full max-w-lg flex-shrink flex-col bg-white p-4 shadow dark:bg-gray-800"
	>
		<div class="mb-2">
			<h3 class="mb-1">{$LL.profile.updateUserInfo()}</h3>
			{#if loading}
				<div class="flex flex-row justify-center">
					<div class="inline-block h-6 w-6"><Spinner /></div>
				</div>
			{:else if userInfo}
				<UserInfoComponent bind:userInfo onUpdate={onUserInfoUpdate} />
			{/if}
		</div>
	</div>
</div>
<div class="flex flex-col justify-end px-4 md:justify-start">
	<div
		class="mx-auto mb-4 flex w-full max-w-lg flex-shrink flex-col bg-white p-4 shadow dark:bg-gray-800"
	>
		<div class="mb-2">
			<h3 class="mb-1">{$LL.auth.resetPassword()}</h3>
			<ResetPassword />
		</div>
	</div>
</div>
