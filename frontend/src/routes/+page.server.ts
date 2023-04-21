import { FetchResultTypes, QueryError, fetch } from '@sapphire/fetch';
import type { Actions } from './$types';
import { PUBLIC_API_URL } from '$env/static/public';
import type { AuthLoginResponse } from '../types/api/auth';
import { fail, redirect } from '@sveltejs/kit';

export const actions = {
	default: async ({ cookies, request }) => {
		try {
			const requestData = await request.formData();

			const username = requestData.get('username');
			const password = requestData.get('password');

			const loginResponse = await fetch<AuthLoginResponse>(
				`${PUBLIC_API_URL}/auth/login`,
				{
					method: 'POST',
					headers: {
						'Content-Type': 'application/json'
					},
					body: JSON.stringify({ username, password })
				},
				FetchResultTypes.JSON
			);

			cookies.set('accessToken', loginResponse.accessToken, {
				httpOnly: true,
				path: '/',
				sameSite: 'strict'
			});

			throw redirect(302, '/');
		} catch (error) {
			if (error instanceof QueryError) {
				return fail(400, JSON.parse(error.body));
			}

			throw error;
		}
	}
} satisfies Actions;
