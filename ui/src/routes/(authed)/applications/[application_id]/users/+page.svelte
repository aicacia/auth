<svelte:options immutable />

<script lang="ts">
	import type { PageData } from './$types';
	import ArrowLeft from 'lucide-svelte/icons/arrow-left';
	import LL from '$lib/i18n/i18n-svelte';
	import { base } from '$app/paths';
	import Users from '$lib/components/Users/Users.svelte';
	import Plus from 'lucide-svelte/icons/plus';
	import type { User } from '$lib/openapi/auth';
	import Modal from '$lib/components/Modal.svelte';
	import NewUser from '$lib/components/Users/NewUser.svelte';
	import { userApi } from '$lib/openapi';

	export let data: PageData;

	$: application = data.application;
	let users: User[] = [];

	let newUserOpen = false;
	function onNewUserOpen() {
		newUserOpen = true;
	}
	async function onNewUser(username: string) {
		const user = await userApi.createUser(application.id, {
			username
		});
		users = [user, ...users];
		newUserOpen = false;
	}
</script>

<svelte:head>
	<title>{application.description}: {$LL.users.title()}</title>
</svelte:head>

<div class="flex flex-col justify-end px-4 md:justify-start">
	<div
		class="mx-auto mt-4 flex w-full max-w-6xl flex-shrink flex-col bg-white p-4 shadow dark:bg-gray-800"
	>
		<div class="flex flex-row justify-between">
			<div class="flex flex-row">
				<a class="btn icon primary me-3" href={`${base}/applications/${application.id}`}>
					<ArrowLeft />
				</a>
				<h4>{application.description} / {$LL.users.title()}</h4>
			</div>
			<div>
				<button class="btn icon primary" on:click={onNewUserOpen}>
					<Plus />
				</button>
			</div>
		</div>
	</div>
</div>

<div class="flex flex-grow flex-col px-4">
	<div
		class="mx-auto mt-4 flex w-full max-w-6xl flex-grow flex-col bg-white p-4 shadow dark:bg-gray-800"
	>
		<Users bind:application bind:users />
	</div>
</div>

<Modal bind:open={newUserOpen}>
	<h4 slot="title">{$LL.users.newUser.title()}</h4>
	<NewUser onCreate={onNewUser} />
</Modal>
