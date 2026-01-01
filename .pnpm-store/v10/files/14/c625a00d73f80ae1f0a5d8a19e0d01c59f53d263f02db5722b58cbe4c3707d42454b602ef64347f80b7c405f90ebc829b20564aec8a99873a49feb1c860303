import type { RuleContext, ASTNode } from '../../types';
import type * as TS from 'typescript';
export type TypeScript = typeof TS;
export type { TS };
export type TSTools = {
    service: {
        esTreeNodeToTSNodeMap: ReadonlyMap<unknown, TS.Node>;
        tsNodeToESTreeNodeMap: ReadonlyMap<TS.Node, ASTNode>;
        program: TS.Program;
        hasFullTypeInformation: boolean;
    };
    ts: TypeScript;
};
/**
 * Get TypeScript tools
 */
export declare function getTypeScriptTools(context: RuleContext): TSTools | null;
/**
 * Get TypeScript tools
 */
export declare function getTypeScript(context: RuleContext): TypeScript | null;
/**
 * Check whether the given type is a truthy literal type or not.
 */
export declare function isTruthyLiteral(type: TS.Type, tsTools: TSTools): boolean;
/**
 * Check whether the given type is a falsy type or not.
 */
export declare function isFalsyType(type: TS.Type, tsTools: TSTools): boolean;
/**
 * Check whether the given type is a nullish type or not.
 */
export declare function isNullishType(type: TS.Type, ts: TypeScript): boolean;
/**
 * Checks whether the given type is nullable or not.
 */
export declare function isNullableType(type: TS.Type, ts: TypeScript): boolean;
/**
 * Check whether the given type is a boolean literal type or not.
 */
export declare function isBooleanLiteralType(type: TS.Type, ts: TypeScript): boolean;
/**
 * Check whether the given type is an object type or not.
 */
export declare function isObjectType(type: TS.Type, ts: TypeScript): type is TS.ObjectType;
/**
 * Check whether the given type is a reference type or not.
 */
export declare function isReferenceObjectType(type: TS.Type, ts: TypeScript): type is TS.TypeReference;
/**
 * Check whether the given type is a tuple type or not.
 */
export declare function isTupleObjectType(type: TS.Type, ts: TypeScript): type is TS.TupleType;
/**
 * Check whether the given type is a tuple type or not.
 * Unlike isTupleObjectType, it also refers to reference types.
 */
export declare function isTupleType(type: TS.Type, ts: TypeScript): boolean;
/**
 * Check whether the given type is an any type or not.
 */
export declare function isAnyType(type: TS.Type, ts: TypeScript): boolean;
/**
 * Check whether the given type is an unknown type or not.
 */
export declare function isUnknownType(type: TS.Type, ts: TypeScript): boolean;
/**
 * Check whether the given type is a never type or not.
 */
export declare function isNeverType(type: TS.Type, ts: TypeScript): boolean;
/**
 * Check whether the given type is an undefined type or not.
 */
export declare function isUndefinedType(type: TS.Type, ts: TypeScript): boolean;
/**
 * Check whether the given type is a void type or not.
 */
export declare function isVoidType(type: TS.Type, ts: TypeScript): boolean;
/**
 * Check whether the given type is a null type or not.
 */
export declare function isNullType(type: TS.Type, ts: TypeScript): boolean;
/**
 * Check whether the given type is a possibly falsy type or not.
 */
export declare function isPossiblyFalsyType(type: TS.Type, ts: TypeScript): boolean;
/**
 * Get the call signatures from the given type.
 * This method is heavily inspired by tsutils. https://github.com/ajafff/tsutils
 * The MIT License (MIT) Copyright (c) 2017 Klaus Meinhardt
 * https://github.com/ajafff/tsutils/blob/master/LICENSE
 */
export declare function getCallSignaturesOfType(type: TS.Type): readonly TS.Signature[];
/**
 * Resolves the given node's type. Will resolve to the type's generic constraint, if it has one.
 * Copied this method from @typescript-eslint/type-utils. https://github.com/typescript-eslint/typescript-eslint/blob/main/packages/type-utils
 * The MIT License (MIT) Copyright (c) 2021 TypeScript ESLint and other contributors
 * https://github.com/typescript-eslint/typescript-eslint/blob/main/packages/type-utils/LICENSE
 */
export declare function getConstrainedTypeAtLocation(checker: TS.TypeChecker, node: TS.Node): TS.Type;
/**
 * Get the type name of a given type.
 *
 * Copied this method from @typescript-eslint/type-utils. https://github.com/typescript-eslint/typescript-eslint/blob/main/packages/type-utils
 * The MIT License (MIT) Copyright (c) 2021 TypeScript ESLint and other contributors
 * https://github.com/typescript-eslint/typescript-eslint/blob/main/packages/type-utils/LICENSE
 *
 * @param type The type to get the name of.
 */
export declare function getTypeName(type: TS.Type, tsTools: TSTools): string;
/**
 * Return the type of the given property in the given type, or undefined if no such property exists
 */
export declare function getTypeOfPropertyOfType(type: TS.Type, name: string, checker: TS.TypeChecker): TS.Type | undefined;
