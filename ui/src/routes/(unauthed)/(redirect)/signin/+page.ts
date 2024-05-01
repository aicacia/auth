import type { PageLoad } from './$types';

export const load: PageLoad = async (event) => {
	const { openIDConfiguration } = await event.parent();

	return {
		openIDConfiguration
	};
};
