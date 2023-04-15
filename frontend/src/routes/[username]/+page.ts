import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';
import type { User } from '../../types/api/user';

async function getUser(fetchData: typeof fetch, username: string): Promise<User> {
	// todo: hardcoded api url
	const res = await fetchData(`http://127.0.0.1:8000/users/${username}`);
	if (res.ok) {
		return await res.json();
	}

	throw error(res.status, res.statusText);
}

export const load = (async ({ fetch, params }) => {
	return {
		user: getUser(fetch, params.username)
	};
}) satisfies PageLoad;
