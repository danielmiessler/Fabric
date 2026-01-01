"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const svelte_store_1 = require("./reference-helpers/svelte-store");
const compat_1 = require("../utils/compat");
exports.default = (0, utils_1.createRule)('require-store-reactive-access', {
    meta: {
        docs: {
            description: 'disallow to use of the store itself as an operand. Need to use $ prefix or get function.',
            category: 'Possible Errors',
            // TODO Switch to recommended in the major version.
            // recommended: true,
            recommended: false
        },
        fixable: 'code',
        schema: [],
        messages: {
            usingRawStoreInText: 'Use the $ prefix or the get function to access reactive values instead of accessing the raw store.'
        },
        type: 'problem'
    },
    create(context) {
        if (!(0, compat_1.getSourceCode)(context).parserServices.isSvelte) {
            return {};
        }
        const isStore = (0, svelte_store_1.createStoreChecker)(context);
        /** Verify for expression node */
        function verifyExpression(node, options) {
            if (!node)
                return;
            if (isStore(node, { consistent: options?.consistent })) {
                context.report({
                    node,
                    messageId: 'usingRawStoreInText',
                    fix: node.type === 'Identifier' && !options?.disableFix
                        ? (fixer) => fixer.insertTextBefore(node, '$')
                        : null
                });
            }
        }
        return {
            SvelteMustacheTag(node) {
                if (canAcceptStoreMustache(node)) {
                    return;
                }
                // Check for <p>{store}<p/>, etc.
                verifyExpression(node.expression);
            },
            SvelteShorthandAttribute(node) {
                if (canAcceptStoreAttributeElement(node.parent.parent)) {
                    return;
                }
                // Check for <div {store} />
                verifyExpression(node.value, { disableFix: true });
            },
            SvelteSpreadAttribute(node) {
                // Check for <Foo {...store} />
                verifyExpression(node.argument);
            },
            SvelteDirective(node) {
                if (node.kind === 'Action' || node.kind === 'Animation' || node.kind === 'Transition') {
                    if (node.key.name.type !== 'Identifier') {
                        return;
                    }
                    // Check for <button use:store />, <div animate:store />, or <div transition:store />
                    verifyExpression(node.key.name);
                }
                else if (node.kind === 'Binding') {
                    if (node.key.name.name !== 'this' && canAcceptStoreAttributeElement(node.parent.parent)) {
                        return;
                    }
                    // Check for <input bind:value={store} />
                    verifyExpression(node.expression, {
                        disableFix: node.shorthand
                    });
                }
                else if (node.kind === 'Class') {
                    // Check for <div class:foo={store} />
                    verifyExpression(node.expression, {
                        disableFix: node.shorthand,
                        consistent: true
                    });
                }
                else if (node.kind === 'EventHandler') {
                    // Check for <button on:click={store} />
                    verifyExpression(node.expression);
                }
            },
            SvelteStyleDirective(node) {
                if (node.shorthand && node.key.name.type === 'Identifier') {
                    // Check for <div style:color />
                    verifyExpression(node.key.name, {
                        disableFix: true
                    });
                }
                // The longform has already been checked in SvelteMustacheTag
            },
            SvelteSpecialDirective(node) {
                if (node.kind === 'this') {
                    // Check for <button this={store} />
                    verifyExpression(node.expression);
                }
            },
            'SvelteIfBlock, SvelteAwaitBlock'(node) {
                // Check for {#if store}, {#await store}
                verifyExpression(node.expression, {
                    consistent: true
                });
            },
            SvelteEachBlock(node) {
                // Check for {#each store}
                verifyExpression(node.expression);
            },
            ['IfStatement, WhileStatement, DoWhileStatement, ConditionalExpression, ForStatement'](node) {
                // Check for `if (store)`, `while (store)`, `do {} while (store)`,
                //   `store ? a : b`, `for (;store;)`
                verifyExpression(node.test, {
                    consistent: true
                });
            },
            'ForInStatement, ForOfStatement'(node) {
                // Check for `for (let foo of store)`, `for (let foo in store)`
                verifyExpression(node.right);
            },
            SwitchStatement(node) {
                // Check for `switch (store)`
                verifyExpression(node.discriminant);
            },
            'CallExpression, NewExpression'(node) {
                if (node.callee.type === 'Super') {
                    return;
                }
                // Check for `store()`
                verifyExpression(node.callee);
            },
            UnaryExpression(node) {
                // Check for `-store`, `+store`, `!store`, `~store`, `typeof store`
                verifyExpression(node.argument, {
                    consistent: node.operator === '!' || node.operator === 'typeof'
                });
            },
            'UpdateExpression, SpreadElement'(node) {
                // Check for `store++`, `store--`, `...store`
                verifyExpression(node.argument);
            },
            AssignmentExpression(node) {
                if (node.operator !== '=') {
                    if (node.left.type !== 'ObjectPattern' && node.left.type !== 'ArrayPattern') {
                        // Check for `store += 1`
                        verifyExpression(node.left);
                    }
                    // Check for `foo += store`
                    verifyExpression(node.right);
                }
            },
            BinaryExpression(node) {
                if (node.left.type !== 'PrivateIdentifier') {
                    // Check for `store+1`
                    verifyExpression(node.left, {
                        consistent: node.operator === '==' ||
                            node.operator === '!=' ||
                            node.operator === '===' ||
                            node.operator === '!=='
                    });
                }
                // Check for `1+store`
                verifyExpression(node.right, {
                    consistent: node.operator === '==' ||
                        node.operator === '!=' ||
                        node.operator === '===' ||
                        node.operator === '!=='
                });
            },
            LogicalExpression(node) {
                // Check for `store && foo`
                verifyExpression(node.left, {
                    consistent: true
                });
            },
            TemplateLiteral(node) {
                for (const expr of node.expressions) {
                    // Check for `${store}`
                    verifyExpression(expr);
                }
            },
            TaggedTemplateExpression(node) {
                // Check for ` store`${foo}` `
                verifyExpression(node.tag);
            },
            'Property, PropertyDefinition, MethodDefinition'(node) {
                if (node.key.type === 'PrivateIdentifier' || !node.computed) {
                    return;
                }
                // Check for `{ [store]: foo}`
                verifyExpression(node.key);
            },
            ImportExpression(node) {
                // Check for `import(store)`
                verifyExpression(node.source);
            },
            AwaitExpression(node) {
                // Check for `await store`
                verifyExpression(node.argument, {
                    consistent: true
                });
            }
        };
        /**
         * Checks whether the given mustache node accepts a store instance.
         */
        function canAcceptStoreMustache(node) {
            if (node.parent.type !== 'SvelteAttribute') {
                // Text interpolation
                // e.g.
                // <p>{store}</p>
                // <input style:color={store} />
                return false;
            }
            const attr = node.parent;
            if (attr.value.length > 1) {
                // Template attribute value
                // e.g.
                // <Foo message="Hello {store}" />
                return false;
            }
            if (attr.key.name.startsWith('--')) {
                // --style-props
                // e.g.
                // <Foo --style-props={store} />
                return false;
            }
            const element = attr.parent.parent;
            return canAcceptStoreAttributeElement(element);
        }
        /**
         * Checks whether the given element node accepts a store instance attribute.
         */
        function canAcceptStoreAttributeElement(node) {
            if (node.type !== 'SvelteElement') {
                // Unknown. Within <script> or <style>
                return false;
            }
            if (node.kind === 'html' ||
                (node.kind === 'special' && node.name.name === 'svelte:element')) {
                // Native HTML attribute value
                // e.g.
                // <div data-message={store} />
                return false;
            }
            // Component props
            // e.g.
            // <Foo data={store} />
            // <Foo {store} />
            return true;
        }
    }
});
