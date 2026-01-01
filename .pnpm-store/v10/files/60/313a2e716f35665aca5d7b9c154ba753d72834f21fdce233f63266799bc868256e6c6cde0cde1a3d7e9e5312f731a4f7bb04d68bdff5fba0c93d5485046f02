"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.getTypeScriptTools = getTypeScriptTools;
exports.getTypeScript = getTypeScript;
exports.isTruthyLiteral = isTruthyLiteral;
exports.isFalsyType = isFalsyType;
exports.isNullishType = isNullishType;
exports.isNullableType = isNullableType;
exports.isBooleanLiteralType = isBooleanLiteralType;
exports.isObjectType = isObjectType;
exports.isReferenceObjectType = isReferenceObjectType;
exports.isTupleObjectType = isTupleObjectType;
exports.isTupleType = isTupleType;
exports.isAnyType = isAnyType;
exports.isUnknownType = isUnknownType;
exports.isNeverType = isNeverType;
exports.isUndefinedType = isUndefinedType;
exports.isVoidType = isVoidType;
exports.isNullType = isNullType;
exports.isPossiblyFalsyType = isPossiblyFalsyType;
exports.getCallSignaturesOfType = getCallSignaturesOfType;
exports.getConstrainedTypeAtLocation = getConstrainedTypeAtLocation;
exports.getTypeName = getTypeName;
exports.getTypeOfPropertyOfType = getTypeOfPropertyOfType;
const load_module_1 = require("../load-module");
const compat_1 = require("../compat");
/**
 * Get TypeScript tools
 */
function getTypeScriptTools(context) {
    const ts = getTypeScript(context);
    if (!ts) {
        return null;
    }
    const sourceCode = (0, compat_1.getSourceCode)(context);
    const { program, esTreeNodeToTSNodeMap, tsNodeToESTreeNodeMap } = sourceCode.parserServices;
    if (!program || !esTreeNodeToTSNodeMap || !tsNodeToESTreeNodeMap) {
        return null;
    }
    const hasFullTypeInformation = sourceCode.parserServices.hasFullTypeInformation ?? true;
    if (!hasFullTypeInformation) {
        // Full type information is required. User must specify parserOptions.project.
        return null;
    }
    return {
        service: {
            esTreeNodeToTSNodeMap,
            tsNodeToESTreeNodeMap,
            hasFullTypeInformation,
            program
        },
        ts
    };
}
let cacheTypeScript = null;
/**
 * Get TypeScript tools
 */
function getTypeScript(context) {
    if (cacheTypeScript) {
        return cacheTypeScript;
    }
    cacheTypeScript = (0, load_module_1.loadModule)(context, 'typescript');
    if (cacheTypeScript) {
        return cacheTypeScript;
    }
    try {
        // eslint-disable-next-line @typescript-eslint/no-require-imports -- ignore
        cacheTypeScript ?? (cacheTypeScript = require('typescript'));
    }
    catch {
        // ignore
    }
    return cacheTypeScript;
}
/**
 * Check whether the given type is a truthy literal type or not.
 */
function isTruthyLiteral(type, tsTools) {
    if (type.isUnion()) {
        return type.types.every((t) => isTruthyLiteral(t, tsTools));
    }
    return ((isBooleanLiteralType(type, tsTools.ts) &&
        tsTools.service.program.getTypeChecker().typeToString(type) === 'true') ||
        (type.isLiteral() && Boolean(type.value)));
}
/**
 * Check whether the given type is a falsy type or not.
 */
function isFalsyType(type, tsTools) {
    if (type.isUnion()) {
        return type.types.every((t) => isFalsyType(t, tsTools));
    }
    if (isUndefinedType(type, tsTools.ts) ||
        isNullType(type, tsTools.ts) ||
        isVoidType(type, tsTools.ts))
        return true;
    if (type.isLiteral())
        return !type.value;
    return (isBooleanLiteralType(type, tsTools.ts) &&
        tsTools.service.program.getTypeChecker().typeToString(type) === 'false');
}
/**
 * Check whether the given type is a nullish type or not.
 */
function isNullishType(type, ts) {
    if (type.isUnion()) {
        return type.types.every((t) => isNullishType(t, ts));
    }
    return isNullType(type, ts) || isUndefinedType(type, ts);
}
/**
 * Checks whether the given type is nullable or not.
 */
function isNullableType(type, ts) {
    if (type.isUnion()) {
        return type.types.some((t) => isNullableType(t, ts));
    }
    return isNullType(type, ts) || isUndefinedType(type, ts);
}
/**
 * Check whether the given type is a boolean literal type or not.
 */
function isBooleanLiteralType(type, ts) {
    return (type.flags & ts.TypeFlags.BooleanLiteral) !== 0;
}
/**
 * Check whether the given type is an object type or not.
 */
function isObjectType(type, ts) {
    return (type.flags & ts.TypeFlags.Object) !== 0;
}
/**
 * Check whether the given type is a reference type or not.
 */
function isReferenceObjectType(type, ts) {
    return isObjectType(type, ts) && (type.objectFlags & ts.ObjectFlags.Reference) !== 0;
}
/**
 * Check whether the given type is a tuple type or not.
 */
function isTupleObjectType(type, ts) {
    return isObjectType(type, ts) && (type.objectFlags & ts.ObjectFlags.Tuple) !== 0;
}
/**
 * Check whether the given type is a tuple type or not.
 * Unlike isTupleObjectType, it also refers to reference types.
 */
function isTupleType(type, ts) {
    return (isTupleObjectType(type, ts) ||
        (isReferenceObjectType(type, ts) && isTupleObjectType(type.target, ts)));
}
/**
 * Check whether the given type is an any type or not.
 */
function isAnyType(type, ts) {
    return (type.flags & ts.TypeFlags.Any) !== 0;
}
/**
 * Check whether the given type is an unknown type or not.
 */
function isUnknownType(type, ts) {
    return (type.flags & ts.TypeFlags.Unknown) !== 0;
}
/**
 * Check whether the given type is a never type or not.
 */
function isNeverType(type, ts) {
    return (type.flags & ts.TypeFlags.Never) !== 0;
}
/**
 * Check whether the given type is an undefined type or not.
 */
function isUndefinedType(type, ts) {
    return (type.flags & ts.TypeFlags.Undefined) !== 0;
}
/**
 * Check whether the given type is a void type or not.
 */
function isVoidType(type, ts) {
    return (type.flags & ts.TypeFlags.Void) !== 0;
}
/**
 * Check whether the given type is a null type or not.
 */
function isNullType(type, ts) {
    return (type.flags & ts.TypeFlags.Null) !== 0;
}
/**
 * Check whether the given type is a possibly falsy type or not.
 */
function isPossiblyFalsyType(type, ts) {
    if (type.isUnion()) {
        return type.types.some((t) => isPossiblyFalsyType(t, ts));
    }
    return (type.flags & ts.TypeFlags.PossiblyFalsy) !== 0;
}
/**
 * Get the call signatures from the given type.
 * This method is heavily inspired by tsutils. https://github.com/ajafff/tsutils
 * The MIT License (MIT) Copyright (c) 2017 Klaus Meinhardt
 * https://github.com/ajafff/tsutils/blob/master/LICENSE
 */
function getCallSignaturesOfType(type) {
    if (type.isUnion()) {
        return type.types.flatMap((t) => getCallSignaturesOfType(t));
    }
    if (type.isIntersection()) {
        let signatures = [];
        for (const t of type.types) {
            const sig = getCallSignaturesOfType(t);
            if (sig.length !== 0) {
                if (signatures.length) {
                    // if more than one type of the intersection has call signatures, none of them is useful for inference
                    return [];
                }
                signatures = sig;
            }
        }
        return signatures;
    }
    return type.getCallSignatures();
}
/**
 * Resolves the given node's type. Will resolve to the type's generic constraint, if it has one.
 * Copied this method from @typescript-eslint/type-utils. https://github.com/typescript-eslint/typescript-eslint/blob/main/packages/type-utils
 * The MIT License (MIT) Copyright (c) 2021 TypeScript ESLint and other contributors
 * https://github.com/typescript-eslint/typescript-eslint/blob/main/packages/type-utils/LICENSE
 */
function getConstrainedTypeAtLocation(checker, node) {
    const nodeType = checker.getTypeAtLocation(node);
    const constrained = checker.getBaseConstraintOfType(nodeType);
    return constrained ?? nodeType;
}
/**
 * Get the type name of a given type.
 *
 * Copied this method from @typescript-eslint/type-utils. https://github.com/typescript-eslint/typescript-eslint/blob/main/packages/type-utils
 * The MIT License (MIT) Copyright (c) 2021 TypeScript ESLint and other contributors
 * https://github.com/typescript-eslint/typescript-eslint/blob/main/packages/type-utils/LICENSE
 *
 * @param type The type to get the name of.
 */
function getTypeName(type, tsTools) {
    const { ts } = tsTools;
    // It handles `string` and string literal types as string.
    if ((type.flags & ts.TypeFlags.StringLike) !== 0) {
        return 'string';
    }
    const typeChecker = tsTools.service.program.getTypeChecker();
    // If the type is a type parameter which extends primitive string types,
    // but it was not recognized as a string like. So check the constraint
    // type of the type parameter.
    if ((type.flags & ts.TypeFlags.TypeParameter) !== 0) {
        // `type.getConstraint()` method doesn't return the constraint type of
        // the type parameter for some reason. So this gets the constraint type
        // via AST.
        const symbol = type.getSymbol();
        const decls = symbol?.getDeclarations();
        const typeParamDecl = decls?.[0];
        if (ts.isTypeParameterDeclaration(typeParamDecl) && typeParamDecl.constraint != null) {
            return getTypeName(typeChecker.getTypeFromTypeNode(typeParamDecl.constraint), tsTools);
        }
    }
    // If the type is a union and all types in the union are string like,
    // return `string`. For example:
    // - `"a" | "b"` is string.
    // - `string | string[]` is not string.
    if (type.isUnion() &&
        type.types.map((value) => getTypeName(value, tsTools)).every((t) => t === 'string')) {
        return 'string';
    }
    // If the type is an intersection and a type in the intersection is string
    // like, return `string`. For example: `string & {__htmlEscaped: void}`
    if (type.isIntersection() &&
        type.types.map((value) => getTypeName(value, tsTools)).some((t) => t === 'string')) {
        return 'string';
    }
    return typeChecker.typeToString(type);
}
/**
 * Return the type of the given property in the given type, or undefined if no such property exists
 */
function getTypeOfPropertyOfType(type, name, checker) {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any -- getTypeOfPropertyOfType is an internal API of TS.
    return checker.getTypeOfPropertyOfType(type, name);
}
