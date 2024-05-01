<svelte:options immutable />

<script lang="ts" context="module">
</script>

<script lang="ts">
	import LL from '$lib/i18n/i18n-svelte';

	import Modal from '../Modal.svelte';
	import DeleteApplicationForm from './DeleteApplicationForm.svelte';

	export let id: number;
	export let originalUri: string;

	let open = false;
	function onOpenDelete() {
		open = true;
	}
</script>

<div>
	<h4>{$LL.application.delete.dangerZone()}</h4>
	<hr class="mb-1" />
	<p>
		{$LL.application.delete.dangerZoneMessage()}
	</p>
	<div class="flex flex-row justify-end">
		<button class="btn danger" on:click={onOpenDelete}
			>{$LL.application.delete.dangerZoneDeleteApplication()}</button
		>
	</div>
</div>

<Modal bind:open>
	<h4 slot="title">{$LL.application.delete.confirmTitle()}</h4>
	<DeleteApplicationForm {id} {originalUri} />
</Modal>
