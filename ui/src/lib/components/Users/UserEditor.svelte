<svelte:options immutable />

<script lang="ts">
	import type { User } from '$lib/openapi/auth';
	import Profile from '../Profile/Profile.svelte';

	export let user: User;
</script>

<Profile bind:user />
