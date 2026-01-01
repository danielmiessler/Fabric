"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const css_utils_1 = require("../utils/css-utils");
exports.default = (0, utils_1.createRule)('no-dupe-style-properties', {
    meta: {
        docs: {
            description: 'disallow duplicate style properties',
            category: 'Possible Errors',
            recommended: true
        },
        schema: [],
        messages: {
            unexpected: "Duplicate property '{{name}}'."
        },
        type: 'problem'
    },
    create(context) {
        return {
            SvelteStartTag(node) {
                const reported = new Set();
                const beforeDeclarations = new Map();
                for (const { decls } of iterateStyleDeclSetFromAttrs(node.attributes)) {
                    for (const decl of decls) {
                        const already = beforeDeclarations.get(decl.prop);
                        if (already) {
                            for (const report of [already, decl].filter((n) => !reported.has(n))) {
                                context.report({
                                    node,
                                    loc: report.loc,
                                    messageId: 'unexpected',
                                    data: { name: report.prop }
                                });
                                reported.add(report);
                            }
                        }
                    }
                    for (const decl of decls) {
                        beforeDeclarations.set(decl.prop, decl);
                    }
                }
            }
        };
        /** Iterate the style decl set from attrs */
        function* iterateStyleDeclSetFromAttrs(attrs) {
            for (const attr of attrs) {
                if (attr.type === 'SvelteStyleDirective') {
                    yield {
                        decls: [{ prop: attr.key.name.name, loc: attr.key.name.loc }]
                    };
                }
                else if (attr.type === 'SvelteAttribute') {
                    if (attr.key.name !== 'style') {
                        continue;
                    }
                    const root = (0, css_utils_1.parseStyleAttributeValue)(attr, context);
                    if (!root) {
                        continue;
                    }
                    yield* iterateStyleDeclSetFromStyleRoot(root);
                }
            }
        }
        /** Iterate the style decl set from style root */
        function* iterateStyleDeclSetFromStyleRoot(root) {
            for (const child of root.nodes) {
                if (child.type === 'decl') {
                    yield {
                        decls: [
                            {
                                prop: child.prop.name,
                                get loc() {
                                    return child.prop.loc;
                                }
                            }
                        ]
                    };
                }
                else if (child.type === 'inline') {
                    const decls = [];
                    for (const root of child.getAllInlineStyles().values()) {
                        for (const set of iterateStyleDeclSetFromStyleRoot(root)) {
                            decls.push(...set.decls);
                        }
                    }
                    yield { decls };
                }
            }
        }
    }
});
