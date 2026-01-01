"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.defineVisitor = defineVisitor;
const commons_1 = require("./commons");
const eslint_utils_1 = require("@eslint-community/eslint-utils");
const ast_utils_1 = require("../../utils/ast-utils");
/**
 * Creates AST event handlers for ES nodes.
 *
 * @param context The rule context.
 * @returns AST event handlers.
 */
function defineVisitor(context) {
    const { sourceCode, offsets, options } = context;
    /**
     * Find the head of chaining nodes.
     */
    function getChainHeadNode(node) {
        let target = node;
        let parent = (0, ast_utils_1.getParent)(target);
        while (parent &&
            (parent.type === 'AssignmentExpression' ||
                parent.type === 'AssignmentPattern' ||
                parent.type === 'BinaryExpression' ||
                parent.type === 'LogicalExpression')) {
            const prevToken = sourceCode.getTokenBefore(target);
            const nextToken = sourceCode.getTokenAfter(target);
            if (prevToken &&
                (0, eslint_utils_1.isOpeningParenToken)(prevToken) &&
                nextToken &&
                (0, eslint_utils_1.isClosingParenToken)(nextToken)) {
                // The chain is broken because it is enclosed in parentheses.
                break;
            }
            target = parent;
            parent = (0, ast_utils_1.getParent)(target);
        }
        return target;
    }
    const visitor = {
        Program(node) {
            for (const body of node.body) {
                if (body.type === 'SvelteText' && !body.value.trim()) {
                    continue;
                }
                offsets.setStartOffsetToken(sourceCode.getFirstToken(body), 0);
            }
        },
        ArrayExpression(node) {
            const firstToken = sourceCode.getFirstToken(node);
            const rightToken = sourceCode.getTokenAfter(node.elements[node.elements.length - 1] || firstToken, { filter: eslint_utils_1.isClosingBracketToken, includeComments: false });
            offsets.setOffsetElementList(node.elements, firstToken, rightToken, 1);
        },
        ArrayPattern(node) {
            visitor.ArrayExpression(node);
        },
        ArrowFunctionExpression(node) {
            const [firstToken, secondToken] = sourceCode.getFirstTokens(node, {
                count: 2,
                includeComments: false
            });
            const leftToken = node.async ? secondToken : firstToken;
            const arrowToken = sourceCode.getTokenBefore(node.body, {
                filter: eslint_utils_1.isArrowToken,
                includeComments: false
            });
            if (node.async) {
                offsets.setOffsetToken(secondToken, 1, firstToken);
            }
            if ((0, eslint_utils_1.isOpeningParenToken)(leftToken)) {
                const rightToken = sourceCode.getTokenAfter(node.params[node.params.length - 1] || leftToken, { filter: eslint_utils_1.isClosingParenToken, includeComments: false });
                offsets.setOffsetElementList(node.params, leftToken, rightToken, 1);
            }
            offsets.setOffsetToken(arrowToken, 1, firstToken);
            const bodyFirstToken = sourceCode.getFirstToken(node.body);
            offsets.setOffsetToken(bodyFirstToken, (0, eslint_utils_1.isOpeningBraceToken)(bodyFirstToken) ? 0 : 1, firstToken);
        },
        AssignmentExpression(node) {
            const baseNode = getChainHeadNode(node);
            const opToken = sourceCode.getTokenAfter(node.left, {
                filter: eslint_utils_1.isNotClosingParenToken,
                includeComments: false
            });
            const baseToken = baseNode.type === 'AssignmentExpression' || baseNode.type === 'AssignmentPattern'
                ? sourceCode.getFirstToken(baseNode)
                : (0, commons_1.getFirstAndLastTokens)(sourceCode, baseNode).firstToken;
            const leftToken = (0, commons_1.getFirstAndLastTokens)(sourceCode, node.left).firstToken;
            const rightToken = (0, commons_1.getFirstAndLastTokens)(sourceCode, node.right).firstToken;
            offsets.setOffsetToken([leftToken === baseToken ? null : leftToken, opToken, rightToken], 1, baseToken);
        },
        AssignmentPattern(node) {
            visitor.AssignmentExpression(node);
        },
        BinaryExpression(node) {
            visitor.AssignmentExpression(node);
        },
        LogicalExpression(node) {
            visitor.AssignmentExpression(node);
        },
        AwaitExpression(node) {
            // `await`, `...`, or UnaryOperator
            const firstToken = sourceCode.getFirstToken(node);
            const nextToken = sourceCode.getTokenAfter(firstToken);
            offsets.setOffsetToken(nextToken, 1, firstToken);
        },
        RestElement(node) {
            visitor.AwaitExpression(node);
        },
        SpreadElement(node) {
            visitor.AwaitExpression(node);
        },
        UnaryExpression(node) {
            visitor.AwaitExpression(node);
        },
        BlockStatement(node) {
            offsets.setOffsetElementList(node.body, sourceCode.getFirstToken(node), sourceCode.getLastToken(node), 1);
        },
        ClassBody(node) {
            visitor.BlockStatement(node);
        },
        BreakStatement(node) {
            if (node.label) {
                const firstToken = sourceCode.getFirstToken(node);
                const nextToken = sourceCode.getTokenAfter(firstToken);
                offsets.setOffsetToken(nextToken, 1, firstToken);
            }
        },
        ContinueStatement(node) {
            visitor.BreakStatement(node);
        },
        CallExpression(node) {
            const typeArguments = node.typeArguments ?? node.typeParameters;
            const firstToken = sourceCode.getFirstToken(node);
            const leftParenToken = sourceCode.getTokenAfter(typeArguments || node.callee, {
                filter: eslint_utils_1.isOpeningParenToken,
                includeComments: false
            });
            const rightParenToken = sourceCode.getLastToken(node);
            if (typeArguments) {
                offsets.setOffsetToken(sourceCode.getFirstToken(typeArguments), 1, firstToken);
            }
            for (const optionalToken of sourceCode.getTokensBetween(sourceCode.getLastToken(typeArguments || node.callee), leftParenToken, { filter: isOptionalToken, includeComments: false })) {
                offsets.setOffsetToken(optionalToken, 1, firstToken);
            }
            offsets.setOffsetToken(leftParenToken, 1, firstToken);
            offsets.setOffsetElementList(node.arguments, leftParenToken, rightParenToken, 1);
        },
        CatchClause(node) {
            const catchToken = sourceCode.getFirstToken(node);
            if (node.param != null) {
                const leftParenToken = sourceCode.getTokenBefore(node.param);
                const rightParenToken = sourceCode.getTokenAfter(node.param);
                offsets.setOffsetToken(leftParenToken, 1, catchToken);
                offsets.setOffsetElementList([node.param], leftParenToken, rightParenToken, 1);
            }
            const bodyToken = sourceCode.getFirstToken(node.body);
            offsets.setOffsetToken(bodyToken, 0, catchToken);
        },
        ClassDeclaration(node) {
            const classToken = sourceCode.getFirstToken(node);
            if (node.id != null) {
                offsets.setOffsetToken(sourceCode.getFirstToken(node.id), 1, classToken);
            }
            if (node.superClass != null) {
                const extendsToken = sourceCode.getTokenBefore(node.superClass);
                const superClassToken = sourceCode.getTokenAfter(extendsToken);
                offsets.setOffsetToken(extendsToken, 1, classToken);
                offsets.setOffsetToken(superClassToken, 1, extendsToken);
            }
            const bodyToken = sourceCode.getFirstToken(node.body);
            offsets.setOffsetToken(bodyToken, 0, classToken);
        },
        ClassExpression(node) {
            visitor.ClassDeclaration(node);
        },
        ConditionalExpression(node) {
            const questionToken = sourceCode.getTokenAfter(node.test, {
                filter: eslint_utils_1.isNotClosingParenToken,
                includeComments: false
            });
            const consequentToken = sourceCode.getTokenAfter(questionToken);
            const colonToken = sourceCode.getTokenAfter(node.consequent, {
                filter: eslint_utils_1.isNotClosingParenToken,
                includeComments: false
            });
            const alternateToken = sourceCode.getTokenAfter(colonToken);
            let baseNode = node;
            let parent = (0, ast_utils_1.getParent)(baseNode);
            while (parent && parent.type === 'ConditionalExpression' && parent.alternate === baseNode) {
                baseNode = parent;
                parent = (0, ast_utils_1.getParent)(baseNode);
            }
            const baseToken = sourceCode.getFirstToken(baseNode);
            offsets.setOffsetToken([questionToken, colonToken], 1, baseToken);
            offsets.setOffsetToken(consequentToken, 1, questionToken);
            offsets.setOffsetToken(alternateToken, 1, colonToken);
        },
        DoWhileStatement(node) {
            const doToken = sourceCode.getFirstToken(node);
            const whileToken = sourceCode.getTokenAfter(node.body, {
                filter: eslint_utils_1.isNotClosingParenToken,
                includeComments: false
            });
            const leftParenToken = sourceCode.getTokenAfter(whileToken);
            const rightParenToken = sourceCode.getTokenAfter(node.test);
            const bodyFirstToken = sourceCode.getFirstToken(node.body);
            offsets.setOffsetToken(bodyFirstToken, (0, eslint_utils_1.isOpeningBraceToken)(bodyFirstToken) ? 0 : 1, doToken);
            offsets.setOffsetToken(whileToken, 0, doToken);
            offsets.setOffsetToken(leftParenToken, 1, whileToken);
            offsets.setOffsetElementList([node.test], leftParenToken, rightParenToken, 1);
        },
        ExportAllDeclaration(node) {
            const exportToken = sourceCode.getFirstToken(node);
            const tokens = sourceCode.getTokensBetween(exportToken, node.source);
            const fromIndex = tokens.findIndex((t) => t.value === 'from');
            const fromToken = tokens[fromIndex];
            const beforeTokens = tokens.slice(0, fromIndex);
            const afterTokens = [...tokens.slice(fromIndex + 1), sourceCode.getFirstToken(node.source)];
            // eslint-disable-next-line @typescript-eslint/no-explicit-any -- type bug?
            if (!node.exported) {
                // export * from "mod"
                offsets.setOffsetToken(beforeTokens, 1, exportToken);
            }
            else {
                // export * as foo from "mod"
                const asIndex = beforeTokens.findIndex((t) => t.value === 'as');
                offsets.setOffsetToken(beforeTokens.slice(0, asIndex), 1, exportToken);
                offsets.setOffsetToken(beforeTokens.slice(asIndex), 1, beforeTokens[asIndex - 1]);
            }
            offsets.setOffsetToken(fromToken, 0, exportToken);
            offsets.setOffsetToken(afterTokens, 1, fromToken);
            // assertions
            const lastToken = sourceCode.getLastToken(node, {
                filter: eslint_utils_1.isNotSemicolonToken,
                includeComments: false
            });
            const assertionTokens = sourceCode.getTokensBetween(node.source, lastToken);
            if (assertionTokens.length) {
                const assertToken = assertionTokens.shift();
                offsets.setOffsetToken(assertToken, 0, exportToken);
                const assertionOpen = assertionTokens.shift();
                if (assertionOpen) {
                    offsets.setOffsetToken(assertionOpen, 1, assertToken);
                    offsets.setOffsetElementList(assertionTokens, assertionOpen, lastToken, 1);
                }
            }
        },
        ExportDefaultDeclaration(node) {
            const exportToken = sourceCode.getFirstToken(node);
            const declarationToken = (0, commons_1.getFirstAndLastTokens)(sourceCode, node.declaration).firstToken;
            const defaultTokens = sourceCode.getTokensBetween(exportToken, declarationToken);
            offsets.setOffsetToken([...defaultTokens, declarationToken], 1, exportToken);
        },
        ExportNamedDeclaration(node) {
            const exportToken = sourceCode.getFirstToken(node);
            if (node.declaration) {
                // export var foo = 1;
                const declarationToken = sourceCode.getFirstToken(node.declaration);
                offsets.setOffsetToken(declarationToken, 1, exportToken);
            }
            else {
                const firstSpecifier = node.specifiers[0];
                if (!firstSpecifier || firstSpecifier.type === 'ExportSpecifier') {
                    // export {foo, bar}; or export {foo, bar} from "mod";
                    const leftBraceTokens = firstSpecifier
                        ? sourceCode.getTokensBetween(exportToken, firstSpecifier)
                        : [sourceCode.getTokenAfter(exportToken)];
                    const rightBraceToken = node.source
                        ? sourceCode.getTokenBefore(node.source, {
                            filter: eslint_utils_1.isClosingBraceToken,
                            includeComments: false
                        })
                        : sourceCode.getLastToken(node, {
                            filter: eslint_utils_1.isClosingBraceToken,
                            includeComments: false
                        });
                    offsets.setOffsetToken(leftBraceTokens, 0, exportToken);
                    offsets.setOffsetElementList(node.specifiers, leftBraceTokens[leftBraceTokens.length - 1], rightBraceToken, 1);
                    if (node.source) {
                        const [fromToken, ...tokens] = sourceCode.getTokensBetween(rightBraceToken, node.source);
                        offsets.setOffsetToken(fromToken, 0, exportToken);
                        offsets.setOffsetToken([...tokens, sourceCode.getFirstToken(node.source)], 1, fromToken);
                        // assertions
                        const lastToken = sourceCode.getLastToken(node, {
                            filter: eslint_utils_1.isNotSemicolonToken,
                            includeComments: false
                        });
                        const assertionTokens = sourceCode.getTokensBetween(node.source, lastToken);
                        if (assertionTokens.length) {
                            const assertToken = assertionTokens.shift();
                            offsets.setOffsetToken(assertToken, 0, exportToken);
                            const assertionOpen = assertionTokens.shift();
                            if (assertionOpen) {
                                offsets.setOffsetToken(assertionOpen, 1, assertToken);
                                offsets.setOffsetElementList(assertionTokens, assertionOpen, lastToken, 1);
                            }
                        }
                    }
                }
                else {
                    // maybe babel-eslint
                }
            }
        },
        ExportSpecifier(node) {
            const tokens = sourceCode.getTokens(node);
            let firstToken = tokens.shift();
            if (firstToken.value === 'type') {
                const typeToken = firstToken;
                firstToken = tokens.shift();
                offsets.setOffsetToken(firstToken, 0, typeToken);
            }
            offsets.setOffsetToken(tokens, 1, firstToken);
        },
        ForInStatement(node) {
            const forToken = sourceCode.getFirstToken(node);
            const awaitToken = (node.type === 'ForOfStatement' && node.await && sourceCode.getTokenAfter(forToken)) ||
                null;
            const leftParenToken = sourceCode.getTokenAfter(awaitToken || forToken);
            const leftToken = sourceCode.getFirstToken(node.left);
            const inOrOfToken = sourceCode.getTokenAfter(node.left, {
                filter: eslint_utils_1.isNotClosingParenToken,
                includeComments: false
            });
            const rightToken = sourceCode.getTokenAfter(inOrOfToken);
            const rightParenToken = sourceCode.getTokenBefore(node.body, {
                filter: eslint_utils_1.isNotOpeningParenToken,
                includeComments: false
            });
            if (awaitToken != null) {
                offsets.setOffsetToken(awaitToken, 0, forToken);
            }
            offsets.setOffsetToken(leftParenToken, 1, forToken);
            offsets.setOffsetToken(leftToken, 1, leftParenToken);
            offsets.setOffsetToken([inOrOfToken, rightToken], 1, leftToken);
            offsets.setOffsetToken(rightParenToken, 0, leftParenToken);
            const bodyFirstToken = sourceCode.getFirstToken(node.body);
            offsets.setOffsetToken(bodyFirstToken, (0, eslint_utils_1.isOpeningBraceToken)(bodyFirstToken) ? 0 : 1, forToken);
        },
        ForOfStatement(node) {
            visitor.ForInStatement(node);
        },
        ForStatement(node) {
            const forToken = sourceCode.getFirstToken(node);
            const leftParenToken = sourceCode.getTokenAfter(forToken);
            const rightParenToken = sourceCode.getTokenBefore(node.body, {
                filter: eslint_utils_1.isNotOpeningParenToken,
                includeComments: false
            });
            offsets.setOffsetToken(leftParenToken, 1, forToken);
            offsets.setOffsetElementList([node.init, node.test, node.update], leftParenToken, rightParenToken, 1);
            const bodyFirstToken = sourceCode.getFirstToken(node.body);
            offsets.setOffsetToken(bodyFirstToken, (0, eslint_utils_1.isOpeningBraceToken)(bodyFirstToken) ? 0 : 1, forToken);
        },
        FunctionDeclaration(node) {
            const firstToken = sourceCode.getFirstToken(node);
            const leftParenToken = sourceCode.getTokenBefore(node.params[0] ||
                node.returnType ||
                sourceCode.getTokenBefore(node.body), {
                filter: eslint_utils_1.isOpeningParenToken,
                includeComments: false
            });
            let bodyBaseToken = null;
            if (firstToken.type === 'Punctuator') {
                // method
                bodyBaseToken = sourceCode.getFirstToken((0, ast_utils_1.getParent)(node));
            }
            else {
                let tokenOffset = 0;
                for (const token of sourceCode.getTokensBetween(firstToken, leftParenToken)) {
                    if (token.value === '<') {
                        break;
                    }
                    if (token.value === '*' || (node.id && token.range[0] === node.id.range[0])) {
                        tokenOffset = 1;
                    }
                    offsets.setOffsetToken(token, tokenOffset, firstToken);
                }
                bodyBaseToken = firstToken;
            }
            const rightParenToken = sourceCode.getTokenAfter(node.params[node.params.length - 1] || leftParenToken, { filter: eslint_utils_1.isClosingParenToken, includeComments: false });
            offsets.setOffsetToken(leftParenToken, 1, bodyBaseToken);
            offsets.setOffsetElementList(node.params, leftParenToken, rightParenToken, 1);
            const bodyToken = sourceCode.getFirstToken(node.body);
            offsets.setOffsetToken(bodyToken, 0, bodyBaseToken);
        },
        FunctionExpression(node) {
            visitor.FunctionDeclaration(node);
        },
        IfStatement(node) {
            const [ifToken, ifLeftParenToken] = sourceCode.getFirstTokens(node, {
                count: 2,
                includeComments: false
            });
            const ifRightParenToken = sourceCode.getTokenBefore(node.consequent, {
                filter: eslint_utils_1.isClosingParenToken,
                includeComments: false
            });
            offsets.setOffsetToken(ifLeftParenToken, 1, ifToken);
            offsets.setOffsetToken(ifRightParenToken, 0, ifLeftParenToken);
            const consequentFirstToken = sourceCode.getFirstToken(node.consequent);
            offsets.setOffsetToken(consequentFirstToken, (0, eslint_utils_1.isOpeningBraceToken)(consequentFirstToken) ? 0 : 1, ifToken);
            if (node.alternate != null) {
                const elseToken = sourceCode.getTokenAfter(node.consequent, {
                    filter: eslint_utils_1.isNotClosingParenToken,
                    includeComments: false
                });
                offsets.setOffsetToken(elseToken, 0, ifToken);
                const alternateFirstToken = sourceCode.getFirstToken(node.alternate);
                offsets.setOffsetToken(alternateFirstToken, (0, eslint_utils_1.isOpeningBraceToken)(alternateFirstToken) ? 0 : 1, elseToken);
            }
        },
        ImportDeclaration(node) {
            const importToken = sourceCode.getFirstToken(node);
            const tokens = sourceCode.getTokensBetween(importToken, node.source);
            const fromIndex = tokens.map((t) => t.value).lastIndexOf('from');
            const { fromToken, beforeTokens, afterTokens } = fromIndex >= 0
                ? {
                    fromToken: tokens[fromIndex],
                    beforeTokens: tokens.slice(0, fromIndex),
                    afterTokens: [...tokens.slice(fromIndex + 1), sourceCode.getFirstToken(node.source)]
                }
                : {
                    fromToken: null,
                    beforeTokens: [...tokens, sourceCode.getFirstToken(node.source)],
                    afterTokens: []
                };
            const namedSpecifiers = [];
            for (const specifier of node.specifiers) {
                if (specifier.type === 'ImportSpecifier') {
                    namedSpecifiers.push(specifier);
                }
                else {
                    const removeTokens = sourceCode.getTokens(specifier);
                    removeTokens.shift();
                    for (const token of removeTokens) {
                        const i = beforeTokens.indexOf(token);
                        if (i >= 0) {
                            beforeTokens.splice(i, 1);
                        }
                    }
                }
            }
            if (namedSpecifiers.length) {
                const leftBrace = sourceCode.getTokenBefore(namedSpecifiers[0]);
                const rightBrace = sourceCode.getTokenAfter(namedSpecifiers[namedSpecifiers.length - 1], {
                    filter: eslint_utils_1.isClosingBraceToken,
                    includeComments: false
                });
                offsets.setOffsetElementList(namedSpecifiers, leftBrace, rightBrace, 1);
                for (const token of [...sourceCode.getTokensBetween(leftBrace, rightBrace), rightBrace]) {
                    const i = beforeTokens.indexOf(token);
                    if (i >= 0) {
                        beforeTokens.splice(i, 1);
                    }
                }
            }
            if (beforeTokens.every((t) => (0, eslint_utils_1.isOpeningBraceToken)(t) || (0, eslint_utils_1.isClosingBraceToken)(t))) {
                offsets.setOffsetToken(beforeTokens, 0, importToken);
            }
            else {
                offsets.setOffsetToken(beforeTokens, 1, importToken);
            }
            if (fromToken) {
                offsets.setOffsetToken(fromToken, 0, importToken);
                offsets.setOffsetToken(afterTokens, 1, fromToken);
            }
            // assertions
            const lastToken = sourceCode.getLastToken(node, {
                filter: eslint_utils_1.isNotSemicolonToken,
                includeComments: false
            });
            const assertionTokens = sourceCode.getTokensBetween(node.source, lastToken);
            if (assertionTokens.length) {
                const assertToken = assertionTokens.shift();
                offsets.setOffsetToken(assertToken, 0, importToken);
                const assertionOpen = assertionTokens.shift();
                if (assertionOpen) {
                    offsets.setOffsetToken(assertionOpen, 1, assertToken);
                    offsets.setOffsetElementList(assertionTokens, assertionOpen, lastToken, 1);
                }
            }
        },
        ImportExpression(node) {
            const firstToken = sourceCode.getFirstToken(node);
            const rightToken = sourceCode.getLastToken(node);
            const leftToken = sourceCode.getTokenAfter(firstToken, {
                filter: eslint_utils_1.isOpeningParenToken,
                includeComments: false
            });
            offsets.setOffsetToken(leftToken, 1, firstToken);
            offsets.setOffsetElementList([node.source], leftToken, rightToken, 1);
        },
        ImportNamespaceSpecifier(node) {
            const tokens = sourceCode.getTokens(node);
            const firstToken = tokens.shift();
            offsets.setOffsetToken(tokens, 1, firstToken);
        },
        ImportSpecifier(node) {
            visitor.ExportSpecifier(node);
        },
        LabeledStatement(node) {
            const labelToken = sourceCode.getFirstToken(node);
            const colonToken = sourceCode.getTokenAfter(labelToken);
            const bodyToken = sourceCode.getTokenAfter(colonToken);
            offsets.setOffsetToken([colonToken, bodyToken], 1, labelToken);
        },
        SvelteReactiveStatement(node) {
            visitor.LabeledStatement(node);
        },
        MemberExpression(node) {
            const objectToken = sourceCode.getFirstToken(node);
            if (node.type === 'MemberExpression' && node.computed) {
                const leftBracketToken = sourceCode.getTokenBefore(node.property, {
                    filter: eslint_utils_1.isOpeningBracketToken,
                    includeComments: false
                });
                const rightBracketToken = sourceCode.getTokenAfter(node.property, {
                    filter: eslint_utils_1.isClosingBracketToken,
                    includeComments: false
                });
                for (const optionalToken of sourceCode.getTokensBetween(sourceCode.getLastToken(node.object), leftBracketToken, { filter: isOptionalToken, includeComments: false })) {
                    offsets.setOffsetToken(optionalToken, 1, objectToken);
                }
                offsets.setOffsetToken(leftBracketToken, 1, objectToken);
                offsets.setOffsetElementList([node.property], leftBracketToken, rightBracketToken, 1);
            }
            else {
                const dotToken = sourceCode.getTokenBefore(node.property);
                const propertyToken = sourceCode.getTokenAfter(dotToken);
                offsets.setOffsetToken([dotToken, propertyToken], 1, objectToken);
            }
        },
        MetaProperty(node) {
            visitor.MemberExpression(node);
        },
        MethodDefinition(node) {
            const firstToken = sourceCode.getFirstToken(node);
            const keyTokens = (0, commons_1.getFirstAndLastTokens)(sourceCode, node.key);
            const prefixTokens = sourceCode.getTokensBetween(firstToken, keyTokens.firstToken);
            if (node.computed) {
                prefixTokens.pop(); // pop [
            }
            offsets.setOffsetToken(prefixTokens, 0, firstToken);
            let lastKeyToken;
            if (node.computed) {
                const leftBracketToken = sourceCode.getTokenBefore(keyTokens.firstToken);
                const rightBracketToken = (lastKeyToken = sourceCode.getTokenAfter(keyTokens.lastToken));
                offsets.setOffsetToken(leftBracketToken, 0, firstToken);
                offsets.setOffsetElementList([node.key], leftBracketToken, rightBracketToken, 1);
            }
            else {
                offsets.setOffsetToken(keyTokens.firstToken, 0, firstToken);
                lastKeyToken = keyTokens.lastToken;
            }
            if (node.value) {
                const initToken = sourceCode.getFirstToken(node.value);
                offsets.setOffsetToken([...sourceCode.getTokensBetween(lastKeyToken, initToken), initToken], 1, lastKeyToken);
            }
        },
        Property(node) {
            visitor.MethodDefinition(node);
        },
        NewExpression(node) {
            const typeArguments = node.typeArguments ?? node.typeParameters;
            const newToken = sourceCode.getFirstToken(node);
            const calleeTokens = (0, commons_1.getFirstAndLastTokens)(sourceCode, node.callee);
            offsets.setOffsetToken(calleeTokens.firstToken, 1, newToken);
            if (typeArguments) {
                offsets.setOffsetToken(sourceCode.getFirstToken(typeArguments), 1, calleeTokens.firstToken);
            }
            const leftParenBefore = typeArguments || calleeTokens.lastToken;
            if (node.arguments.length || leftParenBefore.range[1] < node.range[1]) {
                const rightParenToken = sourceCode.getLastToken(node);
                const leftParenToken = sourceCode.getTokenAfter(leftParenBefore);
                offsets.setOffsetToken(leftParenToken, 1, calleeTokens.firstToken);
                offsets.setOffsetElementList(node.arguments, leftParenToken, rightParenToken, 1);
            }
        },
        ObjectExpression(node) {
            const firstToken = sourceCode.getFirstToken(node);
            const rightToken = sourceCode.getTokenAfter(node.properties[node.properties.length - 1] || firstToken, { filter: eslint_utils_1.isClosingBraceToken, includeComments: false });
            offsets.setOffsetElementList(node.properties, firstToken, rightToken, 1);
        },
        ObjectPattern(node) {
            visitor.ObjectExpression(node);
        },
        PropertyDefinition(node) {
            visitor.MethodDefinition(node);
        },
        ReturnStatement(node) {
            if (node.argument) {
                const firstToken = sourceCode.getFirstToken(node);
                const nextToken = sourceCode.getTokenAfter(firstToken);
                offsets.setOffsetToken(nextToken, 1, firstToken);
            }
        },
        ThrowStatement(node) {
            visitor.ReturnStatement(node);
        },
        SequenceExpression(node) {
            const firstToken = sourceCode.getFirstToken(node);
            offsets.setOffsetElementList(node.expressions, firstToken, null, 0);
        },
        SwitchCase(node) {
            const caseToken = sourceCode.getFirstToken(node);
            if (node.test != null) {
                const testTokens = (0, commons_1.getFirstAndLastTokens)(sourceCode, node.test);
                const colonToken = sourceCode.getTokenAfter(testTokens.lastToken);
                offsets.setOffsetToken([testTokens.firstToken, colonToken], 1, caseToken);
            }
            else {
                const colonToken = sourceCode.getTokenAfter(caseToken);
                offsets.setOffsetToken(colonToken, 1, caseToken);
            }
            if (node.consequent.length === 1 && node.consequent[0].type === 'BlockStatement') {
                offsets.setOffsetToken(sourceCode.getFirstToken(node.consequent[0]), 0, caseToken);
            }
            else {
                for (const statement of node.consequent) {
                    offsets.setOffsetToken((0, commons_1.getFirstAndLastTokens)(sourceCode, statement).firstToken, 1, caseToken);
                }
            }
        },
        SwitchStatement(node) {
            const switchToken = sourceCode.getFirstToken(node);
            const { firstToken: leftParenToken, lastToken: rightParenToken } = (0, commons_1.getFirstAndLastTokens)(sourceCode, node.discriminant);
            const leftBraceToken = sourceCode.getTokenAfter(rightParenToken);
            const rightBraceToken = sourceCode.getLastToken(node);
            offsets.setOffsetToken(leftParenToken, 1, switchToken);
            offsets.setOffsetElementList([node.discriminant], leftParenToken, rightParenToken, 1);
            offsets.setOffsetToken(leftBraceToken, 0, switchToken);
            offsets.setOffsetElementList(node.cases, leftBraceToken, rightBraceToken, options.switchCase);
        },
        TaggedTemplateExpression(node) {
            const tagTokens = (0, commons_1.getFirstAndLastTokens)(sourceCode, node.tag);
            offsets.setOffsetToken(sourceCode.getFirstToken(node.quasi), 1, tagTokens.firstToken);
        },
        TemplateLiteral(node) {
            const firstToken = sourceCode.getFirstToken(node);
            const quasiTokens = node.quasis.slice(1).map((n) => sourceCode.getFirstToken(n));
            const expressionToken = node.quasis.slice(0, -1).map((n) => sourceCode.getTokenAfter(n));
            offsets.setOffsetToken(quasiTokens, 0, firstToken);
            offsets.setOffsetToken(expressionToken, 1, firstToken);
        },
        TryStatement(node) {
            const tryToken = sourceCode.getFirstToken(node);
            const tryBlockToken = sourceCode.getFirstToken(node.block);
            offsets.setOffsetToken(tryBlockToken, 0, tryToken);
            if (node.handler != null) {
                const catchToken = sourceCode.getFirstToken(node.handler);
                offsets.setOffsetToken(catchToken, 0, tryToken);
            }
            if (node.finalizer != null) {
                const finallyToken = sourceCode.getTokenBefore(node.finalizer);
                const finallyBlockToken = sourceCode.getFirstToken(node.finalizer);
                offsets.setOffsetToken([finallyToken, finallyBlockToken], 0, tryToken);
            }
        },
        UpdateExpression(node) {
            const firstToken = sourceCode.getFirstToken(node);
            const nextToken = sourceCode.getTokenAfter(firstToken);
            offsets.setOffsetToken(nextToken, 1, firstToken);
        },
        VariableDeclaration(node) {
            offsets.setOffsetElementList(node.declarations, sourceCode.getFirstToken(node), null, 1);
        },
        VariableDeclarator(node) {
            if (node.init != null) {
                const idToken = sourceCode.getFirstToken(node);
                const eqToken = sourceCode.getTokenAfter(node.id);
                const initToken = sourceCode.getTokenAfter(eqToken);
                offsets.setOffsetToken([eqToken, initToken], 1, idToken);
            }
        },
        WhileStatement(node) {
            const firstToken = sourceCode.getFirstToken(node);
            const leftParenToken = sourceCode.getTokenAfter(firstToken);
            const rightParenToken = sourceCode.getTokenBefore(node.body, {
                filter: eslint_utils_1.isClosingParenToken,
                includeComments: false
            });
            offsets.setOffsetToken(leftParenToken, 1, firstToken);
            offsets.setOffsetToken(rightParenToken, 0, leftParenToken);
            const bodyFirstToken = sourceCode.getFirstToken(node.body);
            offsets.setOffsetToken(bodyFirstToken, (0, eslint_utils_1.isOpeningBraceToken)(bodyFirstToken) ? 0 : 1, firstToken);
        },
        WithStatement(node) {
            visitor.WhileStatement(node);
        },
        YieldExpression(node) {
            if (node.argument != null) {
                const [yieldToken, secondToken] = sourceCode.getFirstTokens(node, {
                    count: 2,
                    includeComments: false
                });
                offsets.setOffsetToken(secondToken, 1, yieldToken);
                if (node.delegate) {
                    offsets.setOffsetToken(sourceCode.getTokenAfter(secondToken), 1, yieldToken);
                }
            }
        },
        // ----------------------------------------------------------------------
        // SINGLE TOKEN NODES
        // ----------------------------------------------------------------------
        DebuggerStatement() {
            // noop
        },
        Identifier() {
            // noop
        },
        ImportDefaultSpecifier() {
            // noop
        },
        Literal() {
            // noop
        },
        PrivateIdentifier() {
            // noop
        },
        Super() {
            // noop
        },
        TemplateElement() {
            // noop
        },
        ThisExpression() {
            // noop
        },
        // ----------------------------------------------------------------------
        // WRAPPER NODES
        // ----------------------------------------------------------------------
        ExpressionStatement() {
            // noop
        },
        ChainExpression() {
            // noop
        },
        EmptyStatement() {
            // noop
        }
    };
    // eslint-disable-next-line @typescript-eslint/no-explicit-any -- ignore
    const commonVisitor = {
        ':statement, PropertyDefinition'(node) {
            const firstToken = sourceCode.getFirstToken(node);
            const lastToken = sourceCode.getLastToken(node);
            if ((0, eslint_utils_1.isSemicolonToken)(lastToken) && firstToken !== lastToken) {
                const next = sourceCode.getTokenAfter(lastToken);
                if (!next || lastToken.loc.start.line < next.loc.start.line) {
                    // End of line semicolons
                    offsets.setOffsetToken(lastToken, 0, firstToken);
                }
            }
        },
        ':expression'(node) {
            // Proc parentheses.
            let leftToken = sourceCode.getTokenBefore(node);
            let rightToken = sourceCode.getTokenAfter(node);
            let firstToken = sourceCode.getFirstToken(node);
            while (leftToken &&
                (0, eslint_utils_1.isOpeningParenToken)(leftToken) &&
                rightToken &&
                (0, eslint_utils_1.isClosingParenToken)(rightToken)) {
                offsets.setOffsetToken(firstToken, 1, leftToken);
                offsets.setOffsetToken(rightToken, 0, leftToken);
                firstToken = leftToken;
                leftToken = sourceCode.getTokenBefore(leftToken);
                rightToken = sourceCode.getTokenAfter(rightToken);
            }
        }
    };
    const v = visitor;
    return {
        ...v,
        ...commonVisitor
    };
}
/**
 * Checks whether given text is known button type
 */
function isOptionalToken(token) {
    return token.type === 'Punctuator' && token.value === '?.';
}
