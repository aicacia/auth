import type { PageLoad } from './$types';

export const load: PageLoad = async (event) => {
	const { user } = await event.parent();
	return {
		user
	};
};
