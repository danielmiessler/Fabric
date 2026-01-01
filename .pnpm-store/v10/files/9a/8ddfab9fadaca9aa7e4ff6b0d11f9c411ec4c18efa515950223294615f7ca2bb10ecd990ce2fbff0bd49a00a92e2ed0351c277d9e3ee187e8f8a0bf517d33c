"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.resolveSvelteConfigFromOption = resolveSvelteConfigFromOption;
const path_1 = __importDefault(require("path"));
const fs_1 = __importDefault(require("fs"));
const parser_1 = require("./parser");
const caches = new Map();
/**
 * Resolves svelte.config.
 */
function resolveSvelteConfigFromOption(options) {
    if (options === null || options === void 0 ? void 0 : options.svelteConfig) {
        return options.svelteConfig;
    }
    return resolveSvelteConfig(options === null || options === void 0 ? void 0 : options.filePath);
}
/**
 * Resolves `svelte.config.js`.
 * It searches the parent directories of the given file to find `svelte.config.js`,
 * and returns the static analysis result for it.
 */
function resolveSvelteConfig(filePath) {
    let cwd = filePath && fs_1.default.existsSync(filePath) ? path_1.default.dirname(filePath) : null;
    if (cwd == null) {
        if (typeof process === "undefined")
            return null;
        cwd = process.cwd();
    }
    const configFilePath = findConfigFilePath(cwd);
    if (!configFilePath)
        return null;
    if (caches.has(configFilePath)) {
        return caches.get(configFilePath) || null;
    }
    const code = fs_1.default.readFileSync(configFilePath, "utf8");
    const config = (0, parser_1.parseConfig)(code);
    caches.set(configFilePath, config);
    return config;
}
/**
 * Searches from the current working directory up until finding the config filename.
 * @param {string} cwd The current working directory to search from.
 * @returns {string|undefined} The file if found or `undefined` if not.
 */
function findConfigFilePath(cwd) {
    let directory = path_1.default.resolve(cwd);
    const { root } = path_1.default.parse(directory);
    const stopAt = path_1.default.resolve(directory, root);
    while (directory !== stopAt) {
        const target = path_1.default.resolve(directory, "svelte.config.js");
        const stat = fs_1.default.existsSync(target)
            ? fs_1.default.statSync(target, {
                throwIfNoEntry: false,
            })
            : null;
        if (stat === null || stat === void 0 ? void 0 : stat.isFile()) {
            return target;
        }
        const next = path_1.default.dirname(directory);
        if (next === directory)
            break;
        directory = next;
    }
    return null;
}
