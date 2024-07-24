<svelte:options immutable />

<script lang="ts">
	import type { User, TOTP, Tenent } from '$lib/openapi/auth';
	import TOTPComponent from './TOTP.svelte';

	export let user: User;
	export let totps: TOTP[];
	export let tenents: Tenent[];

	$: totpsByTenentId = totps.reduce(
		(acc, totp) => {
			acc[totp.tenentId] = totp;
			return acc;
		},
		{} as { [tenentId: number]: TOTP }
	);

	function onTOTPUpdate(totp: TOTP) {
		const index = totps.findIndex((t) => t.id === totp.id);
		if (index === -1) {
			totps = [totp, ...totps];
		} else {
			const newTOTPs = totps.slice();
			newTOTPs[index] = totp;
			totps = newTOTPs;
		}
	}
	function onTOTPDelete(tenent: Tenent) {
		const index = totps.findIndex((totp) => totp.tenentId === tenent.id);
		if (index !== -1) {
			const newTOTPs = totps.slice();
			newTOTPs.splice(index, 1);
			totps = newTOTPs;
		}
	}
</script>

<div class="flex flex-col">
	{#each tenents as tenent (tenent.id)}
		{@const totp = totpsByTenentId[tenent.id]}
		<TOTPComponent {user} {totp} {tenent} onUpdate={onTOTPUpdate} onDelete={onTOTPDelete} />
	{/each}
</div>
