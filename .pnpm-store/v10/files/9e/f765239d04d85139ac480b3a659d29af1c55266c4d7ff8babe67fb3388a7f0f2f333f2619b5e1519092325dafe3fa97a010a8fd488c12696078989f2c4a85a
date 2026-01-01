/**
 * Simple cache manager.
 *
 * refer: https://github.com/mysticatea/eslint-plugin-node/blob/f45c6149be7235c0f7422d1179c25726afeecd83/lib/util/cache.js
 */
/**
 * The cache will dispose of each value if the value has not been accessed
 * during 5 seconds.
 * @returns getter and setter ofr the cache.
 */
export declare function createCache<T>(): {
    get: (key: string) => T | null;
    set: (key: string, value: T) => void;
};
