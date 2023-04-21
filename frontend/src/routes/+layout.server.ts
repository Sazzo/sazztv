import type { LayoutServerLoad } from './$types';

export const load = (({ locals }) => {
	return {
		accessToken: locals.accessToken,
		currentUser: locals.currentUser
	};
}) satisfies LayoutServerLoad;
