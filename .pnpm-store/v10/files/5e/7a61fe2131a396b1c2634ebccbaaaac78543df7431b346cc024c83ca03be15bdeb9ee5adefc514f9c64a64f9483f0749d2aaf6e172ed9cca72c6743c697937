/** @param {import('estree').Node} expression */
export function analyze(expression: import('estree').Node): {
    map: WeakMap<import("estree").Node, Scope>;
    scope: Scope;
    globals: Map<string, import("estree").Node>;
};
/**
 * @param {import('estree').Node} param
 * @returns {string[]}
 */
export function extract_names(param: import('estree').Node): string[];
/**
 * @param {import('estree').Node} param
 * @param {import('estree').Identifier[]} nodes
 * @returns {import('estree').Identifier[]}
 */
export function extract_identifiers(param: import('estree').Node, nodes?: import('estree').Identifier[]): import('estree').Identifier[];
export class Scope {
    /**
     * @param {Scope | null} parent
     * @param {boolean} block
     */
    constructor(parent: Scope | null, block: boolean);
    /** @type {Scope | null} */
    parent: Scope | null;
    /** @type {boolean} */
    block: boolean;
    /** @type {Map<string, import('estree').Node>} */
    declarations: Map<string, import('estree').Node>;
    /** @type {Set<string>} */
    initialised_declarations: Set<string>;
    /** @type {Set<string>} */
    references: Set<string>;
    /**
     * @param {import('estree').VariableDeclaration | import('estree').ClassDeclaration} node
     */
    add_declaration(node: import('estree').VariableDeclaration | import('estree').ClassDeclaration): void;
    /**
     * @param {string} name
     * @returns {Scope | null}
     */
    find_owner(name: string): Scope | null;
    /**
     * @param {string} name
     * @returns {boolean}
     */
    has(name: string): boolean;
}
