declare namespace stableStringify {
    type Key = string | number;

    type NonArrayNode = Record<Key, unknown>;
    type Node = unknown[] | NonArrayNode;

    type Getter = { get(key: Key): unknown };

    type Comparator = (a: { key: string, value: unknown }, b: { key: string, value: unknown }, getter: Getter) => number;

    type StableStringifyOptions = {
        cmp?: Comparator;
        collapseEmpty?: boolean;
        cycles?: boolean;
        replacer?: (this: Node, key: Key, value: unknown) => unknown;
        space?: string | number;
    };
}

declare function stableStringify(
    obj: unknown,
    options?: (stableStringify.Comparator & stableStringify.StableStringifyOptions) | stableStringify.StableStringifyOptions,
): string | undefined;

export = stableStringify;
