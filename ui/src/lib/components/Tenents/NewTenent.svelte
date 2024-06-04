<svelte:options immutable />

<script lang="ts" context="module">
	import { create, only, test, enforce } from 'vest';

	export interface NewTenentForm {
		description: string;
		uri: string;
		authorizationWebsite: string;
	}

	const createSuite = (LL: TranslationFunctions) =>
		create((data: Partial<NewTenentForm> = {}, fields: string[]) => {
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
			test('authorizationWebsite', LL.errors.message.required(), () => {
				enforce(data.authorizationWebsite).isNotBlank();
			});
		});
</script>

<script lang="ts">
	import classNames from 'vest/classnames';
	import Spinner from '$lib/components/Spinner.svelte';
	import { handleError } from '$lib/errors';
	import { debounce } from '@aicacia/debounce';
	import InputResults from '$lib/components/InputResults.svelte';
	import type { TranslationFunctions } from '$lib/i18n/i18n-types';
	import LL from '$lib/i18n/i18n-svelte';
	import type { CreateTenent } from '$lib/openapi/auth';

	export let onCreate: (create: CreateTenent) => Promise<void>;

	let description = '';
	let uri = '';
	let authorizationWebsite = '';

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
		suite({ description, uri, authorizationWebsite }, Array.from(fields)).done((r) => {
			result = r;
		});
		fields.clear();
	}, 300);
	function validateAll() {
		fields.add('description');
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
				await onCreate({ description, uri, authorizationWebsite });
				description = uri = '';
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
		<input
			class="w-full {cn('description')}"
			type="text"
			name="description"
			placeholder={$LL.tenents.descriptionPlaceholder()}
			bind:value={description}
			on:input={onChange}
		/>
		<InputResults name="description" {result} />
	</div>
	<div class="mb-2">
		<input
			class="w-full {cn('uri')}"
			type="text"
			name="uri"
			placeholder={$LL.tenents.uriPlaceholder()}
			bind:value={uri}
			on:input={onChange}
		/>
		<InputResults name="uri" {result} />
	</div>
	<div class="mb-2">
		<label for="authorizationWebsite">{$LL.tenents.authorizationWebsite()}</label>
		<input
			class="w-full {cn('authorizationWebsite')}"
			type="text"
			name="authorizationWebsite"
			placeholder={$LL.tenents.authorizationWebsitePlaceholder()}
			bind:value={authorizationWebsite}
			on:input={onChange}
		/>
		<InputResults name="authorizationWebsite" {result} />
	</div>
	<div class="flex flex-row justify-end">
		<button type="submit" class="btn primary flex flex-shrink" {disabled}>
			{#if loading}<div class="mr-2 flex flex-row justify-center">
					<div class="inline-block h-6 w-6"><Spinner /></div>
				</div>{/if}
			{$LL.tenents.newTenent.button()}
		</button>
	</div>
</form>
