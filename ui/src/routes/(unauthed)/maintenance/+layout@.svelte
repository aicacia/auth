<svelte:options immutable />

<div class="flex h-full w-full flex-col">
	<div class="flex flex-grow flex-col overflow-y-auto overflow-x-hidden">
		<slot />
	</div>
</div>
