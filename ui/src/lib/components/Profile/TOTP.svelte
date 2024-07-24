<svelte:options immutable />

<script lang="ts">
	import { currentUserApi } from '$lib/openapi';
	import type { TOTP, Tenent, TOTPWithSecret, User } from '$lib/openapi/auth';
	import QRCode from 'qrcode';
	import Plus from 'lucide-svelte/icons/plus';
	import Trash from 'lucide-svelte/icons/trash';
	import Modal from '../Modal.svelte';
	import LL from '$lib/i18n/i18n-svelte';

	export let user: User;
	export let totp: TOTP | undefined = undefined;
	export let tenent: Tenent;
	export let onUpdate: (totp: TOTPWithSecret) => void;
	export let onDelete: (tenent: Tenent) => void;

	let creating = false;
	let secretUrl: string | undefined;
	let secretDataUrl: string | undefined;
	async function onCreateTOTP() {
		try {
			creating = true;
			const newTOTP = await currentUserApi.createTotp(tenent.id);
			totp = newTOTP;
			secretUrl = `otpauth://totp/${encodeURIComponent(tenent.description)}:${encodeURIComponent(user.username)}?secret=${encodeURIComponent(newTOTP.secret)}&issuer=${encodeURIComponent(tenent.description)}&algorithm=SHA1&digits=6&period=30`;
			secretDataUrl = await QRCode.toDataURL(secretUrl);
			onUpdate(newTOTP);
		} finally {
			creating = false;
		}
	}
	let deleteOpen = false;
	function onOpenDeleteTOTP() {
		deleteOpen = true;
	}
	let deleting = false;
	async function onDeleteTOTP() {
		try {
			deleting = true;
			await currentUserApi.deleteTotp(tenent.id);
			totp = secretUrl = secretDataUrl = undefined;
			deleteOpen = false;
			onDelete(tenent);
		} finally {
			deleting = false;
		}
	}

	let toggingEnabled = false;
	async function onToggleEnabled() {
		try {
			toggingEnabled = true;
			const newTOTP = await currentUserApi.enableTotp(tenent.id);
			totp = newTOTP;
			onUpdate(newTOTP);
		} finally {
			toggingEnabled = false;
		}
	}
</script>

<div>
	<h5>{tenent.description}</h5>
	{#if secretDataUrl && secretUrl}
		<img src={secretDataUrl} alt={secretUrl} />
	{/if}
</div>
<div class="flex flex-row justify-between">
	{#if totp}
		<div>
			<label for="{totp.id}-enabled"
				>{#if totp.enabled}{$LL.profile.mfa.enabled()}{:else}{$LL.profile.mfa.disabled()}{/if}</label
			>
			<input
				name="{totp.id}-enabled"
				type="checkbox"
				checked={totp.enabled}
				disabled={toggingEnabled}
				on:change={onToggleEnabled}
			/>
		</div>
		<button class="btn danger icon" on:click={onOpenDeleteTOTP} disabled={deleting}>
			<Trash />
		</button>
	{:else}
		<div></div>
		<button class="btn primary icon" on:click={onCreateTOTP} disabled={creating}>
			<Plus />
		</button>
	{/if}
</div>

<Modal bind:open={deleteOpen}>
	<h4 slot="title">{$LL.profile.totps.deleteTOTP(tenent.description)}</h4>
	<form on:submit|preventDefault={onDeleteTOTP}>
		<div class="mt-2 flex flex-row justify-end">
			<button class="btn danger" type="submit">{$LL.profile.phoneNumbers.delete()}</button>
		</div>
	</form>
</Modal>
