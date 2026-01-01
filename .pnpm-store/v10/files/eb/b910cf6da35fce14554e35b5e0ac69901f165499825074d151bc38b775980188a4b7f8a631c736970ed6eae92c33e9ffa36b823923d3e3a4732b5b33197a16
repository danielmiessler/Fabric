"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.prepareContent = void 0;
const strip_indent_1 = __importDefault(require("strip-indent"));
// todo: could use magig-string and generate some sourcemaps 🗺
function prepareContent({ options, content, }) {
    if (typeof options !== 'object') {
        return content;
    }
    if (options.stripIndent) {
        content = (0, strip_indent_1.default)(content);
    }
    if (options.prependData) {
        content = `${options.prependData}\n${content}`;
    }
    return content;
}
exports.prepareContent = prepareContent;
