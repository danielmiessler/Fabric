"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const ast_utils_1 = require("../utils/ast-utils");
const compat_1 = require("../utils/compat");
exports.default = (0, utils_1.createRule)('no-dupe-use-directives', {
    meta: {
        docs: {
            description: 'disallow duplicate `use:` directives',
            category: 'Possible Errors',
            recommended: false
        },
        schema: [],
        messages: {
            duplication: 'This `{{keyText}}` directive is the same and duplicate directives in L{{lineNo}}.'
        },
        type: 'problem'
    },
    create(context) {
        const sourceCode = (0, compat_1.getSourceCode)(context);
        const directiveDataMap = new Map();
        return {
            SvelteDirective(node) {
                if (node.kind !== 'Action')
                    return;
                const keyText = (0, ast_utils_1.getAttributeKeyText)(node, context);
                const directiveDataList = directiveDataMap.get(keyText);
                if (!directiveDataList) {
                    directiveDataMap.set(keyText, [
                        {
                            expression: node.expression,
                            nodes: [node]
                        }
                    ]);
                    return;
                }
                const directiveData = directiveDataList.find((data) => {
                    if (!data.expression || !node.expression) {
                        return data.expression === node.expression;
                    }
                    return (0, ast_utils_1.equalTokens)(data.expression, node.expression, sourceCode);
                });
                if (!directiveData) {
                    directiveDataList.push({
                        expression: node.expression,
                        nodes: [node]
                    });
                    return;
                }
                directiveData.nodes.push(node);
            },
            'SvelteStartTag:exit'() {
                for (const [keyText, directiveDataList] of directiveDataMap) {
                    for (const { nodes } of directiveDataList) {
                        if (nodes.length < 2) {
                            continue;
                        }
                        for (const node of nodes) {
                            context.report({
                                node,
                                messageId: 'duplication',
                                data: {
                                    keyText,
                                    lineNo: String((nodes[0] !== node ? nodes[0] : nodes[1]).loc.start.line)
                                }
                            });
                        }
                    }
                }
                directiveDataMap.clear();
            }
        };
    }
});
