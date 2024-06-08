<svelte:options immutable />

<script lang="ts">
	import type { Application } from '$lib/openapi/auth';

	export let application: Application;
</script>
