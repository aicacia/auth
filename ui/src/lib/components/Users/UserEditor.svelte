<svelte:options immutable />

<script lang="ts">
	import type { User, UserInfo } from '$lib/openapi/auth';
	import { currentUser, updateCurrentUser } from '$lib/stores/user';
	import { userApi } from '$lib/openapi';
	import { createNotification } from '$lib/stores/notifications';
	import { get } from 'svelte/store';
	import LL from '$lib/i18n/i18n-svelte';
	import { invalidateAll } from '$app/navigation';
	import { onMount } from 'svelte';
	import Spinner from '../Spinner.svelte';
	import Username from '../Profile/Username.svelte';
	import UserInfoComponent from '../Profile/UserInfo.svelte';
	import Emails from '../Profile/Emails.svelte';
	import PhoneNumbers from '../Profile/PhoneNumbers.svelte';

	export let user: User;
	export let onUpdate: (user: User) => void;

	async function onUsernameUpdate(username: string) {
		await userApi.updateUserById(user.applicationId, user.id, { username });
		if (user.id === get(currentUser)?.id) {
			user.username = username;
			updateCurrentUser(user);
		}
		createNotification(get(LL).profile.notification.usernameChangedSuccess(), 'success');
		onUpdate(user);
		await invalidateAll();
	}

	async function onUserInfoUpdate(data: UserInfo) {
		const userInfo = await userApi.updateUserInfo(user.applicationId, user.id, data);
		createNotification(get(LL).profile.notification.userInfoChangedSuccess(), 'success');
		return userInfo;
	}

	async function onEmailUpdate(user: User) {
		onUpdate(user);
		await invalidateAll();
	}

	async function onPhoneNumberUpdate(user: User) {
		onUpdate(user);
		await invalidateAll();
	}

	let loading = true;
	let userInfo: UserInfo;
	onMount(async () => {
		try {
			userInfo = await userApi.userInfo(user.applicationId, user.id);
		} finally {
			loading = false;
		}
	});
</script>

<div class="flex flex-col justify-end md:justify-start">
	<div class="mb-2">
		<h3 class="mb-1">{$LL.profile.updateUsername()}</h3>
		<Username bind:user onUpdate={onUsernameUpdate} />
	</div>
</div>
<div class="flex flex-col justify-end md:justify-start">
	<div class="mb-2">
		<h3 class="mb-1">{$LL.profile.updateEmails()}</h3>
		<Emails bind:user onUpdate={onEmailUpdate} />
	</div>
</div>
<div class="flex flex-col justify-end md:justify-start">
	<div class="mb-2">
		<h3 class="mb-1">{$LL.profile.updatePhoneNumbers()}</h3>
		<PhoneNumbers bind:user onUpdate={onPhoneNumberUpdate} />
	</div>
</div>
<div class="flex flex-col justify-end md:justify-start">
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
