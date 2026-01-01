"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.JAVASCRIPT_RESERVED_KEYWORD_SET = exports.setProp = exports.findUp = exports.isValidLocalPath = exports.hasDepInstalled = exports.getIncludePaths = exports.concat = void 0;
const fs_1 = require("fs");
const path_1 = require("path");
function concat(...arrs) {
    return arrs.reduce((acc, a) => {
        if (a)
            return acc.concat(a);
        return acc;
    }, []);
}
exports.concat = concat;
/** Paths used by preprocessors to resolve @imports */
function getIncludePaths(fromFilename, base = []) {
    if (fromFilename == null)
        return [];
    return [
        ...new Set([...base, 'node_modules', process.cwd(), (0, path_1.dirname)(fromFilename)]),
    ];
}
exports.getIncludePaths = getIncludePaths;
const depCheckCache = {};
/**
 * Checks if a package is installed.
 *
 * @export
 * @param {string} dep
 * @returns boolean
 */
async function hasDepInstalled(dep) {
    if (depCheckCache[dep] != null) {
        return depCheckCache[dep];
    }
    let result = false;
    try {
        await Promise.resolve(`${dep}`).then(s => __importStar(require(s)));
        result = true;
    }
    catch (e) {
        result = false;
    }
    return (depCheckCache[dep] = result);
}
exports.hasDepInstalled = hasDepInstalled;
function isValidLocalPath(path) {
    return path.startsWith('.');
}
exports.isValidLocalPath = isValidLocalPath;
// finds a existing path up the tree
function findUp({ what, from }) {
    const { root, dir } = (0, path_1.parse)(from);
    let cur = dir;
    try {
        while (cur !== root) {
            const possiblePath = (0, path_1.join)(cur, what);
            if ((0, fs_1.existsSync)(possiblePath)) {
                return possiblePath;
            }
            cur = (0, path_1.dirname)(cur);
        }
    }
    catch (e) {
        console.error(e);
    }
    return null;
}
exports.findUp = findUp;
// set deep property in object
function setProp(obj, keyList, value) {
    let i = 0;
    for (; i < keyList.length - 1; i++) {
        const key = keyList[i];
        if (typeof obj[key] !== 'object') {
            obj[key] = {};
        }
        obj = obj[key];
    }
    obj[keyList[i]] = value;
}
exports.setProp = setProp;
exports.JAVASCRIPT_RESERVED_KEYWORD_SET = new Set([
    'arguments',
    'await',
    'break',
    'case',
    'catch',
    'class',
    'const',
    'continue',
    'debugger',
    'default',
    'delete',
    'do',
    'else',
    'enum',
    'eval',
    'export',
    'extends',
    'false',
    'finally',
    'for',
    'function',
    'if',
    'implements',
    'import',
    'in',
    'instanceof',
    'interface',
    'let',
    'new',
    'null',
    'package',
    'private',
    'protected',
    'public',
    'return',
    'static',
    'super',
    'switch',
    'this',
    'throw',
    'true',
    'try',
    'typeof',
    'var',
    'void',
    'while',
    'with',
    'yield',
]);
