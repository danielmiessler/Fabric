"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const safe_parser_1 = __importDefault(require("postcss-safe-parser/lib/safe-parser"));
const template_tokenize_1 = __importDefault(require("./template-tokenize"));
class TemplateSafeParser extends safe_parser_1.default {
    createTokenizer() {
        this.tokenizer = (0, template_tokenize_1.default)(this.input, { ignoreErrors: true });
    }
}
exports.default = TemplateSafeParser;
