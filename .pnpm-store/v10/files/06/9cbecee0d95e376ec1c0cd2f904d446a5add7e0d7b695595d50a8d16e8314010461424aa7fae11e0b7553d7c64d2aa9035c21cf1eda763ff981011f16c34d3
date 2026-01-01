"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.isParserObject = isParserObject;
exports.isEnhancedParserObject = isEnhancedParserObject;
exports.isBasicParserObject = isBasicParserObject;
exports.maybeTSESLintParserObject = maybeTSESLintParserObject;
exports.isTSESLintParserObject = isTSESLintParserObject;
/** Checks whether given object is ParserObject */
function isParserObject(value) {
    return isEnhancedParserObject(value) || isBasicParserObject(value);
}
/** Checks whether given object is EnhancedParserObject */
function isEnhancedParserObject(value) {
    return Boolean(value && typeof value.parseForESLint === "function");
}
/** Checks whether given object is BasicParserObject */
function isBasicParserObject(value) {
    return Boolean(value && typeof value.parse === "function");
}
/** Checks whether given object maybe "@typescript-eslint/parser" */
function maybeTSESLintParserObject(value) {
    return (isEnhancedParserObject(value) &&
        isBasicParserObject(value) &&
        typeof value.createProgram === "function" &&
        typeof value.clearCaches === "function" &&
        typeof value.version === "string");
}
/** Checks whether given object is "@typescript-eslint/parser" */
function isTSESLintParserObject(value) {
    if (!isEnhancedParserObject(value))
        return false;
    try {
        const result = value.parseForESLint("", {});
        const services = result.services;
        return Boolean(services &&
            services.esTreeNodeToTSNodeMap &&
            services.tsNodeToESTreeNodeMap);
    }
    catch (_a) {
        return false;
    }
}
