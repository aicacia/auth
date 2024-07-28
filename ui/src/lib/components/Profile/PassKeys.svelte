<svelte:options immutable />

<script lang="ts">
	import { passkeyApi } from '$lib/openapi';
	import { base64URLToUint8Array } from '$lib/util';
	import Plus from 'lucide-svelte/icons/plus';

	async function onCreatePassKey() {
		const publicKey = await passkeyApi.passkeyBeginRegistration();
		const credential = await navigator.credentials.create({
			publicKey: {
				...publicKey,
				challenge: new Uint8Array(publicKey.challenge || []).buffer,
				user: {
					...publicKey.user,
					id: base64URLToUint8Array(publicKey.user.id as string)
				}
			}
		});
	}
</script>

<div class="flex flex-col"></div>

<div class="flex flex-row justify-end">
	<button class="btn primary icon" on:click={onCreatePassKey}>
		<Plus />
	</button>
</div>
