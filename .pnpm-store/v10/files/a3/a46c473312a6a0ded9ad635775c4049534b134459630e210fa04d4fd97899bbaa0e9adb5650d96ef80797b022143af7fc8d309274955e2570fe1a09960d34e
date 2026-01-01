"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.convertIfBlock = convertIfBlock;
exports.convertEachBlock = convertEachBlock;
exports.convertAwaitBlock = convertAwaitBlock;
exports.convertKeyBlock = convertKeyBlock;
exports.convertSnippetBlock = convertSnippetBlock;
const element_1 = require("./element");
const common_1 = require("./common");
const compat_1 = require("../compat");
/** Get start index of block */
function startBlockIndex(code, endIndex, block) {
    return (0, common_1.lastIndexOf)(code, (c, index) => {
        if (c !== "{") {
            return false;
        }
        for (let next = index + 1; next < code.length; next++) {
            const nextC = code[next];
            if (!nextC.trim()) {
                continue;
            }
            return code.startsWith(block, next);
        }
        return false;
    }, endIndex);
}
function startIndexFromFragment(fragment, getBeforeEndIndex) {
    if (fragment.start != null) {
        return fragment.start;
    }
    const children = (0, compat_1.getChildren)(fragment);
    return children.length ? children[0].start : getBeforeEndIndex();
}
function endIndexFromFragment(fragment, getBeforeEndIndex) {
    if (fragment.end != null) {
        return fragment.end;
    }
    const children = (0, compat_1.getChildren)(fragment);
    return children.length
        ? children[children.length - 1].end
        : getBeforeEndIndex();
}
function endIndexFromBlock(fragment, lastExpression, ctx) {
    return endIndexFromFragment(fragment, () => {
        return ctx.code.indexOf("}", (0, common_1.getWithLoc)(lastExpression).end) + 1;
    });
}
/** Convert for IfBlock */
function convertIfBlock(node, parent, ctx, elseifContext) {
    var _a;
    // {#if expr} {:else} {/if}
    // {:else if expr} {/if}
    const elseif = Boolean(elseifContext);
    const nodeStart = startBlockIndex(ctx.code, (_a = elseifContext === null || elseifContext === void 0 ? void 0 : elseifContext.start) !== null && _a !== void 0 ? _a : node.start, elseif ? ":else" : "#if");
    const ifBlock = Object.assign({ type: "SvelteIfBlock", elseif: Boolean(elseif), expression: null, children: [], else: null, parent }, ctx.getConvertLocation({ start: nodeStart, end: node.end }));
    const test = (0, compat_1.getTestFromIfBlock)(node);
    ctx.scriptLet.nestIfBlock(test, ifBlock, (es) => {
        ifBlock.expression = es;
    });
    const consequent = (0, compat_1.getConsequentFromIfBlock)(node);
    ifBlock.children.push(...(0, element_1.convertChildren)({
        nodes: 
        // Adjust for Svelte v5
        (0, compat_1.trimChildren)((0, compat_1.getChildren)(consequent)),
    }, ifBlock, ctx));
    ctx.scriptLet.closeScope();
    if (elseif) {
        const index = ctx.code.indexOf("if", nodeStart);
        ctx.addToken("MustacheKeyword", { start: index, end: index + 2 });
    }
    extractMustacheBlockTokens(ifBlock, ctx, { startOnly: elseif });
    const elseFragment = (0, compat_1.getAlternateFromIfBlock)(node);
    if (!elseFragment) {
        return ifBlock;
    }
    const elseStart = startBlockIndexForElse(elseFragment, consequent, test, ctx);
    const elseChildren = (0, compat_1.getChildren)(elseFragment);
    if (elseChildren.length === 1) {
        const c = elseChildren[0];
        if (c.type === "IfBlock" && c.elseif) {
            const elseBlock = Object.assign({ type: "SvelteElseBlock", elseif: true, children: [], parent: ifBlock }, ctx.getConvertLocation({
                start: elseStart,
                end: c.end,
            }));
            ifBlock.else = elseBlock;
            const elseIfBlock = convertIfBlock(c, elseBlock, ctx, {
                start: elseStart,
            });
            // adjust loc
            elseBlock.range[1] = elseIfBlock.range[1];
            elseBlock.loc.end = {
                line: elseIfBlock.loc.end.line,
                column: elseIfBlock.loc.end.column,
            };
            elseBlock.children = [elseIfBlock];
            return ifBlock;
        }
    }
    const elseBlock = Object.assign({ type: "SvelteElseBlock", elseif: false, children: [], parent: ifBlock }, ctx.getConvertLocation({
        start: elseStart,
        end: endIndexFromFragment(elseFragment, () => ctx.code.indexOf("}", elseStart + 5) + 1),
    }));
    ifBlock.else = elseBlock;
    ctx.scriptLet.nestBlock(elseBlock);
    elseBlock.children.push(...(0, element_1.convertChildren)({
        nodes: 
        // Adjust for Svelte v5
        (0, compat_1.trimChildren)(elseChildren),
    }, elseBlock, ctx));
    ctx.scriptLet.closeScope();
    extractMustacheBlockTokens(elseBlock, ctx, { startOnly: true });
    return ifBlock;
}
function startBlockIndexForElse(elseFragment, beforeFragment, lastExpression, ctx) {
    const elseChildren = (0, compat_1.getChildren)(elseFragment);
    if (elseChildren.length > 0) {
        const c = elseChildren[0];
        if (c.type === "IfBlock" && c.elseif) {
            const contentStart = (0, common_1.getWithLoc)((0, compat_1.getTestFromIfBlock)(c)).start;
            if (contentStart <= c.start) {
                return startBlockIndex(ctx.code, contentStart - 1, ":else");
            }
        }
        return startBlockIndex(ctx.code, c.start, ":else");
    }
    const beforeEnd = endIndexFromFragment(beforeFragment, () => {
        return ctx.code.indexOf("}", (0, common_1.getWithLoc)(lastExpression).end) + 1;
    });
    return startBlockIndex(ctx.code, beforeEnd, ":else");
}
/** Convert for EachBlock */
function convertEachBlock(node, parent, ctx) {
    // {#each expr as item, index (key)} {/each}
    const nodeStart = startBlockIndex(ctx.code, node.start, "#each");
    const eachBlock = Object.assign({ type: "SvelteEachBlock", expression: null, context: null, index: null, key: null, children: [], else: null, parent }, ctx.getConvertLocation({ start: nodeStart, end: node.end }));
    let indexRange = null;
    if (node.index) {
        const start = ctx.code.indexOf(node.index, (0, common_1.getWithLoc)(node.context).end);
        indexRange = {
            start,
            end: start + node.index.length,
        };
    }
    ctx.scriptLet.nestEachBlock(node.expression, node.context, indexRange, eachBlock, (expression, context, index) => {
        eachBlock.expression = expression;
        eachBlock.context = context;
        eachBlock.index = index;
    });
    const asStart = ctx.code.indexOf("as", (0, common_1.getWithLoc)(node.expression).end);
    ctx.addToken("Keyword", {
        start: asStart,
        end: asStart + 2,
    });
    if (node.key) {
        ctx.scriptLet.addExpression(node.key, eachBlock, null, (key) => {
            eachBlock.key = key;
        });
    }
    const body = (0, compat_1.getBodyFromEachBlock)(node);
    eachBlock.children.push(...(0, element_1.convertChildren)({
        nodes: 
        // Adjust for Svelte v5
        (0, compat_1.trimChildren)((0, compat_1.getChildren)(body)),
    }, eachBlock, ctx));
    ctx.scriptLet.closeScope();
    extractMustacheBlockTokens(eachBlock, ctx);
    const fallbackFragment = (0, compat_1.getFallbackFromEachBlock)(node);
    if (!fallbackFragment) {
        return eachBlock;
    }
    const elseStart = startBlockIndexForElse(fallbackFragment, body, node.key || indexRange || node.context, ctx);
    const elseBlock = Object.assign({ type: "SvelteElseBlock", elseif: false, children: [], parent: eachBlock }, ctx.getConvertLocation({
        start: elseStart,
        end: endIndexFromFragment(fallbackFragment, () => elseStart),
    }));
    eachBlock.else = elseBlock;
    ctx.scriptLet.nestBlock(elseBlock);
    elseBlock.children.push(...(0, element_1.convertChildren)({
        nodes: 
        // Adjust for Svelte v5
        (0, compat_1.trimChildren)((0, compat_1.getChildren)(fallbackFragment)),
    }, elseBlock, ctx));
    ctx.scriptLet.closeScope();
    extractMustacheBlockTokens(elseBlock, ctx, { startOnly: true });
    return eachBlock;
}
/** Convert for AwaitBlock */
function convertAwaitBlock(node, parent, ctx) {
    const nodeStart = startBlockIndex(ctx.code, node.start, "#await");
    const awaitBlock = Object.assign({ type: "SvelteAwaitBlock", expression: null, kind: "await", pending: null, then: null, catch: null, parent }, ctx.getConvertLocation({ start: nodeStart, end: node.end }));
    ctx.scriptLet.addExpression(node.expression, awaitBlock, null, (expression) => {
        awaitBlock.expression = expression;
    });
    const pending = (0, compat_1.getPendingFromAwaitBlock)(node);
    if (pending) {
        const pendingBlock = Object.assign({ type: "SvelteAwaitPendingBlock", children: [], parent: awaitBlock }, ctx.getConvertLocation({
            start: awaitBlock.range[0],
            end: endIndexFromBlock(pending, node.expression, ctx),
        }));
        ctx.scriptLet.nestBlock(pendingBlock);
        pendingBlock.children.push(...(0, element_1.convertChildren)(pending, pendingBlock, ctx));
        awaitBlock.pending = pendingBlock;
        ctx.scriptLet.closeScope();
    }
    const then = (0, compat_1.getThenFromAwaitBlock)(node);
    if (then) {
        const awaitThen = !pending;
        if (awaitThen) {
            awaitBlock.kind = "await-then";
        }
        const thenStart = awaitBlock.pending
            ? startBlockIndex(ctx.code, node.value
                ? (0, common_1.getWithLoc)(node.value).start
                : startIndexFromFragment(then, () => {
                    return awaitBlock.pending.range[1];
                }), ":then")
            : nodeStart;
        const thenBlock = Object.assign({ type: "SvelteAwaitThenBlock", awaitThen, value: null, children: [], parent: awaitBlock }, ctx.getConvertLocation({
            start: thenStart,
            end: endIndexFromFragment(then, () => {
                return (ctx.code.indexOf("}", node.value ? (0, common_1.getWithLoc)(node.value).end : thenStart + 5) + 1);
            }),
        }));
        if (node.value) {
            const baseParam = {
                node: node.value,
                parent: thenBlock,
                callback(value) {
                    thenBlock.value = value;
                },
                typing: "any",
            };
            ctx.scriptLet.nestBlock(thenBlock, (typeCtx) => {
                if (!typeCtx) {
                    return {
                        param: baseParam,
                    };
                }
                const expression = ctx.getText(node.expression);
                if (node.expression.type === "Literal") {
                    return {
                        param: Object.assign(Object.assign({}, baseParam), { typing: expression }),
                    };
                }
                const idAwaitThenValue = typeCtx.generateUniqueId("AwaitThenValue");
                if (node.expression.type === "Identifier" &&
                    // We cannot use type annotations like `(x: Foo<x>)` if they have the same identifier name.
                    !hasIdentifierFor(node.expression.name, baseParam.node)) {
                    return {
                        preparationScript: [generateAwaitThenValueType(idAwaitThenValue)],
                        param: Object.assign(Object.assign({}, baseParam), { typing: `${idAwaitThenValue}<(typeof ${expression})>` }),
                    };
                }
                const id = typeCtx.generateUniqueId(expression);
                return {
                    preparationScript: [
                        {
                            script: `const ${id} = ${expression};`,
                            nodeType: "VariableDeclaration",
                        },
                        generateAwaitThenValueType(idAwaitThenValue),
                    ],
                    param: Object.assign(Object.assign({}, baseParam), { typing: `${idAwaitThenValue}<(typeof ${id})>` }),
                };
            });
        }
        else {
            ctx.scriptLet.nestBlock(thenBlock);
        }
        thenBlock.children.push(...(0, element_1.convertChildren)(then, thenBlock, ctx));
        if (awaitBlock.pending) {
            extractMustacheBlockTokens(thenBlock, ctx, { startOnly: true });
        }
        else {
            const thenIndex = ctx.code.indexOf("then", (0, common_1.getWithLoc)(node.expression).end);
            ctx.addToken("MustacheKeyword", {
                start: thenIndex,
                end: thenIndex + 4,
            });
        }
        awaitBlock.then = thenBlock;
        ctx.scriptLet.closeScope();
    }
    const catchFragment = (0, compat_1.getCatchFromAwaitBlock)(node);
    if (catchFragment) {
        const awaitCatch = !pending && !then;
        if (awaitCatch) {
            awaitBlock.kind = "await-catch";
        }
        const catchStart = awaitBlock.then || awaitBlock.pending
            ? startBlockIndex(ctx.code, node.error
                ? (0, common_1.getWithLoc)(node.error).start
                : startIndexFromFragment(catchFragment, () => {
                    return (awaitBlock.then || awaitBlock.pending).range[1];
                }), ":catch")
            : nodeStart;
        const catchBlock = Object.assign({ type: "SvelteAwaitCatchBlock", awaitCatch, error: null, children: [], parent: awaitBlock }, ctx.getConvertLocation({
            start: catchStart,
            end: endIndexFromFragment(catchFragment, () => {
                return (ctx.code.indexOf("}", node.error ? (0, common_1.getWithLoc)(node.error).end : catchStart + 6) + 1);
            }),
        }));
        if (node.error) {
            ctx.scriptLet.nestBlock(catchBlock, [
                {
                    node: node.error,
                    parent: catchBlock,
                    typing: "Error",
                    callback: (error) => {
                        catchBlock.error = error;
                    },
                },
            ]);
        }
        else {
            ctx.scriptLet.nestBlock(catchBlock);
        }
        catchBlock.children.push(...(0, element_1.convertChildren)(catchFragment, catchBlock, ctx));
        if (awaitBlock.pending || awaitBlock.then) {
            extractMustacheBlockTokens(catchBlock, ctx, { startOnly: true });
        }
        else {
            const catchIndex = ctx.code.indexOf("catch", (0, common_1.getWithLoc)(node.expression).end);
            ctx.addToken("MustacheKeyword", {
                start: catchIndex,
                end: catchIndex + 5,
            });
        }
        awaitBlock.catch = catchBlock;
        ctx.scriptLet.closeScope();
    }
    extractMustacheBlockTokens(awaitBlock, ctx);
    return awaitBlock;
}
/** Convert for KeyBlock */
function convertKeyBlock(node, parent, ctx) {
    const nodeStart = startBlockIndex(ctx.code, node.start, "#key");
    const keyBlock = Object.assign({ type: "SvelteKeyBlock", expression: null, children: [], parent }, ctx.getConvertLocation({ start: nodeStart, end: node.end }));
    ctx.scriptLet.addExpression(node.expression, keyBlock, null, (expression) => {
        keyBlock.expression = expression;
    });
    ctx.scriptLet.nestBlock(keyBlock);
    keyBlock.children.push(...(0, element_1.convertChildren)({
        nodes: 
        // Adjust for Svelte v5
        (0, compat_1.trimChildren)((0, compat_1.getChildren)((0, compat_1.getFragment)(node))),
    }, keyBlock, ctx));
    ctx.scriptLet.closeScope();
    extractMustacheBlockTokens(keyBlock, ctx);
    return keyBlock;
}
/** Convert for SnippetBlock */
function convertSnippetBlock(node, parent, ctx) {
    var _a;
    // {#snippet x(args)}...{/snippet}
    const nodeStart = startBlockIndex(ctx.code, node.start, "#snippet");
    const snippetBlock = Object.assign({ type: "SvelteSnippetBlock", id: null, params: [], children: [], parent }, ctx.getConvertLocation({ start: nodeStart, end: node.end }));
    let beforeClosingParen;
    if (node.parameters.length > 0) {
        const lastParam = node.parameters[node.parameters.length - 1];
        beforeClosingParen = (_a = lastParam.typeAnnotation) !== null && _a !== void 0 ? _a : lastParam;
    }
    else {
        beforeClosingParen = node.expression;
    }
    const closeParenIndex = ctx.code.indexOf(")", (0, common_1.getWithLoc)(beforeClosingParen).end);
    const scopeKind = parent.type === "Program"
        ? "snippet"
        : // use currentScriptScopeKind
            null;
    ctx.scriptLet.nestSnippetBlock(node.expression, closeParenIndex, snippetBlock, scopeKind, (id, params) => {
        snippetBlock.id = id;
        snippetBlock.params = params;
    });
    snippetBlock.children.push(...(0, element_1.convertChildren)({
        nodes: 
        // Adjust for Svelte v5
        (0, compat_1.trimChildren)(node.body.nodes),
    }, snippetBlock, ctx));
    ctx.scriptLet.closeScope();
    extractMustacheBlockTokens(snippetBlock, ctx);
    ctx.snippets.push(snippetBlock);
    return snippetBlock;
}
/** Extract mustache block tokens */
function extractMustacheBlockTokens(node, ctx, option) {
    const startSectionNameStart = (0, common_1.indexOf)(ctx.code, (c) => Boolean(c.trim()), node.range[0] + 1);
    const startSectionNameEnd = (0, common_1.indexOf)(ctx.code, (c) => c === "}" || !c.trim(), startSectionNameStart + 1);
    ctx.addToken("MustacheKeyword", {
        start: startSectionNameStart,
        end: startSectionNameEnd,
    });
    if (option === null || option === void 0 ? void 0 : option.startOnly) {
        return;
    }
    const endSectionNameEnd = (0, common_1.lastIndexOf)(ctx.code, (c) => Boolean(c.trim()), node.range[1] - 2) + 1;
    const endSectionNameStart = (0, common_1.lastIndexOf)(ctx.code, (c) => c === "{" || c === "/" || !c.trim(), endSectionNameEnd - 1);
    ctx.addToken("MustacheKeyword", {
        start: endSectionNameStart,
        end: endSectionNameEnd,
    });
}
/** Generate Awaited like type code */
function generateAwaitThenValueType(id) {
    return {
        script: `type ${id}<T> = T extends null | undefined
    ? T
    : T extends { then(value: infer F): any }
    ? F extends (value: infer V, ...args: any) => any
        ? ${id}<V>
        : never
        : T;`,
        nodeType: "TSTypeAliasDeclaration",
    };
}
/** Checks whether the given name identifier is exists or not. */
function hasIdentifierFor(name, node) {
    if (node.type === "Identifier") {
        return node.name === name;
    }
    if (node.type === "ObjectPattern") {
        return node.properties.some((property) => property.type === "Property"
            ? hasIdentifierFor(name, property.value)
            : hasIdentifierFor(name, property));
    }
    if (node.type === "ArrayPattern") {
        return node.elements.some((element) => element && hasIdentifierFor(name, element));
    }
    if (node.type === "RestElement") {
        return hasIdentifierFor(name, node.argument);
    }
    if (node.type === "AssignmentPattern") {
        return hasIdentifierFor(name, node.left);
    }
    return false;
}
