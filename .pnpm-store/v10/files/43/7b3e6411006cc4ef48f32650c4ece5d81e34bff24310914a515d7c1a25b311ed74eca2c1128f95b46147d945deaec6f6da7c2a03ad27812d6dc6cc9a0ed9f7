"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.parseScriptInSvelte = parseScriptInSvelte;
exports.parseScript = parseScript;
exports.parseScriptWithoutAnalyzeScope = parseScriptWithoutAnalyzeScope;
const analyze_scope_1 = require("./analyze-scope");
const traverse_1 = require("../traverse");
const resolve_parser_1 = require("./resolve-parser");
const parser_object_1 = require("./parser-object");
/**
 * Parse for <script>
 */
function parseScriptInSvelte(code, attrs, parserOptions) {
    const result = parseScript(code, attrs, parserOptions);
    (0, traverse_1.traverseNodes)(result.ast, {
        visitorKeys: result.visitorKeys,
        enterNode(node, parent) {
            node.parent = parent;
            if (node.type === "LabeledStatement" && node.label.name === "$") {
                if ((parent === null || parent === void 0 ? void 0 : parent.type) === "Program") {
                    // Transform node type
                    node.type = "SvelteReactiveStatement";
                }
            }
        },
        leaveNode() {
            //
        },
    });
    return result;
}
/**
 * Parse for script
 */
function parseScript(code, attrs, parserOptions) {
    const result = parseScriptWithoutAnalyzeScopeFromVCode(code, attrs, parserOptions);
    if (!result.scopeManager) {
        const scopeManager = (0, analyze_scope_1.analyzeScope)(result.ast, parserOptions);
        result.scopeManager = scopeManager;
    }
    return result;
}
/**
 * Parse for script without analyze scope
 */
function parseScriptWithoutAnalyzeScope(code, attrs, options) {
    const parser = (0, resolve_parser_1.getParser)(attrs, options.parser);
    const result = (0, parser_object_1.isEnhancedParserObject)(parser)
        ? parser.parseForESLint(code, options)
        : parser.parse(code, options);
    if ("ast" in result && result.ast != null) {
        return result;
    }
    return { ast: result };
}
/**
 * Parse for script without analyze scope
 */
function parseScriptWithoutAnalyzeScopeFromVCode(code, attrs, options) {
    const result = parseScriptWithoutAnalyzeScope(code, attrs, options);
    result._virtualScriptCode = code;
    return result;
}
