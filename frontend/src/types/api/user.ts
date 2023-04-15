import type { Stream } from './stream';

export interface User {
	id: string;
	username: string;
	last_stream_at: string;
	stream: Stream | null;
}
