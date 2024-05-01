<svelte:options immutable />

<script lang="ts" context="module">
	export type Position =
		| 'top-left'
		| 'top-right'
		| 'bottom-right'
		| 'bottom-left'
		| 'top-center'
		| 'bottom-center';
</script>

<script lang="ts">
	import { clickoutside } from '@svelte-put/clickoutside';
	import { onMount, tick } from 'svelte';

	export let anchor: Element;
	export let anchorPosition: Position = 'bottom-right';
	export let position: Position = 'top-right';
	export let open = false;
	export let closeOnClickOutside = true;

	let child: HTMLElement;
	let offsetWidth: number;

	$: if (anchor && child && open) {
		resize();
	}

	async function resize() {
		const anchorRect = anchor.getBoundingClientRect();
		child.style.top = child.style.left = child.style.bottom = child.style.right = '';
		await tick();
		let isRightAligned = false;
		switch (anchorPosition) {
			case 'top-left': {
				child.style.top = `${anchorRect.top}px`;
				child.style.left = `${anchorRect.left}px`;
				break;
			}
			case 'top-right': {
				child.style.top = `${anchorRect.top}px`;
				child.style.right = `${anchor.ownerDocument.body.offsetWidth - anchorRect.right}px`;
				isRightAligned = true;
				break;
			}
			case 'top-center': {
				child.style.top = `${anchorRect.top}px`;
				child.style.left = `${anchorRect.left + anchorRect.width * 0.5 - offsetWidth * 0.5}px`;
				break;
			}
			case 'bottom-right': {
				child.style.top = `${anchorRect.bottom}px`;
				child.style.right = `${anchor.ownerDocument.body.offsetWidth - anchorRect.right}px`;
				isRightAligned = true;
				break;
			}
			case 'bottom-left': {
				child.style.top = `${anchorRect.bottom}px`;
				child.style.left = `${anchorRect.left}px`;
				break;
			}
			case 'bottom-center': {
				child.style.top = `${anchorRect.bottom}px`;
				child.style.left = `${anchorRect.left + anchorRect.width * 0.5 - offsetWidth * 0.5}px`;
				break;
			}
		}
		await tick();
		const rect = child.getBoundingClientRect();
		if (rect.left < 0) {
			child.style.right = '';
			child.style.left = '0px';
		} else if (rect.right > anchor.ownerDocument.body.offsetWidth) {
			child.style.right = '0px';
			child.style.left = '';
		}
		if (rect.top < 0) {
			child.style.top = '0px';
		} else if (rect.top > anchor.ownerDocument.body.offsetHeight) {
			child.style.top = '';
			child.style.bottom = '0px';
		}
	}

	onMount(() => {
		window.addEventListener('resize', resize);
		return () => {
			window.removeEventListener('resize', resize);
		};
	});

	function onClickOutside(e: Event) {
		e.stopPropagation();
		if (open) {
			open = false;
		}
	}
</script>

<div
	class="absolute z-[1001] max-h-full max-w-full border-gray-300 dark:border-gray-600 bg-gray-50 dark:bg-gray-700 shadow-md transition-transform duration-75 focus:outline-none"
	bind:this={child}
	bind:offsetWidth
	use:clickoutside={{ event: 'pointerdown', enabled: closeOnClickOutside }}
	on:clickoutside={onClickOutside}
	role="menu"
	class:scale-0={!open}
	class:scale-100={open}
	class:origin-top-right={position === 'top-right'}
	class:origin-top-left={position === 'top-left'}
	class:origin-top={position === 'top-center'}
	class:origin-bottom-right={position === 'bottom-right'}
	class:origin-bottom-left={position === 'bottom-left'}
	class:origin-bottom={position === 'bottom-center'}
	tabindex="-1"
>
	<slot />
</div>
