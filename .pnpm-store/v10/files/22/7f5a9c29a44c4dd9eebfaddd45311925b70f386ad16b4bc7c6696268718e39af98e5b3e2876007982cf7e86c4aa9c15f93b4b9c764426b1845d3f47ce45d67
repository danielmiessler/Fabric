"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.convertAttributes = convertAttributes;
const common_1 = require("./common");
const mustache_1 = require("./mustache");
const text_1 = require("./text");
const errors_1 = require("../../errors");
const svelte_version_1 = require("../svelte-version");
const utils_1 = require("../../utils");
const compat_1 = require("../compat");
/** Convert for Attributes */
function* convertAttributes(attributes, parent, ctx) {
    for (const attr of attributes) {
        if (attr.type === "Attribute") {
            yield convertAttribute(attr, parent, ctx);
            continue;
        }
        if (attr.type === "SpreadAttribute" || attr.type === "Spread") {
            yield convertSpreadAttribute(attr, parent, ctx);
            continue;
        }
        if (attr.type === "BindDirective" || attr.type === "Binding") {
            yield convertBindingDirective(attr, parent, ctx);
            continue;
        }
        if (attr.type === "OnDirective" || attr.type === "EventHandler") {
            yield convertEventHandlerDirective(attr, parent, ctx);
            continue;
        }
        if (attr.type === "ClassDirective" || attr.type === "Class") {
            yield convertClassDirective(attr, parent, ctx);
            continue;
        }
        if (attr.type === "StyleDirective") {
            yield convertStyleDirective(attr, parent, ctx);
            continue;
        }
        if (attr.type === "TransitionDirective" || attr.type === "Transition") {
            yield convertTransitionDirective(attr, parent, ctx);
            continue;
        }
        if (attr.type === "AnimateDirective" || attr.type === "Animation") {
            yield convertAnimationDirective(attr, parent, ctx);
            continue;
        }
        if (attr.type === "UseDirective" || attr.type === "Action") {
            yield convertActionDirective(attr, parent, ctx);
            continue;
        }
        if (attr.type === "LetDirective") {
            yield convertLetDirective(attr, parent, ctx);
            continue;
        }
        if (attr.type === "Let") {
            yield convertLetDirective(attr, parent, ctx);
            continue;
        }
        if (attr.type === "Ref") {
            throw new errors_1.ParseError("Ref are not supported.", attr.start, ctx);
        }
        if (attr.type === "Style") {
            throw new errors_1.ParseError(`Svelte v3.46.0 is no longer supported. Please use Svelte>=v3.46.1.`, attr.start, ctx);
        }
        throw new errors_1.ParseError(`Unknown directive or attribute (${attr.type}) are not supported.`, attr.start, ctx);
    }
}
/** Convert for Attribute */
function convertAttribute(node, parent, ctx) {
    const attribute = Object.assign({ type: "SvelteAttribute", boolean: false, key: null, value: [], parent }, ctx.getConvertLocation(node));
    const keyStart = ctx.code.indexOf(node.name, node.start);
    const keyRange = { start: keyStart, end: keyStart + node.name.length };
    attribute.key = Object.assign({ type: "SvelteName", name: node.name, parent: attribute }, ctx.getConvertLocation(keyRange));
    if (node.value === true) {
        // Boolean attribute
        attribute.boolean = true;
        ctx.addToken("HTMLIdentifier", keyRange);
        return attribute;
    }
    const value = node.value;
    const shorthand = isAttributeShorthandForSvelteV4(value) ||
        isAttributeShorthandForSvelteV5(value);
    if (shorthand) {
        const key = Object.assign(Object.assign({}, attribute.key), { type: "Identifier" });
        const sAttr = {
            type: "SvelteShorthandAttribute",
            key,
            value: key,
            parent,
            loc: attribute.loc,
            range: attribute.range,
        };
        key.parent = sAttr;
        ctx.scriptLet.addObjectShorthandProperty(attribute.key, sAttr, (es) => {
            if (
            // FIXME: Older parsers may use the same node. In that case, do not replace.
            // We will drop support for ESLint v7 in the next major version and remove this branch.
            es.key !== es.value) {
                sAttr.key = es.key;
            }
            sAttr.value = es.value;
        });
        return sAttr;
    }
    // Not required for shorthands. Therefore, register the token here.
    ctx.addToken("HTMLIdentifier", keyRange);
    processAttributeValue(value, attribute, parent, ctx);
    return attribute;
    function isAttributeShorthandForSvelteV4(value) {
        var _a;
        return Array.isArray(value) && ((_a = value[0]) === null || _a === void 0 ? void 0 : _a.type) === "AttributeShorthand";
    }
    function isAttributeShorthandForSvelteV5(value) {
        return (!Array.isArray(value) &&
            value.type === "ExpressionTag" &&
            ctx.code[node.start] === "{" &&
            ctx.code[node.end - 1] === "}");
    }
}
/** Common process attribute value */
function processAttributeValue(nodeValue, attribute, attributeParent, ctx) {
    const nodes = (Array.isArray(nodeValue) ? nodeValue : [nodeValue])
        .filter((v) => v.type !== "Text" ||
        // ignore empty
        // https://github.com/sveltejs/svelte/pull/6539
        v.start < v.end)
        .map((v, index, array) => {
        if (v.type === "Text") {
            const next = array[index + 1];
            if (next && next.start < v.end) {
                // Maybe bug in Svelte can cause the completion index to shift.
                return Object.assign(Object.assign({}, v), { end: next.start });
            }
        }
        return v;
    });
    if (nodes.length === 1 &&
        (nodes[0].type === "ExpressionTag" || nodes[0].type === "MustacheTag") &&
        attribute.type === "SvelteAttribute") {
        const typing = buildAttributeType(attributeParent.parent, attribute.key.name, ctx);
        const mustache = (0, mustache_1.convertMustacheTag)(nodes[0], attribute, typing, ctx);
        attribute.value.push(mustache);
        return;
    }
    for (const v of nodes) {
        if (v.type === "Text") {
            attribute.value.push((0, text_1.convertTextToLiteral)(v, attribute, ctx));
            continue;
        }
        if (v.type === "ExpressionTag" || v.type === "MustacheTag") {
            const mustache = (0, mustache_1.convertMustacheTag)(v, attribute, null, ctx);
            attribute.value.push(mustache);
            continue;
        }
        const u = v;
        throw new errors_1.ParseError(`Unknown attribute value (${u.type}) are not supported.`, u.start, ctx);
    }
}
/** Build attribute type */
function buildAttributeType(element, attrName, ctx) {
    if (svelte_version_1.svelteVersion.gte(5) &&
        attrName.startsWith("on") &&
        (element.type !== "SvelteElement" || element.kind === "html")) {
        return buildEventHandlerType(element, attrName.slice(2), ctx);
    }
    if (element.type !== "SvelteElement" || element.kind !== "component") {
        return null;
    }
    const elementName = ctx.elements.get(element).name;
    const componentPropsType = `import('svelte').ComponentProps<typeof ${elementName}>`;
    return conditional({
        check: `'${attrName}'`,
        extends: `infer PROP`,
        true: conditional({
            check: `PROP`,
            extends: `keyof ${componentPropsType}`,
            true: `${componentPropsType}[PROP]`,
            false: `never`,
        }),
        false: `never`,
    });
    /** Generate `C extends E ? T : F` type. */
    function conditional(types) {
        return `${types.check} extends ${types.extends}?(${types.true}):(${types.false})`;
    }
}
/** Convert for Spread */
function convertSpreadAttribute(node, parent, ctx) {
    const attribute = Object.assign({ type: "SvelteSpreadAttribute", argument: null, parent }, ctx.getConvertLocation(node));
    const spreadStart = ctx.code.indexOf("...", node.start);
    ctx.addToken("Punctuator", {
        start: spreadStart,
        end: spreadStart + 3,
    });
    ctx.scriptLet.addExpression(node.expression, attribute, null, (es) => {
        attribute.argument = es;
    });
    return attribute;
}
/** Convert for Binding Directive */
function convertBindingDirective(node, parent, ctx) {
    const directive = Object.assign({ type: "SvelteDirective", kind: "Binding", key: null, shorthand: false, expression: null, parent }, ctx.getConvertLocation(node));
    processDirective(node, directive, ctx, {
        processExpression(expression, shorthand) {
            directive.shorthand = shorthand;
            return ctx.scriptLet.addExpression(expression, directive, null, (es, { getScope }) => {
                directive.expression = es;
                const scope = getScope(es);
                const reference = scope.references.find((ref) => ref.identifier === es);
                if (reference) {
                    // The bind directive does read and write.
                    reference.isWrite = () => true;
                    reference.isWriteOnly = () => false;
                    reference.isReadWrite = () => true;
                    reference.isReadOnly = () => false;
                    reference.isRead = () => true;
                }
            });
        },
    });
    return directive;
}
/** Convert for EventHandler Directive */
function convertEventHandlerDirective(node, parent, ctx) {
    const directive = Object.assign({ type: "SvelteDirective", kind: "EventHandler", key: null, expression: null, parent }, ctx.getConvertLocation(node));
    const typing = buildEventHandlerType(parent.parent, node.name, ctx);
    processDirective(node, directive, ctx, {
        processExpression: buildProcessExpressionForExpression(directive, ctx, typing),
    });
    return directive;
}
/** Build event handler type */
function buildEventHandlerType(element, eventName, ctx) {
    const nativeEventHandlerType = `(e:${conditional({
        check: `'${eventName}'`,
        extends: `infer EVT`,
        true: conditional({
            check: `EVT`,
            extends: `keyof HTMLElementEventMap`,
            true: `HTMLElementEventMap[EVT]`,
            false: `CustomEvent<any>`,
        }),
        false: `never`,
    })})=>void`;
    if (element.type !== "SvelteElement") {
        return nativeEventHandlerType;
    }
    const elementName = ctx.elements.get(element).name;
    if (element.kind === "component") {
        const componentEventsType = `import('svelte').ComponentEvents<${elementName}>`;
        return `(e:${conditional({
            check: `0`,
            extends: `(1 & ${componentEventsType})`,
            // `componentEventsType` is `any`
            //   `@typescript-eslint/parser` currently cannot parse `*.svelte` import types correctly.
            //   So if we try to do a correct type parsing, it's argument type will be `any`.
            //   A workaround is to inject the type directly, as `CustomEvent<any>` is better than `any`.
            true: `CustomEvent<any>`,
            // `componentEventsType` has an exact type.
            false: conditional({
                check: `'${eventName}'`,
                extends: `infer EVT`,
                true: conditional({
                    check: `EVT`,
                    extends: `keyof ${componentEventsType}`,
                    true: `${componentEventsType}[EVT]`,
                    false: `CustomEvent<any>`,
                }),
                false: `never`,
            }),
        })})=>void`;
    }
    if (element.kind === "special") {
        if (elementName === "svelte:component")
            return `(e:CustomEvent<any>)=>void`;
        return nativeEventHandlerType;
    }
    const attrName = `on:${eventName}`;
    const svelteHTMLElementsType = "import('svelte/elements').SvelteHTMLElements";
    return conditional({
        check: `'${elementName}'`,
        extends: "infer EL",
        true: conditional({
            check: `EL`,
            extends: `keyof ${svelteHTMLElementsType}`,
            true: conditional({
                check: `'${attrName}'`,
                extends: "infer ATTR",
                true: conditional({
                    check: `ATTR`,
                    extends: `keyof ${svelteHTMLElementsType}[EL]`,
                    true: `${svelteHTMLElementsType}[EL][ATTR]`,
                    false: nativeEventHandlerType,
                }),
                false: `never`,
            }),
            false: nativeEventHandlerType,
        }),
        false: `never`,
    });
    /** Generate `C extends E ? T : F` type. */
    function conditional(types) {
        return `${types.check} extends ${types.extends}?(${types.true}):(${types.false})`;
    }
}
/** Convert for Class Directive */
function convertClassDirective(node, parent, ctx) {
    const directive = Object.assign({ type: "SvelteDirective", kind: "Class", key: null, shorthand: false, expression: null, parent }, ctx.getConvertLocation(node));
    processDirective(node, directive, ctx, {
        processExpression(expression, shorthand) {
            directive.shorthand = shorthand;
            return ctx.scriptLet.addExpression(expression, directive);
        },
    });
    return directive;
}
/** Convert for Style Directive */
function convertStyleDirective(node, parent, ctx) {
    const directive = Object.assign({ type: "SvelteStyleDirective", key: null, shorthand: false, value: [], parent }, ctx.getConvertLocation(node));
    processDirectiveKey(node, directive, ctx);
    const keyName = directive.key.name;
    if (node.value === true) {
        const shorthandDirective = directive;
        shorthandDirective.shorthand = true;
        ctx.scriptLet.addExpression(keyName, shorthandDirective.key, null, (expression) => {
            if (expression.type !== "Identifier") {
                throw new errors_1.ParseError(`Expected JS identifier or attribute value.`, expression.range[0], ctx);
            }
            shorthandDirective.key.name = expression;
        });
        return shorthandDirective;
    }
    ctx.addToken("HTMLIdentifier", {
        start: keyName.range[0],
        end: keyName.range[1],
    });
    processAttributeValue(node.value, directive, parent, ctx);
    return directive;
}
/** Convert for Transition Directive */
function convertTransitionDirective(node, parent, ctx) {
    const directive = Object.assign({ type: "SvelteDirective", kind: "Transition", intro: node.intro, outro: node.outro, key: null, expression: null, parent }, ctx.getConvertLocation(node));
    processDirective(node, directive, ctx, {
        processExpression: buildProcessExpressionForExpression(directive, ctx, null),
        processName: (name) => ctx.scriptLet.addExpression(name, directive.key, null, buildExpressionTypeChecker(["Identifier"], ctx)),
    });
    return directive;
}
/** Convert for Animation Directive */
function convertAnimationDirective(node, parent, ctx) {
    const directive = Object.assign({ type: "SvelteDirective", kind: "Animation", key: null, expression: null, parent }, ctx.getConvertLocation(node));
    processDirective(node, directive, ctx, {
        processExpression: buildProcessExpressionForExpression(directive, ctx, null),
        processName: (name) => ctx.scriptLet.addExpression(name, directive.key, null, buildExpressionTypeChecker(["Identifier"], ctx)),
    });
    return directive;
}
/** Convert for Action Directive */
function convertActionDirective(node, parent, ctx) {
    const directive = Object.assign({ type: "SvelteDirective", kind: "Action", key: null, expression: null, parent }, ctx.getConvertLocation(node));
    processDirective(node, directive, ctx, {
        processExpression: buildProcessExpressionForExpression(directive, ctx, `Parameters<typeof ${node.name}>[1]`),
        processName: (name) => ctx.scriptLet.addExpression(name, directive.key, null, buildExpressionTypeChecker(["Identifier", "MemberExpression"], ctx)),
    });
    return directive;
}
/** Convert for Let Directive */
function convertLetDirective(node, parent, ctx) {
    const directive = Object.assign({ type: "SvelteDirective", kind: "Let", key: null, expression: null, parent }, ctx.getConvertLocation(node));
    processDirective(node, directive, ctx, {
        processPattern(pattern) {
            return ctx.letDirCollections
                .getCollection()
                .addPattern(pattern, directive, buildLetDirectiveType(parent.parent, node.name, ctx));
        },
        processName: node.expression
            ? undefined
            : (name) => {
                // shorthand
                ctx.letDirCollections
                    .getCollection()
                    .addPattern(name, directive, buildLetDirectiveType(parent.parent, node.name, ctx), (es) => {
                    directive.expression = es;
                });
                return [];
            },
    });
    return directive;
}
/** Build let directive param type */
function buildLetDirectiveType(element, letName, ctx) {
    if (element.type !== "SvelteElement") {
        return "any";
    }
    let slotName = "default";
    let componentName;
    const svelteNode = ctx.elements.get(element);
    const slotAttr = svelteNode.attributes.find((attr) => {
        return attr.type === "Attribute" && attr.name === "slot";
    });
    if (slotAttr) {
        if (Array.isArray(slotAttr.value) &&
            slotAttr.value.length === 1 &&
            slotAttr.value[0].type === "Text") {
            slotName = slotAttr.value[0].data;
        }
        else {
            return "any";
        }
        const parent = findParentComponent(element);
        if (parent == null)
            return "any";
        componentName = ctx.elements.get(parent).name;
    }
    else {
        if (element.kind === "component") {
            componentName = svelteNode.name;
        }
        else {
            const parent = findParentComponent(element);
            if (parent == null)
                return "any";
            componentName = ctx.elements.get(parent).name;
        }
    }
    return `${String(componentName)}['$$slot_def'][${JSON.stringify(slotName)}][${JSON.stringify(letName)}]`;
    /** Find parent component element */
    function findParentComponent(node) {
        let parent = node.parent;
        while (parent && parent.type !== "SvelteElement") {
            parent = parent.parent;
        }
        if (!parent || parent.kind !== "component") {
            return null;
        }
        return parent;
    }
}
/** Common process for directive */
function processDirective(node, directive, ctx, processors) {
    processDirectiveKey(node, directive, ctx);
    processDirectiveExpression(node, directive, ctx, processors);
}
/** Common process for directive key */
function processDirectiveKey(node, directive, ctx) {
    const colonIndex = ctx.code.indexOf(":", directive.range[0]);
    ctx.addToken("HTMLIdentifier", {
        start: directive.range[0],
        end: colonIndex,
    });
    const nameIndex = ctx.code.indexOf(node.name, colonIndex + 1);
    const nameRange = {
        start: nameIndex,
        end: nameIndex + node.name.length,
    };
    let keyEndIndex = nameRange.end;
    // modifiers
    if (ctx.code[nameRange.end] === "|") {
        let nextStart = nameRange.end + 1;
        let nextEnd = (0, common_1.indexOf)(ctx.code, (c) => c === "=" || c === ">" || c === "/" || c === "|" || !c.trim(), nextStart);
        ctx.addToken("HTMLIdentifier", { start: nextStart, end: nextEnd });
        while (ctx.code[nextEnd] === "|") {
            nextStart = nextEnd + 1;
            nextEnd = (0, common_1.indexOf)(ctx.code, (c) => c === "=" || c === ">" || c === "/" || c === "|" || !c.trim(), nextStart);
            ctx.addToken("HTMLIdentifier", { start: nextStart, end: nextEnd });
        }
        keyEndIndex = nextEnd;
    }
    const key = (directive.key = Object.assign({ type: "SvelteDirectiveKey", name: null, modifiers: (0, compat_1.getModifiers)(node), parent: directive }, ctx.getConvertLocation({ start: node.start, end: keyEndIndex })));
    // put name
    key.name = Object.assign({ type: "SvelteName", name: node.name, parent: key }, ctx.getConvertLocation(nameRange));
}
/** Common process for directive expression */
function processDirectiveExpression(node, directive, ctx, processors) {
    const key = directive.key;
    const keyName = key.name;
    let shorthand = false;
    if (node.expression) {
        shorthand =
            node.expression.type === "Identifier" &&
                node.expression.name === node.name &&
                (0, common_1.getWithLoc)(node.expression).start === keyName.range[0];
        if (shorthand && (0, common_1.getWithLoc)(node.expression).end !== keyName.range[1]) {
            // The identifier location may be incorrect in some edge cases.
            // e.g. bind:value=""
            (0, common_1.getWithLoc)(node.expression).end = keyName.range[1];
        }
        if (processors.processExpression) {
            processors.processExpression(node.expression, shorthand).push((es) => {
                if (node.expression && es.type !== node.expression.type) {
                    throw new errors_1.ParseError(`Expected ${node.expression.type}, but ${es.type} found.`, es.range[0], ctx);
                }
                directive.expression = es;
            });
        }
        else {
            processors.processPattern(node.expression, shorthand).push((es) => {
                directive.expression = es;
            });
        }
    }
    if (!shorthand) {
        if (processors.processName) {
            processors.processName(keyName).push((es) => {
                key.name = es;
            });
        }
        else {
            ctx.addToken("HTMLIdentifier", {
                start: keyName.range[0],
                end: keyName.range[1],
            });
        }
    }
}
/** Build processExpression for Expression */
function buildProcessExpressionForExpression(directive, ctx, typing) {
    return (expression) => {
        if ((0, utils_1.hasTypeInfo)(expression)) {
            return ctx.scriptLet.addExpression(expression, directive, null);
        }
        return ctx.scriptLet.addExpression(expression, directive, typing);
    };
}
/** Build expression type checker to script let callbacks */
function buildExpressionTypeChecker(expected, ctx) {
    return (node) => {
        if (!expected.includes(node.type)) {
            throw new errors_1.ParseError(`Expected JS ${expected.join(", or ")}, but ${node.type} found.`, node.range[0], ctx);
        }
    };
}
