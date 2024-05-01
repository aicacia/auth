import type { LayoutLoad } from './$types';
import { base } from '$app/paths';
import { tryGetCurrentUser } from '$lib/stores/user';
import { redirect } from '@sveltejs/kit';
import { createNotification } from '$lib/stores/notifications';

export const prerender = false;
export const ssr = false;

export const load: LayoutLoad = async (event) => {
	await event.parent();

	const currentUser = await tryGetCurrentUser();

	if (currentUser) {
		return {
			user: currentUser
		};
	} else {
		createNotification('Not signed in.');
		redirect(302, `${base}/signin`);
	}
};
