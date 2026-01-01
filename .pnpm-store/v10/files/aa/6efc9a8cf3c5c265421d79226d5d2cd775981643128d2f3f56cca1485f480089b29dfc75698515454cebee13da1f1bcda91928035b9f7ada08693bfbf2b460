"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.getEspree = getEspree;
const module_1 = __importDefault(require("module"));
const path_1 = __importDefault(require("path"));
const createRequire = 
// Added in v12.2.0
module_1.default.createRequire ||
    // Added in v10.12.0, but deprecated in v12.2.0.
    // @ts-expect-error -- old type
    module_1.default.createRequireFromPath ||
    // Polyfill - This is not executed on the tests on node@>=10.
    /* istanbul ignore next */
    ((modName) => {
        const mod = new module_1.default(modName);
        mod.filename = modName;
        mod.paths = module_1.default._nodeModulePaths(path_1.default.dirname(modName));
        mod._compile("module.exports = require;", modName);
        return mod.exports;
    });
let espreeCache = null;
/** Checks if given path is linter path */
function isLinterPath(p) {
    return (
    // ESLint 6 and above
    p.includes(`eslint${path_1.default.sep}lib${path_1.default.sep}linter${path_1.default.sep}linter.js`) ||
        // ESLint 5
        p.includes(`eslint${path_1.default.sep}lib${path_1.default.sep}linter.js`));
}
/**
 * Load `espree` from the loaded ESLint.
 * If the loaded ESLint was not found, just returns `require("espree")`.
 */
function getEspree() {
    if (!espreeCache) {
        // Lookup the loaded eslint
        const linterPath = Object.keys(require.cache || {}).find(isLinterPath);
        if (linterPath) {
            try {
                espreeCache = createRequire(linterPath)("espree");
            }
            catch (_a) {
                // ignore
            }
        }
        if (!espreeCache) {
            // eslint-disable-next-line @typescript-eslint/no-require-imports -- ignore
            espreeCache = require("espree");
        }
    }
    return espreeCache;
}
