import type * as Compiler from "svelte/compiler";
export type SvelteConfig = {
    compilerOptions?: Compiler.CompileOptions;
    extensions?: string[];
    kit?: KitConfig;
    preprocess?: unknown;
    vitePlugin?: unknown;
    onwarn?: (warning: Compiler.Warning, defaultHandler: (warning: Compiler.Warning) => void) => void;
    [key: string]: unknown;
};
interface KitConfig {
    adapter?: unknown;
    alias?: Record<string, string>;
    appDir?: string;
    csp?: {
        mode?: "hash" | "nonce" | "auto";
        directives?: unknown;
        reportOnly?: unknown;
    };
    csrf?: {
        checkOrigin?: boolean;
    };
    embedded?: boolean;
    env?: {
        dir?: string;
        publicPrefix?: string;
        privatePrefix?: string;
    };
    files?: {
        assets?: string;
        hooks?: {
            client?: string;
            server?: string;
            universal?: string;
        };
        lib?: string;
        params?: string;
        routes?: string;
        serviceWorker?: string;
        appTemplate?: string;
        errorTemplate?: string;
    };
    inlineStyleThreshold?: number;
    moduleExtensions?: string[];
    outDir?: string;
    output?: {
        preloadStrategy?: "modulepreload" | "preload-js" | "preload-mjs";
    };
    paths?: {
        assets?: "" | `http://${string}` | `https://${string}`;
        base?: "" | `/${string}`;
        relative?: boolean;
    };
    prerender?: {
        concurrency?: number;
        crawl?: boolean;
        entries?: ("*" | `/${string}`)[];
        handleHttpError?: unknown;
        handleMissingId?: unknown;
        handleEntryGeneratorMismatch?: unknown;
        origin?: string;
    };
    serviceWorker?: {
        register?: boolean;
        files?(filepath: string): boolean;
    };
    typescript?: {
        config?: (config: Record<string, any>) => Record<string, any> | void;
    };
    version?: {
        name?: string;
        pollInterval?: number;
    };
}
/**
 * Resolves svelte.config.
 */
export declare function resolveSvelteConfigFromOption(options: any): SvelteConfig | null;
export {};
