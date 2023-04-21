// See https://kit.svelte.dev/docs/types#app

import type { User } from './types/api/user';

// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		interface Locals {
			accessToken?: string;
			currentUser?: User;
		}
		// interface PageData {}
		// interface Platform {}
	}
}

export {};
