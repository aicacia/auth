<svelte:options immutable />

<script lang="ts">
	import TOTP from './TOTP.svelte';

	export let type: 'totp';
</script>

{#if type === 'totp'}
	<TOTP />
{/if}
