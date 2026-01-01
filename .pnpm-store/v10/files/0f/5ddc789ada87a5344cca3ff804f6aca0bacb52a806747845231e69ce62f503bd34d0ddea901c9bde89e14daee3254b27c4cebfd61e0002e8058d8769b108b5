"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const ast_utils_1 = require("../utils/ast-utils");
const eslint_utils_1 = require("@eslint-community/eslint-utils");
const DOM_MANIPULATING_METHODS = new Set([
    'appendChild', // https://developer.mozilla.org/en-US/docs/Web/API/Node/appendChild
    'insertBefore', // https://developer.mozilla.org/en-US/docs/Web/API/Node/insertBefore
    'normalize', // https://developer.mozilla.org/en-US/docs/Web/API/Node/normalize
    'removeChild', // https://developer.mozilla.org/en-US/docs/Web/API/Node/removeChild
    'replaceChild', // https://developer.mozilla.org/en-US/docs/Web/API/Node/replaceChild
    'after', // https://developer.mozilla.org/en-US/docs/Web/API/Element/after
    'append', // https://developer.mozilla.org/en-US/docs/Web/API/Element/append
    'before', // https://developer.mozilla.org/en-US/docs/Web/API/Element/before
    'insertAdjacentElement', // https://developer.mozilla.org/en-US/docs/Web/API/Element/insertAdjacentElement
    'insertAdjacentHTML', // https://developer.mozilla.org/en-US/docs/Web/API/Element/insertAdjacentHTML
    'insertAdjacentText', // https://developer.mozilla.org/en-US/docs/Web/API/Element/insertAdjacentText
    'prepend', // https://developer.mozilla.org/en-US/docs/Web/API/Element/prepend
    'remove', // https://developer.mozilla.org/en-US/docs/Web/API/Element/remove
    'replaceChildren', // https://developer.mozilla.org/en-US/docs/Web/API/Element/replaceChildren
    'replaceWith' // https://developer.mozilla.org/en-US/docs/Web/API/Element/replaceWith
]);
const DOM_MANIPULATING_PROPERTIES = new Set([
    'textContent', // https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent
    'innerHTML', // https://developer.mozilla.org/en-US/docs/Web/API/Element/innerHTML
    'outerHTML', // https://developer.mozilla.org/en-US/docs/Web/API/Element/outerHTML
    'innerText', // https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/innerText
    'outerText' // https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/outerText
]);
exports.default = (0, utils_1.createRule)('no-dom-manipulating', {
    meta: {
        docs: {
            description: 'disallow DOM manipulating',
            category: 'Possible Errors',
            recommended: false
        },
        schema: [],
        messages: {
            disallowManipulateDOM: "Don't manipulate the DOM directly. The Svelte runtime can get confused if there is a difference between the actual DOM and the DOM expected by the Svelte runtime."
        },
        type: 'problem'
    },
    create(context) {
        const domVariables = new Set();
        /**
         * Verify DOM variable identifier node
         */
        function verifyIdentifier(node) {
            const member = node.parent;
            if (member?.type !== 'MemberExpression' || member.object !== node) {
                return;
            }
            const name = (0, eslint_utils_1.getPropertyName)(member);
            if (!name) {
                return;
            }
            let target = member;
            let parent = target.parent;
            while (parent?.type === 'ChainExpression') {
                target = parent;
                parent = parent.parent;
            }
            if (!parent) {
                return;
            }
            if (parent.type === 'CallExpression') {
                if (parent.callee !== target || !DOM_MANIPULATING_METHODS.has(name)) {
                    return;
                }
            }
            else if (parent.type === 'AssignmentExpression') {
                if (parent.left !== target || !DOM_MANIPULATING_PROPERTIES.has(name)) {
                    return;
                }
            }
            else {
                return;
            }
            context.report({
                node: member,
                messageId: 'disallowManipulateDOM'
            });
        }
        return {
            "SvelteDirective[kind='Binding']"(node) {
                if (node.key.name.name !== 'this' ||
                    !node.expression ||
                    node.expression.type !== 'Identifier') {
                    // not bind:this={id}
                    return;
                }
                const element = node.parent.parent;
                if (element.type !== 'SvelteElement' || !isHTMLElement(element)) {
                    // not HTML element
                    return;
                }
                const variable = (0, ast_utils_1.findVariable)(context, node.expression);
                if (!variable || (variable.scope.type !== 'module' && variable.scope.type !== 'global')) {
                    return;
                }
                domVariables.add(variable);
            },
            'Program:exit'() {
                for (const variable of domVariables) {
                    for (const reference of variable.references) {
                        verifyIdentifier(reference.identifier);
                    }
                }
            }
        };
        /**
         * Checks whether the given node is a HTML element or not.
         */
        function isHTMLElement(node) {
            return (node.kind === 'html' || (node.kind === 'special' && (0, ast_utils_1.getNodeName)(node) === 'svelte:element'));
        }
    }
});
