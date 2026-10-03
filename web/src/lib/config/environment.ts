/**
 * Environment configuration for the Fabric web app.
 * All environment variable reads are in this file.
 */

const DEFAULT_FABRIC_BASE_URL = 'http://localhost:8080';

/**
 * Get the Fabric base URL from environment variable or default
 * This function works in both server and client contexts
 */
export function getFabricBaseUrl(): string {
  if (typeof process !== 'undefined' && process.env) {
    return process.env.FABRIC_BASE_URL || DEFAULT_FABRIC_BASE_URL;
  }

  // vite.config.ts defines __FABRIC_CONFIG__ for the client build.
  if (typeof window !== 'undefined' && (window as any).__FABRIC_CONFIG__) {
    return (window as any).__FABRIC_CONFIG__.FABRIC_BASE_URL || DEFAULT_FABRIC_BASE_URL;
  }

  return DEFAULT_FABRIC_BASE_URL;
}

/**
 * Get the Fabric API base URL (adds /api if not present)
 */
export function getFabricApiUrl(): string {
  const baseUrl = getFabricBaseUrl();

  const cleanBaseUrl = baseUrl.replace(/\/$/, '');

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

export interface FabricConfig {
  FABRIC_BASE_URL: string;
}

declare global {
  interface Window {
    __FABRIC_CONFIG__?: FabricConfig;
  }
}
