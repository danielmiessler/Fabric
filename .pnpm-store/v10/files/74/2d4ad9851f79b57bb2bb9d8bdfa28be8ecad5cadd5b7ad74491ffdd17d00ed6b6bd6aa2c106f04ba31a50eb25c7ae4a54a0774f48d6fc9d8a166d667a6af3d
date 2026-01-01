"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
Object.defineProperty(exports, "__esModule", { value: true });
exports.equalTokens = equalTokens;
exports.getStringIfConstant = getStringIfConstant;
exports.needParentheses = needParentheses;
exports.isHTMLElementLike = isHTMLElementLike;
exports.findAttribute = findAttribute;
exports.findShorthandAttribute = findShorthandAttribute;
exports.findBindDirective = findBindDirective;
exports.getStaticAttributeValue = getStaticAttributeValue;
exports.getLangValue = getLangValue;
exports.findVariable = findVariable;
exports.iterateIdentifiers = iterateIdentifiers;
exports.getScope = getScope;
exports.getParent = getParent;
exports.getAttributeValueQuoteAndRange = getAttributeValueQuoteAndRange;
exports.getMustacheTokens = getMustacheTokens;
exports.getAttributeKeyText = getAttributeKeyText;
exports.getDirectiveName = getDirectiveName;
exports.getNodeName = getNodeName;
exports.isVoidHtmlElement = isVoidHtmlElement;
exports.isForeignElement = isForeignElement;
exports.isExpressionIdentifier = isExpressionIdentifier;
const eslintUtils = __importStar(require("@eslint-community/eslint-utils"));
const element_types_1 = require("./element-types");
const compat_1 = require("./compat");
/**
 * Checks whether or not the tokens of two given nodes are same.
 * @param left A node 1 to compare.
 * @param right A node 2 to compare.
 * @param sourceCode The ESLint source code object.
 * @returns  the source code for the given node.
 */
function equalTokens(left, right, sourceCode) {
    const tokensL = sourceCode.getTokens(left);
    const tokensR = sourceCode.getTokens(right);
    if (tokensL.length !== tokensR.length) {
        return false;
    }
    for (let i = 0; i < tokensL.length; ++i) {
        if (tokensL[i].type !== tokensR[i].type || tokensL[i].value !== tokensR[i].value) {
            return false;
        }
    }
    return true;
}
/**
 * Get the value of a given node if it's a literal or a template literal.
 */
function getStringIfConstant(node) {
    if (node.type === 'Literal') {
        if (typeof node.value === 'string')
            return node.value;
    }
    else if (node.type === 'TemplateLiteral') {
        let str = '';
        const quasis = [...node.quasis];
        const expressions = [...node.expressions];
        let quasi, expr;
        while ((quasi = quasis.shift())) {
            str += quasi.value.cooked;
            expr = expressions.shift();
            if (expr) {
                const exprStr = getStringIfConstant(expr);
                if (exprStr == null) {
                    return null;
                }
                str += exprStr;
            }
        }
        return str;
    }
    else if (node.type === 'BinaryExpression') {
        if (node.operator === '+') {
            const left = getStringIfConstant(node.left);
            if (left == null) {
                return null;
            }
            const right = getStringIfConstant(node.right);
            if (right == null) {
                return null;
            }
            return left + right;
        }
    }
    return null;
}
/**
 * Check if it need parentheses.
 */
function needParentheses(node, kind) {
    if (node.type === 'ArrowFunctionExpression' ||
        node.type === 'AssignmentExpression' ||
        node.type === 'BinaryExpression' ||
        node.type === 'ConditionalExpression' ||
        node.type === 'LogicalExpression' ||
        node.type === 'SequenceExpression' ||
        node.type === 'UnaryExpression' ||
        node.type === 'UpdateExpression')
        return true;
    if (kind === 'logical') {
        return node.type === 'FunctionExpression';
    }
    return false;
}
/** Checks whether the given node is the html element node or <svelte:element> node. */
function isHTMLElementLike(node) {
    if (node.type !== 'SvelteElement') {
        return false;
    }
    switch (node.kind) {
        case 'html':
            return true;
        case 'special':
            return node.name.name === 'svelte:element';
        default:
            return false;
    }
}
/**
 * Find the attribute from the given element node
 */
function findAttribute(node, name) {
    const startTag = node.type === 'SvelteStartTag' ? node : node.startTag;
    for (const attr of startTag.attributes) {
        if (attr.type === 'SvelteAttribute') {
            if (attr.key.name === name) {
                return attr;
            }
        }
    }
    return null;
}
/**
 * Find the shorthand attribute from the given element node
 */
function findShorthandAttribute(node, name) {
    const startTag = node.type === 'SvelteStartTag' ? node : node.startTag;
    for (const attr of startTag.attributes) {
        if (attr.type === 'SvelteShorthandAttribute') {
            if (attr.key.name === name) {
                return attr;
            }
        }
    }
    return null;
}
/**
 * Find the bind directive from the given element node
 */
function findBindDirective(node, name) {
    const startTag = node.type === 'SvelteStartTag' ? node : node.startTag;
    for (const attr of startTag.attributes) {
        if (attr.type === 'SvelteDirective') {
            if (attr.kind === 'Binding' && attr.key.name.name === name) {
                return attr;
            }
        }
    }
    return null;
}
/**
 * Get the static attribute value from given attribute
 */
function getStaticAttributeValue(node) {
    let str = '';
    for (const value of node.value) {
        if (value.type === 'SvelteLiteral') {
            str += value.value;
        }
        else {
            return null;
        }
    }
    return str;
}
/**
 * Get the static attribute value from given attribute
 */
function getLangValue(node) {
    const langAttr = findAttribute(node, 'lang');
    return langAttr && getStaticAttributeValue(langAttr);
}
/**
 * Find the variable of a given name.
 */
function findVariable(context, node) {
    const initialScope = eslintUtils.getInnermostScope(getScope(context, node), node);
    const variable = eslintUtils.findVariable(initialScope, node);
    if (variable) {
        return variable;
    }
    if (!node.name.startsWith('$')) {
        return variable;
    }
    // Remove the $ and search for the variable again, as it may be a store access variable.
    return eslintUtils.findVariable(initialScope, node.name.slice(1));
}
/**
 * Iterate the identifiers of a given pattern node.
 */
function* iterateIdentifiers(node) {
    const buffer = [node];
    let pattern;
    while ((pattern = buffer.shift())) {
        if (pattern.type === 'Identifier') {
            yield pattern;
        }
        else if (pattern.type === 'ArrayPattern') {
            for (const element of pattern.elements) {
                if (element) {
                    buffer.push(element);
                }
            }
        }
        else if (pattern.type === 'ObjectPattern') {
            for (const property of pattern.properties) {
                if (property.type === 'Property') {
                    buffer.push(property.value);
                }
                else if (property.type === 'RestElement') {
                    buffer.push(property);
                }
            }
        }
        else if (pattern.type === 'AssignmentPattern') {
            buffer.push(pattern.left);
        }
        else if (pattern.type === 'RestElement') {
            buffer.push(pattern.argument);
        }
        else if (pattern.type === 'MemberExpression') {
            // noop
        }
    }
}
/**
 * Gets the scope for the current node
 */
function getScope(context, currentNode) {
    const scopeManager = (0, compat_1.getSourceCode)(context).scopeManager;
    let node = currentNode;
    for (; node; node = node.parent || null) {
        const scope = scopeManager.acquire(node, false);
        if (scope) {
            if (scope.type === 'function-expression-name') {
                return scope.childScopes[0];
            }
            return scope;
        }
    }
    return scopeManager.scopes[0];
}
/** Get the parent node from the given node */
function getParent(node) {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any -- ignore
    return node.parent || null;
}
/** Get the quote and range from given attribute values */
function getAttributeValueQuoteAndRange(attr, sourceCode) {
    const valueTokens = getAttributeValueRangeTokens(attr, sourceCode);
    if (valueTokens == null) {
        return null;
    }
    const { firstToken: valueFirstToken, lastToken: valueLastToken } = valueTokens;
    const eqToken = sourceCode.getTokenAfter(attr.key);
    if (!eqToken || eqToken.value !== '=' || valueFirstToken.range[0] < eqToken.range[1]) {
        // invalid
        return null;
    }
    const beforeTokens = sourceCode.getTokensBetween(eqToken, valueFirstToken);
    if (beforeTokens.length === 0) {
        return {
            quote: 'unquoted',
            range: [valueFirstToken.range[0], valueLastToken.range[1]],
            firstToken: valueFirstToken,
            lastToken: valueLastToken
        };
    }
    else if (beforeTokens.length > 1 ||
        (beforeTokens[0].value !== '"' && beforeTokens[0].value !== "'")) {
        // invalid
        return null;
    }
    const beforeToken = beforeTokens[0];
    const afterToken = sourceCode.getTokenAfter(valueLastToken);
    if (!afterToken || afterToken.value !== beforeToken.value) {
        // invalid
        return null;
    }
    return {
        quote: beforeToken.value === '"' ? 'double' : 'single',
        range: [beforeToken.range[0], afterToken.range[1]],
        firstToken: beforeToken,
        lastToken: afterToken
    };
}
/** Get the mustache tokens from given node */
function getMustacheTokens(node, sourceCode) {
    if (isWrappedInBraces(node)) {
        const openToken = sourceCode.getFirstToken(node);
        const closeToken = sourceCode.getLastToken(node);
        return {
            openToken,
            closeToken
        };
    }
    if (node.expression == null) {
        return null;
    }
    if (node.key.range[0] <= node.expression.range[0] &&
        node.expression.range[1] <= node.key.range[1]) {
        // shorthand
        return null;
    }
    let openToken = sourceCode.getTokenBefore(node.expression);
    let closeToken = sourceCode.getTokenAfter(node.expression);
    while (openToken &&
        closeToken &&
        eslintUtils.isOpeningParenToken(openToken) &&
        eslintUtils.isClosingParenToken(closeToken)) {
        openToken = sourceCode.getTokenBefore(openToken);
        closeToken = sourceCode.getTokenAfter(closeToken);
    }
    if (!openToken ||
        !closeToken ||
        eslintUtils.isNotOpeningBraceToken(openToken) ||
        eslintUtils.isNotClosingBraceToken(closeToken)) {
        return null;
    }
    return {
        openToken,
        closeToken
    };
}
function isWrappedInBraces(node) {
    return (node.type === 'SvelteMustacheTag' ||
        node.type === 'SvelteShorthandAttribute' ||
        node.type === 'SvelteSpreadAttribute' ||
        node.type === 'SvelteDebugTag' ||
        node.type === 'SvelteRenderTag');
}
/** Get attribute key text */
function getAttributeKeyText(node, context) {
    switch (node.type) {
        case 'SvelteAttribute':
        case 'SvelteShorthandAttribute':
        case 'SvelteGenericsDirective':
            return node.key.name;
        case 'SvelteStyleDirective':
            return `style:${node.key.name.name}`;
        case 'SvelteSpecialDirective':
            return node.kind;
        case 'SvelteDirective': {
            const dir = getDirectiveName(node);
            return `${dir}:${getSimpleNameFromNode(node.key.name, context)}${node.key.modifiers.length ? `|${node.key.modifiers.join('|')}` : ''}`;
        }
        default:
            throw new Error(`Unknown node type: ${
            // eslint-disable-next-line @typescript-eslint/no-explicit-any -- error
            node.type}`);
    }
}
/** Get directive name */
function getDirectiveName(node) {
    switch (node.kind) {
        case 'Action':
            return 'use';
        case 'Animation':
            return 'animate';
        case 'Binding':
            return 'bind';
        case 'Class':
            return 'class';
        case 'EventHandler':
            return 'on';
        case 'Let':
            return 'let';
        case 'Transition':
            return node.intro && node.outro ? 'transition' : node.intro ? 'in' : 'out';
        case 'Ref':
            return 'ref';
        default:
            throw new Error('Unknown directive kind');
    }
}
/** Get the value tokens from given attribute */
function getAttributeValueRangeTokens(attr, sourceCode) {
    if (attr.type === 'SvelteAttribute' || attr.type === 'SvelteStyleDirective') {
        if (!attr.value.length) {
            return null;
        }
        const first = attr.value[0];
        const last = attr.value[attr.value.length - 1];
        return {
            firstToken: sourceCode.getFirstToken(first),
            lastToken: sourceCode.getLastToken(last)
        };
    }
    const tokens = getMustacheTokens(attr, sourceCode);
    if (!tokens) {
        return null;
    }
    return {
        firstToken: tokens.openToken,
        lastToken: tokens.closeToken
    };
}
/**
 * Returns name of SvelteElement
 */
function getNodeName(node) {
    return getSimpleNameFromNode(node.name);
}
/**
 * Returns true if element is known void element
 * {@link https://developer.mozilla.org/en-US/docs/Glossary/Empty_element}
 */
function isVoidHtmlElement(node) {
    return element_types_1.voidElements.includes(getNodeName(node));
}
/**
 * Returns true if element is known foreign (SVG or MathML) element
 */
function isForeignElement(node) {
    return element_types_1.svgElements.includes(getNodeName(node)) || element_types_1.mathmlElements.includes(getNodeName(node));
}
/** Checks whether the given identifier node is used as an expression. */
function isExpressionIdentifier(node) {
    const parent = node.parent;
    if (!parent) {
        return true;
    }
    if (parent.type === 'MemberExpression') {
        return !parent.computed || parent.property !== node;
    }
    if (parent.type === 'Property' ||
        parent.type === 'MethodDefinition' ||
        parent.type === 'PropertyDefinition') {
        return !parent.computed || parent.key !== node;
    }
    if (parent.type === 'FunctionDeclaration' ||
        parent.type === 'FunctionExpression' ||
        parent.type === 'ClassDeclaration' ||
        parent.type === 'ClassExpression') {
        return parent.id !== node;
    }
    if (parent.type === 'LabeledStatement' ||
        parent.type === 'BreakStatement' ||
        parent.type === 'ContinueStatement') {
        return parent.label !== node;
    }
    if (parent.type === 'MetaProperty') {
        return parent.property !== node;
    }
    if (parent.type === 'ImportSpecifier') {
        return parent.imported !== node;
    }
    if (parent.type === 'ExportSpecifier') {
        return parent.exported !== node;
    }
    return true;
}
/** Get simple name from give node */
function getSimpleNameFromNode(node, context) {
    if (node.type === 'Identifier' || node.type === 'SvelteName') {
        return node.name;
    }
    if (node.type === 'SvelteMemberExpressionName' ||
        (node.type === 'MemberExpression' && !node.computed)) {
        return `${getSimpleNameFromNode(node.object, context)}.${getSimpleNameFromNode(node.property, context)}`;
    }
    // No nodes other than those listed above are currently expected to be used in names.
    if (!context) {
        throw new Error('Rule context is required');
    }
    return (0, compat_1.getSourceCode)(context).getText(node);
}
