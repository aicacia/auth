<svelte:options immutable />

<script lang="ts" context="module">
	import { create, test, enforce, only, omitWhen } from 'vest';

	export type TenentKeyForm = {
		algorithm: string;
		publicKey: string;
		privateKey: string;
	};

	const defaultPrivateKey = {};

	function isHMAC(algorithm?: string) {
		return algorithm === 'HS256' || algorithm === 'HS384' || algorithm === 'HS512';
	}

	const createSuite = (LL: TranslationFunctions) =>
		create((data: Partial<TenentKeyForm> = {}, fields: string[]) => {
			if (!fields.length) {
				return;
			}
			only(fields);

			test('algorithm', LL.errors.message.required(), () => {
				enforce(data.algorithm).isNotBlank();
			});
			omitWhen(
				() => isHMAC(data.algorithm),
				() => {
					test('publicKey', LL.errors.message.required(), () => {
						enforce(data.publicKey).isNotBlank();
					});
				}
			);
			test('privateKey', LL.errors.message.required(), () => {
				enforce(data.privateKey).isNotBlank();
			});
		});
</script>

<script lang="ts">
	import type { TranslationFunctions } from '$lib/i18n/i18n-types';
	import LL from '$lib/i18n/i18n-svelte';
	import RefreshCCW from 'lucide-svelte/icons/refresh-ccw';
	import Undo2 from 'lucide-svelte/icons/undo-2';
	import EyeOff from 'lucide-svelte/icons/eye-off';
	import Eye from 'lucide-svelte/icons/eye';
	import { debounce } from '@aicacia/debounce';
	import classNames from 'vest/classnames';
	import InputResults from '../InputResults.svelte';
	import { handleError } from '$lib/errors';
	import { generateRandomBase64String } from '$lib/util';
	import Spinner from '../Spinner.svelte';
	import deepEqual from 'deep-equal';

	export let getPrivateKey: () => Promise<string | undefined>;
	export let onUpdate: (data: TenentKeyForm) => void;

	export let algorithm = 'HS256';
	export let publicKey = '';
	let privateKey = defaultPrivateKey as string;

	const original = {
		algorithm,
		publicKey,
		privateKey
	};
	$: updates = {
		algorithm,
		publicKey,
		privateKey
	};

	$: suite = createSuite($LL);
	$: result = suite.get();
	$: hasUpdates = !deepEqual(updates, original);
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
		suite({ algorithm, publicKey, privateKey }, Array.from(fields)).done((r) => {
			result = r;
		});
		fields.clear();
		hasUpdates = !deepEqual(updates, original);
	}, 300);
	function validateAll() {
		fields.add('algorithm');
		fields.add('publicKey');
		fields.add('privateKey');
		validate();
		validate.flush();
	}
	function onChange(
		e: Event & { currentTarget: HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement }
	) {
		fields.add(e.currentTarget.name);
		validate();
	}
	function onAlgorithmChange(e: Event & { currentTarget: HTMLSelectElement }) {
		onChange(e);
		algorithm = e.currentTarget.value;
		switch (algorithm) {
			case 'HS256':
			case 'HS384':
			case 'HS512':
				generateNewPrivateKey();
				break;
			default:
				privateKey = publicKey = '';
				break;
		}
	}

	let loading = false;
	async function onSubmit() {
		try {
			loading = true;
			validateAll();
			if (result.isValid()) {
				await onUpdate({ algorithm, publicKey, privateKey });
				suite.reset();
				result = suite.get();
			}
		} catch (error) {
			await handleError(error);
		} finally {
			loading = false;
		}
	}

	function generateNewPrivateKey() {
		let size = 32;
		switch (algorithm) {
			case 'HS384':
				size = 48;
				break;
			case 'HS512':
				size = 64;
				break;
		}
		privateKey = generateRandomBase64String(size);
	}

	$: showPrivateKey = algorithm !== original.algorithm;
	async function togglePrivateKey() {
		if (!showPrivateKey) {
			await internalGetPrivateKey();
		}
		showPrivateKey = !showPrivateKey;
	}
	let gettingPrivateKey = false;
	let fetchedPrivateKey = false;
	async function internalGetPrivateKey() {
		if (fetchedPrivateKey || gettingPrivateKey) {
			return;
		}
		try {
			gettingPrivateKey = true;
			const pk = await getPrivateKey();
			if (pk) {
				original.privateKey = pk;
				if (algorithm === original.algorithm) {
					privateKey = pk;
				}
				fetchedPrivateKey = true;
			}
		} finally {
			gettingPrivateKey = false;
		}
	}

	function reset() {
		algorithm = original.algorithm;
		publicKey = original.publicKey;
		privateKey = original.privateKey;
	}
</script>

<form on:submit|preventDefault={onSubmit}>
	<div class="mb-2">
		<select
			class="w-full {cn('algorithm')}"
			name="algorithm"
			bind:value={algorithm}
			on:input={onAlgorithmChange}
		>
			{#each Object.entries($LL.algorithms) as [value, translationFn] (value)}
				<option {value}>{translationFn()}</option>
			{/each}
		</select>
		<InputResults name="description" {result} />
	</div>
	{#if !isHMAC(algorithm)}
		<div class="mb-2">
			<label for="publicKey">{$LL.tenents.publicKey()}</label>
			<textarea
				class="min-h-10 w-full {cn('publicKey')}"
				name="publicKey"
				bind:value={publicKey}
				on:input={onChange}
			/>
			<InputResults name="publicKey" {result} />
		</div>
	{/if}
	<div class="mb-2">
		<label for="privateKey">{$LL.tenents.privateKey()}</label>
		<div class="flex flex-row">
			<div class="relative flex-grow">
				{#if showPrivateKey}
					<textarea
						class="min-h-10 w-full pe-8 {cn('privateKey')}"
						name="privateKey"
						bind:value={privateKey}
						on:input={onChange}
					/>
				{:else}
					<input
						class="min-h-10 w-full pe-8 {cn('privateKey')}"
						type="password"
						name="privateKey"
						bind:value={privateKey}
						on:input={onChange}
					/>
				{/if}
				{#if algorithm === original.algorithm}
					<button
						class="btn icon absolute inset-y-0 right-0 h-min"
						disabled={gettingPrivateKey}
						on:click|preventDefault={togglePrivateKey}
					>
						{#if showPrivateKey}
							<Eye />
						{:else}
							<EyeOff />
						{/if}
					</button>
				{/if}
			</div>
			{#if isHMAC(algorithm)}
				<div class="flex flex-col">
					<button
						class="btn primary"
						title={$LL.tenents.edit.regenerateKey()}
						on:click|preventDefault={generateNewPrivateKey}
					>
						<RefreshCCW />
					</button>
				</div>
			{/if}
		</div>
		<InputResults name="privateKey" {result} />
	</div>
	<div class="flex flex-row justify-between">
		<button class="btn primary" title={$LL.tenents.edit.reset()} on:click|preventDefault={reset}>
			<Undo2 />
		</button>
		{#if hasUpdates}
			<button type="submit" class="btn primary flex flex-shrink" {disabled}>
				{#if loading}<div class="mr-2 flex flex-row justify-center">
						<div class="inline-block h-6 w-6"><Spinner /></div>
					</div>{/if}
				{$LL.tenents.edit.confirm()}
			</button>
		{/if}
	</div>
</form>
