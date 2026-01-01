import type { AST } from 'svelte-eslint-parser';
import type { RuleContext } from '../../types';
export type IgnoreItemWithoutCode = {
    range: [number, number];
    code: null;
    token: AST.Token | AST.Comment;
};
export type IgnoreItem = {
    range: [number, number];
    code: string;
    codeForV5: string;
    token: AST.Token | AST.Comment;
};
/** Extract all svelte-ignore comment items */
export declare function getSvelteIgnoreItems(context: RuleContext): (IgnoreItem | IgnoreItemWithoutCode)[];
