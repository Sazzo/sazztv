import type { StreamSettings } from './stream';

export interface User {
	id: string;
	username: string;
	stream_settings: StreamSettings;
	is_admin: boolean;
	is_live: boolean;
	last_stream_at: string;
	created_at: string;
	updated_at: string;
}
