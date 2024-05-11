<svelte:options immutable />

<script lang="ts">
	import { goto } from '$app/navigation';
	import { base } from '$app/paths';
	import Dropdown from '$lib/components/Dropdown.svelte';
	import { signOut, signedIn } from '$lib/stores/user';
	import LogOut from 'lucide-svelte/icons/log-out';
	import Menu from 'lucide-svelte/icons/menu';
	import AppWindow from 'lucide-svelte/icons/app-window';
	import User from 'lucide-svelte/icons/user';
	import { page } from '$app/stores';
	import LL from '$lib/i18n/i18n-svelte';

	async function onSignOut() {
		signOut();
		onGoto();
		await goto(`${base}/signin`);
	}

	let open = false;
	function onGoto() {
		open = false;
	}
</script>

<div class="flex flex-shrink flex-row justify-between bg-white shadow dark:bg-gray-800">
	<div class="ms-2 flex flex-shrink flex-row">
		<a class="btn text-lg" href={`${base}/`}>{$LL.header.title()}</a>
	</div>
	<div class="me-2 flex flex-shrink flex-row">
		<div class="flex flex-col content-center justify-center">
			{#if $signedIn}
				<Dropdown bind:open>
					<Menu slot="button" />
					<a
						href={`${base}/applications`}
						class="default flex cursor-pointer flex-row justify-between p-2 hover:bg-gray-200 dark:hover:bg-gray-600"
						class:active={$page.route.id === '/(authed)/applications'}
						on:click={onGoto}
					>
						<AppWindow /><span class="ms-4">{$LL.header.applications()}</span>
					</a>
					<a
						href={`${base}/profile`}
						class="default flex cursor-pointer flex-row justify-between p-2 hover:bg-gray-200 dark:hover:bg-gray-600"
						class:active={$page.route.id === '/(authed)/profile'}
						on:click={onGoto}
					>
						<User /><span class="ms-4">{$LL.header.profile()}</span>
					</a>
					<hr />
					<!-- svelte-ignore a11y-click-events-have-key-events -->
					<!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
					<li
						class="flex cursor-pointer flex-row justify-between p-2 hover:bg-gray-200 dark:hover:bg-gray-600"
						on:click={onSignOut}
					>
						<LogOut /><span class="ms-4">{$LL.header.signOut()}</span>
					</li>
				</Dropdown>
			{:else}
				<a href={`${base}/signin`}>{$LL.header.signIn()}</a>
			{/if}
		</div>
	</div>
</div>

<style lang="postcss">
	li.active,
	a.active {
		@apply bg-gray-200 dark:bg-gray-600;
	}
</style>
