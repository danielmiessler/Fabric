"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.loadModule = loadModule;
exports.loadModulesForBrowser = loadModulesForBrowser;
const module_1 = __importDefault(require("module"));
const path_1 = __importDefault(require("path"));
const compat_1 = require("./compat");
const cache = new WeakMap();
const cache4b = new Map();
/**
 * Load module
 */
function loadModule(context, name) {
    const key = (0, compat_1.getSourceCode)(context).ast;
    let modules = cache.get(key);
    if (!modules) {
        modules = {};
        cache.set(key, modules);
    }
    const mod = modules[name] || cache4b.get(name);
    if (mod)
        return mod;
    try {
        // load from cwd
        const cwd = (0, compat_1.getCwd)(context);
        const relativeTo = path_1.default.join(cwd, '__placeholder__.js');
        return (modules[name] = module_1.default.createRequire(relativeTo)(name));
    }
    catch {
        // ignore
    }
    for (const relativeTo of [
        // load from lint file name
        (0, compat_1.getFilename)(context),
        // load from lint file name (physical)
        (0, compat_1.getPhysicalFilename)(context),
        // load from this plugin module
        typeof __filename !== 'undefined' ? __filename : ''
    ]) {
        if (relativeTo) {
            try {
                return (modules[name] = module_1.default.createRequire(relativeTo)(name));
            }
            catch {
                // ignore
            }
        }
    }
    return null;
}
/**  Load modules for browser */
async function loadModulesForBrowser() {
    const [sass, typescript] = await Promise.all([import('sass'), import('typescript')]);
    cache4b.set('sass', sass);
    cache4b.set('typescript', typescript);
}
