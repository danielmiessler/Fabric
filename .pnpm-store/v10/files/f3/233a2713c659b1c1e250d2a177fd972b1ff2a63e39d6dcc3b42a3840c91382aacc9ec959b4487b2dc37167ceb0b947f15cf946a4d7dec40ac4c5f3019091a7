"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.convertSvelteRoot = convertSvelteRoot;
const element_1 = require("./element");
const attr_1 = require("./attr");
const script_1 = require("../script");
const fix_locations_1 = require("../../context/fix-locations");
const compat_1 = require("../compat");
const sort_1 = require("../sort");
/**
 * Convert root
 */
function convertSvelteRoot(svelteAst, ctx) {
    var _a;
    const ast = Object.assign({ type: "Program", body: [], comments: ctx.comments, sourceType: "module", tokens: ctx.tokens, parent: null }, ctx.getConvertLocation({ start: 0, end: ctx.code.length }));
    const body = ast.body;
    const snippetChildren = [];
    const fragment = (0, compat_1.getFragmentFromRoot)(svelteAst);
    if (fragment) {
        let children = (0, compat_1.getChildren)(fragment);
        const options = (0, compat_1.getOptionsFromRoot)(svelteAst);
        if (options) {
            children = [...children];
            if (!children.some((node, idx) => {
                if (options.end <= node.start) {
                    children.splice(idx, 0, options);
                    return true;
                }
                return false;
            })) {
                children.push(options);
            }
        }
        const nonSnippetChildren = [];
        for (const child of children) {
            if (child.type === "SnippetBlock") {
                snippetChildren.push(child);
            }
            else {
                nonSnippetChildren.push(child);
            }
        }
        body.push(...(0, element_1.convertChildren)({ nodes: nonSnippetChildren }, ast, ctx));
    }
    let script = null;
    const instance = (0, compat_1.getInstanceFromRoot)(svelteAst);
    if (instance) {
        script = Object.assign({ type: "SvelteScriptElement", name: null, startTag: null, body: [], endTag: null, parent: ast }, ctx.getConvertLocation(instance));
        extractAttributes(script, ctx);
        (0, element_1.extractElementTags)(script, ctx, {
            buildNameNode: (openTokenRange) => {
                ctx.addToken("HTMLIdentifier", openTokenRange);
                const name = Object.assign({ type: "SvelteName", name: "script", parent: script }, ctx.getConvertLocation(openTokenRange));
                return name;
            },
        });
        body.push(script);
    }
    const module = (0, compat_1.getModuleFromRoot)(svelteAst);
    if (module) {
        const script = Object.assign({ type: "SvelteScriptElement", name: null, startTag: null, body: [], endTag: null, parent: ast }, ctx.getConvertLocation(module));
        extractAttributes(script, ctx);
        (0, element_1.extractElementTags)(script, ctx, {
            buildNameNode: (openTokenRange) => {
                ctx.addToken("HTMLIdentifier", openTokenRange);
                const name = Object.assign({ type: "SvelteName", name: "script", parent: script }, ctx.getConvertLocation(openTokenRange));
                return name;
            },
        });
        body.push(script);
    }
    if (svelteAst.css) {
        const style = Object.assign({ type: "SvelteStyleElement", name: null, startTag: null, children: [], endTag: null, parent: ast }, ctx.getConvertLocation(svelteAst.css));
        extractAttributes(style, ctx);
        (0, element_1.extractElementTags)(style, ctx, {
            buildNameNode: (openTokenRange) => {
                ctx.addToken("HTMLIdentifier", openTokenRange);
                const name = Object.assign({ type: "SvelteName", name: "style", parent: style }, ctx.getConvertLocation(openTokenRange));
                return name;
            },
        });
        if (style.endTag && style.startTag.range[1] < style.endTag.range[0]) {
            const contentRange = {
                start: style.startTag.range[1],
                end: style.endTag.range[0],
            };
            ctx.addToken("HTMLText", contentRange);
            style.children = [
                Object.assign({ type: "SvelteText", value: ctx.code.slice(contentRange.start, contentRange.end), parent: style }, ctx.getConvertLocation(contentRange)),
            ];
        }
        body.push(style);
    }
    body.push(...(0, element_1.convertChildren)({ nodes: snippetChildren }, ast, ctx));
    if (script && ((_a = ctx.parserOptions.svelteFeatures) === null || _a === void 0 ? void 0 : _a.experimentalGenerics))
        convertGenericsAttribute(script, ctx);
    // Set the scope of the Program node.
    ctx.scriptLet.addProgramRestore((node, _tokens, _comments, { scopeManager, registerNodeToScope, addPostProcess }) => {
        const scopes = [];
        for (const scope of scopeManager.scopes) {
            if (scope.block === node) {
                registerNodeToScope(ast, scope);
                scopes.push(scope);
            }
        }
        addPostProcess(() => {
            // Reverts the node indicated by `block` to the original Program node.
            // This state is incorrect, but `eslint-utils`'s `referenceTracker.iterateEsmReferences()` tracks import statements
            // from Program nodes set to `block` in global scope. This can only be handled by the original Program node.
            scopeManager.globalScope.block = node;
        });
    });
    (0, sort_1.sortNodes)(body);
    return ast;
}
/** Extract attrs */
function extractAttributes(element, ctx) {
    element.startTag = {
        type: "SvelteStartTag",
        attributes: [],
        selfClosing: false,
        parent: element,
        range: [element.range[0], null],
        loc: {
            start: {
                line: element.loc.start.line,
                column: element.loc.start.column,
            },
            end: null,
        },
    };
    const block = ctx.findBlock(element);
    if (block) {
        element.startTag.attributes.push(...(0, attr_1.convertAttributes)(block.attrs, element.startTag, ctx));
    }
}
/** Convert generics attribute */
function convertGenericsAttribute(script, ctx) {
    const lang = ctx.sourceCode.scripts.attrs.lang;
    if (lang !== "ts" && lang !== "typescript") {
        return;
    }
    const genericsAttribute = script.startTag.attributes.find((attr) => {
        return (attr.type === "SvelteAttribute" &&
            attr.key.name === "generics" &&
            attr.value.length === 1 &&
            attr.value[0].type === "SvelteLiteral");
    });
    if (!genericsAttribute) {
        return;
    }
    const value = genericsAttribute.value[0];
    const genericValueCode = ctx.code.slice(value.range[0], value.range[1]);
    const scriptLet = `void function<${genericValueCode}>(){}`;
    let result;
    try {
        result = (0, script_1.parseScriptWithoutAnalyzeScope)(scriptLet, ctx.sourceCode.scripts.attrs, Object.assign(Object.assign({}, ctx.parserOptions), { 
            // Without typings
            project: null }));
    }
    catch (_a) {
        // ignore
        return;
    }
    delete genericsAttribute.boolean;
    delete genericsAttribute.value;
    // Remove value token indexes
    (0, sort_1.sortNodes)(ctx.tokens);
    const firstTokenIndex = ctx.tokens.findIndex((token) => value.range[0] <= token.range[0] && token.range[1] <= value.range[1]);
    const lastTokenCount = ctx.tokens
        .slice(firstTokenIndex)
        .findIndex((token) => value.range[1] <= token.range[0]);
    ctx.tokens.splice(firstTokenIndex, lastTokenCount >= 0 ? lastTokenCount : Infinity);
    const generics = genericsAttribute;
    generics.type = "SvelteGenericsDirective";
    generics.params = [];
    result.ast.tokens.shift(); // void
    result.ast.tokens.shift(); // function
    result.ast.tokens.shift(); // <
    result.ast.tokens.pop(); // }
    result.ast.tokens.pop(); // {
    result.ast.tokens.pop(); // )
    result.ast.tokens.pop(); // (
    result.ast.tokens.pop(); // >
    (0, fix_locations_1.fixLocations)(result.ast, result.ast.tokens, result.ast.comments, value.range[0] - 14, result.visitorKeys, ctx);
    const { ast } = result;
    const statement = ast.body[0];
    const rawExpression = statement.expression;
    const fnDecl = rawExpression.argument;
    const typeParameters = fnDecl
        .typeParameters;
    const params = typeParameters.params;
    // Replace tokens
    for (const tokensKey of ["tokens", "comments"]) {
        for (const token of result.ast[tokensKey]) {
            if (params.every((param) => token.range[1] <= param.range[0] ||
                param.range[1] <= token.range[0])) {
                ctx[tokensKey].push(token);
            }
        }
    }
    for (const param of params) {
        param.parent = generics;
        generics.params.push(param);
        ctx.scriptLet.addGenericTypeAliasDeclaration(param, (id, typeNode) => {
            param.name = id;
            if (param.constraint) {
                param.constraint = typeNode;
            }
        }, (typeNode) => {
            param.default = typeNode;
        });
    }
}
