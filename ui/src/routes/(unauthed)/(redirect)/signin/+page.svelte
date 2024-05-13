<svelte:options immutable />

<script lang="ts" context="module">
	import { create, test, enforce, only } from 'vest';

	type SignInForm = {
		username: string;
		password: string;
	};

	const createSuite = (LL: TranslationFunctions) =>
		create((data: Partial<SignInForm> = {}, fields: string[]) => {
			if (!fields.length) {
				return;
			}
			only(fields);

			test('username', LL.errors.message.required(), () => {
				enforce(data.username).isNotBlank();
			});
			test('password', LL.errors.message.required(), () => {
				enforce(data.password).isNotBlank();
			});
		});
</script>

<script lang="ts">
	import LL from '$lib/i18n/i18n-svelte';
	import classNames from 'vest/classnames';
	import { goto } from '$app/navigation';
	import { base } from '$app/paths';
	import Spinner from '$lib/components/Spinner.svelte';
	import { signIn } from '$lib/stores/user';
	import { handleError } from '$lib/errors';
	import { debounce } from '@aicacia/debounce';
	import InputResults from '$lib/components/InputResults.svelte';
	import type { PageData } from './$types';
	import type { TranslationFunctions } from '$lib/i18n/i18n-types';

	export let data: PageData;

	let username = '';
	let password = '';
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
		suite({ username, password }, Array.from(fields)).done((r) => {
			result = r;
		});
		fields.clear();
	}, 300);
	function validateAll() {
		fields.add('username');
		fields.add('password');
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
				await signIn(username, password);
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
	<title>{$LL.auth.signIn()}</title>
</svelte:head>

<div class="flex flex-grow flex-col justify-end md:justify-start">
	<div
		class="mx-auto my-10 flex w-full flex-shrink flex-col bg-white p-4 shadow dark:bg-gray-800 md:w-72"
	>
		<h1 class="mb-1">{$LL.auth.signIn()}</h1>
		{#if data.openIDConfiguration?.grantTypesSupported.includes('password')}
			<p class="py-2">
				<span>{$LL.auth.notAMember()}</span>
				<a href={`${base}/signup`} class="text-blue-500 underline">{$LL.auth.signUp()}</a>
			</p>
		{/if}
		<form on:submit|preventDefault={onSubmit}>
			<div class="mb-2">
				<input
					class="w-full {cn('username')}"
					type="text"
					name="username"
					autocomplete="username"
					placeholder={$LL.auth.usernameOrEmailPlaceholder()}
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
					autocomplete="current-password"
					placeholder={$LL.auth.passwordPlaceholder()}
					bind:value={password}
					on:input={onChange}
				/>
				<InputResults name="password" {result} />
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
	</div>
</div>
