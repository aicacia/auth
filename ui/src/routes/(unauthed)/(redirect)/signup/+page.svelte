<svelte:options immutable />

<script lang="ts" context="module">
	import { create, test, enforce, only } from 'vest';

	type SignUpForm = {
		username: string;
		password: string;
		passwordConfirmation: string;
		email: string;
	};

	const createSuite = (LL: TranslationFunctions) =>
		create((data: Partial<SignUpForm> = {}, fields: Set<string>) => {
			if (!fields.size) {
				return;
			}
			only(Array.from(fields));

			test('email', LL.errors.message.required(), () => {
				enforce(data.email).isNotBlank();
			});
			test('username', LL.errors.message.required(), () => {
				enforce(data.username).isNotBlank();
			});
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
	import LL from '$lib/i18n/i18n-svelte';
	import classNames from 'vest/classnames';
	import { goto } from '$app/navigation';
	import { base } from '$app/paths';
	import Spinner from '$lib/components/Spinner.svelte';
	import { signUp } from '$lib/stores/user';
	import { handleError } from '$lib/errors';
	import { debounce } from '@aicacia/debounce';
	import InputResults from '$lib/components/InputResults.svelte';
	import type { TranslationFunctions } from '$lib/i18n/i18n-types';

	let username = '';
	let email = '';
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
		suite({ username, password, passwordConfirmation, email }, fields).done((r) => {
			result = r;
		});
		fields.clear();
	}, 300);
	function validateAll() {
		fields.add('username');
		fields.add('password');
		fields.add('passwordConfirmation');
		fields.add('email');
		fields.add('countryId');
		validate();
		validate.flush();
	}
	function onChange(e: Event & { currentTarget: HTMLInputElement | HTMLSelectElement }) {
		e.currentTarget.value = e.currentTarget.value.trim();
		fields.add(e.currentTarget.name);
		validate();
	}

	let loading = false;
	async function onSubmit() {
		try {
			loading = true;
			validateAll();
			if (result.isValid()) {
				await signUp(username, password, passwordConfirmation, email);
				await goto(`${base}/`);
			}
		} catch (error) {
			await handleError(error);
		} finally {
			loading = false;
		}
	}
</script>

<svelte:head>
	<title>{$LL.auth.signUp()}</title>
</svelte:head>

<div class="flex flex-col flex-grow justify-end md:justify-start">
	<div
		class="flex flex-col flex-shrink md:w-72 w-full mx-auto my-10 bg-white dark:bg-gray-800 shadow p-4"
	>
		<h1 class="mb-1">{$LL.auth.signUp()}</h1>
		<p class="py-2">
			<span>{$LL.auth.alreadyAMember()}</span>
			<a href={`${base}/signin`} class="underline text-blue-500">{$LL.auth.signIn()}</a>
		</p>
		<form on:submit|preventDefault={onSubmit}>
			<div class="mb-2">
				<input
					class="w-full {cn('email')}"
					type="email"
					name="email"
					autocomplete="email"
					placeholder={$LL.auth.emailPlaceholder()}
					bind:value={email}
					on:input={onChange}
				/>
				<InputResults name="email" {result} />
			</div>
			<div class="mb-2">
				<input
					class="w-full {cn('username')}"
					type="text"
					name="username"
					autocomplete="username"
					placeholder={$LL.auth.usernamePlaceholder()}
					bind:value={username}
					on:input={onChange}
				/>
				<InputResults name="username" {result} />
			</div>
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
				<button type="submit" class="btn primary flex flex-shrink" {disabled}>
					{#if loading}<div class="flex flex-row justify-center mr-2">
							<div class="inline-block w-6 h-6"><Spinner /></div>
						</div>{/if}
					{$LL.auth.signUp()}
				</button>
			</div>
		</form>
	</div>
</div>
