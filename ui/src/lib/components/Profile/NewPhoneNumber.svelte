<svelte:options immutable />

<script lang="ts" context="module">
	import { create, test, enforce, only } from 'vest';

	type NewPhoneNumberForm = {
		phoneNumber: string;
	};

	const createSuite = (LL: TranslationFunctions) =>
		create((data: Partial<NewPhoneNumberForm> = {}, fields: string[]) => {
			if (!fields.length) {
				return;
			}
			only(fields);

			test('phoneNumber', LL.errors.message.required(), () => {
				enforce(data.phoneNumber).isNotBlank();
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
	import type { TranslationFunctions } from '$lib/i18n/i18n-types';
	import LL from '$lib/i18n/i18n-svelte';

	export let onCreate: (phoneNumber: string) => Promise<void>;

	let phoneNumber = '';

	$: suite = createSuite($LL);
	$: result = suite.get();
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
		suite({ phoneNumber }, Array.from(fields)).done((r) => {
			result = r;
		});
		fields.clear();
	}, 300);
	function validateAll() {
		fields.add('phoneNumber');
		validate();
		validate.flush();
	}
	function onChange(e: Event & { currentTarget: HTMLInputElement | HTMLSelectElement }) {
		fields.add(e.currentTarget.name);
		validate();
	}

	let adding = false;
	function onAddPhoneNumber() {
		adding = true;
	}
	function onCancelPhoneNumber() {
		adding = false;
	}

	let loading = false;
	async function onSubmit() {
		try {
			loading = true;
			validateAll();
			if (result.isValid()) {
				await onCreate(phoneNumber);
				adding = false;
				phoneNumber = '';
			}
		} catch (error) {
			await handleError(error);
		} finally {
			loading = false;
		}
	}
</script>

<div class="mt-4 flex flex-grow flex-row justify-end">
	<form class="flex flex-grow flex-col" on:submit|preventDefault={onSubmit} class:hidden={!adding}>
		<div class="mb-2">
			<input
				class="w-full {cn('phoneNumber')}"
				type="phoneNumber"
				name="phoneNumber"
				autocomplete="cc-number"
				placeholder={$LL.profile.phoneNumbers.newPlaceholder()}
				bind:value={phoneNumber}
				on:input={onChange}
			/>
			<InputResults name="phoneNumber" {result} />
		</div>
		<div class="flex flex-row justify-end">
			<button class="btn secondary flex flex-shrink" on:click={onCancelPhoneNumber} {disabled}>
				{$LL.profile.phoneNumbers.cancel()}
			</button>
			<button type="submit" class="btn primary flex flex-shrink" {disabled}>
				{#if loading}<div class="mr-2 flex flex-row justify-center">
						<div class="inline-block h-6 w-6"><Spinner /></div>
					</div>{/if}
				{$LL.profile.phoneNumbers.add()}
			</button>
		</div>
	</form>
	<button class="btn primary icon" on:click={onAddPhoneNumber} class:hidden={adding}>
		<Plus />
	</button>
</div>
