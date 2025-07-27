import { writable } from 'svelte/store';

/**
 * List of available strategies fetched from backend.
 * Each strategy has a name and description.
 */
export const strategies = writable<Array<{ name: string; description: string }>>([]);

/**
 * Currently selected strategy name.
 * Default is empty string meaning "None".
 */
export const selectedStrategy = writable<string>("");

/**
 * Fetches available strategies from the backend `/strategies` endpoint.
 * Populates the `strategies` store.
 */
export async function fetchStrategies() {
  try {
    const response = await fetch('/strategies/strategies.json');
    if (!response.ok) {
      console.error('Failed to fetch strategies:', response.statusText);
      return;
    }
    const data = await response.json();
    // Expecting an array of { name, description }
    strategies.set(data);
  } catch (error) {
    console.error('Error fetching strategies:', error);
  }
}
