<svelte:options immutable />

<script lang="ts">
	import type { Tenent } from '$lib/openapi/auth';

	export let tenent: Tenent;
	export let onUpdate: (tenent: Tenent) => void;
</script>

<div class="flex flex-col justify-end md:justify-start">
	<div class="mb-2"></div>
</div>
