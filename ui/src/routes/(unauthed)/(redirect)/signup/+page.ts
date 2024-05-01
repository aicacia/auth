import { redirect } from '@sveltejs/kit';
import type { PageLoad } from './$types';
import { base } from '$app/paths';
import { createNotification } from '$lib/stores/notifications';

export const load: PageLoad = async (event) => {
	const { openIDConfiguration } = await event.parent();

	if (!openIDConfiguration?.grant_types_supported.includes('password')) {
		createNotification('Sign up is not enabled.');
		redirect(302, `${base}/signin`);
	} else {
		return {
			openIDConfiguration
		};
	}
};
