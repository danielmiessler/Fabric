/**
 * refer: https://github.com/mysticatea/eslint-plugin-node/blob/f45c6149be7235c0f7422d1179c25726afeecd83/lib/util/get-package-json.js
 */
type PackageJson = {
    name?: unknown;
    dependencies?: {
        [key in string]?: unknown;
    };
    devDependencies?: {
        [key in string]?: unknown;
    };
    filePath: string;
};
/**
 * Gets a `package.json` data.
 * The data is cached if found, then it's used after.
 * @param startPath A file path to lookup.
 * @returns A found `package.json` data or `null`.
 *      This object have additional property `filePath`.
 */
export declare function getPackageJson(startPath?: string): PackageJson | null;
export {};
