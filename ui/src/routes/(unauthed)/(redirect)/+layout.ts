import type { LayoutLoad } from './$types';
import { base } from '$app/paths';
import { isSignedIn } from '$lib/stores/user';
import { redirect } from '@sveltejs/kit';
import { waitForOpenIDConfiguration } from '$lib/stores/system';

export const load: LayoutLoad = async (event) => {
	await event.parent();

	if (isSignedIn()) {
		redirect(302, `${base}/`);
	} else {
		const openIDConfiguration = await waitForOpenIDConfiguration();
		return {
			openIDConfiguration
		};
	}
};
