"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const eslint_utils_1 = require("@eslint-community/eslint-utils");
const compat_1 = require("../utils/compat");
const ast_utils_1 = require("../utils/ast-utils");
exports.default = (0, utils_1.createRule)('no-goto-without-base', {
    meta: {
        docs: {
            description: 'disallow using goto() without the base path',
            category: 'SvelteKit',
            recommended: false
        },
        schema: [],
        messages: {
            isNotPrefixedWithBasePath: "Found a goto() call with a url that isn't prefixed with the base path."
        },
        type: 'suggestion'
    },
    create(context) {
        return {
            Program() {
                const referenceTracker = new eslint_utils_1.ReferenceTracker((0, compat_1.getSourceCode)(context).scopeManager.globalScope);
                const basePathNames = extractBasePathReferences(referenceTracker, context);
                for (const gotoCall of extractGotoReferences(referenceTracker)) {
                    if (gotoCall.arguments.length < 1) {
                        continue;
                    }
                    const path = gotoCall.arguments[0];
                    switch (path.type) {
                        case 'BinaryExpression':
                            checkBinaryExpression(context, path, basePathNames);
                            break;
                        case 'Literal':
                            checkLiteral(context, path);
                            break;
                        case 'TemplateLiteral':
                            checkTemplateLiteral(context, path, basePathNames);
                            break;
                        default:
                            context.report({ loc: path.loc, messageId: 'isNotPrefixedWithBasePath' });
                    }
                }
            }
        };
    }
});
function checkBinaryExpression(context, path, basePathNames) {
    if (path.left.type !== 'Identifier' || !basePathNames.has(path.left)) {
        context.report({ loc: path.loc, messageId: 'isNotPrefixedWithBasePath' });
    }
}
function checkTemplateLiteral(context, path, basePathNames) {
    const startingIdentifier = extractStartingIdentifier(path);
    if (startingIdentifier === undefined || !basePathNames.has(startingIdentifier)) {
        context.report({ loc: path.loc, messageId: 'isNotPrefixedWithBasePath' });
    }
}
function checkLiteral(context, path) {
    const absolutePathRegex = /^(?:[+a-z]+:)?\/\//i;
    if (!absolutePathRegex.test(path.value?.toString() ?? '')) {
        context.report({ loc: path.loc, messageId: 'isNotPrefixedWithBasePath' });
    }
}
function extractStartingIdentifier(templateLiteral) {
    const literalParts = [...templateLiteral.expressions, ...templateLiteral.quasis].sort((a, b) => a.range[0] < b.range[0] ? -1 : 1);
    for (const part of literalParts) {
        if (part.type === 'TemplateElement' && part.value.raw === '') {
            // Skip empty quasi in the begining
            continue;
        }
        if (part.type === 'Identifier') {
            return part;
        }
        return undefined;
    }
    return undefined;
}
function extractGotoReferences(referenceTracker) {
    return Array.from(referenceTracker.iterateEsmReferences({
        '$app/navigation': {
            [eslint_utils_1.ReferenceTracker.ESM]: true,
            goto: {
                [eslint_utils_1.ReferenceTracker.CALL]: true
            }
        }
    }), ({ node }) => node);
}
function extractBasePathReferences(referenceTracker, context) {
    const set = new Set();
    for (const { node } of referenceTracker.iterateEsmReferences({
        '$app/paths': {
            [eslint_utils_1.ReferenceTracker.ESM]: true,
            base: {
                [eslint_utils_1.ReferenceTracker.READ]: true
            }
        }
    })) {
        const variable = (0, ast_utils_1.findVariable)(context, node.local);
        if (!variable)
            continue;
        for (const reference of variable.references) {
            if (reference.identifier.type === 'Identifier')
                set.add(reference.identifier);
        }
    }
    return set;
}
