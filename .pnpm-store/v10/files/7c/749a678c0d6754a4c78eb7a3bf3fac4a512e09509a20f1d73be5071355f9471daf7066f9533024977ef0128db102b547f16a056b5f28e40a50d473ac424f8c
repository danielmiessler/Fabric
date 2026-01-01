"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const ast_utils_1 = require("../utils/ast-utils");
const compat_1 = require("../utils/compat");
const TYPE_MESSAGES = {
    normal: 'HTML elements',
    void: 'HTML void elements',
    foreign: 'foreign (SVG or MathML) elements',
    component: 'Svelte custom components',
    svelte: 'Svelte special elements'
};
exports.default = (0, utils_1.createRule)('html-self-closing', {
    meta: {
        docs: {
            description: 'enforce self-closing style',
            category: 'Stylistic Issues',
            recommended: false,
            conflictWithPrettier: true
        },
        type: 'layout',
        fixable: 'code',
        messages: {
            requireClosing: 'Require self-closing on {{type}}.',
            disallowClosing: 'Disallow self-closing on {{type}}.'
        },
        schema: [
            {
                anyOf: [
                    {
                        properties: {
                            void: {
                                enum: ['never', 'always', 'ignore']
                            },
                            normal: {
                                enum: ['never', 'always', 'ignore']
                            },
                            foreign: {
                                enum: ['never', 'always', 'ignore']
                            },
                            component: {
                                enum: ['never', 'always', 'ignore']
                            },
                            svelte: {
                                enum: ['never', 'always', 'ignore']
                            }
                        },
                        additionalProperties: false
                    },
                    {
                        enum: ['all', 'html', 'none']
                    }
                ]
            }
        ]
    },
    create(context) {
        let options = {
            void: 'always',
            normal: 'always',
            foreign: 'always',
            component: 'always',
            svelte: 'always'
        };
        const option = context.options?.[0];
        switch (option) {
            case 'none':
                options = {
                    void: 'never',
                    normal: 'never',
                    foreign: 'never',
                    component: 'never',
                    svelte: 'never'
                };
                break;
            case 'html':
                options = {
                    void: 'always',
                    normal: 'never',
                    foreign: 'always',
                    component: 'never',
                    svelte: 'always'
                };
                break;
            default:
                if (typeof option !== 'object' || option === null)
                    break;
                options = {
                    ...options,
                    ...option
                };
                break;
        }
        /**
         * Get SvelteElement type.
         * If element is custom component "component" is returned
         * If element is svelte special element such as svelte:self "svelte" is returned
         * If element is void element "void" is returned
         * otherwise "normal" is returned
         */
        function getElementType(node) {
            if (node.kind === 'component')
                return 'component';
            if (node.kind === 'special')
                return 'svelte';
            if ((0, ast_utils_1.isVoidHtmlElement)(node))
                return 'void';
            if ((0, ast_utils_1.isForeignElement)(node))
                return 'foreign';
            return 'normal';
        }
        /**
         * Returns true if element has no children, or has only whitespace text
         */
        function isElementEmpty(node) {
            if (node.children.length <= 0)
                return true;
            for (const child of node.children) {
                if (child.type !== 'SvelteText')
                    return false;
                if (!/^\s*$/.test(child.value))
                    return false;
            }
            return true;
        }
        /**
         * Report
         */
        function report(node, shouldBeClosed) {
            const elementType = getElementType(node);
            context.report({
                node,
                loc: {
                    start: (0, compat_1.getSourceCode)(context).getLocFromIndex(node.startTag.range[1] - (node.startTag.selfClosing ? 2 : 1)),
                    end: node.loc.end
                },
                messageId: shouldBeClosed ? 'requireClosing' : 'disallowClosing',
                data: {
                    type: TYPE_MESSAGES[elementType]
                },
                *fix(fixer) {
                    if (shouldBeClosed) {
                        for (const child of node.children) {
                            yield fixer.removeRange(child.range);
                        }
                        yield fixer.insertTextBeforeRange([node.startTag.range[1] - 1, node.startTag.range[1]], '/');
                        if (node.endTag)
                            yield fixer.removeRange(node.endTag.range);
                    }
                    else {
                        yield fixer.removeRange([node.startTag.range[1] - 2, node.startTag.range[1] - 1]);
                        if (!(0, ast_utils_1.isVoidHtmlElement)(node))
                            yield fixer.insertTextAfter(node, `</${(0, ast_utils_1.getNodeName)(node)}>`);
                    }
                }
            });
        }
        return {
            SvelteElement(node) {
                if (!isElementEmpty(node))
                    return;
                const elementType = getElementType(node);
                const elementTypeOptions = options[elementType];
                if (elementTypeOptions === 'ignore')
                    return;
                const shouldBeClosed = elementTypeOptions === 'always';
                if (shouldBeClosed && !node.startTag.selfClosing) {
                    report(node, true);
                }
                else if (!shouldBeClosed && node.startTag.selfClosing) {
                    report(node, false);
                }
            }
        };
    }
});
