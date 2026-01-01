"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.getParserForLang = getParserForLang;
exports.getParser = getParser;
const espree_1 = require("./espree");
const parser_object_1 = require("./parser-object");
/** Get parser for script lang */
function getParserForLang(lang, parser) {
    if (parser) {
        if (typeof parser === "string" || (0, parser_object_1.isParserObject)(parser)) {
            return parser;
        }
        if (typeof parser === "object") {
            const value = parser[lang || "js"];
            if (typeof value === "string" || (0, parser_object_1.isParserObject)(value)) {
                return value;
            }
        }
    }
    return "espree";
}
/** Get parser */
function getParser(attrs, parser) {
    const parserValue = getParserForLang(attrs.lang, parser);
    if ((0, parser_object_1.isParserObject)(parserValue)) {
        return parserValue;
    }
    if (parserValue !== "espree") {
        // eslint-disable-next-line @typescript-eslint/no-require-imports -- ignore
        return require(parserValue);
    }
    return (0, espree_1.getEspree)();
}
