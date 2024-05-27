<svelte:options immutable />

<script lang="ts" context="module">
</script>

<script lang="ts">
	import { page } from '$app/stores';
	import MFA from '$lib/components/MFA/MFA.svelte';
	import LL from '$lib/i18n/i18n-svelte';
	import { onMount } from 'svelte';

	let mfaType: 'totp' | null = null;
	onMount(() => {
		mfaType = $page.url.searchParams.get('type') as any;
	});
</script>

<svelte:head>
	<title>{$LL.mfa.title()}</title>
</svelte:head>

<div class="flex flex-grow flex-col justify-end md:justify-start">
	<div class="mx-auto flex w-full flex-shrink flex-col p-4 py-10 md:w-96">
		<div class="flex flex-grow flex-col bg-white p-4 shadow dark:bg-gray-800">
			<h4 class="mb-1">{$LL.mfa.title()}</h4>
			{#if mfaType}
				<MFA type={mfaType} />
			{/if}
		</div>
	</div>
</div>
