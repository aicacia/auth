<svelte:options immutable />

<script lang="ts" context="module">
	import { create, only } from 'vest';

	const createSuite = (LL: TranslationFunctions) =>
		create((data: Partial<UserInfo> = {}, fields: string[]) => {
			if (!fields.length) {
				return;
			}
			only(fields);
		});
</script>

<script lang="ts">
	import classNames from 'vest/classnames';
	import Spinner from '$lib/components/Spinner.svelte';
	import { handleError } from '$lib/errors';
	import { debounce } from '@aicacia/debounce';
	import InputResults from '$lib/components/InputResults.svelte';
	import type { UserInfo } from '$lib/openapi/auth';
	import type { TranslationFunctions } from '$lib/i18n/i18n-types';
	import LL from '$lib/i18n/i18n-svelte';

	export let userInfo: UserInfo;
	export let onUpdate: (data: Partial<UserInfo>) => Promise<void>;

	$: newUserInfo = { ...userInfo };

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
		suite(newUserInfo, Array.from(fields)).done((r) => {
			result = r;
		});
		fields.clear();
	}, 300);
	function validateAll() {
		fields.add('name');
		fields.add('givenName');
		fields.add('middleName');
		fields.add('familyName');
		fields.add('nickname');
		fields.add('birthdate');
		fields.add('gender');
		fields.add('picture');
		fields.add('profile');
		fields.add('website');
		fields.add('locale');
		fields.add('zoneinfo');
		fields.add('address.country');
		fields.add('address.locality');
		fields.add('address.postalCode');
		fields.add('address.region');
		fields.add('address.streetAddress');
		validate();
		validate.flush();
	}
	function onBirthdateChange(e: Event & { currentTarget: HTMLInputElement | HTMLSelectElement }) {
		newUserInfo.birthdate = new Date(e.currentTarget.value.trim());
		fields.add('birthdate');
		validate();
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
				await onUpdate(newUserInfo);
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
		<label for="name">Name</label>
		<input
			class="w-full {cn('name')}"
			type="text"
			name="name"
			autocomplete="name"
			placeholder="Name"
			bind:value={newUserInfo.name}
			on:input={onChange}
		/>
		<InputResults name="name" {result} />
	</div>
	<div class="mb-2">
		<label for="givenName">Given Name</label>
		<input
			class="w-full {cn('givenName')}"
			type="text"
			name="givenName"
			autocomplete="name"
			placeholder="Given Name"
			bind:value={newUserInfo.givenName}
			on:input={onChange}
		/>
		<InputResults name="givenName" {result} />
	</div>
	<div class="mb-2">
		<label for="middleName">Middle Name</label>
		<input
			class="w-full {cn('middleName')}"
			type="text"
			name="middleName"
			autocomplete="name"
			placeholder="Middle Name"
			bind:value={newUserInfo.middleName}
			on:input={onChange}
		/>
		<InputResults name="middleName" {result} />
	</div>
	<div class="mb-2">
		<label for="familyName">Family Name</label>
		<input
			class="w-full {cn('familyName')}"
			type="text"
			name="familyName"
			autocomplete="name"
			placeholder="Family Name"
			bind:value={newUserInfo.familyName}
			on:input={onChange}
		/>
		<InputResults name="familyName" {result} />
	</div>
	<div class="mb-2">
		<label for="nickname">Nickname</label>
		<input
			class="w-full {cn('nickname')}"
			type="text"
			name="nickname"
			autocomplete="name"
			placeholder="Nickname"
			bind:value={newUserInfo.nickname}
			on:input={onChange}
		/>
		<InputResults name="nickname" {result} />
	</div>
	<div class="mb-2">
		<label for="birthdate">Birthdate</label>
		<input
			class="w-full {cn('birthdate')}"
			type="date"
			name="birthdate"
			placeholder="Birthdate"
			on:input={onBirthdateChange}
		/>
		<InputResults name="birthdate" {result} />
	</div>
	<div class="mb-2">
		<label for="gender">Gender</label>
		<select
			class="w-full {cn('gender')}"
			name="gender"
			autocomplete="gender"
			bind:value={newUserInfo.gender}
			on:input={onChange}
		>
			<option value="male">Male</option>
			<option value="female">Female</option>
		</select>
		<InputResults name="gender" {result} />
	</div>
	<div class="mb-2">
		<label for="picture">Picture</label>
		<input
			class="w-full {cn('picture')}"
			type="text"
			name="picture"
			autocomplete="picture"
			placeholder="Picture"
			bind:value={newUserInfo.picture}
			on:input={onChange}
		/>
		<InputResults name="picture" {result} />
	</div>
	<div class="mb-2">
		<label for="website">Website</label>
		<input
			class="w-full {cn('website')}"
			type="url"
			name="website"
			autocomplete="name"
			placeholder="Website"
			bind:value={newUserInfo.website}
			on:input={onChange}
		/>
		<InputResults name="website" {result} />
	</div>
	<div class="mb-2">
		<label for="locale">Locale</label>
		<select
			class="w-full {cn('locale')}"
			name="locale"
			autocomplete="locale"
			bind:value={newUserInfo.locale}
			on:input={onChange}
		>
			<option value="en-US">English (US)</option>
		</select>
		<InputResults name="name" {result} />
	</div>
	<div class="mb-2">
		<label for="zoneinfo">Zone Info</label>
		<select
			class="w-full {cn('zoneinfo')}"
			name="zoneinfo"
			autocomplete="zoneinfo"
			bind:value={newUserInfo.locale}
			on:input={onChange}
		>
			<option value="American/New_York">American/New York</option>
		</select>
		<InputResults name="name" {result} />
	</div>
	<div class="mb-2">
		<label for="address.country">Country</label>
		<input
			class="w-full {cn('address.country')}"
			type="url"
			name="address.country"
			autocomplete="country"
			placeholder="Country"
			bind:value={newUserInfo.address.country}
			on:input={onChange}
		/>
		<InputResults name="address.country" {result} />
	</div>
	<div class="mb-2">
		<label for="address.locality">Locality</label>
		<input
			class="w-full {cn('address.locality')}"
			type="url"
			name="address.locality"
			autocomplete="locality"
			placeholder="Locality"
			bind:value={newUserInfo.address.locality}
			on:input={onChange}
		/>
		<InputResults name="address.region" {result} />
	</div>
	<div class="mb-2">
		<label for="address.region">Region</label>
		<input
			class="w-full {cn('address.region')}"
			type="url"
			name="address.region"
			autocomplete="name"
			placeholder="Region"
			bind:value={newUserInfo.address.region}
			on:input={onChange}
		/>
		<InputResults name="address.region" {result} />
	</div>
	<div class="mb-2">
		<label for="address.country">Street Address</label>
		<input
			class="w-full {cn('address.streetAddress')}"
			type="url"
			name="address.streetAddress"
			autocomplete="street-address"
			placeholder="Street Address"
			bind:value={newUserInfo.address.streetAddress}
			on:input={onChange}
		/>
		<InputResults name="address.streetAddress" {result} />
	</div>
	<div class="flex flex-row justify-end">
		<button type="submit" class="btn primary flex flex-shrink" {disabled}>
			{#if loading}<div class="flex flex-row justify-center mr-2">
					<div class="inline-block w-6 h-6"><Spinner /></div>
				</div>{/if}
			Update
		</button>
	</div>
</form>
