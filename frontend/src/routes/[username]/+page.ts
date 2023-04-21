import type { PageLoad } from './$types';
import type { User } from '../../types/api/user';
import { fetch, FetchResultTypes } from '@sapphire/fetch';

import { env } from '$env/dynamic/public';

async function getUser(username: string): Promise<User> {
	return await fetch<User>(`${env.PUBLIC_API_URL}/users/${username}`, FetchResultTypes.JSON);
}

export const load = (async ({ params }) => {
	return {
		user: getUser(params.username)
	};
}) satisfies PageLoad;
