"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const ast_utils_1 = require("../utils/ast-utils");
const compat_1 = require("../utils/compat");
exports.default = (0, utils_1.createRule)('prefer-class-directive', {
    meta: {
        docs: {
            description: 'require class directives instead of ternary expressions',
            category: 'Stylistic Issues',
            recommended: false,
            conflictWithPrettier: false
        },
        fixable: 'code',
        schema: [
            {
                type: 'object',
                properties: {
                    prefer: { enum: ['always', 'empty'] }
                },
                additionalProperties: false
            }
        ],
        messages: {
            unexpected: 'Unexpected class using the ternary operator.'
        },
        type: 'suggestion'
    },
    create(context) {
        const sourceCode = (0, compat_1.getSourceCode)(context);
        const preferEmpty = context.options[0]?.prefer !== 'always';
        /**
         * Returns a map of expressions and strings from ConditionalExpression.
         * Returns null if it has an unknown string.
         */
        function parseConditionalExpression(node) {
            const result = new Map();
            if (!processItems({
                node: node.test
            }, node.consequent)) {
                return null;
            }
            if (!processItems({
                not: true,
                node: node.test
            }, node.alternate)) {
                return null;
            }
            return result;
            /** Process items */
            function processItems(key, e) {
                if (e.type === 'ConditionalExpression') {
                    const sub = parseConditionalExpression(e);
                    if (sub == null) {
                        return false;
                    }
                    for (const [expr, str] of sub) {
                        result.set({
                            ...key,
                            chains: expr
                        }, str);
                    }
                }
                else {
                    const str = (0, ast_utils_1.getStringIfConstant)(e);
                    if (str == null) {
                        return false;
                    }
                    result.set(key, str);
                }
                return true;
            }
        }
        /**
         * Expr to string
         */
        function exprToString({ node, not }) {
            let text = sourceCode.text.slice(...node.range);
            // *Currently not supported.
            // if (chains) {
            //   if (needParentheses(node, "logical")) {
            //     text = `(${text})`
            //   }
            //   let chainsText = exprToString(chains)
            //   const needParenForChains =
            //     !/^[!(]/u.test(chainsText) && needParentheses(chains.node, "logical")
            //   if (needParenForChains) {
            //     chainsText = `(${chainsText})`
            //   }
            //   text = `${text} && ${chainsText}`
            //   if (not) {
            //     text = `!(${text})`
            //   }
            //   return text
            // }
            if (not) {
                if (node.type === 'BinaryExpression') {
                    if (node.operator === '===' ||
                        node.operator === '==' ||
                        node.operator === '!==' ||
                        node.operator === '!=') {
                        const left = sourceCode.text.slice(...node.left.range);
                        const op = sourceCode.text.slice(node.left.range[1], node.right.range[0]);
                        const right = sourceCode.text.slice(...node.right.range);
                        return `${left}${node.operator === '===' || node.operator === '=='
                            ? op.replace(/[=](={1,2})/g, '!$1')
                            : op.replace(/!(={1,2})/g, '=$1')}${right}`;
                    }
                }
                else if (node.type === 'UnaryExpression') {
                    if (node.operator === '!' && node.prefix) {
                        return sourceCode.text.slice(...node.argument.range);
                    }
                }
                if ((0, ast_utils_1.needParentheses)(node, 'not')) {
                    text = `(${text})`;
                }
                text = `!${text}`;
            }
            return text;
        }
        /**
         * Returns all possible strings.
         */
        function getStrings(node) {
            if (node.type === 'SvelteLiteral') {
                return [node.value];
            }
            if (node.expression.type === 'ConditionalExpression') {
                const values = parseConditionalExpression(node.expression);
                if (values == null) {
                    // unknown
                    return null;
                }
                return [...values.values()];
            }
            const str = (0, ast_utils_1.getStringIfConstant)(node.expression);
            if (str == null) {
                // unknown
                return null;
            }
            return [str];
        }
        /**
         * Checks if the last character is a non word.
         */
        function endsWithNonWord(node, index) {
            for (let i = index; i >= 0; i--) {
                const valueNode = node.value[i];
                const strings = getStrings(valueNode);
                if (strings == null) {
                    // unknown
                    return false;
                }
                for (const str of strings) {
                    if (str) {
                        return !str[str.length - 1].trim();
                    }
                }
                // If the string is empty, check the previous string.
            }
            return true;
        }
        /**
         * Checks if the first character is a non word.
         */
        function startsWithNonWord(node, index) {
            for (let i = index; i < node.value.length; i++) {
                const valueNode = node.value[i];
                const strings = getStrings(valueNode);
                if (strings == null) {
                    // unknown
                    return false;
                }
                for (const str of strings) {
                    if (str) {
                        return !str[0].trim();
                    }
                }
                // If the string is empty, check the previous string.
            }
            return true;
        }
        /** Report */
        function report(node, map, attr) {
            context.report({
                node,
                messageId: 'unexpected',
                *fix(fixer) {
                    const classDirectives = [];
                    let space = ' ';
                    for (const [expr, className] of map) {
                        const trimmedClassName = className.trim();
                        if (trimmedClassName) {
                            classDirectives.push(`class:${trimmedClassName}={${exprToString(expr)}}`);
                        }
                        else {
                            space = className;
                        }
                    }
                    const fixesBuffer = [];
                    const index = attr.value.indexOf(node);
                    const beforeAttrValues = attr.value.slice(0, index);
                    const afterAttrValues = attr.value.slice(index + 1);
                    let valueNode;
                    while ((valueNode = beforeAttrValues[beforeAttrValues.length - 1])) {
                        if (valueNode.type === 'SvelteLiteral') {
                            if (!valueNode.value.trim()) {
                                // Before spaces
                                beforeAttrValues.pop();
                                fixesBuffer.push(fixer.remove(valueNode));
                                continue;
                            }
                            if (valueNode.value.trimEnd() !== valueNode.value) {
                                // Before spaces
                                fixesBuffer.push(fixer.replaceText(valueNode, valueNode.value.trimEnd()));
                            }
                        }
                        break;
                    }
                    while ((valueNode = afterAttrValues[0])) {
                        if (valueNode.type === 'SvelteLiteral') {
                            if (!valueNode.value.trim()) {
                                // After spaces
                                afterAttrValues.shift();
                                fixesBuffer.push(fixer.remove(valueNode));
                                continue;
                            }
                            if (valueNode.value.trimStart() !== valueNode.value) {
                                // After spaces
                                fixesBuffer.push(fixer.replaceText(valueNode, valueNode.value.trimStart()));
                            }
                        }
                        break;
                    }
                    if (!beforeAttrValues.length && !afterAttrValues.length) {
                        yield fixer.replaceText(attr, classDirectives.join(' '));
                    }
                    else {
                        yield* fixesBuffer;
                        if (beforeAttrValues.length && afterAttrValues.length) {
                            yield fixer.replaceText(node, space || ' ');
                        }
                        else {
                            yield fixer.remove(node);
                        }
                        yield fixer.insertTextAfterRange([attr.range[1], attr.range[1]], ` ${classDirectives.join(' ')}`);
                    }
                }
            });
        }
        /** verify */
        function verify(node, index, attr) {
            if (node.expression.type !== 'ConditionalExpression') {
                return;
            }
            const map = parseConditionalExpression(node.expression);
            if (map == null) {
                // has unknown
                return;
            }
            if (map.size > 2) {
                // It's too complicated.
                return;
            }
            if (preferEmpty && [...map.values()].every((x) => x.trim())) {
                // We prefer directives when there's an empty string, but they're all not empty
                return;
            }
            const prevIsWord = !startsWithNonWord(attr, index + 1);
            const nextIsWord = !endsWithNonWord(attr, index - 1);
            let canTransform = true;
            for (const className of map.values()) {
                if (className) {
                    if (!/^[\w-]*$/u.test(className.trim())) {
                        // Cannot be transformed to an attribute.
                        canTransform = false;
                        break;
                    }
                    if ((className[0].trim() && prevIsWord) ||
                        (className[className.length - 1].trim() && nextIsWord)) {
                        // The previous or next may be connected to this element.
                        canTransform = false;
                        break;
                    }
                }
                else if (prevIsWord && nextIsWord) {
                    // The previous and next may be connected.
                    canTransform = false;
                    break;
                }
            }
            if (!canTransform) {
                return;
            }
            report(node, map, attr);
        }
        return {
            'SvelteStartTag > SvelteAttribute'(node) {
                if (!(0, ast_utils_1.isHTMLElementLike)(node.parent.parent) || node.key.name !== 'class') {
                    return;
                }
                for (let index = 0; index < node.value.length; index++) {
                    const valueElement = node.value[index];
                    if (valueElement.type !== 'SvelteMustacheTag') {
                        continue;
                    }
                    verify(valueElement, index, node);
                }
            }
        };
    }
});
