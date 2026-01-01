"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const eslint_utils_1 = require("@eslint-community/eslint-utils");
const esutils_1 = require("esutils");
const utils_1 = require("../utils");
const ast_utils_1 = require("../utils/ast-utils");
const compat_1 = require("../utils/compat");
exports.default = (0, utils_1.createRule)('prefer-destructured-store-props', {
    meta: {
        docs: {
            description: 'destructure values from object stores for better change tracking & fewer redraws',
            category: 'Best Practices',
            recommended: false
        },
        hasSuggestions: true,
        schema: [],
        messages: {
            useDestructuring: `Destructure {{property}} from {{store}} for better change tracking & fewer redraws`,
            fixUseDestructuring: `Using destructuring like $: ({ {{property}} } = {{store}}); will run faster`,
            fixUseVariable: `Using the predefined reactive variable {{variable}}`
        },
        type: 'suggestion'
    },
    create(context) {
        let mainScript = null;
        // Store off instances of probably-destructurable statements
        const reports = [];
        let inScriptElement = false;
        const storeMemberAccessStack = [];
        /** Find for defined reactive variables. */
        function* findReactiveVariable(object, propName) {
            const storeVar = (0, ast_utils_1.findVariable)(context, object);
            if (!storeVar) {
                return;
            }
            for (const reference of storeVar.references) {
                const id = reference.identifier;
                if (id.name !== object.name)
                    continue;
                if (isReactiveVariableDefinitionWithMemberExpression(id)) {
                    // $: target = $store.prop
                    yield id.parent.parent.left;
                }
                else if (isReactiveVariableDefinitionWithDestructuring(id)) {
                    const prop = id.parent.left.properties.find((prop) => prop.type === 'Property' &&
                        prop.value.type === 'Identifier' &&
                        (0, eslint_utils_1.getPropertyName)(prop) === propName);
                    if (prop) {
                        // $: ({prop: target} = $store)
                        yield prop.value;
                    }
                }
            }
            /** Checks whether the given node is reactive variable definition with member expression. */
            function isReactiveVariableDefinitionWithMemberExpression(node) {
                return (node.type === 'Identifier' &&
                    node.parent?.type === 'MemberExpression' &&
                    node.parent.object === node &&
                    (0, eslint_utils_1.getPropertyName)(node.parent) === propName &&
                    node.parent.parent?.type === 'AssignmentExpression' &&
                    node.parent.parent.right === node.parent &&
                    node.parent.parent.left.type === 'Identifier' &&
                    node.parent.parent.parent?.type === 'ExpressionStatement' &&
                    node.parent.parent.parent.parent?.type ===
                        'SvelteReactiveStatement');
            }
            /** Checks whether the given node is reactive variable definition with destructuring. */
            function isReactiveVariableDefinitionWithDestructuring(node) {
                return (node.type === 'Identifier' &&
                    node.parent?.type === 'AssignmentExpression' &&
                    node.parent.right === node &&
                    node.parent.left.type === 'ObjectPattern' &&
                    node.parent.parent?.type === 'ExpressionStatement' &&
                    node.parent.parent.parent?.type ===
                        'SvelteReactiveStatement');
            }
        }
        /** Checks whether the given name is already defined as a variable. */
        function hasTopLevelVariable(name) {
            const scopeManager = (0, compat_1.getSourceCode)(context).scopeManager;
            if (scopeManager.globalScope?.set.has(name)) {
                return true;
            }
            const moduleScope = scopeManager.globalScope?.childScopes.find((s) => s.type === 'module');
            return moduleScope?.set.has(name) || false;
        }
        return {
            SvelteScriptElement(node) {
                inScriptElement = true;
                const scriptContext = (0, ast_utils_1.findAttribute)(node, 'context');
                const contextValue = scriptContext?.value.length === 1 && scriptContext.value[0];
                if (contextValue &&
                    contextValue.type === 'SvelteLiteral' &&
                    contextValue.value === 'module') {
                    // It is <script context="module">
                    return;
                }
                mainScript = node;
            },
            'SvelteScriptElement:exit'() {
                inScriptElement = false;
            },
            // {$foo.bar}
            // should be
            // $: ({ bar } = $foo);
            // {bar}
            // Same with {$foo["bar"]}
            "MemberExpression[object.type='Identifier'][object.name=/^\\$[^\\$]/]"(node) {
                if (inScriptElement)
                    return; // Within a script tag
                storeMemberAccessStack.unshift({ node, identifiers: [] });
            },
            Identifier(node) {
                storeMemberAccessStack[0]?.identifiers.push(node);
            },
            "MemberExpression[object.type='Identifier'][object.name=/^\\$[^\\$]/]:exit"(node) {
                if (storeMemberAccessStack[0]?.node !== node)
                    return;
                const { identifiers } = storeMemberAccessStack.shift();
                for (const id of identifiers) {
                    if (!(0, ast_utils_1.isExpressionIdentifier)(id))
                        continue;
                    const variable = (0, ast_utils_1.findVariable)(context, id);
                    const isTopLevel = !variable || variable.scope.type === 'module' || variable.scope.type === 'global';
                    if (!isTopLevel) {
                        // Member expressions may use variables defined with {#each} etc.
                        return;
                    }
                }
                reports.push(node);
            },
            'Program:exit'() {
                const scriptEndTag = mainScript && mainScript.endTag;
                for (const node of reports) {
                    const store = node.object.name;
                    const suggest = [];
                    if (
                    // Avoid suggestions for:
                    //  dynamic accesses like {$foo[bar]}
                    !node.computed) {
                        for (const variable of new Set(findReactiveVariable(node.object, node.property.name))) {
                            suggest.push({
                                messageId: 'fixUseVariable',
                                data: {
                                    variable: variable.name
                                },
                                fix(fixer) {
                                    return fixer.replaceText(node, variable.name);
                                }
                            });
                        }
                        if (
                        // Avoid suggestions for:
                        //  no <script> tag
                        //  no <script> ending
                        scriptEndTag) {
                            suggest.push({
                                messageId: 'fixUseDestructuring',
                                data: {
                                    store,
                                    property: node.property.name
                                },
                                fix(fixer) {
                                    const propName = node.property.name;
                                    let varName = propName;
                                    if (varName.startsWith('$')) {
                                        varName = varName.slice(1);
                                    }
                                    const baseName = varName;
                                    let suffix = 0;
                                    if (esutils_1.keyword.isReservedWordES6(varName, true) ||
                                        esutils_1.keyword.isRestrictedWord(varName)) {
                                        varName = `${baseName}${++suffix}`;
                                    }
                                    while (hasTopLevelVariable(varName)) {
                                        varName = `${baseName}${++suffix}`;
                                    }
                                    return [
                                        fixer.insertTextAfterRange([scriptEndTag.range[0], scriptEndTag.range[0]], `$: ({ ${propName}${propName !== varName ? `: ${varName}` : ''} } = ${store});\n`),
                                        fixer.replaceText(node, varName)
                                    ];
                                }
                            });
                        }
                    }
                    context.report({
                        node,
                        messageId: 'useDestructuring',
                        data: {
                            store,
                            property: !node.computed
                                ? node.property.name
                                : (0, compat_1.getSourceCode)(context).getText(node.property).replace(/\s+/g, ' ')
                        },
                        suggest
                    });
                }
            }
        };
    }
});
