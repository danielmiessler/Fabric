import type { ASTNode, SourceCode } from '../../types';
import type { AnyToken, IndentOptions, MaybeNode } from './commons';
declare const enum OffsetDataType {
    normal = 0,
    align = 1,
    start = 2
}
type OffsetData = {
    type: OffsetDataType.normal;
    base: number;
    offset: number;
    expectedIndent?: number;
} | {
    type: OffsetDataType.align;
    base: number;
    alignIndent: number;
    expectedIndent?: number;
} | {
    type: OffsetDataType.start;
    offset: number;
    expectedIndent?: number;
};
export declare class OffsetContext {
    private readonly sourceCode;
    private readonly options;
    private readonly offsets;
    private readonly ignoreRanges;
    constructor(arg: {
        sourceCode: SourceCode;
        options: IndentOptions;
    });
    /**
     * Set offset to the given index.
     */
    setOffsetIndex(index: number, offset: number, base: number): void;
    /**
     * Set align indent to the given index.
     */
    private setAlignIndent;
    /**
     * Set offset to the given tokens.
     */
    setOffsetToken(token: AnyToken | null | undefined | (AnyToken | null | undefined)[], offset: number, baseToken: AnyToken): void;
    /**
     * Copy offset to the given index from srcIndex.
     */
    copyOffset(index: number, srcIndex: number): void;
    /**
     * Set start offset to the given index.
     */
    setStartOffsetIndex(index: number, offset: number): void;
    /**
     * Set start offset to the given token.
     */
    setStartOffsetToken(token: AnyToken | null | undefined | (AnyToken | null | undefined)[], offset: number): void;
    setOffsetElementList(nodes: (ASTNode | AnyToken | MaybeNode | null | undefined)[], baseNodeOrToken: ASTNode | AnyToken | MaybeNode, lastNodeOrToken: ASTNode | AnyToken | MaybeNode | null, offset: number, align?: boolean): void;
    private _setOffsetElementList;
    /**
     * Ignore range of the given node.
     */
    ignore(node: ASTNode): void;
    getOffsetCalculator(): OffsetCalculator;
}
export declare class OffsetCalculator {
    private readonly options;
    private readonly offsets;
    private readonly ignoreRanges;
    constructor(arg: {
        offsets: Map<number, OffsetData>;
        options: IndentOptions;
        ignoreRanges: [number, number][];
    });
    /**
     * Calculate correct indentation of the given index.
     */
    private getExpectedIndentFromIndex;
    /**
     * Calculate correct indentation of the given token.
     */
    getExpectedIndentFromToken(token: AnyToken): number | null;
    /**
     * Calculate correct indentation of the line of the given tokens.
     */
    getExpectedIndentFromTokens(tokens: AnyToken[]): null | number;
    /** Save expected indent to give tokens */
    saveExpectedIndent(tokens: AnyToken[], expectedIndent: number): void;
}
export {};
