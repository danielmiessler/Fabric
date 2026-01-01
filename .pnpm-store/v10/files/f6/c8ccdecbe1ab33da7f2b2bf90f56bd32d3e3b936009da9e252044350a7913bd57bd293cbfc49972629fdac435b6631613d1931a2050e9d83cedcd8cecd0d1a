"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.ScriptLetContext = void 0;
const common_1 = require("../parser/converts/common");
const scope_1 = require("../scope");
const traverse_1 = require("../traverse");
const unique_1 = require("./unique");
const fix_locations_1 = require("./fix-locations");
/**
 * Get node range
 */
function getNodeRange(node) {
    let start = null;
    let end = null;
    if (node.leadingComments) {
        start = (0, common_1.getWithLoc)(node.leadingComments[0]).start;
    }
    if (node.trailingComments) {
        end = (0, common_1.getWithLoc)(node.trailingComments[node.trailingComments.length - 1]).end;
    }
    const loc = "range" in node
        ? { start: node.range[0], end: node.range[1] }
        : (0, common_1.getWithLoc)(node);
    return [
        start ? Math.min(start, loc.start) : loc.start,
        end ? Math.max(end, loc.end) : loc.end,
    ];
}
/**
 * A class that handles script fragments.
 * The script fragment AST node remaps and connects to the original directive AST node.
 */
class ScriptLetContext {
    constructor(ctx) {
        this.restoreCallbacks = [];
        this.programRestoreCallbacks = [];
        this.closeScopeCallbacks = [];
        this.unique = new unique_1.UniqueIdGenerator();
        this.currentScriptScopeKind = "render";
        this.script = ctx.sourceCode.scripts;
        this.ctx = ctx;
    }
    addExpression(expression, parent, typing, ...callbacks) {
        const range = getNodeRange(expression);
        return this.addExpressionFromRange(range, parent, typing, ...callbacks);
    }
    addExpressionFromRange(range, parent, typing, ...callbacks) {
        const part = this.ctx.code.slice(...range);
        const isTS = typing && this.ctx.isTypeScript();
        this.appendScript(`(${part})${isTS ? `as (${typing})` : ""};`, range[0] - 1, this.currentScriptScopeKind, "ExpressionStatement", (st, tokens, comments, result) => {
            const exprSt = st;
            const tsAs = isTS
                ? exprSt.expression
                : null;
            const node = (tsAs === null || tsAs === void 0 ? void 0 : tsAs.expression) || exprSt.expression;
            // Process for nodes
            for (const callback of callbacks) {
                callback(node, result);
            }
            tokens.shift(); // (
            tokens.pop(); // )
            tokens.pop(); // ;
            if (isTS) {
                for (const scope of extractTypeNodeScopes(tsAs.typeAnnotation, result)) {
                    (0, scope_1.removeScope)(result.scopeManager, scope);
                }
                this.remapNodes([
                    {
                        offset: range[0] - 1,
                        range,
                        newNode: node,
                    },
                ], tokens, comments, result.visitorKeys);
            }
            node.parent = parent;
            // Disconnect the tree structure.
            exprSt.expression = null;
        });
        return callbacks;
    }
    addObjectShorthandProperty(identifier, parent, ...callbacks) {
        const range = getNodeRange(identifier);
        const part = this.ctx.code.slice(...range);
        this.appendScript(`({${part}});`, range[0] - 2, this.currentScriptScopeKind, "ExpressionStatement", (st, tokens, _comments, result) => {
            const exprSt = st;
            const objectExpression = exprSt.expression;
            const node = objectExpression.properties[0];
            // Process for nodes
            for (const callback of callbacks) {
                callback(node, result);
            }
            node.key.parent = parent;
            node.value.parent = parent;
            tokens.shift(); // (
            tokens.shift(); // {
            tokens.pop(); // }
            tokens.pop(); // )
            tokens.pop(); // ;
            // Disconnect the tree structure.
            exprSt.expression = null;
        });
    }
    addVariableDeclarator(declarator, parent, ...callbacks) {
        const range = declarator.type === "VariableDeclarator"
            ? // As of Svelte v5-next.65, VariableDeclarator nodes do not have location information.
                [getNodeRange(declarator.id)[0], getNodeRange(declarator.init)[1]]
            : getNodeRange(declarator);
        const part = this.ctx.code.slice(...range);
        this.appendScript(`const ${part};`, range[0] - 6, this.currentScriptScopeKind, "VariableDeclaration", (st, tokens, _comments, result) => {
            const decl = st;
            const node = decl.declarations[0];
            // Process for nodes
            for (const callback of callbacks) {
                callback(node, result);
            }
            const scope = result.getScope(decl);
            for (const variable of scope.variables) {
                for (const def of variable.defs) {
                    if (def.parent === decl) {
                        def.parent = parent;
                    }
                }
            }
            node.parent = parent;
            tokens.shift(); // const
            tokens.pop(); // ;
            // Disconnect the tree structure.
            decl.declarations = [];
        });
        return callbacks;
    }
    addGenericTypeAliasDeclaration(param, callbackId, callbackDefault) {
        var _a, _b;
        const ranges = getTypeParameterRanges(this.ctx.code, param);
        let scriptLet = `type ${this.ctx.code.slice(...ranges.idRange)}`;
        if (ranges.constraintRange) {
            scriptLet += ` =     ${this.ctx.code.slice(...ranges.constraintRange)};`;
            //           |extends|
        }
        else {
            scriptLet += " = unknown;";
        }
        this.appendScript(scriptLet, ranges.idRange[0] - 5, "generics", "TSTypeAliasDeclaration", (st, tokens, _comments, result) => {
            const decl = st;
            const id = decl.id;
            const typeAnnotation = decl.typeAnnotation;
            // Process for nodes
            callbackId(id, typeAnnotation);
            const scope = result.getScope(decl);
            for (const variable of scope.variables) {
                for (const def of variable.defs) {
                    if (def.node === decl) {
                        def.node = param;
                    }
                }
            }
            id.parent = param;
            typeAnnotation.parent = param;
            tokens.shift(); // type
            if (ranges.constraintRange) {
                const eqToken = tokens[1];
                eqToken.type = "Keyword";
                eqToken.value = "extends";
                eqToken.range[0] = eqToken.range[0] - 1;
                eqToken.range[1] = eqToken.range[0] + 7;
                tokens.pop(); // ;
            }
            else {
                tokens.pop(); // ;
                tokens.pop(); // unknown
                tokens.pop(); // =
            }
            // Disconnect the tree structure.
            delete decl.id;
            delete decl.typeAnnotation;
            delete decl.typeParameters;
        });
        if (ranges.defaultRange) {
            const eqDefaultType = this.ctx.code.slice((_b = (_a = ranges.constraintRange) === null || _a === void 0 ? void 0 : _a[1]) !== null && _b !== void 0 ? _b : ranges.idRange[1], ranges.defaultRange[1]);
            const id = this.generateUniqueId(eqDefaultType);
            const scriptLet = `type ${id}${eqDefaultType};`;
            this.appendScript(scriptLet, ranges.defaultRange[0] - 5 - id.length - 1, "generics", "TSTypeAliasDeclaration", (st, tokens, _comments, result) => {
                const decl = st;
                const typeAnnotation = decl.typeAnnotation;
                // Process for nodes
                callbackDefault(typeAnnotation);
                const scope = result.getScope(decl);
                (0, scope_1.removeIdentifierVariable)(decl.id, scope);
                typeAnnotation.parent = param;
                tokens.shift(); // type
                tokens.shift(); // ${id}
                tokens.pop(); // ;
                // Disconnect the tree structure.
                delete decl.id;
                delete decl.typeAnnotation;
                delete decl.typeParameters;
            });
        }
    }
    nestIfBlock(expression, ifBlock, callback) {
        const range = getNodeRange(expression);
        const part = this.ctx.code.slice(...range);
        const restore = this.appendScript(`if(${part}){`, range[0] - 3, this.currentScriptScopeKind, "IfStatement", (st, tokens, _comments, result) => {
            const ifSt = st;
            const node = ifSt.test;
            const scope = result.getScope(ifSt.consequent);
            // Process for nodes
            callback(node, result);
            node.parent = ifBlock;
            // Process for scope
            result.registerNodeToScope(ifBlock, scope);
            tokens.shift(); // if
            tokens.shift(); // (
            tokens.pop(); // )
            tokens.pop(); // {
            tokens.pop(); // }
            // Disconnect the tree structure.
            ifSt.test = null;
            ifSt.consequent = null;
        });
        this.pushScope(restore, "}", this.currentScriptScopeKind);
    }
    nestEachBlock(expression, context, indexRange, eachBlock, callback) {
        const exprRange = getNodeRange(expression);
        const ctxRange = getNodeRange(context);
        let source = "Array.from(";
        const exprOffset = source.length;
        source += `${this.ctx.code.slice(...exprRange)}).forEach((`;
        const ctxOffset = source.length;
        source += this.ctx.code.slice(...ctxRange);
        let idxOffset = null;
        if (indexRange) {
            source += ",";
            idxOffset = source.length;
            source += this.ctx.code.slice(indexRange.start, indexRange.end);
        }
        source += ")=>{";
        const restore = this.appendScript(source, exprRange[0] - exprOffset, this.currentScriptScopeKind, "ExpressionStatement", (st, tokens, comments, result) => {
            var _a;
            const expSt = st;
            const call = expSt.expression;
            const fn = call.arguments[0];
            const callArrayFrom = call.callee
                .object;
            const expr = callArrayFrom.arguments[0];
            const ctx = fn.params[0];
            const idx = ((_a = fn.params[1]) !== null && _a !== void 0 ? _a : null);
            const scope = result.getScope(fn.body);
            // Process for nodes
            callback(expr, ctx, idx);
            // Process for scope
            result.registerNodeToScope(eachBlock, scope);
            for (const v of scope.variables) {
                for (const def of v.defs) {
                    if (def.node === fn) {
                        def.node = eachBlock;
                    }
                }
            }
            // remove Array reference
            const arrayId = callArrayFrom.callee
                .object;
            const ref = scope.upper.references.find((r) => r.identifier === arrayId);
            if (ref) {
                (0, scope_1.removeReference)(ref, scope.upper);
            }
            expr.parent = eachBlock;
            ctx.parent = eachBlock;
            if (idx) {
                idx.parent = eachBlock;
            }
            tokens.shift(); // Array
            tokens.shift(); // .
            tokens.shift(); // from
            tokens.shift(); // (
            tokens.pop(); // )
            tokens.pop(); // =>
            tokens.pop(); // {
            tokens.pop(); // }
            tokens.pop(); // )
            tokens.pop(); // ;
            const map = [
                {
                    offset: exprOffset,
                    range: exprRange,
                    newNode: expr,
                },
                {
                    offset: ctxOffset,
                    range: ctxRange,
                    newNode: ctx,
                },
            ];
            if (indexRange) {
                map.push({
                    offset: idxOffset,
                    range: [indexRange.start, indexRange.end],
                    newNode: idx,
                });
            }
            this.remapNodes(map, tokens, comments, result.visitorKeys);
            // Disconnect the tree structure.
            expSt.expression = null;
        });
        this.pushScope(restore, "});", this.currentScriptScopeKind);
    }
    nestSnippetBlock(id, closeParentIndex, snippetBlock, 
    // If set to null, `currentScriptScopeKind` will be used.
    kind, callback) {
        const scopeKind = kind || this.currentScriptScopeKind;
        const idRange = getNodeRange(id);
        const part = this.ctx.code.slice(idRange[0], closeParentIndex + 1);
        const restore = this.appendScript(`function ${part}{`, idRange[0] - 9, scopeKind, "FunctionDeclaration", (st, tokens, _comments, result) => {
            const fnDecl = st;
            const idNode = fnDecl.id;
            const params = [...fnDecl.params];
            const scope = result.getScope(fnDecl);
            // Process for nodes
            callback(idNode, params);
            idNode.parent = snippetBlock;
            for (const param of params) {
                param.parent = snippetBlock;
            }
            // Process for scope
            result.registerNodeToScope(snippetBlock, scope);
            tokens.shift(); // function
            tokens.pop(); // {
            tokens.pop(); // }
            // Disconnect the tree structure.
            fnDecl.id = null;
            fnDecl.params = [];
        });
        this.pushScope(restore, "}", scopeKind);
    }
    nestBlock(block, params) {
        let resolvedParams;
        if (typeof params === "function") {
            if (this.ctx.isTypeScript()) {
                const generatedTypes = params({
                    generateUniqueId: (base) => this.generateUniqueId(base),
                });
                resolvedParams = [generatedTypes.param];
                if (generatedTypes.preparationScript) {
                    for (const preparationScript of generatedTypes.preparationScript) {
                        this.appendScriptWithoutOffset(preparationScript.script, this.currentScriptScopeKind, preparationScript.nodeType, (node, tokens, comments, result) => {
                            tokens.length = 0;
                            comments.length = 0;
                            (0, scope_1.removeAllScopeAndVariableAndReference)(node, result);
                        });
                    }
                }
            }
            else {
                const generatedTypes = params(null);
                resolvedParams = [generatedTypes.param];
            }
        }
        else {
            resolvedParams = params;
        }
        if (!resolvedParams || resolvedParams.length === 0) {
            const restore = this.appendScript(`{`, block.range[0], this.currentScriptScopeKind, "BlockStatement", (st, tokens, _comments, result) => {
                const blockSt = st;
                // Process for scope
                const scope = result.getScope(blockSt);
                result.registerNodeToScope(block, scope);
                tokens.length = 0; // clear tokens
                // Disconnect the tree structure.
                blockSt.body = null;
            });
            this.pushScope(restore, "}", this.currentScriptScopeKind);
        }
        else {
            const sortedParams = [...resolvedParams]
                .map((d) => {
                return Object.assign(Object.assign({}, d), { range: getNodeRange(d.node) });
            })
                .sort((a, b) => {
                const [pA] = a.range;
                const [pB] = b.range;
                return pA - pB;
            });
            const maps = [];
            let source = "";
            for (let index = 0; index < sortedParams.length; index++) {
                const param = sortedParams[index];
                const range = param.range;
                if (source) {
                    source += ",";
                }
                const offset = source.length + 1; /* ( */
                source += this.ctx.code.slice(...range);
                maps.push({
                    index: maps.length,
                    offset,
                    range,
                });
                if (this.ctx.isTypeScript()) {
                    source += `: (${param.typing})`;
                }
            }
            const restore = this.appendScript(`(${source})=>{`, maps[0].range[0] - 1, this.currentScriptScopeKind, "ExpressionStatement", (st, tokens, comments, result) => {
                const exprSt = st;
                const fn = exprSt.expression;
                const scope = result.getScope(fn.body);
                // Process for nodes
                for (let index = 0; index < fn.params.length; index++) {
                    const p = fn.params[index];
                    sortedParams[index].callback(p, result);
                    if (this.ctx.isTypeScript()) {
                        const typeAnnotation = p.typeAnnotation;
                        delete p.typeAnnotation;
                        p.range[1] = typeAnnotation.range[0];
                        p.loc.end = {
                            line: typeAnnotation.loc.start.line,
                            column: typeAnnotation.loc.start.column,
                        };
                        (0, scope_1.removeAllScopeAndVariableAndReference)(typeAnnotation, result);
                    }
                    p.parent = sortedParams[index].parent;
                }
                // Process for scope
                result.registerNodeToScope(block, scope);
                for (const v of scope.variables) {
                    for (const def of v.defs) {
                        if (def.node === fn) {
                            def.node = block;
                        }
                    }
                }
                tokens.shift(); // (
                tokens.pop(); // )
                tokens.pop(); // =>
                tokens.pop(); // {
                tokens.pop(); // }
                tokens.pop(); // ;
                this.remapNodes(maps.map((m) => {
                    return {
                        offset: m.offset,
                        range: m.range,
                        newNode: fn.params[m.index],
                    };
                }), tokens, comments, result.visitorKeys);
                // Disconnect the tree structure.
                exprSt.expression = null;
            });
            this.pushScope(restore, "};", this.currentScriptScopeKind);
        }
    }
    closeScope() {
        this.closeScopeCallbacks.pop()();
    }
    addProgramRestore(callback) {
        this.programRestoreCallbacks.push(callback);
    }
    appendScript(text, offset, kind, nodeType, callback) {
        const resultCallback = this.appendScriptWithoutOffset(text, kind, nodeType, (node, tokens, comments, result) => {
            (0, fix_locations_1.fixLocations)(node, tokens, comments, offset - resultCallback.start, result.visitorKeys, this.ctx);
            callback(node, tokens, comments, result);
        });
        return resultCallback;
    }
    appendScriptWithoutOffset(text, kind, nodeType, callback) {
        const { start: startOffset, end: endOffset } = this.script.addLet(text, kind);
        const restoreCallback = {
            start: startOffset,
            end: endOffset,
            nodeType,
            callback,
        };
        this.restoreCallbacks.push(restoreCallback);
        return restoreCallback;
    }
    pushScope(restoreCallback, closeToken, kind) {
        const upper = this.currentScriptScopeKind;
        this.currentScriptScopeKind = kind;
        this.closeScopeCallbacks.push(() => {
            this.script.addLet(closeToken, kind);
            this.currentScriptScopeKind = upper;
            restoreCallback.end = this.script.getCurrentVirtualCodeLength();
        });
    }
    /**
     * Restore AST nodes
     */
    restore(result) {
        const nodeToScope = getNodeToScope(result.scopeManager);
        const postprocessList = [];
        const callbackOption = {
            getScope,
            registerNodeToScope,
            scopeManager: result.scopeManager,
            visitorKeys: result.visitorKeys,
            addPostProcess: (cb) => postprocessList.push(cb),
        };
        this.restoreNodes(result, callbackOption);
        this.restoreProgram(result, callbackOption);
        postprocessList.forEach((p) => p());
        // Helpers
        /** Get scope */
        function getScope(node) {
            return (0, scope_1.getScopeFromNode)(result.scopeManager, node);
        }
        /** Register node to scope */
        function registerNodeToScope(node, scope) {
            // If we replace the `scope.block` at this time,
            // the scope restore calculation will not work, so we will replace the `scope.block` later.
            postprocessList.push(() => {
                var _a, _b;
                const beforeBlock = scope.block;
                scope.block = node;
                for (const variable of [
                    ...scope.variables,
                    ...((_b = (_a = scope.upper) === null || _a === void 0 ? void 0 : _a.variables) !== null && _b !== void 0 ? _b : []),
                ]) {
                    for (const def of variable.defs) {
                        if (def.node === beforeBlock) {
                            def.node = node;
                        }
                    }
                }
            });
            const scopes = nodeToScope.get(node);
            if (scopes) {
                scopes.push(scope);
            }
            else {
                nodeToScope.set(node, [scope]);
            }
        }
    }
    /**
     * Restore AST nodes
     */
    restoreNodes(result, callbackOption) {
        let orderedRestoreCallback = this.restoreCallbacks.shift();
        if (!orderedRestoreCallback) {
            return;
        }
        const separateIndexes = this.script.separateIndexes;
        const tokens = result.ast.tokens;
        const processedTokens = [];
        const comments = result.ast.comments;
        const processedComments = [];
        let tok;
        while ((tok = tokens.shift())) {
            if (separateIndexes.includes(tok.range[0]) && tok.value === ";") {
                break;
            }
            if (orderedRestoreCallback.start <= tok.range[0]) {
                tokens.unshift(tok);
                break;
            }
            processedTokens.push(tok);
        }
        while ((tok = comments.shift())) {
            if (orderedRestoreCallback.start <= tok.range[0]) {
                comments.unshift(tok);
                break;
            }
            processedComments.push(tok);
        }
        const targetNodes = new Map();
        const removeStatements = [];
        (0, traverse_1.traverseNodes)(result.ast, {
            visitorKeys: result.visitorKeys,
            enterNode: (node) => {
                while (node.range && separateIndexes.includes(node.range[1] - 1)) {
                    node.range[1]--;
                    node.loc.end.column--;
                }
                if (node.loc.end.column < 0) {
                    node.loc.end = this.ctx.getLocFromIndex(node.range[1]);
                }
                if (node.parent === result.ast &&
                    separateIndexes[0] <= node.range[0]) {
                    removeStatements.push(node);
                }
                if (!orderedRestoreCallback) {
                    return;
                }
                if (orderedRestoreCallback.nodeType === node.type &&
                    orderedRestoreCallback.start <= node.range[0] &&
                    node.range[1] <= orderedRestoreCallback.end) {
                    targetNodes.set(node, orderedRestoreCallback);
                    orderedRestoreCallback = this.restoreCallbacks.shift();
                }
                //
            },
            leaveNode(node) {
                const restoreCallback = targetNodes.get(node);
                if (!restoreCallback) {
                    return;
                }
                const startIndex = {
                    token: tokens.findIndex((t) => restoreCallback.start <= t.range[0]),
                    comment: comments.findIndex((t) => restoreCallback.start <= t.range[0]),
                };
                if (startIndex.comment === -1) {
                    startIndex.comment = comments.length;
                }
                const endIndex = {
                    token: tokens.findIndex((t) => restoreCallback.end < t.range[1], startIndex.token),
                    comment: comments.findIndex((t) => restoreCallback.end < t.range[1], startIndex.comment),
                };
                if (endIndex.token === -1) {
                    endIndex.token = tokens.length;
                }
                if (endIndex.comment === -1) {
                    endIndex.comment = comments.length;
                }
                const targetTokens = tokens.splice(startIndex.token, endIndex.token - startIndex.token);
                const targetComments = comments.splice(startIndex.comment, endIndex.comment - startIndex.comment);
                restoreCallback.callback(node, targetTokens, targetComments, callbackOption);
                processedTokens.push(...targetTokens);
                processedComments.push(...targetComments);
            },
        });
        for (const st of removeStatements) {
            const index = result.ast.body.indexOf(st);
            result.ast.body.splice(index, 1);
        }
        result.ast.tokens = processedTokens;
        result.ast.comments = processedComments;
    }
    /**
     * Restore program node
     */
    restoreProgram(result, callbackOption) {
        for (const callback of this.programRestoreCallbacks) {
            callback(result.ast, result.ast.tokens, result.ast.comments, callbackOption);
        }
    }
    remapNodes(maps, tokens, comments, visitorKeys) {
        const targetMaps = [...maps];
        const first = targetMaps.shift();
        let tokenIndex = 0;
        for (; tokenIndex < tokens.length; tokenIndex++) {
            const token = tokens[tokenIndex];
            if (first.range[1] <= token.range[0]) {
                break;
            }
        }
        for (const map of targetMaps) {
            const startOffset = map.offset - first.offset + first.range[0];
            const endOffset = startOffset + (map.range[1] - map.range[0]);
            let removeCount = 0;
            let target = tokens[tokenIndex];
            while (target && target.range[1] <= startOffset) {
                removeCount++;
                target = tokens[tokenIndex + removeCount];
            }
            if (removeCount) {
                tokens.splice(tokenIndex, removeCount);
            }
            const bufferTokens = [];
            for (; tokenIndex < tokens.length; tokenIndex++) {
                const token = tokens[tokenIndex];
                if (endOffset <= token.range[0]) {
                    break;
                }
                bufferTokens.push(token);
            }
            (0, fix_locations_1.fixLocations)(map.newNode, bufferTokens, comments.filter((t) => startOffset <= t.range[0] && t.range[1] <= endOffset), map.range[0] - startOffset, visitorKeys, this.ctx);
        }
        tokens.splice(tokenIndex);
    }
    generateUniqueId(base) {
        return this.unique.generate(base, this.ctx.code, this.script.getCurrentVirtualCode());
    }
}
exports.ScriptLetContext = ScriptLetContext;
function getTypeParameterRanges(code, param) {
    var _a;
    const idRange = [
        param.range[0],
        param.constraint || param.default ? param.name.range[1] : param.range[1],
    ];
    let constraintRange = null;
    let defaultRange = null;
    if (param.constraint) {
        constraintRange = [
            param.constraint.range[0],
            param.default ? param.constraint.range[1] : param.range[1],
        ];
        const index = getTokenIndex(code, (code, index) => code.startsWith("extends", index), idRange[1], param.constraint.range[0]);
        if (index != null) {
            idRange[1] = index;
            constraintRange[0] = index + 7;
        }
    }
    if (param.default) {
        defaultRange = [param.default.range[0], param.range[1]];
        const index = getTokenIndex(code, (code, index) => code[index] === "=", (_a = constraintRange === null || constraintRange === void 0 ? void 0 : constraintRange[1]) !== null && _a !== void 0 ? _a : idRange[1], param.default.range[0]);
        if (index != null) {
            (constraintRange !== null && constraintRange !== void 0 ? constraintRange : idRange)[1] = index;
            defaultRange[0] = index + 1;
        }
    }
    return { idRange, constraintRange, defaultRange };
}
function getTokenIndex(code, targetToken, start, end) {
    let index = start;
    while (index < end) {
        if (targetToken(code, index)) {
            return index;
        }
        if (code.startsWith("//", index)) {
            const lfIndex = code.indexOf("\n", index);
            if (lfIndex >= 0) {
                index = lfIndex + 1;
                continue;
            }
            return null;
        }
        if (code.startsWith("/*", index)) {
            const endIndex = code.indexOf("*/", index);
            if (endIndex >= 0) {
                index = endIndex + 2;
                continue;
            }
            return null;
        }
        index++;
    }
    return null;
}
/** Get the node to scope map from given scope manager  */
function getNodeToScope(scopeManager) {
    if ("__nodeToScope" in scopeManager) {
        return scopeManager.__nodeToScope;
    }
    // transform scopeManager
    const nodeToScope = new WeakMap();
    for (const scope of scopeManager.scopes) {
        const scopes = nodeToScope.get(scope.block);
        if (scopes) {
            scopes.push(scope);
        }
        else {
            nodeToScope.set(scope.block, [scope]);
        }
    }
    scopeManager.acquire = function (node, inner) {
        /**
         * predicate
         */
        function predicate(testScope) {
            if (testScope.type === "function" && testScope.functionExpressionScope) {
                return false;
            }
            return true;
        }
        const scopes = nodeToScope.get(node);
        if (!scopes || scopes.length === 0) {
            return null;
        }
        // Heuristic selection from all scopes.
        // If you would like to get all scopes, please use ScopeManager#acquireAll.
        if (scopes.length === 1) {
            return scopes[0];
        }
        if (inner) {
            for (let i = scopes.length - 1; i >= 0; --i) {
                const scope = scopes[i];
                if (predicate(scope)) {
                    return scope;
                }
            }
        }
        else {
            for (let i = 0, iz = scopes.length; i < iz; ++i) {
                const scope = scopes[i];
                if (predicate(scope)) {
                    return scope;
                }
            }
        }
        return null;
    };
    return nodeToScope;
}
/** Extract the type scope of the given node. */
function extractTypeNodeScopes(node, result) {
    const scopes = new Set();
    for (const scope of iterateTypeNodeScopes(node)) {
        scopes.add(scope);
    }
    return scopes;
    /** Iterate the type scope of the given node. */
    function* iterateTypeNodeScopes(node) {
        if (node.type === "TSParenthesizedType") {
            // Skip TSParenthesizedType.
            yield* iterateTypeNodeScopes(node.typeAnnotation);
        }
        else if (node.type === "TSConditionalType") {
            yield result.getScope(node);
            // `falseType` of `TSConditionalType` is sibling scope.
            const falseType = node.falseType;
            yield* iterateTypeNodeScopes(falseType);
        }
        else if (node.type === "TSFunctionType" ||
            node.type === "TSMappedType" ||
            node.type === "TSConstructorType") {
            yield result.getScope(node);
        }
        else {
            const typeNode = node;
            for (const key of (0, traverse_1.getKeys)(typeNode, result.visitorKeys)) {
                for (const child of (0, traverse_1.getNodes)(typeNode, key)) {
                    yield* iterateTypeNodeScopes(child);
                }
            }
        }
    }
}
