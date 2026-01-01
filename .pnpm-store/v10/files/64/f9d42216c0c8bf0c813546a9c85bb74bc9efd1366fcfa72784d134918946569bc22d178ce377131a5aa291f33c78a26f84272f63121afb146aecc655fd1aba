"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.parseConfig = parseConfig;
const traverse_1 = require("../traverse");
const espree_1 = require("../parser/espree");
const eslint_scope_1 = require("eslint-scope");
const scope_1 = require("../scope");
function parseConfig(code) {
    const espree = (0, espree_1.getEspree)();
    const ast = espree.parse(code, {
        range: true,
        loc: true,
        ecmaVersion: espree.latestEcmaVersion,
        sourceType: "module",
    });
    // Set parent nodes.
    (0, traverse_1.traverseNodes)(ast, {
        enterNode(node, parent) {
            node.parent = parent;
        },
        leaveNode() {
            /* do nothing */
        },
    });
    // Analyze scopes.
    const scopeManager = (0, eslint_scope_1.analyze)(ast, {
        ignoreEval: true,
        nodejsScope: false,
        ecmaVersion: espree.latestEcmaVersion,
        sourceType: "module",
        fallback: traverse_1.getFallbackKeys,
    });
    return parseAst(ast, scopeManager);
}
function parseAst(ast, scopeManager) {
    const edd = ast.body.find((node) => node.type === "ExportDefaultDeclaration");
    if (!edd)
        return {};
    const decl = edd.declaration;
    if (decl.type === "ClassDeclaration" || decl.type === "FunctionDeclaration")
        return {};
    return parseSvelteConfigExpression(decl, scopeManager);
}
function parseSvelteConfigExpression(node, scopeManager) {
    var _a, _b;
    const evaluated = evaluateExpression(node, scopeManager);
    if ((evaluated === null || evaluated === void 0 ? void 0 : evaluated.type) !== 1 /* EvaluatedType.object */)
        return {};
    const result = {};
    // Returns only known properties.
    const compilerOptions = evaluated.getProperty("compilerOptions");
    if ((compilerOptions === null || compilerOptions === void 0 ? void 0 : compilerOptions.type) === 1 /* EvaluatedType.object */) {
        result.compilerOptions = {};
        const runes = (_a = compilerOptions.getProperty("runes")) === null || _a === void 0 ? void 0 : _a.getStatic();
        if ((runes === null || runes === void 0 ? void 0 : runes.value) != null) {
            result.compilerOptions.runes = Boolean(runes.value);
        }
    }
    const kit = evaluated.getProperty("kit");
    if ((kit === null || kit === void 0 ? void 0 : kit.type) === 1 /* EvaluatedType.object */) {
        result.kit = {};
        const files = (_b = kit.getProperty("files")) === null || _b === void 0 ? void 0 : _b.getStatic();
        if ((files === null || files === void 0 ? void 0 : files.value) != null)
            result.kit.files = files.value;
    }
    return result;
}
class EvaluatedLiteral {
    constructor(value) {
        this.type = 0 /* EvaluatedType.literal */;
        this.value = value;
    }
    getStatic() {
        return this;
    }
}
/** Evaluating an object expression. */
class EvaluatedObject {
    constructor(node, parseExpression) {
        this.type = 1 /* EvaluatedType.object */;
        this.cached = new Map();
        this.node = node;
        this.parseExpression = parseExpression;
    }
    /** Gets the evaluated value of the property with the given name. */
    getProperty(key) {
        return this.withCache(key, () => {
            let unknown = false;
            for (const prop of [...this.node.properties].reverse()) {
                if (prop.type === "Property") {
                    const name = this.getKey(prop);
                    if (name === key)
                        return this.parseExpression(prop.value);
                    if (name == null)
                        unknown = true;
                }
                else if (prop.type === "SpreadElement") {
                    const evaluated = this.parseExpression(prop.argument);
                    if ((evaluated === null || evaluated === void 0 ? void 0 : evaluated.type) === 1 /* EvaluatedType.object */) {
                        const value = evaluated.getProperty(key);
                        if (value)
                            return value;
                    }
                    unknown = true;
                }
            }
            return unknown ? null : new EvaluatedLiteral(undefined);
        });
    }
    getStatic() {
        var _a, _b;
        const object = {};
        for (const prop of this.node.properties) {
            if (prop.type === "Property") {
                const name = this.getKey(prop);
                if (name == null)
                    return null;
                const evaluated = (_a = this.withCache(name, () => this.parseExpression(prop.value))) === null || _a === void 0 ? void 0 : _a.getStatic();
                if (!evaluated)
                    return null;
                object[name] = evaluated.value;
            }
            else if (prop.type === "SpreadElement") {
                const evaluated = (_b = this.parseExpression(prop.argument)) === null || _b === void 0 ? void 0 : _b.getStatic();
                if (!evaluated)
                    return null;
                Object.assign(object, evaluated.value);
            }
        }
        return { value: object };
    }
    withCache(key, parse) {
        if (this.cached.has(key))
            return this.cached.get(key) || null;
        const evaluated = parse();
        this.cached.set(key, evaluated);
        return evaluated;
    }
    getKey(node) {
        var _a;
        if (!node.computed && node.key.type === "Identifier")
            return node.key.name;
        const evaluatedKey = (_a = this.parseExpression(node.key)) === null || _a === void 0 ? void 0 : _a.getStatic();
        if (evaluatedKey)
            return String(evaluatedKey.value);
        return null;
    }
}
function evaluateExpression(node, scopeManager) {
    const tracked = new Map();
    return parseExpression(node);
    function parseExpression(node) {
        if (node.type === "Literal") {
            return new EvaluatedLiteral(node.value);
        }
        if (node.type === "Identifier") {
            return parseIdentifier(node);
        }
        if (node.type === "ObjectExpression") {
            return new EvaluatedObject(node, parseExpression);
        }
        return null;
    }
    function parseIdentifier(node) {
        const defs = getIdentifierDefinitions(node);
        if (defs.length !== 1) {
            if (defs.length === 0 && node.name === "undefined")
                return new EvaluatedLiteral(undefined);
            return null;
        }
        const def = defs[0];
        if (def.type !== "Variable")
            return null;
        if (def.parent.kind !== "const" || !def.node.init)
            return null;
        const evaluated = parseExpression(def.node.init);
        if (!evaluated)
            return null;
        const assigns = parsePatternAssign(def.name, def.node.id);
        let result = evaluated;
        while (assigns.length) {
            const assign = assigns.shift();
            if (assign.type === "member") {
                if (result.type !== 1 /* EvaluatedType.object */)
                    return null;
                const next = result.getProperty(assign.name);
                if (!next)
                    return null;
                result = next;
            }
            else if (assign.type === "assignment") {
                if (result.type === 0 /* EvaluatedType.literal */ &&
                    result.value === undefined) {
                    const next = parseExpression(assign.node.right);
                    if (!next)
                        return null;
                    result = next;
                }
            }
        }
        return result;
    }
    function getIdentifierDefinitions(node) {
        var _a, _b;
        if (tracked.has(node))
            return tracked.get(node);
        tracked.set(node, []);
        const defs = (_a = (0, scope_1.findVariable)(scopeManager, node)) === null || _a === void 0 ? void 0 : _a.defs;
        if (!defs)
            return [];
        tracked.set(node, defs);
        if (defs.length !== 1) {
            const def = defs[0];
            if (def.type === "Variable" &&
                def.parent.kind === "const" &&
                def.node.id.type === "Identifier" &&
                ((_b = def.node.init) === null || _b === void 0 ? void 0 : _b.type) === "Identifier") {
                const newDef = getIdentifierDefinitions(def.node.init);
                tracked.set(node, newDef);
                return newDef;
            }
        }
        return defs;
    }
}
/**
 * Returns the assignment path.
 * For example,
 * `let {a: {target}} = {}`
 *   -> `[{type: "member", name: 'a'}, {type: "member", name: 'target'}]`.
 * `let {a: {target} = foo} = {}`
 *   -> `[{type: "member", name: 'a'}, {type: "assignment"}, {type: "member", name: 'target'}]`.
 */
function parsePatternAssign(node, root) {
    return parse(root) || [];
    function parse(target) {
        if (node === target) {
            return [];
        }
        if (target.type === "Identifier") {
            return null;
        }
        if (target.type === "AssignmentPattern") {
            const left = parse(target.left);
            if (!left)
                return null;
            return [{ type: "assignment", node: target }, ...left];
        }
        if (target.type === "ObjectPattern") {
            for (const prop of target.properties) {
                if (prop.type === "Property") {
                    const name = !prop.computed && prop.key.type === "Identifier"
                        ? prop.key.name
                        : prop.key.type === "Literal"
                            ? String(prop.key.value)
                            : null;
                    if (!name)
                        continue;
                    const value = parse(prop.value);
                    if (!value)
                        return null;
                    return [{ type: "member", name }, ...value];
                }
            }
            return null;
        }
        if (target.type === "ArrayPattern") {
            for (const [index, element] of target.elements.entries()) {
                if (!element)
                    continue;
                const value = parse(element);
                if (!value)
                    return null;
                return [{ type: "member", name: String(index) }, ...value];
            }
            return null;
        }
        return null;
    }
}
