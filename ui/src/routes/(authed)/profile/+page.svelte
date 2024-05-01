<svelte:options immutable />

<script lang="ts">
	import Profile from '$lib/components/Profile/Profile.svelte';
	import LL from '$lib/i18n/i18n-svelte';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<svelte:head>
	<title>{$LL.profile.title()}</title>
</svelte:head>

<Profile user={data.user} />
