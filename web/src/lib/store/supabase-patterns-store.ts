import { writable } from 'svelte/store';

export interface SupabasePattern {
	id: string;
	name: string;
	description?: string | null;
	body: string;
	tags: string[];
	is_system: boolean;
	created_at?: string;
	updated_at?: string;
}

export interface SupabasePatternPayload {
	name: string;
	description?: string | null;
	body: string;
	tags?: string[];
	is_system?: boolean;
}

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL ?? '';

const withBase = (path: string) => (API_BASE_URL ? `${API_BASE_URL}${path}` : path);

export const supabasePatterns = writable<SupabasePattern[]>([]);
export const supabasePatternsLoading = writable<boolean>(false);
export const supabasePatternsError = writable<string | null>(null);

export async function loadSupabasePatterns(fetchFn: typeof fetch = fetch): Promise<void> {
	supabasePatternsLoading.set(true);
	supabasePatternsError.set(null);
	try {
		const response = await fetchFn(withBase('/supabase/patterns'));
		if (!response.ok) {
			throw new Error(`Failed to load patterns: ${response.statusText}`);
		}
		const data: SupabasePattern[] = await response.json();
		supabasePatterns.set(data ?? []);
	} catch (err) {
		supabasePatternsError.set(err instanceof Error ? err.message : 'Unknown error');
		supabasePatterns.set([]);
	} finally {
		supabasePatternsLoading.set(false);
	}
}

export async function createSupabasePattern(payload: SupabasePatternPayload, fetchFn: typeof fetch = fetch): Promise<SupabasePattern | null> {
	const response = await fetchFn(withBase('/supabase/patterns'), {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify(payload)
	});
	if (!response.ok) {
		throw new Error(`Failed to create pattern: ${response.statusText}`);
	}
	const created: SupabasePattern = await response.json();
	supabasePatterns.update((patterns) => [created, ...patterns]);
	return created;
}

export async function updateSupabasePattern(id: string, payload: SupabasePatternPayload, fetchFn: typeof fetch = fetch): Promise<SupabasePattern | null> {
	const response = await fetchFn(withBase(`/supabase/patterns/${id}`), {
		method: 'PUT',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify(payload)
	});
	if (response.status === 404) {
		return null;
	}
	if (!response.ok) {
		throw new Error(`Failed to update pattern: ${response.statusText}`);
	}
	const updated: SupabasePattern = await response.json();
	supabasePatterns.update((patterns) => patterns.map((p) => (p.id === updated.id ? updated : p)));
	return updated;
}

export async function deleteSupabasePattern(id: string, fetchFn: typeof fetch = fetch): Promise<void> {
	const response = await fetchFn(withBase(`/supabase/patterns/${id}`), {
		method: 'DELETE'
	});
	if (response.status !== 204 && !response.ok) {
		throw new Error(`Failed to delete pattern: ${response.statusText}`);
	}
	supabasePatterns.update((patterns) => patterns.filter((pattern) => pattern.id !== id));
}
