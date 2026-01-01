"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformer = void 0;
const postcss_1 = __importDefault(require("postcss"));
const globalifySelector_1 = require("../modules/globalifySelector");
const selectorPattern = /:global(?!\()/;
const globalifyRulePlugin = (root) => {
    root.walkRules(selectorPattern, (rule) => {
        var _a;
        const modifiedSelectors = rule.selectors
            .filter((selector) => selector !== ':global')
            .map((selector) => {
            const [beginning, ...rest] = selector.split(selectorPattern);
            if (rest.length === 0)
                return beginning;
            return [beginning, ...rest.map(globalifySelector_1.globalifySelector)]
                .map((str) => str.trim())
                .join(' ')
                .trim();
        });
        if (modifiedSelectors.length === 0) {
            if (((_a = rule.parent) === null || _a === void 0 ? void 0 : _a.type) === 'atrule' && rule.selector === ':global') {
                rule.replaceWith(...rule.nodes);
            }
            else {
                rule.remove();
            }
            return;
        }
        rule.replaceWith(rule.clone({
            selectors: modifiedSelectors,
        }));
    });
};
const globalAttrPlugin = (root) => {
    root.walkAtRules(/keyframes$/, (atrule) => {
        if (!atrule.params.startsWith('-global-')) {
            atrule.replaceWith(atrule.clone({
                params: `-global-${atrule.params}`,
            }));
        }
    });
    root.walkRules((rule) => {
        var _a, _b;
        // we use endsWith for checking @keyframes and prefixed @-{prefix}-keyframes
        if ((_b = (_a = rule === null || rule === void 0 ? void 0 : rule.parent) === null || _a === void 0 ? void 0 : _a.name) === null || _b === void 0 ? void 0 : _b.endsWith('keyframes')) {
            return;
        }
        rule.replaceWith(rule.clone({
            selectors: rule.selectors.map(globalifySelector_1.globalifySelector),
        }));
    });
};
const transformer = async ({ content, filename, options, map, attributes, }) => {
    const plugins = [
        globalifyRulePlugin,
        (attributes === null || attributes === void 0 ? void 0 : attributes.global) && globalAttrPlugin,
    ].filter(Boolean);
    const { css, map: newMap } = await (0, postcss_1.default)(plugins).process(content, {
        from: filename,
        to: filename,
        map: (options === null || options === void 0 ? void 0 : options.sourceMap) ? { prev: map } : false,
    });
    return { code: css, map: newMap };
};
exports.transformer = transformer;
