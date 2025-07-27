/**
 * Environment configuration for the Fabric web app
 * Centralizes all environment variable handling
 */

// Default values
const DEFAULT_FABRIC_BASE_URL = 'http://localhost:8080';

/**
 * Get the Fabric base URL from environment variable or default
 * This function works in both server and client contexts
 */
export function getFabricBaseUrl(): string {
  // In server context (Node.js), use process.env
  if (typeof process !== 'undefined' && process.env) {
    return process.env.FABRIC_BASE_URL || DEFAULT_FABRIC_BASE_URL;
  }

  // In client context, check if the environment was injected via Vite
  if (typeof window !== 'undefined' && (window as any).__FABRIC_CONFIG__) {
    return (window as any).__FABRIC_CONFIG__.FABRIC_BASE_URL || DEFAULT_FABRIC_BASE_URL;
  }

  // Fallback to default
  return DEFAULT_FABRIC_BASE_URL;
}

/**
 * Get the Fabric API base URL (adds /api if not present)
 */
export function getFabricApiUrl(): string {
  const baseUrl = getFabricBaseUrl();

  // Remove trailing slash if present
  const cleanBaseUrl = baseUrl.replace(/\/$/, '');

  // Check if it already ends with /api
  if (cleanBaseUrl.endsWith('/api')) {
    return cleanBaseUrl;
  }

  return `${cleanBaseUrl}/api`;
}

/**
 * Configuration object for easy access to all environment settings
 */
export const config = {
  fabricBaseUrl: getFabricBaseUrl(),
  fabricApiUrl: getFabricApiUrl(),
} as const;

// Type definitions
export interface FabricConfig {
  FABRIC_BASE_URL: string;
}

declare global {
  interface Window {
    __FABRIC_CONFIG__?: FabricConfig;
  }
}
