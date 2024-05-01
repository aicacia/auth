import { applicationApi } from '$lib/openapi';
import type { PageLoad } from './$types';

export const load: PageLoad = async (event) => {
	const { user } = await event.parent();
	const pagination = await applicationApi.applications();
	return { user, pagination };
};
