import { writable } from 'svelte/store';

export interface SupabaseSessionSummary {
	id: string;
	title: string;
	description?: string | null;
	updated_at?: string;
}

export const supabaseSessions = writable<SupabaseSessionSummary[]>([]);
export const supabaseSessionsLoading = writable<boolean>(false);
export const supabaseSessionsError = writable<string | null>(null);

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL ?? '';

export async function loadSupabaseSessions(fetchFn: typeof fetch = fetch): Promise<void> {
	supabaseSessionsLoading.set(true);
	supabaseSessionsError.set(null);
	try {
		const url = API_BASE_URL ? `${API_BASE_URL}/supabase/sessions` : '/supabase/sessions';
		const response = await fetchFn(url);
		if (response.status === 404) {
			supabaseSessions.set([]);
			return;
		}
		if (!response.ok) {
			throw new Error(`Failed to load sessions: ${response.statusText}`);
		}
		const data: SupabaseSessionSummary[] = await response.json();
		supabaseSessions.set(data ?? []);
	} catch (err) {
		supabaseSessionsError.set(err instanceof Error ? err.message : 'Unknown error');
		supabaseSessions.set([]);
	} finally {
		supabaseSessionsLoading.set(false);
	}
}
