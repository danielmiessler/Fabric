"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.normalizeParserOptions = normalizeParserOptions;
exports.isTypeScript = isTypeScript;
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const parser_object_1 = require("./parser-object");
const resolve_parser_1 = require("./resolve-parser");
/** Normalize parserOptions */
function normalizeParserOptions(options) {
    const parserOptions = Object.assign({ ecmaVersion: 2020, sourceType: "module", loc: true, range: true, raw: true, tokens: true, comment: true, eslintVisitorKeys: true, eslintScopeManager: true }, (options || {}));
    parserOptions.sourceType = "module";
    if (parserOptions.ecmaVersion <= 5 || parserOptions.ecmaVersion == null) {
        parserOptions.ecmaVersion = 2015;
    }
    return parserOptions;
}
const TS_PARSER_NAMES = [
    "@typescript-eslint/parser",
    "typescript-eslint-parser-for-extra-files",
];
function isTypeScript(parserOptions, lang) {
    var _a;
    if (!lang) {
        return false;
    }
    const parserValue = (0, resolve_parser_1.getParserForLang)(lang, parserOptions === null || parserOptions === void 0 ? void 0 : parserOptions.parser);
    if (typeof parserValue !== "string") {
        return ((0, parser_object_1.maybeTSESLintParserObject)(parserValue) ||
            (0, parser_object_1.isTSESLintParserObject)(parserValue));
    }
    const parserName = parserValue;
    if (TS_PARSER_NAMES.includes(parserName)) {
        return true;
    }
    if (TS_PARSER_NAMES.some((nm) => parserName.includes(nm))) {
        let targetPath = parserName;
        while (targetPath) {
            const pkgPath = path_1.default.join(targetPath, "package.json");
            if (fs_1.default.existsSync(pkgPath)) {
                try {
                    return TS_PARSER_NAMES.includes((_a = JSON.parse(fs_1.default.readFileSync(pkgPath, "utf-8"))) === null || _a === void 0 ? void 0 : _a.name);
                }
                catch (_b) {
                    return false;
                }
            }
            const parent = path_1.default.dirname(targetPath);
            if (targetPath === parent) {
                break;
            }
            targetPath = parent;
        }
    }
    return false;
}
