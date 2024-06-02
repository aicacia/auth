<svelte:options immutable />

<script lang="ts" context="module">
	import { create, test, enforce, only } from 'vest';

	type UpdateTenentForm = UpdateTenent;

	const createSuite = (LL: TranslationFunctions) =>
		create((data: Partial<UpdateTenentForm> = {}, fields: string[]) => {
			if (!fields.length) {
				return;
			}
			only(fields);

			test('description', LL.errors.message.required(), () => {
				enforce(data.description).isNotBlank();
			});
			test('uri', LL.errors.message.required(), () => {
				enforce(data.uri).isNotBlank();
			});
		});
</script>

<script lang="ts">
	import type { Tenent, UpdateTenent } from '$lib/openapi/auth';
	import type { TranslationFunctions } from '$lib/i18n/i18n-types';
	import LL from '$lib/i18n/i18n-svelte';
	import { debounce } from '@aicacia/debounce';
	import classNames from 'vest/classnames';
	import InputResults from '../InputResults.svelte';
	import { handleError } from '$lib/errors';
	import Spinner from '../Spinner.svelte';
	import deepEqual from 'deep-equal';
	import { tenentApi } from '$lib/openapi';

	export let tenent: Tenent;
	export let onUpdate: (tenent: Tenent) => void;

	$: tenentUpdates = {
		...tenent
	};

	$: suite = createSuite($LL);
	$: result = suite.get();
	$: hasUpdates = !deepEqual(tenentUpdates, tenent);
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
		suite(tenentUpdates, Array.from(fields)).done((r) => {
			result = r;
		});
		fields.clear();
		hasUpdates = !deepEqual(tenentUpdates, tenent);
	}, 300);
	function validateAll() {
		for (const field of Object.keys(tenentUpdates)) {
			fields.add(field);
		}
		validate();
		validate.flush();
	}
	function onChange(e: Event & { currentTarget: HTMLInputElement | HTMLSelectElement }) {
		fields.add(e.currentTarget.name);
		validate();
	}

	let loading = false;
	async function onSubmit() {
		try {
			loading = true;
			validateAll();
			if (result.isValid()) {
				tenent = await tenentApi.updateTenent(tenent.applicationId, tenent.id, tenentUpdates);
				await onUpdate(tenent);
				suite.reset();
				result = suite.get();
			}
		} catch (error) {
			await handleError(error);
		} finally {
			loading = false;
		}
	}
</script>

<form on:submit|preventDefault={onSubmit}>
	<div class="mb-2">
		<label for="description">{$LL.tenents.description()}</label>
		<input
			class="w-full {cn('description')}"
			type="text"
			name="description"
			placeholder={$LL.tenents.descriptionPlaceholder()}
			bind:value={tenentUpdates.description}
			on:input={onChange}
		/>
		<InputResults name="description" {result} />
	</div>
	<div class="mb-2">
		<label for="uri">{$LL.tenents.uri()}</label>
		<input
			class="w-full {cn('uri')}"
			type="text"
			name="uri"
			placeholder={$LL.tenents.uriPlaceholder()}
			bind:value={tenentUpdates.uri}
			on:input={onChange}
		/>
		<InputResults name="uri" {result} />
	</div>
	<div class="flex flex-row justify-end">
		{#if hasUpdates}
			<button type="submit" class="btn primary flex flex-shrink" {disabled}>
				{#if loading}<div class="mr-2 flex flex-row justify-center">
						<div class="inline-block h-6 w-6"><Spinner /></div>
					</div>{/if}
				Update
			</button>
		{/if}
	</div>
</form>
