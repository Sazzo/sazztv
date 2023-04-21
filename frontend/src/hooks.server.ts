import { PUBLIC_API_URL } from '$env/static/public';
import { FetchResultTypes, fetch } from '@sapphire/fetch';
import type { Handle } from '@sveltejs/kit';
import type { User } from './types/api/user';

export const handle: Handle = async ({ event, resolve }) => {
	const accessToken = event.cookies.get('accessToken');

	if (accessToken) {
		event.locals.accessToken = accessToken;

		const currentUser = await fetch<User>(
			`${PUBLIC_API_URL}/users/@me`,
			{
				headers: {
					Authorization: `Bearer ${accessToken}`
				}
			},
			FetchResultTypes.JSON
		);

		event.locals.currentUser = currentUser;
	}

	return await resolve(event);
};
