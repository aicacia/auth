<svelte:options immutable />

<script lang="ts">
	import LL from '$lib/i18n/i18n-svelte';
	import type { Application } from '$lib/openapi/auth';
	import Resources from './Resources.svelte';
	import Roles from './Roles.svelte';

	export let application: Application;

	let tab: 'roles' | 'resources' = 'roles';
	function createOnTab(newTab: typeof tab) {
		return () => {
			tab = newTab;
		};
	}
</script>

<div class="flex flex-grow flex-col">
	<div class="flex flex-shrink flex-row">
		<button class="btn primary me-2" class:active={tab === 'roles'} on:click={createOnTab('roles')}
			>{$LL.rbac.roles.title()}</button
		>
		<button
			class="btn primary"
			class:active={tab === 'resources'}
			on:click={createOnTab('resources')}>{$LL.rbac.resources.title()}</button
		>
	</div>
	<hr class="my-2" />
	<div class="flex flex-grow flex-col">
		{#if tab === 'roles'}
			<Roles {application} />
		{:else if tab === 'resources'}
			<Resources {application} />
		{/if}
	</div>
</div>
