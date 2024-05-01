<svelte:options immutable />

<script lang="ts" context="module">
	import { create, test, enforce, only } from 'vest';

	type ApplicationForm = {
		description: string;
		uri: string;
	};

	const createSuite = (LL: TranslationFunctions) =>
		create((data: Partial<ApplicationForm> = {}, fields: string[]) => {
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
	import classNames from 'vest/classnames';
	import { handleError } from '$lib/errors';
	import { debounce } from '@aicacia/debounce';
	import InputResults from '$lib/components/InputResults.svelte';
	import Spinner from '$lib/components/Spinner.svelte';
	import { applicationApi } from '$lib/openapi';
	import type { Application } from '$lib/openapi/auth';
	import { toURLSafe } from '$lib/util';
	import { invalidate } from '$app/navigation';
	import { base } from '$app/paths';
	import type { TranslationFunctions } from '$lib/i18n/i18n-types';
	import LL from '$lib/i18n/i18n-svelte';

	export let id: number | undefined = undefined;
	export let description: string = '';
	export let uri: string = '';
	export let onDone: (application: Application) => void;

	let initalDescription = description;
	let initalURI = uri;

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
		suite({ description, uri }, Array.from(fields)).done((r) => {
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
	function onNameChange(e: Event & { currentTarget: HTMLInputElement | HTMLSelectElement }) {
		uri = toURLSafe(description);
		onChange(e);
	}
	function onURIChange(e: Event & { currentTarget: HTMLInputElement | HTMLSelectElement }) {
		uri = toURLSafe(uri);
		onChange(e);
	}

	let loading = false;
	async function onSubmit() {
		try {
			loading = true;
			validateAll();
			if (result.isValid()) {
				const application =
					id == null
						? await applicationApi.createApplication({ description, uri })
						: await applicationApi.updateApplication(id, { description, uri });
				await invalidate(`${base}/applications`);
				initalDescription = application.description;
				initalURI = application.uri;
				onDone(application);
				suite.reset();
			}
		} catch (error) {
			await handleError(error);
		} finally {
			loading = false;
		}
	}
</script>

<form class="flex flex-col flex-grow" on:submit|preventDefault={onSubmit}>
	<div class="mb-2">
		<label for="application-description">{$LL.application.description()}</label>
		<input
			id="application-description"
			class="w-full {cn('description')}"
			type="text"
			name="description"
			placeholder={$LL.application.description()}
			bind:value={description}
			on:input={onNameChange}
		/>
		<InputResults name="description" {result} />
	</div>
	<div class="mb-2">
		<label for="application-uri">{$LL.application.uri()}</label>
		<input
			id="application-uri"
			class="w-full {cn('uri')}"
			type="text"
			name="uri"
			placeholder={$LL.application.uri()}
			bind:value={uri}
			on:input={onURIChange}
		/>
		<InputResults name="uri" {result} />
	</div>
	{#if initalDescription !== description || initalURI !== uri}
		<div class="flex flex-row justify-end">
			<button type="submit" class="btn primary flex flex-shrink" {disabled}>
				{#if loading}<div class="flex flex-row justify-center mr-2">
						<div class="inline-block w-6 h-6"><Spinner /></div>
					</div>{/if}
				{#if id == null}{$LL.application.create()}{:else}{$LL.application.update()}{/if}
			</button>
		</div>
	{/if}
</form>
