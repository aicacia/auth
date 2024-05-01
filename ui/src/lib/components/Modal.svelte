<svelte:options immutable />

<script lang="ts">
	import { portal } from 'svelte-portal';
	import X from 'lucide-svelte/icons/x';
	import { fade, fly } from 'svelte/transition';

	export let onClose: () => void = () => undefined;
	export let small = false;
	export let open = false;
	export let backdrop = true;
	export let backgroundClose = backdrop;

	function close() {
		open = false;
		onClose();
	}

	let container: HTMLElement;
	function onClickOutside(e: Event) {
		if (!backgroundClose) {
			return;
		}
		if (container.contains(e.target as Node)) {
			e.stopPropagation();
			return;
		}
		if (open) {
			open = false;
			onClose();
		}
	}
</script>

<div use:portal class="relative" role="dialog" aria-modal="true">
	{#if open}
		<div
			class="fixed inset-0 z-[10000] bg-black bg-opacity-25"
			class:hidden={backdrop ? !open : true}
			transition:fade={{ duration: 150 }}
		/>
	{/if}
	{#if open}
		<div
			class="fixed inset-0 z-[10000] overflow-y-auto"
			on:pointerdown={onClickOutside}
			transition:fly={{ duration: 150, y: 300 }}
		>
			<div class="flex min-h-full items-end justify-center p-4 sm:items-center sm:p-0">
				<div
					class="relative w-full rounded bg-white dark:bg-gray-800 shadow-lg sm:my-8 sm:max-w-lg"
					class:m-auto={small}
					bind:this={container}
				>
					<div class="flex flex-shrink flex-row items-start justify-between px-4 pt-4">
						<div class="flex flex-grow flex-col">
							<slot name="title" />
						</div>
						<button class="btn primary icon" on:click|stopPropagation={close}>
							<div class="h-6 w-6"><X /></div>
						</button>
					</div>
					<div class="relative flex-grow flex-col px-4 pb-4 pt-0">
						<slot />
					</div>
				</div>
			</div>
		</div>
	{/if}
</div>
