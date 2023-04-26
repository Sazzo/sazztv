import { PUBLIC_API_URL } from '$env/static/public';
import { FetchResultTypes, QueryError, fetch } from '@sapphire/fetch';
import type { Handle } from '@sveltejs/kit';
import type { User } from './types/api/user';

export const handle: Handle = async ({ event, resolve }) => {
	const accessToken = event.cookies.get('accessToken');

	if (accessToken) {
		event.locals.accessToken = accessToken;

		try {
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
		} catch (error) {
			if (error instanceof QueryError && error.code === 401) {
				event.locals.accessToken = undefined;
				event.cookies.delete('accessToken');

				return resolve(event);
			}

			throw error;
		}
	}

	return await resolve(event);
};
