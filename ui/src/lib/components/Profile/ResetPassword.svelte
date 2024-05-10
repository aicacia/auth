<svelte:options immutable />

<script lang="ts" context="module">
	import { create, test, enforce, only } from 'vest';

	type ProfileForm = {
		password: string;
		passwordConfirmation: string;
	};

	const createSuite = (LL: TranslationFunctions) =>
		create((data: Partial<ProfileForm> = {}, fields: string[]) => {
			if (!fields.length) {
				return;
			}
			only(fields);

			test('password', LL.errors.message.required(), () => {
				enforce(data.password).isNotBlank();
			});
			test('passwordConfirmation', LL.errors.message.required(), () => {
				enforce(data.passwordConfirmation).isNotBlank();
			});
			test('passwordConfirmation', LL.errors.message.mismatch(), () => {
				enforce(data.passwordConfirmation).equals(data.password);
			});
		});
</script>

<script lang="ts">
	import classNames from 'vest/classnames';
	import Spinner from '$lib/components/Spinner.svelte';
	import { handleError } from '$lib/errors';
	import { debounce } from '@aicacia/debounce';
	import InputResults from '$lib/components/InputResults.svelte';
	import { currentUserApi } from '$lib/openapi';
	import { createNotification } from '$lib/stores/notifications';
	import type { TranslationFunctions } from '$lib/i18n/i18n-types';
	import LL from '$lib/i18n/i18n-svelte';
	import { get } from 'svelte/store';

	let password = '';
	let passwordConfirmation = '';

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
		suite({ password, passwordConfirmation }, Array.from(fields)).done((r) => {
			result = r;
		});
		fields.clear();
	}, 300);
	function validateAll() {
		fields.add('password');
		fields.add('passwordConfirmation');
		validate();
		validate.flush();
	}
	function onChange(e: Event & { currentTarget: HTMLInputElement | HTMLSelectElement }) {
		fields.add(e.currentTarget.name);
		validate();
	}

	let loading = false;
	async function onSubmit(e: SubmitEvent) {
		try {
			loading = true;
			validateAll();
			if (result.isValid()) {
				await currentUserApi.resetPassword({
					password,
					passwordConfirmation: passwordConfirmation
				});
				password = '';
				passwordConfirmation = '';
				suite.reset();
				result = suite.get();
				createNotification(get(LL).profile.notification.passwordResetSuccess(), 'success');
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
			class="w-full {cn('password')}"
			type="password"
			name="password"
			autocomplete="new-password"
			placeholder={$LL.auth.passwordPlaceholder()}
			bind:value={password}
			on:input={onChange}
		/>
		<InputResults name="password" {result} />
	</div>
	<div class="mb-2">
		<input
			class="w-full {cn('passwordConfirmation')}"
			type="password"
			name="passwordConfirmation"
			autocomplete="new-password"
			placeholder={$LL.auth.passwordConfirmationPlaceholder()}
			bind:value={passwordConfirmation}
			on:input={onChange}
		/>
		<InputResults name="passwordConfirmation" {result} />
	</div>
	<div class="flex flex-row justify-end">
		{#if password && passwordConfirmation}
			<button type="submit" class="btn primary flex flex-shrink" {disabled}>
				{#if loading}<div class="flex flex-row justify-center mr-2">
						<div class="inline-block w-6 h-6"><Spinner /></div>
					</div>{/if}
				{$LL.auth.reset()}
			</button>
		{/if}
	</div>
</form>
