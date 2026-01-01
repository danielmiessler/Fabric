"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const css_utils_1 = require("../utils/css-utils");
exports.default = (0, utils_1.createRule)('no-shorthand-style-property-overrides', {
    meta: {
        docs: {
            description: 'disallow shorthand style properties that override related longhand properties',
            category: 'Possible Errors',
            recommended: true
        },
        schema: [],
        messages: {
            unexpected: "Unexpected shorthand '{{shorthand}}' after '{{original}}'."
        },
        type: 'problem'
    },
    create(context) {
        return {
            SvelteStartTag(node) {
                const beforeDeclarations = new Set();
                for (const { decls } of iterateStyleDeclSetFromAttrs(node.attributes)) {
                    for (const decl of decls) {
                        const normalized = (0, css_utils_1.stripVendorPrefix)(decl.prop);
                        const prefix = (0, css_utils_1.getVendorPrefix)(decl.prop);
                        const longhandProps = css_utils_1.SHORTHAND_PROPERTIES.get(normalized);
                        if (!longhandProps) {
                            continue;
                        }
                        for (const longhandProp of longhandProps) {
                            const longhandPropWithPrefix = prefix + longhandProp;
                            if (!beforeDeclarations.has(longhandPropWithPrefix)) {
                                continue;
                            }
                            context.report({
                                node,
                                loc: decl.loc,
                                messageId: 'unexpected',
                                data: {
                                    shorthand: decl.prop,
                                    original: longhandPropWithPrefix
                                }
                            });
                        }
                    }
                    for (const decl of decls) {
                        beforeDeclarations.add(decl.prop);
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
