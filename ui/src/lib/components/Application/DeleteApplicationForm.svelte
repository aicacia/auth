<svelte:options immutable />

<script lang="ts" context="module">
	import { create, test, enforce, only } from 'vest';

	type URIForm = {
		uri: string;
		originalUri: string;
	};

	const createSuite = (LL: TranslationFunctions) =>
		create((data: Partial<URIForm> = {}, fields: string[]) => {
			if (!fields.length) {
				return;
			}
			only(fields);

			test('uri', LL.errors.message.required(), () => {
				enforce(data.uri).isNotBlank();
			});
			test('uri', LL.errors.message.mismatch(), () => {
				enforce(data.uri).equals(data.originalUri);
			});
		});
</script>

<script lang="ts">
	import classNames from 'vest/classnames';
	import { handleError } from '$lib/errors';
	import { debounce } from '@aicacia/debounce';
	import InputResults from '$lib/components/InputResults.svelte';
	import Spinner from '$lib/components/Spinner.svelte';
	import { applicationApi } from '$lib/openapi';
	import { goto, invalidate } from '$app/navigation';
	import { base } from '$app/paths';
	import LL from '$lib/i18n/i18n-svelte';
	import type { TranslationFunctions } from '$lib/i18n/i18n-types';

	export let id: number;
	export let originalUri: string;

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

	let uri = '';

	const fields = new Set<string>();
	const validate = debounce(() => {
		suite({ uri, originalUri }, Array.from(fields)).done((r) => {
			result = r;
		});
		fields.clear();
	}, 300);
	function validateAll() {
		fields.add('uri');
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
				await applicationApi.deleteApplication(id);
				await invalidate(`${base}/applications`);
				await goto(`${base}/applications`);
			}
		} catch (error) {
			await handleError(error);
		} finally {
			loading = false;
		}
	}
</script>

<form class="flex flex-col flex-grow" on:submit|preventDefault={onSubmit}>
	<label for="uri">{$LL.application.delete.confirmMessage(originalUri)}</label>
	<div class="flex flex-row">
		<div class="flex flex-grow relative">
			<input
				id="uri"
				class="flex flex-grow {cn('uri')}"
				type="text"
				name="uri"
				placeholder="URI"
				bind:value={uri}
				on:input={onChange}
			/>
		</div>
		<div class="flex flex-shrink">
			<button type="submit" class="btn danger flex flex-shrink" {disabled}>
				{#if loading}<div class="flex flex-row justify-center mr-2">
						<div class="inline-block w-6 h-6"><Spinner /></div>
					</div>{/if}
				{$LL.application.delete.confirm()}
			</button>
		</div>
	</div>
	<InputResults name="uri" {result} />
</form>
