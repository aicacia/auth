<svelte:options immutable />

<script lang="ts" context="module">
	import { create, test, enforce, only } from 'vest';

	type TOTPForm = {
		code: string;
	};

	const createSuite = (LL: TranslationFunctions) =>
		create((data: Partial<TOTPForm> = {}, fields: string[]) => {
			if (!fields.length) {
				return;
			}
			only(fields);

			test('code', LL.errors.message.required(), () => {
				enforce(data.code).isNotBlank();
			});
		});
</script>

<script lang="ts">
	import LL from '$lib/i18n/i18n-svelte';
	import classNames from 'vest/classnames';
	import { goto } from '$app/navigation';
	import { base } from '$app/paths';
	import Spinner from '$lib/components/Spinner.svelte';
	import { signInWithToken } from '$lib/stores/user';
	import { handleError } from '$lib/errors';
	import { debounce } from '@aicacia/debounce';
	import InputResults from '$lib/components/InputResults.svelte';
	import type { TranslationFunctions } from '$lib/i18n/i18n-types';
	import { tokenApi } from '$lib/openapi';

	let code = '';
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
		suite({ code }, Array.from(fields)).done((r) => {
			result = r;
		});
		fields.clear();
	}, 300);
	function validateAll() {
		fields.add('code');
		validate();
		validate.flush();
	}
	function onChange(e: Event & { currentTarget: HTMLInputElement | HTMLSelectElement }) {
		e.currentTarget.value = e.currentTarget.value.trim();
		fields.add(e.currentTarget.name);
		validate();
	}

	let loading = false;
	async function onSubmit(e: SubmitEvent) {
		try {
			loading = true;
			validateAll();
			if (result.isValid()) {
				const token = await tokenApi.validateMfa({
					code
				});
				await signInWithToken(token);
				await goto(`${base}/`);
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
			class="w-full {cn('code')}"
			type="text"
			name="code"
			placeholder={$LL.mfa.totp.codePlaceHolder()}
			bind:value={code}
			on:input={onChange}
		/>
		<InputResults name="code" {result} />
	</div>
	<div class="flex flex-row justify-end">
		<button type="submit" class="btn primary flex flex-shrink" {disabled}>
			{#if loading}<div class="mr-2 flex flex-row justify-center">
					<div class="inline-block h-6 w-6"><Spinner /></div>
				</div>{/if}
			{$LL.auth.signIn()}
		</button>
	</div>
</form>
