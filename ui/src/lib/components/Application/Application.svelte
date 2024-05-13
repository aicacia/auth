<svelte:options immutable />

<script lang="ts">
	import type { Application } from '$lib/openapi/auth';
	import ApplicationEditor from './ApplicationEditor.svelte';
	import DeleteApplication from './DeleteApplication.svelte';

	export let application: Application;

	function onUpdate(updatedApplication: Application) {
		application = updatedApplication;
	}
</script>

<div
	class="mx-auto mt-4 flex w-full max-w-6xl flex-shrink flex-col bg-white p-4 shadow dark:bg-gray-800"
>
	<ApplicationEditor
		id={application.id}
		description={application.description}
		uri={application.uri}
		onDone={onUpdate}
	/>
</div>

<div
	class="mx-auto mt-4 flex w-full max-w-6xl flex-shrink flex-col bg-white p-4 shadow dark:bg-gray-800"
>
	<DeleteApplication id={application.id} originalUri={application.uri} />
</div>
