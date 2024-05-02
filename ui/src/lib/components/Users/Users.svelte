<svelte:options immutable />

<script lang="ts">
	import { userApi } from '$lib/openapi';
	import type { Application, User } from '$lib/openapi/auth';
	import { onMount } from 'svelte';

	export let application: Application;

	let users: User[] = [];
	onMount(async () => {
		const pagination = await userApi.users(application.id, 20, 0);
		users = pagination.items;
	});
</script>

<table class="w-full">
	<thead class="sticky top-0">
		<tr class="text-left">
			<th>Id</th>
			<th>Username</th>
		</tr>
	</thead>
	<tbody class="h-full overflow-y-auto">
		{#each users as user (user.id)}
			<tr>
				<td>{user.id}</td>
				<td>{user.username}</td>
			</tr>
		{/each}
	</tbody>
</table>
