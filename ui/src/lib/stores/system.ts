import { browser, building } from '$app/environment';
import { handleError } from '$lib/errors';
import { wellKnownApi } from '$lib/openapi';
import type { OpenIDConfiguration } from '$lib/openapi/auth';
import EventEmitter from 'eventemitter3';
import { get, derived, writable } from 'svelte/store';

const openIDConfigurationWritable = writable<OpenIDConfiguration | null>(null);
const emitter = new EventEmitter<{ openIDConfiguration: OpenIDConfiguration }>();
let openIDConfigurationLoaded = false;

export const openIDConfiguration = derived(openIDConfigurationWritable, (state) => state);

export function waitForOpenIDConfiguration() {
	if (building || openIDConfigurationLoaded) {
		return Promise.resolve(get(openIDConfiguration));
	} else {
		return new Promise<OpenIDConfiguration>((resolve) =>
			emitter.once('openIDConfiguration', resolve)
		);
	}
}

async function fetchOpenIDConfiguration() {
	try {
		const openIDConfiguration = await wellKnownApi.openidConfiguration();
		openIDConfigurationWritable.set(openIDConfiguration);
		openIDConfigurationLoaded = true;
		emitter.emit('openIDConfiguration', openIDConfiguration);
	} catch (e) {
		await handleError(e);
	}
}

if (!building || browser) {
	fetchOpenIDConfiguration();
}
