import type ESTree from "estree";
/** indexOf */
export declare function indexOf(str: string, search: (c: string, index: number) => boolean, start: number, end?: number): number;
/** lastIndexOf */
export declare function lastIndexOf(str: string, search: (c: string, index: number) => boolean, end: number): number;
export declare function getWithLoc<N extends ESTree.Comment>(node: N): N & {
    start: number;
    end: number;
};
export declare function getWithLoc<N extends ESTree.Node | {
    start: number;
    end: number;
}>(node: N): N & {
    start: number;
    end: number;
};
export declare function getWithLoc<N extends ESTree.Node | {
    start: number;
    end: number;
}>(node: N | null | undefined): (N & {
    start: number;
    end: number;
}) | null | undefined;
