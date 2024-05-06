<svelte:options immutable />

<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import { currentUserApi } from '$lib/openapi';
	import type { User } from '$lib/openapi/auth';
	import Username from '../Profile/Username.svelte';

	export let user: User;
	export let onUpdate: (user: User) => void;

	async function onUsernameUpdate(username: string) {
		const newUser = { ...user };
		newUser.username = username;
		await currentUserApi.updateUsername({ username });
		await invalidateAll();
		onUpdate(newUser);
	}
</script>

<Username bind:user onUpdate={onUsernameUpdate} />
