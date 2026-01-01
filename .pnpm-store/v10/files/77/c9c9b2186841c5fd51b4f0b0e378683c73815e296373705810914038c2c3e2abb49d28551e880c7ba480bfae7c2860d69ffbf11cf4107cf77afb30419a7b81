"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.parseForESLint = parseForESLint;
const visitor_keys_1 = require("../visitor-keys");
const context_1 = require("../context");
const eslint_scope_1 = require("eslint-scope");
const script_1 = require("./script");
const sort_1 = require("./sort");
const template_1 = require("./template");
const analyze_scope_1 = require("./analyze-scope");
const errors_1 = require("../errors");
const typescript_1 = require("./typescript");
const scope_1 = require("../scope");
const style_context_1 = require("./style-context");
const globals_1 = require("./globals");
const parser_options_1 = require("./parser-options");
const compat_1 = require("./compat");
const svelte_parse_context_1 = require("./svelte-parse-context");
const svelte_config_1 = require("../svelte-config");
/**
 * Parse source code
 */
function parseForESLint(code, options) {
    const svelteConfig = (0, svelte_config_1.resolveSvelteConfigFromOption)(options);
    const parserOptions = (0, parser_options_1.normalizeParserOptions)(options);
    if ((0, svelte_parse_context_1.isEnableRunes)(svelteConfig, parserOptions) &&
        parserOptions.filePath &&
        !parserOptions.filePath.endsWith(".svelte") &&
        // If no `filePath` is set in ESLint, "<input>" will be specified.
        parserOptions.filePath !== "<input>") {
        const trimmed = code.trim();
        if (!trimmed.startsWith("<") && !trimmed.endsWith(">")) {
            const svelteParseContext = (0, svelte_parse_context_1.resolveSvelteParseContextForSvelteScript)(svelteConfig, parserOptions);
            return parseAsScript(code, parserOptions, svelteParseContext);
        }
    }
    return parseAsSvelte(code, svelteConfig, parserOptions);
}
/**
 * Parse source code as svelte component
 */
function parseAsSvelte(code, svelteConfig, parserOptions) {
    const ctx = new context_1.Context(code, parserOptions);
    const resultTemplate = (0, template_1.parseTemplate)(ctx.sourceCode.template, ctx, parserOptions);
    const svelteParseContext = (0, svelte_parse_context_1.resolveSvelteParseContextForSvelte)(svelteConfig, parserOptions, resultTemplate.svelteAst);
    const scripts = ctx.sourceCode.scripts;
    const resultScript = ctx.isTypeScript()
        ? (0, typescript_1.parseTypeScriptInSvelte)(scripts.getCurrentVirtualCodeInfo(), scripts.attrs, parserOptions, { slots: ctx.slots, svelteParseContext })
        : (0, script_1.parseScriptInSvelte)(scripts.getCurrentVirtualCode(), scripts.attrs, parserOptions);
    ctx.scriptLet.restore(resultScript);
    ctx.tokens.push(...resultScript.ast.tokens);
    ctx.comments.push(...resultScript.ast.comments);
    (0, sort_1.sortNodes)(ctx.comments);
    (0, sort_1.sortNodes)(ctx.tokens);
    extractTokens(ctx);
    (0, analyze_scope_1.analyzeStoreScope)(resultScript.scopeManager);
    (0, analyze_scope_1.analyzeReactiveScope)(resultScript.scopeManager);
    (0, analyze_scope_1.analyzeStoreScope)(resultScript.scopeManager); // for reactive vars
    (0, analyze_scope_1.analyzeSnippetsScope)(ctx.snippets, resultScript.scopeManager);
    // Add $$xxx variable
    addGlobalVariables(resultScript.scopeManager, (0, globals_1.getGlobalsForSvelte)(svelteParseContext));
    const ast = resultTemplate.ast;
    const statements = [...resultScript.ast.body];
    ast.sourceType = resultScript.ast.sourceType;
    const scriptElements = ast.body.filter((b) => b.type === "SvelteScriptElement");
    for (let index = 0; index < scriptElements.length; index++) {
        const body = scriptElements[index];
        let statement = statements[0];
        while (statement &&
            body.range[0] <= statement.range[0] &&
            (statement.range[1] <= body.range[1] ||
                index === scriptElements.length - 1)) {
            statement.parent = body;
            body.body.push(statement);
            statements.shift();
            statement = statements[0];
        }
        if (!body.startTag.attributes.some((attr) => attr.type === "SvelteAttribute" &&
            attr.key.name === "context" &&
            attr.value.length === 1 &&
            attr.value[0].type === "SvelteLiteral" &&
            attr.value[0].value === "module")) {
            (0, analyze_scope_1.analyzePropsScope)(body, resultScript.scopeManager, svelteParseContext);
        }
    }
    if (statements.length) {
        throw new errors_1.ParseError("The script is unterminated", statements[0].range[1], ctx);
    }
    const styleElement = ast.body.find((b) => b.type === "SvelteStyleElement");
    let styleContext = null;
    resultScript.ast = ast;
    resultScript.services = Object.assign(resultScript.services || {}, {
        isSvelte: true,
        isSvelteScript: false,
        getSvelteHtmlAst() {
            return (0, compat_1.getFragmentFromRoot)(resultTemplate.svelteAst);
        },
        getStyleContext() {
            if (styleContext === null) {
                styleContext = (0, style_context_1.parseStyleContext)(styleElement, ctx);
            }
            return styleContext;
        },
        styleNodeLoc: style_context_1.styleNodeLoc,
        styleNodeRange: style_context_1.styleNodeRange,
        svelteParseContext,
    });
    resultScript.visitorKeys = Object.assign({}, visitor_keys_1.KEYS, resultScript.visitorKeys);
    return resultScript;
}
/**
 * Parse source code as script
 */
function parseAsScript(code, parserOptions, svelteParseContext) {
    var _a;
    const lang = (_a = parserOptions.filePath) === null || _a === void 0 ? void 0 : _a.split(".").pop();
    const resultScript = (0, parser_options_1.isTypeScript)(parserOptions, lang)
        ? (0, typescript_1.parseTypeScript)(code, { lang }, parserOptions, svelteParseContext)
        : (0, script_1.parseScript)(code, { lang }, parserOptions);
    // Add $$xxx variable
    addGlobalVariables(resultScript.scopeManager, (0, globals_1.getGlobalsForSvelteScript)(svelteParseContext));
    resultScript.services = Object.assign(resultScript.services || {}, {
        isSvelte: false,
        isSvelteScript: true,
        svelteParseContext,
    });
    resultScript.visitorKeys = Object.assign({}, visitor_keys_1.KEYS, resultScript.visitorKeys);
    return resultScript;
}
function addGlobalVariables(scopeManager, globals) {
    const globalScope = scopeManager.globalScope;
    for (const globalName of globals) {
        if (globalScope.set.has(globalName))
            continue;
        const variable = new eslint_scope_1.Variable();
        variable.name = globalName;
        variable.scope = globalScope;
        globalScope.variables.push(variable);
        globalScope.set.set(globalName, variable);
        globalScope.through = globalScope.through.filter((reference) => {
            if (reference.identifier.name === globalName) {
                // Links the variable and the reference.
                // And this reference is removed from `Scope#through`.
                reference.resolved = variable;
                (0, scope_1.addReference)(variable.references, reference);
                return false;
            }
            return true;
        });
    }
}
/** Extract tokens */
function extractTokens(ctx) {
    const useRanges = (0, sort_1.sortNodes)([...ctx.tokens, ...ctx.comments]).map((t) => t.range);
    let range = useRanges.shift();
    for (let index = 0; index < ctx.sourceCode.template.length; index++) {
        while (range && range[1] <= index) {
            range = useRanges.shift();
        }
        if (range && range[0] <= index) {
            index = range[1] - 1;
            continue;
        }
        const c = ctx.sourceCode.template[index];
        if (!c.trim()) {
            continue;
        }
        if (isPunctuator(c)) {
            ctx.addToken("Punctuator", { start: index, end: index + 1 });
        }
        else {
            // unknown
            // It is may be a bug.
            ctx.addToken("Identifier", { start: index, end: index + 1 });
        }
    }
    (0, sort_1.sortNodes)(ctx.comments);
    (0, sort_1.sortNodes)(ctx.tokens);
    /**
     * Checks if the given char is punctuator
     */
    function isPunctuator(c) {
        return /^[^\w$]$/iu.test(c);
    }
}
