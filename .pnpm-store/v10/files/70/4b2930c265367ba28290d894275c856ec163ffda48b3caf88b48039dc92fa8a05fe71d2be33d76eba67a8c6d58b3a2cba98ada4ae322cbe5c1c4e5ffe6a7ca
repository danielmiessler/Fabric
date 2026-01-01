import type { DepOptimizationOptions, SSROptions, UserConfig } from 'vite'

export interface CrawlFrameworkPkgsOptions {
  /**
   * Path to the root of the project that contains the `package.json`
   */
  root: string
  /**
   * Whether we're currently in a Vite build
   */
  isBuild: boolean
  /**
   * Optional. If a Vite user config is passed, the output Vite config will respect the
   * set `optimizeDeps` and `ssr` options so it doesn't override it
   */
  viteUserConfig?: UserConfig
  /**
   * Whether this is a framework package by checking it's `package.json`.
   * A framework package is one that exports special files that can't be processed
   * by esbuild natively. For example, exporting `.framework` files.
   *
   * @example
   * ```ts
   * return pkgJson.keywords?.includes('my-framework')
   * ```
   */
  isFrameworkPkgByJson?: (pkgJson: Record<string, any>) => boolean
  /**
   * Whether this is a framework package by checking it's name. This is
   * usually used as a fast path. Return `true` or `false` if you know 100%
   * if it's a framework package or not. Return `undefined` to fallback to
   * `isFrameworkPkgByJson`.
   *
   * @example
   * ```ts
   * return SPECIAL_PACKAGES.includes(pkgName) || undefined
   * ```
   */
  isFrameworkPkgByName?: (pkgName: string) => boolean | undefined
  /**
   * Whether this is a semi-framework package by checking it's `package.json`.
   * A semi-framework package is one that **doesn't** export special files but
   * consumes other APIs of the framework. For example, it only does
   * `import { debounce } from 'my-framework/utils'`.
   *
   * @example
   * ```ts
   * return Object.keys(pkgJson.dependencies || {}).includes('my-framework')
   * ```
   */
  isSemiFrameworkPkgByJson?: (pkgJson: Record<string, any>) => boolean
  /**
   * Whether this is a semi-framework package by checking it's name. This is
   * usually used as a fast path. Return `true` or `false` if you know 100%
   * if it's a semi-framework package or not. Return `undefined` to fallback to
   * `isSemiFrameworkPkgByJson`.
   *
   * @example
   * ```ts
   * return SPECIAL_SEMI_PACKAGES.includes(pkgName) || undefined
   * ```
   */
  isSemiFrameworkPkgByName?: (pkgName: string) => boolean | undefined
}

export interface CrawlFrameworkPkgsResult {
  optimizeDeps: {
    include: string[]
    exclude: string[]
  }
  ssr: {
    noExternal: string[]
    external: string[]
  }
}

/**
 * Crawls for framework packages starting from `<root>/package.json` to build
 * out a partial Vite config. See the source code for details of how this is built.
 */
export declare function crawlFrameworkPkgs(
  options: CrawlFrameworkPkgsOptions
): Promise<CrawlFrameworkPkgsResult>

/**
 * Find the `package.json` of a dep, starting from the parent, e.g. `process.cwd()`.
 * A simplified implementation of https://nodejs.org/api/esm.html#resolver-algorithm-specification
 * (PACKAGE_RESOLVE) for `package.json` specifically.
 */
export declare function findDepPkgJsonPath(
  dep: string,
  parent: string
): Promise<string | undefined>

/**
 * Find the closest `package.json` path by walking `dir` upwards.
 *
 * Pass a function to `predicate` to check if the current `package.json` is the
 * one you're looking for. For example, finding `package.json` that has the
 * `name` field only. Throwing inside the `predicate` is safe and acts the same
 *  as returning false.
 */
export declare function findClosestPkgJsonPath(
  dir: string,
  predicate?: (pkgJsonPath: string) => boolean | Promise<boolean>
): Promise<string | undefined>

/**
 * Check if a package needs to be optimized by Vite, aka if it's CJS-only
 */
export declare function pkgNeedsOptimization(
  pkgJson: Record<string, any>,
  pkgJsonPath: string
): Promise<boolean>

/**
 * Check if a dependency is part of an existing `optimizeDeps.exclude` config
 * @param dep Dependency to be included
 * @param optimizeDepsExclude Existing `optimizeDeps.exclude` config
 * @example
 * ```ts
 * optimizeDeps: {
 *   include: includesToAdd.filter((dep) => !isDepExcluded(dep, existingExclude))
 * }
 * ```
 */
export declare function isDepExcluded(
  dep: string,
  optimizeDepsExclude: NonNullable<DepOptimizationOptions['exclude']>
): boolean

/**
 * Check if a dependency is part of an existing `optimizeDeps.include` config
 * @param dep Dependency to be excluded
 * @param optimizeDepsInclude Existing `optimizeDeps.include` config
 * @example
 * ```ts
 * optimizeDeps: {
 *   exclude: excludesToAdd.filter((dep) => !isDepIncluded(dep, existingInclude))
 * }
 * ```
 */
export declare function isDepIncluded(
  dep: string,
  optimizeDepsInclude: NonNullable<DepOptimizationOptions['include']>
): boolean

/**
 * Check if a dependency is part of an existing `ssr.noExternal` config
 * @param dep Dependency to be excluded
 * @param ssrNoExternal Existing `ssr.noExternal` config
 * @example
 * ```ts
 * ssr: {
 *   external: externalsToAdd.filter((dep) => !isDepNoExternal(dep, existingNoExternal))
 * }
 * ```
 */
export declare function isDepNoExternaled(
  dep: string,
  ssrNoExternal: NonNullable<SSROptions['noExternal']>
): boolean

/**
 * Check if a dependency is part of an existing `ssr.external` config
 * @param dep Dependency to be noExternaled
 * @param ssrExternal Existing `ssr.external` config
 * @example
 * ```ts
 * ssr: {
 *   noExternal: noExternalsToAdd.filter((dep) => !isDepExternal(dep, existingExternal))
 * }
 * ```
 */
export declare function isDepExternaled(
  dep: string,
  ssrExternal: NonNullable<SSROptions['external']>
): boolean
