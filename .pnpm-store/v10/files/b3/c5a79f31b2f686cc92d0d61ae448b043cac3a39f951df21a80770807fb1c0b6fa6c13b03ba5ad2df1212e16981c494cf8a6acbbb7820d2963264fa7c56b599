import type ESTree from "estree";
import type { BaseNode } from "./base";
export type SvelteScriptNode = SvelteReactiveStatement;
/** Node of `$` statement. */
export interface SvelteReactiveStatement extends BaseNode {
    type: "SvelteReactiveStatement";
    label: ESTree.Identifier & {
        name: "$";
    };
    body: ESTree.Statement;
    parent: ESTree.Node;
}
