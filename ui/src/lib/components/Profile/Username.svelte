<svelte:options immutable />

<script lang="ts" context="module">
	import { create, test, enforce, only } from 'vest';

	type ProfileForm = {
		username: string;
		initialUsername: string;
	};

	const createSuite = (LL: TranslationFunctions) =>
		create((data: Partial<ProfileForm> = {}, fields: string[]) => {
			if (!fields.length) {
				return;
			}
			only(fields);

			test('username', LL.errors.message.required(), () => {
				enforce(data.username).isNotBlank();
			});
			test('username', LL.errors.message.noChange(), () => {
				enforce(data.username).notEquals(data.initialUsername);
			});
		});
</script>

<script lang="ts">
	import classNames from 'vest/classnames';
	import Spinner from '$lib/components/Spinner.svelte';
	import { handleError } from '$lib/errors';
	import { debounce } from '@aicacia/debounce';
	import InputResults from '$lib/components/InputResults.svelte';
	import { createNotification } from '$lib/stores/notifications';
	import type { User } from '$lib/openapi/auth';
	import { currentUserApi } from '$lib/openapi';
	import { invalidateAll } from '$app/navigation';
	import { updateCurrentUser } from '$lib/stores/user';
	import type { TranslationFunctions } from '$lib/i18n/i18n-types';
	import LL from '$lib/i18n/i18n-svelte';
	import { get } from 'svelte/store';

	export let user: User;

	$: initialUsername = user.username;
	$: username = initialUsername;

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
		suite({ username, initialUsername }, Array.from(fields)).done((r) => {
			result = r;
		});
		fields.clear();
	}, 300);
	function validateAll() {
		fields.add('username');
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
				await currentUserApi.updateUsername({ username });
				user.username = username;
				updateCurrentUser(user);
				suite.reset();
				result = suite.get();
				createNotification(get(LL).profile.notification.usernameChangedSuccess(), 'success');
				await invalidateAll();
			}
		} catch (error) {
			await handleError(error);
		} finally {
			loading = false;
		}
	}
</script>

<h3 class="mb-1">{$LL.profile.updateUsername()}</h3>
<form on:submit|preventDefault={onSubmit}>
	<div class="mb-2">
		<input
			class="w-full {cn('username')}"
			type="username"
			name="username"
			autocomplete="username"
			placeholder="Username"
			bind:value={username}
			on:input={onChange}
		/>
		<InputResults name="username" {result} />
	</div>
	{#if initialUsername !== username}
		<div class="flex flex-row justify-end">
			<button type="submit" class="btn primary flex flex-shrink" {disabled}>
				{#if loading}<div class="flex flex-row justify-center mr-2">
						<div class="inline-block w-6 h-6"><Spinner /></div>
					</div>{/if}
				{$LL.profile.submitUpdateUsername()}
			</button>
		</div>
	{/if}
</form>
