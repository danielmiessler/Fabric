"use strict";
/**
 * refer: https://github.com/mysticatea/eslint-plugin-node/blob/f45c6149be7235c0f7422d1179c25726afeecd83/lib/util/get-package-json.js
 */
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.isKitPageComponent = isKitPageComponent;
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const get_package_json_1 = require("./get-package-json");
const compat_1 = require("./compat");
const isRunOnBrowser = !fs_1.default.readFileSync;
/**
 * return true if it's a SvelteKit page component.
 * @param context
 * @returns
 */
function isKitPageComponent(context) {
    // Hack: if it runs on browser, it regards as SvelteKit project.
    if (isRunOnBrowser)
        return true;
    if (!hasSvelteKit((0, compat_1.getFilename)(context)))
        return false;
    const routes = (context.settings?.svelte?.kit?.files?.routes ??
        (0, compat_1.getSourceCode)(context).parserServices.svelteParseContext?.svelteConfig?.kit?.files?.routes)?.replace(/^\//, '') ?? 'src/routes';
    const filePath = (0, compat_1.getFilename)(context);
    const projectRootDir = getProjectRootDir((0, compat_1.getFilename)(context)) ?? '';
    const fileName = path_1.default.basename(filePath);
    return (filePath.startsWith(path_1.default.join(projectRootDir, routes)) &&
        // MEMO: check only `+` and file extension for maintainability
        Boolean(/^\+.+\.svelte$/.test(fileName)));
}
/**
 * Check givin file is under SvelteKit project.
 *
 * If it runs on browser, it always returns true.
 *
 * @param filePath A file path.
 * @returns
 */
function hasSvelteKit(filePath) {
    // Hack: if it runs on browser, it regards as SvelteKit project.
    if (isRunOnBrowser)
        return true;
    try {
        const packageJson = (0, get_package_json_1.getPackageJson)(filePath);
        if (!packageJson)
            return false;
        if (packageJson.name === 'eslint-plugin-svelte')
            // Hack: CI removes `@sveltejs/kit` and it returns false and test failed.
            // So always it returns true if it runs on the package.
            return true;
        return Boolean(packageJson.dependencies?.['@sveltejs/kit'] ?? packageJson.devDependencies?.['@sveltejs/kit']);
    }
    catch (_e) {
        return false;
    }
}
/**
 * Gets a  project root folder path.
 * @param filePath A file path to lookup.
 * @returns A found project root folder path or null.
 */
function getProjectRootDir(filePath) {
    if (isRunOnBrowser)
        return null;
    const packageJsonFilePath = (0, get_package_json_1.getPackageJson)(filePath)?.filePath;
    if (!packageJsonFilePath)
        return null;
    return path_1.default.dirname(path_1.default.resolve(packageJsonFilePath));
}
