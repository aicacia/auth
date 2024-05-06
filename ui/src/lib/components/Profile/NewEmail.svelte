<svelte:options immutable />

<script lang="ts" context="module">
	import { create, test, enforce, only } from 'vest';

	type NewEmailForm = {
		email: string;
	};

	const createSuite = () =>
		create((data: Partial<NewEmailForm> = {}, fields: string[]) => {
			if (!fields.length) {
				return;
			}
			only(fields);

			test('email', 'required', () => {
				enforce(data.email).isNotBlank();
			});
		});
</script>

<script lang="ts">
	import Plus from 'lucide-svelte/icons/plus';
	import classNames from 'vest/classnames';
	import { handleError } from '$lib/errors';
	import { debounce } from '@aicacia/debounce';
	import InputResults from '$lib/components/InputResults.svelte';
	import Spinner from '$lib/components/Spinner.svelte';

	export let onCreate: (email: string) => Promise<void>;

	const suite = createSuite();

	let email = '';

	let result = suite.get();
	$: disabled = loading;
	$: cn = classNames(result, {
		untested: 'untested',
		tested: 'tested',
		invalid: 'invalid',
		valid: 'valid',
		warning: 'warning'
	});

	const fields = new Set<string>();
	const validate = debounce(() => {
		suite({ email }, Array.from(fields)).done((r) => {
			result = r;
		});
		fields.clear();
	}, 300);
	function validateAll() {
		fields.add('email');
		validate();
		validate.flush();
	}
	function onChange(e: Event & { currentTarget: HTMLInputElement | HTMLSelectElement }) {
		fields.add(e.currentTarget.name);
		validate();
	}

	let adding = false;
	function onAddEmail() {
		adding = true;
	}
	function onCancelEmail() {
		adding = false;
	}

	let loading = false;
	async function onSubmit() {
		try {
			loading = true;
			validateAll();
			if (result.isValid()) {
				await onCreate(email);
				adding = false;
				email = '';
			}
		} catch (error) {
			await handleError(error);
		} finally {
			loading = false;
		}
	}
</script>

<div class="flex flex-row flex-grow justify-end mt-4">
	<form class="flex flex-col flex-grow" on:submit|preventDefault={onSubmit} class:hidden={!adding}>
		<div class="mb-2">
			<input
				class="w-full {cn('email')}"
				type="email"
				name="email"
				autocomplete="email"
				placeholder="New Email"
				bind:value={email}
				on:input={onChange}
			/>
			<InputResults name="email" {result} />
		</div>
		<div class="flex flex-row justify-end">
			<button class="btn secondary flex flex-shrink" on:click={onCancelEmail} {disabled}>
				Cancel
			</button>
			<button type="submit" class="btn primary flex flex-shrink" {disabled}>
				{#if loading}<div class="flex flex-row justify-center mr-2">
						<div class="inline-block w-6 h-6"><Spinner /></div>
					</div>{/if}
				Add
			</button>
		</div>
	</form>
	<button class="btn primary icon" on:click={onAddEmail} class:hidden={adding}>
		<Plus />
	</button>
</div>
