"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.analyzeTypeScriptInSvelte = analyzeTypeScriptInSvelte;
exports.analyzeTypeScript = analyzeTypeScript;
const scope_1 = require("../../../scope");
const utils_1 = require("../../../utils");
const script_1 = require("../../script");
const context_1 = require("../context");
const set_parent_1 = require("../set-parent");
const globals_1 = require("../../globals");
/**
 * Analyze TypeScript source code in <script>.
 * Generate virtual code to provide correct type information for Svelte store reference names, scopes, and runes.
 * See https://github.com/sveltejs/svelte-eslint-parser/blob/main/docs/internal-mechanism.md#scope-types
 */
function analyzeTypeScriptInSvelte(code, attrs, parserOptions, context) {
    const ctx = new context_1.VirtualTypeScriptContext(code.script + code.render + code.rootScope);
    ctx.appendOriginal(/^\s*/u.exec(code.script)[0].length);
    const result = (0, script_1.parseScriptWithoutAnalyzeScope)(code.script + code.render + code.rootScope, attrs, Object.assign(Object.assign({}, parserOptions), { 
        // Without typings
        project: null }));
    ctx._beforeResult = result;
    analyzeStoreReferenceNames(result, context.svelteParseContext, ctx);
    analyzeDollarDollarVariables(result, ctx, context.svelteParseContext, context.slots);
    analyzeRuneVariables(result, ctx, context.svelteParseContext);
    applyTransforms([
        ...analyzeReactiveScopes(result),
        ...analyzeDollarDerivedScopes(result, context.svelteParseContext),
    ], ctx);
    analyzeRenderScopes(code, ctx);
    ctx.appendOriginalToEnd();
    return ctx;
}
/**
 * Analyze TypeScript source code.
 * Generate virtual code to provide correct type information for Svelte runes.
 * See https://github.com/sveltejs/svelte-eslint-parser/blob/main/docs/internal-mechanism.md#scope-types
 */
function analyzeTypeScript(code, attrs, parserOptions, svelteParseContext) {
    const ctx = new context_1.VirtualTypeScriptContext(code);
    ctx.appendOriginal(/^\s*/u.exec(code)[0].length);
    const result = (0, script_1.parseScriptWithoutAnalyzeScope)(code, attrs, Object.assign(Object.assign({}, parserOptions), { 
        // Without typings
        project: null }));
    ctx._beforeResult = result;
    analyzeRuneVariables(result, ctx, svelteParseContext);
    applyTransforms([...analyzeDollarDerivedScopes(result, svelteParseContext)], ctx);
    ctx.appendOriginalToEnd();
    return ctx;
}
/**
 * Analyze the store reference names.
 * Insert type definitions code to provide correct type information for variables that begin with `$`.
 */
function analyzeStoreReferenceNames(result, svelteParseContext, ctx) {
    const globals = (0, globals_1.getGlobalsForSvelte)(svelteParseContext);
    const scopeManager = result.scopeManager;
    const programScope = (0, scope_1.getProgramScope)(scopeManager);
    const maybeStoreRefNames = new Set();
    for (const reference of scopeManager.globalScope.through) {
        if (
        // Begin with `$`.
        reference.identifier.name.startsWith("$") &&
            // Ignore globals
            !globals.includes(reference.identifier.name) &&
            // Ignore if it is already defined.
            !programScope.set.has(reference.identifier.name)) {
            maybeStoreRefNames.add(reference.identifier.name);
        }
    }
    if (maybeStoreRefNames.size) {
        const storeValueTypeName = ctx.generateUniqueId("StoreValueType");
        ctx.appendVirtualScript(`type ${storeValueTypeName}<T> = T extends null | undefined
? T
: T extends object & { subscribe(run: infer F, ...args: any): any }
? F extends (value: infer V, ...args: any) => any
? V
: never
: T;`);
        ctx.restoreContext.addRestoreStatementProcess((node, result) => {
            if (node.type !== "TSTypeAliasDeclaration" ||
                node.id.name !== storeValueTypeName) {
                return false;
            }
            const program = result.ast;
            program.body.splice(program.body.indexOf(node), 1);
            const scopeManager = result.scopeManager;
            // Remove `type` scope
            (0, scope_1.removeAllScopeAndVariableAndReference)(node, {
                visitorKeys: result.visitorKeys,
                scopeManager,
            });
            return true;
        });
        for (const nm of maybeStoreRefNames) {
            const realName = nm.slice(1);
            ctx.appendVirtualScript(`declare let ${nm}: ${storeValueTypeName}<typeof ${realName}>;`);
            ctx.restoreContext.addRestoreStatementProcess((node, result) => {
                if (node.type !== "VariableDeclaration" ||
                    !node.declare ||
                    node.declarations.length !== 1 ||
                    node.declarations[0].id.type !== "Identifier" ||
                    node.declarations[0].id.name !== nm) {
                    return false;
                }
                const program = result.ast;
                program.body.splice(program.body.indexOf(node), 1);
                const scopeManager = result.scopeManager;
                // Remove `declare` variable
                (0, scope_1.removeAllScopeAndVariableAndReference)(node, {
                    visitorKeys: result.visitorKeys,
                    scopeManager,
                });
                return true;
            });
        }
    }
}
/**
 * Analyze `$$slots`, `$$props`, and `$$restProps` .
 * Insert type definitions code to provide correct type information for `$$slots`, `$$props`, and `$$restProps`.
 */
function analyzeDollarDollarVariables(result, ctx, svelteParseContext, slots) {
    const globals = (0, globals_1.getGlobalsForSvelte)(svelteParseContext);
    const scopeManager = result.scopeManager;
    for (const globalName of globals) {
        if (!scopeManager.globalScope.through.some((reference) => reference.identifier.name === globalName)) {
            continue;
        }
        switch (globalName) {
            case "$$props":
                appendDeclareVirtualScript(globalName, `{ [index: string]: any }`);
                break;
            case "$$restProps":
                appendDeclareVirtualScript(globalName, `{ [index: string]: any }`);
                break;
            case "$$slots": {
                const nameTypes = new Set();
                for (const slot of slots) {
                    const nameAttr = slot.startTag.attributes.find((attr) => attr.type === "SvelteAttribute" && attr.key.name === "name");
                    if (!nameAttr || nameAttr.value.length === 0) {
                        nameTypes.add('"default"');
                        continue;
                    }
                    if (nameAttr.value.length === 1) {
                        const value = nameAttr.value[0];
                        if (value.type === "SvelteLiteral") {
                            nameTypes.add(JSON.stringify(value.value));
                        }
                        else {
                            nameTypes.add("string");
                        }
                        continue;
                    }
                    nameTypes.add(`\`${nameAttr.value
                        .map((value) => value.type === "SvelteLiteral"
                        ? value.value.replace(/([$`])/gu, "\\$1")
                        : "${string}")
                        .join("")}\``);
                }
                appendDeclareVirtualScript(globalName, `Record<${nameTypes.size > 0 ? [...nameTypes].join(" | ") : "any"}, boolean>`);
                break;
            }
            case "$state":
            case "$derived":
            case "$effect":
            case "$props":
            case "$bindable":
            case "$inspect":
            case "$host":
                // Processed by `analyzeRuneVariables`.
                break;
            default: {
                const _ = globalName;
                throw Error(`Unknown global: ${_}`);
            }
        }
    }
    /** Append declare virtual script */
    function appendDeclareVirtualScript(name, type) {
        ctx.appendVirtualScript(`declare let ${name}: ${type};`);
        ctx.restoreContext.addRestoreStatementProcess((node, result) => {
            if (node.type !== "VariableDeclaration" ||
                !node.declare ||
                node.declarations.length !== 1 ||
                node.declarations[0].id.type !== "Identifier" ||
                node.declarations[0].id.name !== name) {
                return false;
            }
            const program = result.ast;
            program.body.splice(program.body.indexOf(node), 1);
            const scopeManager = result.scopeManager;
            // Remove `declare` variable
            (0, scope_1.removeAllScopeAndVariableAndReference)(node, {
                visitorKeys: result.visitorKeys,
                scopeManager,
            });
            return true;
        });
    }
}
/**
 * Analyze Runes.
 * Insert type definitions code to provide correct type information for Runes.
 */
function analyzeRuneVariables(result, ctx, svelteParseContext) {
    if (!svelteParseContext.runes)
        return;
    const scopeManager = result.scopeManager;
    for (const globalName of globals_1.globalsForRunes) {
        if (!scopeManager.globalScope.through.some((reference) => reference.identifier.name === globalName)) {
            continue;
        }
        switch (globalName) {
            // See https://github.com/sveltejs/svelte/blob/41b5cd6f5daae3970a9927e062f42b6b62440d16/packages/svelte/types/index.d.ts#L2299
            case "$state": {
                appendDeclareFunctionVirtualScripts(globalName, [
                    "<T>(initial: T): T",
                    "<T>(): T | undefined",
                ]);
                appendDeclareNamespaceVirtualScripts(globalName, [
                    "export function raw<T>(initial: T): T;",
                    "export function raw<T>(): T | undefined;",
                ]);
                appendDeclareNamespaceVirtualScripts(globalName, [
                    "export function snapshot<T>(state: T): T;",
                ]);
                break;
            }
            // See https://github.com/sveltejs/svelte/blob/41b5cd6f5daae3970a9927e062f42b6b62440d16/packages/svelte/types/index.d.ts#L2453
            case "$derived": {
                appendDeclareFunctionVirtualScripts(globalName, [
                    "<T>(expression: T): T",
                ]);
                appendDeclareNamespaceVirtualScripts(globalName, [
                    "export function by<T>(fn: () => T): T;",
                ]);
                break;
            }
            // See https://github.com/sveltejs/svelte/blob/41b5cd6f5daae3970a9927e062f42b6b62440d16/packages/svelte/types/index.d.ts#L2513
            case "$effect": {
                appendDeclareFunctionVirtualScripts(globalName, [
                    "(fn: () => void | (() => void)): void",
                ]);
                appendDeclareNamespaceVirtualScripts(globalName, [
                    "export function pre(fn: () => void | (() => void)): void;",
                    "export function tracking(): boolean;",
                    "export function root(fn: () => void | (() => void)): () => void;",
                ]);
                break;
            }
            // See https://github.com/sveltejs/svelte/blob/41b5cd6f5daae3970a9927e062f42b6b62440d16/packages/svelte/types/index.d.ts#L2615
            case "$props": {
                // Use type parameters to avoid `@typescript-eslint/no-unsafe-assignment` errors.
                appendDeclareFunctionVirtualScripts(globalName, ["<T>(): T"]);
                break;
            }
            // See https://github.com/sveltejs/svelte/blob/41b5cd6f5daae3970a9927e062f42b6b62440d16/packages/svelte/types/index.d.ts#L2626
            case "$bindable": {
                appendDeclareFunctionVirtualScripts(globalName, [
                    "<T>(fallback?: T): T",
                ]);
                break;
            }
            // See https://github.com/sveltejs/svelte/blob/41b5cd6f5daae3970a9927e062f42b6b62440d16/packages/svelte/types/index.d.ts#L2646
            case "$inspect": {
                appendDeclareFunctionVirtualScripts(globalName, [
                    `<T extends any[]>(...values: T): { with: (fn: (type: 'init' | 'update', ...values: T) => void) => void }`,
                ]);
                break;
            }
            // See https://github.com/sveltejs/svelte/blob/41b5cd6f5daae3970a9927e062f42b6b62440d16/packages/svelte/types/index.d.ts#L2669
            case "$host": {
                appendDeclareFunctionVirtualScripts(globalName, [
                    `<El extends HTMLElement = HTMLElement>(): El`,
                ]);
                break;
            }
            default: {
                const _ = globalName;
                throw Error(`Unknown global: ${_}`);
            }
        }
    }
    /** Append declare virtual script */
    function appendDeclareFunctionVirtualScripts(name, types) {
        for (const type of types) {
            ctx.appendVirtualScript(`declare function ${name}${type};`);
            ctx.restoreContext.addRestoreStatementProcess((node, result) => {
                var _a;
                if (node.type !== "TSDeclareFunction" ||
                    !node.declare ||
                    ((_a = node.id) === null || _a === void 0 ? void 0 : _a.type) !== "Identifier" ||
                    node.id.name !== name) {
                    return false;
                }
                const program = result.ast;
                program.body.splice(program.body.indexOf(node), 1);
                const scopeManager = result.scopeManager;
                // Remove `declare` variable
                (0, scope_1.removeAllScopeAndVariableAndReference)(node, {
                    visitorKeys: result.visitorKeys,
                    scopeManager,
                });
                return true;
            });
        }
    }
    function appendDeclareNamespaceVirtualScripts(name, scripts) {
        for (const script of scripts) {
            ctx.appendVirtualScript(`declare namespace ${name} { ${script} }`);
            ctx.restoreContext.addRestoreStatementProcess((node, result) => {
                var _a;
                if (node.type !== "TSModuleDeclaration" ||
                    !node.declare ||
                    ((_a = node.id) === null || _a === void 0 ? void 0 : _a.type) !== "Identifier" ||
                    node.id.name !== name) {
                    return false;
                }
                const program = result.ast;
                program.body.splice(program.body.indexOf(node), 1);
                const scopeManager = result.scopeManager;
                // Remove `declare` variable
                (0, scope_1.removeAllScopeAndVariableAndReference)(node, {
                    visitorKeys: result.visitorKeys,
                    scopeManager,
                });
                return true;
            });
        }
    }
}
/**
 * Analyze the reactive scopes.
 * Transform source code to provide the correct type information in the `$:` statements.
 */
function* analyzeReactiveScopes(result) {
    const scopeManager = result.scopeManager;
    const throughIds = scopeManager.globalScope.through.map((reference) => reference.identifier);
    for (const statement of result.ast.body) {
        if (statement.type === "LabeledStatement" && statement.label.name === "$") {
            if (statement.body.type === "ExpressionStatement" &&
                statement.body.expression.type === "AssignmentExpression" &&
                statement.body.expression.operator === "=" &&
                // Must be a pattern that can be used in the LHS of variable declarations.
                // https://github.com/sveltejs/svelte-eslint-parser/issues/213
                (statement.body.expression.left.type === "Identifier" ||
                    statement.body.expression.left.type === "ArrayPattern" ||
                    statement.body.expression.left.type === "ObjectPattern")) {
                const left = statement.body.expression.left;
                if (throughIds.some((id) => left.range[0] <= id.range[0] && id.range[1] <= left.range[1])) {
                    const node = statement;
                    const expression = statement.body.expression;
                    yield {
                        node,
                        transform: (ctx) => transformForDeclareReactiveVar(node, left, expression, result.ast.tokens, ctx),
                    };
                    continue;
                }
            }
            yield {
                node: statement,
                transform: (ctx) => transformForReactiveStatement(statement, ctx),
            };
        }
    }
}
/**
 * Analyze the $derived scopes.
 * Transform source code to provide the correct type information in the `$derived(...)` expression.
 */
function* analyzeDollarDerivedScopes(result, svelteParseContext) {
    var _a;
    if (!svelteParseContext.runes)
        return;
    const scopeManager = result.scopeManager;
    const derivedReferences = scopeManager.globalScope.through.filter((reference) => reference.identifier.name === "$derived");
    if (!derivedReferences.length) {
        return;
    }
    (0, set_parent_1.setParent)(result);
    for (const ref of derivedReferences) {
        const derived = ref.identifier;
        if (derived.parent.type === "CallExpression" &&
            derived.parent.callee === derived &&
            ((_a = derived.parent.arguments[0]) === null || _a === void 0 ? void 0 : _a.type) !== "SpreadElement") {
            const node = derived.parent;
            yield {
                node,
                transform: (ctx) => transformForDollarDerived(node, ctx),
            };
        }
    }
}
/**
 * Analyze the render scopes.
 * Transform source code to provide the correct type information in the HTML templates.
 */
function analyzeRenderScopes(code, ctx) {
    ctx.appendOriginal(code.script.length);
    const renderFunctionName = ctx.generateUniqueId("render");
    ctx.appendVirtualScript(`function ${renderFunctionName}(){`);
    ctx.appendOriginal(code.script.length + code.render.length);
    ctx.appendVirtualScript(`}`);
    ctx.restoreContext.addRestoreStatementProcess((node, result) => {
        if (node.type !== "FunctionDeclaration" ||
            node.id.name !== renderFunctionName) {
            return false;
        }
        const program = result.ast;
        program.body.splice(program.body.indexOf(node), 1, ...node.body.body);
        for (const body of node.body.body) {
            body.parent = program;
        }
        const scopeManager = result.scopeManager;
        removeFunctionScope(node, scopeManager);
        return true;
    });
}
/**
 * Applies the given transforms.
 * Note that intersecting transformations are not applied.
 */
function applyTransforms(transforms, ctx) {
    transforms.sort((a, b) => a.node.range[0] - b.node.range[0]);
    let offset = 0;
    for (const transform of transforms) {
        const range = transform.node.range;
        if (offset <= range[0]) {
            transform.transform(ctx);
        }
        offset = range[1];
    }
}
/**
 * Transform for `$: id = ...` to `$: let id = ...`
 */
function transformForDeclareReactiveVar(statement, id, expression, tokens, ctx) {
    // e.g.
    //  From:
    //  $: id = x + y;
    //
    //  To:
    //  $: let id = fn()
    //  function fn () { let tmp; return (tmp = x + y); }
    //
    //
    //  From:
    //  $: ({id} = foo);
    //
    //  To:
    //  $: let {id} = fn()
    //  function fn () { let tmp; return (tmp = foo); }
    /**
     * The opening paren tokens for
     * `$: ({id} = foo);`
     *     ^
     */
    const openParens = [];
    /**
     * The equal token for
     * `$: ({id} = foo);`
     *           ^
     */
    let eq = null;
    /**
     * The closing paren tokens for
     * `$: ({id} = foo);`
     *                ^
     */
    const closeParens = [];
    /**
     * The closing paren token for
     * `$: id = (foo);`
     *              ^
     */
    let expressionCloseParen = null;
    const startIndex = (0, utils_1.sortedLastIndex)(tokens, (target) => target.range[0] - statement.range[0]);
    for (let index = startIndex; index < tokens.length; index++) {
        const token = tokens[index];
        if (statement.range[1] <= token.range[0]) {
            break;
        }
        if (token.range[1] <= statement.range[0]) {
            continue;
        }
        if (token.value === "(" && token.range[1] <= expression.range[0]) {
            openParens.push(token);
        }
        if (token.value === "=" &&
            expression.left.range[1] <= token.range[0] &&
            token.range[1] <= expression.right.range[0]) {
            eq = token;
        }
        if (token.value === ")") {
            if (expression.range[1] <= token.range[0]) {
                closeParens.push(token);
            }
            else if (expression.right.range[1] <= token.range[0]) {
                expressionCloseParen = token;
            }
        }
    }
    const functionId = ctx.generateUniqueId("reactiveVariableScopeFunction");
    const tmpVarId = ctx.generateUniqueId("tmpVar");
    for (const token of openParens) {
        ctx.appendOriginal(token.range[0]);
        ctx.skipOriginalOffset(token.range[1] - token.range[0]);
    }
    ctx.appendOriginal(expression.range[0]);
    ctx.skipUntilOriginalOffset(id.range[0]);
    ctx.appendVirtualScript("let ");
    ctx.appendOriginal(eq ? eq.range[1] : expression.right.range[0]);
    ctx.appendVirtualScript(`${functionId}();\nfunction ${functionId}(){let ${tmpVarId};return (${tmpVarId} = `);
    ctx.appendOriginal(expression.right.range[1]);
    ctx.appendVirtualScript(`)`);
    for (const token of closeParens) {
        ctx.appendOriginal(token.range[0]);
        ctx.skipOriginalOffset(token.range[1] - token.range[0]);
    }
    ctx.appendOriginal(statement.range[1]);
    ctx.appendVirtualScript(`}`);
    ctx.restoreContext.addRestoreStatementProcess((node, result) => {
        var _a;
        if (node.type !== "SvelteReactiveStatement") {
            return false;
        }
        const reactiveStatement = node;
        if (reactiveStatement.body.type !== "VariableDeclaration" ||
            reactiveStatement.body.kind !== "let" ||
            reactiveStatement.body.declarations.length !== 1) {
            return false;
        }
        const [idDecl] = reactiveStatement.body.declarations;
        if (idDecl.type !== "VariableDeclarator" ||
            idDecl.id.type !== id.type ||
            ((_a = idDecl.init) === null || _a === void 0 ? void 0 : _a.type) !== "CallExpression" ||
            idDecl.init.callee.type !== "Identifier" ||
            idDecl.init.callee.name !== functionId) {
            return false;
        }
        const program = result.ast;
        const nextIndex = program.body.indexOf(node) + 1;
        const fnDecl = program.body[nextIndex];
        if (!fnDecl ||
            fnDecl.type !== "FunctionDeclaration" ||
            fnDecl.id.name !== functionId ||
            fnDecl.body.body.length !== 2 ||
            fnDecl.body.body[0].type !== "VariableDeclaration" ||
            fnDecl.body.body[1].type !== "ReturnStatement") {
            return false;
        }
        const tmpVarDeclaration = fnDecl.body.body[0];
        if (tmpVarDeclaration.declarations.length !== 1 ||
            tmpVarDeclaration.declarations[0].type !== "VariableDeclarator") {
            return false;
        }
        const tempVarDeclId = tmpVarDeclaration.declarations[0].id;
        if (tempVarDeclId.type !== "Identifier" ||
            tempVarDeclId.name !== tmpVarId) {
            return false;
        }
        const returnStatement = fnDecl.body.body[1];
        const assignment = returnStatement.argument;
        if ((assignment === null || assignment === void 0 ? void 0 : assignment.type) !== "AssignmentExpression" ||
            assignment.left.type !== "Identifier" ||
            assignment.right.type !== expression.right.type) {
            return false;
        }
        const tempLeft = assignment.left;
        // Remove function declaration
        program.body.splice(nextIndex, 1);
        // Restore expression statement
        assignment.left = idDecl.id;
        assignment.loc = {
            start: idDecl.id.loc.start,
            end: expressionCloseParen
                ? expressionCloseParen.loc.end
                : assignment.right.loc.end,
        };
        assignment.range = [
            idDecl.id.range[0],
            expressionCloseParen
                ? expressionCloseParen.range[1]
                : assignment.right.range[1],
        ];
        idDecl.id.parent = assignment;
        const newBody = {
            type: "ExpressionStatement",
            expression: assignment,
            directive: undefined,
            loc: statement.body.loc,
            range: statement.body.range,
            parent: reactiveStatement,
        };
        assignment.parent = newBody;
        reactiveStatement.body = newBody;
        // Restore statement end location
        reactiveStatement.range[1] = returnStatement.range[1];
        reactiveStatement.loc.end.line = returnStatement.loc.end.line;
        reactiveStatement.loc.end.column = returnStatement.loc.end.column;
        // Restore tokens
        (0, utils_1.addElementsToSortedArray)(program.tokens, [...openParens, ...closeParens], (a, b) => a.range[0] - b.range[0]);
        const scopeManager = result.scopeManager;
        (0, scope_1.removeAllScopeAndVariableAndReference)(tmpVarDeclaration, {
            visitorKeys: result.visitorKeys,
            scopeManager,
        });
        removeFunctionScope(fnDecl, scopeManager);
        const scope = (0, scope_1.getProgramScope)(scopeManager);
        for (const reference of (0, scope_1.getAllReferences)(idDecl.id, scope)) {
            reference.writeExpr = assignment.right;
        }
        (0, scope_1.removeIdentifierReference)(tempLeft, scope);
        (0, scope_1.removeIdentifierVariable)(tempVarDeclId, scope);
        (0, scope_1.removeIdentifierReference)(idDecl.init.callee, scope);
        (0, scope_1.removeIdentifierVariable)(idDecl.id, scope);
        return true;
    });
}
/**
 * Transform for `$: ...` to `$: function foo(){...}`
 */
function transformForReactiveStatement(statement, ctx) {
    const functionId = ctx.generateUniqueId("reactiveStatementScopeFunction");
    const originalBody = statement.body;
    ctx.appendOriginal(originalBody.range[0]);
    ctx.appendVirtualScript(`function ${functionId}(){`);
    ctx.appendOriginal(originalBody.range[1]);
    ctx.appendVirtualScript(`}`);
    ctx.appendOriginal(statement.range[1]);
    ctx.restoreContext.addRestoreStatementProcess((node, result) => {
        if (node.type !== "SvelteReactiveStatement") {
            return false;
        }
        const reactiveStatement = node;
        const body = reactiveStatement.body;
        if (body.type !== "FunctionDeclaration" || body.id.name !== functionId) {
            return false;
        }
        reactiveStatement.body = body.body.body[0];
        reactiveStatement.body.parent = reactiveStatement;
        const scopeManager = result.scopeManager;
        removeFunctionScope(body, scopeManager);
        return true;
    });
}
/**
 * Transform for `$derived(expr)` to `$derived((()=>{ return fn(); function fn () { return expr } })())`
 */
function transformForDollarDerived(derivedCall, ctx) {
    const functionId = ctx.generateUniqueId("$derivedArgument");
    const expression = derivedCall.arguments[0];
    ctx.appendOriginal(expression.range[0]);
    ctx.appendVirtualScript(`(()=>{return ${functionId}();function ${functionId}(){return `);
    ctx.appendOriginal(expression.range[1]);
    ctx.appendVirtualScript(`}})()`);
    ctx.restoreContext.addRestoreExpressionProcess({
        target: "CallExpression",
        restore: (node, result) => {
            var _a;
            if (node.callee.type !== "Identifier" ||
                node.callee.name !== "$derived") {
                return false;
            }
            const arg = node.arguments[0];
            if (!arg ||
                arg.type !== "CallExpression" ||
                arg.arguments.length !== 0 ||
                arg.callee.type !== "ArrowFunctionExpression" ||
                arg.callee.body.type !== "BlockStatement" ||
                arg.callee.body.body.length !== 2 ||
                arg.callee.body.body[0].type !== "ReturnStatement" ||
                ((_a = arg.callee.body.body[0].argument) === null || _a === void 0 ? void 0 : _a.type) !== "CallExpression" ||
                arg.callee.body.body[0].argument.callee.type !== "Identifier" ||
                arg.callee.body.body[0].argument.callee.name !== functionId ||
                arg.callee.body.body[1].type !== "FunctionDeclaration" ||
                arg.callee.body.body[1].id.name !== functionId) {
                return false;
            }
            const fnNode = arg.callee.body.body[1];
            if (fnNode.body.body.length !== 1 ||
                fnNode.body.body[0].type !== "ReturnStatement" ||
                !fnNode.body.body[0].argument) {
                return false;
            }
            const expr = fnNode.body.body[0].argument;
            node.arguments[0] = expr;
            expr.parent = node;
            const scopeManager = result.scopeManager;
            removeFunctionScope(arg.callee.body.body[1], scopeManager);
            (0, scope_1.removeIdentifierReference)(arg.callee.body.body[0].argument.callee, scopeManager.acquire(arg.callee));
            removeFunctionScope(arg.callee, scopeManager);
            return true;
        },
    });
}
/** Remove function scope and marge child scopes to upper scope */
function removeFunctionScope(node, scopeManager) {
    const scope = scopeManager.acquire(node);
    const upper = scope.upper;
    // Remove render function variable
    if (node.id) {
        (0, scope_1.removeIdentifierVariable)(node.id, upper);
        (0, scope_1.removeIdentifierReference)(node.id, upper);
    }
    (0, scope_1.replaceScope)(scopeManager, scope, scope.childScopes);
    // Marge scope
    // * marge variables
    for (const variable of scope.variables) {
        if (variable.name === "arguments" && variable.defs.length === 0) {
            continue;
        }
        const upperVariable = upper.set.get(variable.name);
        if (upperVariable) {
            (0, utils_1.addElementsToSortedArray)(upperVariable.identifiers, variable.identifiers, (a, b) => a.range[0] - b.range[0]);
            (0, utils_1.addElementsToSortedArray)(upperVariable.defs, variable.defs, (a, b) => a.node.range[0] - b.node.range[0]);
            (0, scope_1.addAllReferences)(upperVariable.references, variable.references);
        }
        else {
            upper.set.set(variable.name, variable);
            (0, scope_1.addVariable)(upper.variables, variable);
            variable.scope = upper;
        }
        for (const reference of variable.references) {
            if (reference.from === scope) {
                reference.from = upper;
            }
            reference.resolved = upperVariable || variable;
        }
    }
    // * marge references
    (0, scope_1.addAllReferences)(upper.references, scope.references);
    for (const reference of scope.references) {
        if (reference.from === scope) {
            reference.from = upper;
        }
    }
}
