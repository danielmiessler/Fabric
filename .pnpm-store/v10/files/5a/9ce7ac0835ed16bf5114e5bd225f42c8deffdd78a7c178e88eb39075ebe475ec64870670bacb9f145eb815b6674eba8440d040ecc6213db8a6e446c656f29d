/**
 * Visit nodes, with ancestral information.
 *
 * This algorithm performs *depth-first* *tree traversal* in *preorder*
 * (**NLR**) or if `reverse` is given, in *reverse preorder* (**NRL**).
 *
 * You can choose for which nodes `visitor` is called by passing a `test`.
 * For complex tests, you should test yourself in `visitor`, as it will be
 * faster and will have improved type information.
 *
 * Walking the tree is an intensive task.
 * Make use of the return values of the visitor when possible.
 * Instead of walking a tree multiple times, walk it once, use `unist-util-is`
 * to check if a node matches, and then perform different operations.
 *
 * You can change the tree.
 * See `Visitor` for more info.
 *
 * @overload
 * @param {Tree} tree
 * @param {Check} check
 * @param {BuildVisitor<Tree, Check>} visitor
 * @param {boolean | null | undefined} [reverse]
 * @returns {undefined}
 *
 * @overload
 * @param {Tree} tree
 * @param {BuildVisitor<Tree>} visitor
 * @param {boolean | null | undefined} [reverse]
 * @returns {undefined}
 *
 * @param {UnistNode} tree
 *   Tree to traverse.
 * @param {Visitor | Test} test
 *   `unist-util-is`-compatible test
 * @param {Visitor | boolean | null | undefined} [visitor]
 *   Handle each node.
 * @param {boolean | null | undefined} [reverse]
 *   Traverse in reverse preorder (NRL) instead of the default preorder (NLR).
 * @returns {undefined}
 *   Nothing.
 *
 * @template {UnistNode} Tree
 *   Node type.
 * @template {Test} Check
 *   `unist-util-is`-compatible test.
 */
export function visitParents<Tree extends UnistNode, Check extends Test>(tree: Tree, check: Check, visitor: BuildVisitor<Tree, Check>, reverse?: boolean | null | undefined): undefined;
/**
 * Visit nodes, with ancestral information.
 *
 * This algorithm performs *depth-first* *tree traversal* in *preorder*
 * (**NLR**) or if `reverse` is given, in *reverse preorder* (**NRL**).
 *
 * You can choose for which nodes `visitor` is called by passing a `test`.
 * For complex tests, you should test yourself in `visitor`, as it will be
 * faster and will have improved type information.
 *
 * Walking the tree is an intensive task.
 * Make use of the return values of the visitor when possible.
 * Instead of walking a tree multiple times, walk it once, use `unist-util-is`
 * to check if a node matches, and then perform different operations.
 *
 * You can change the tree.
 * See `Visitor` for more info.
 *
 * @overload
 * @param {Tree} tree
 * @param {Check} check
 * @param {BuildVisitor<Tree, Check>} visitor
 * @param {boolean | null | undefined} [reverse]
 * @returns {undefined}
 *
 * @overload
 * @param {Tree} tree
 * @param {BuildVisitor<Tree>} visitor
 * @param {boolean | null | undefined} [reverse]
 * @returns {undefined}
 *
 * @param {UnistNode} tree
 *   Tree to traverse.
 * @param {Visitor | Test} test
 *   `unist-util-is`-compatible test
 * @param {Visitor | boolean | null | undefined} [visitor]
 *   Handle each node.
 * @param {boolean | null | undefined} [reverse]
 *   Traverse in reverse preorder (NRL) instead of the default preorder (NLR).
 * @returns {undefined}
 *   Nothing.
 *
 * @template {UnistNode} Tree
 *   Node type.
 * @template {Test} Check
 *   `unist-util-is`-compatible test.
 */
export function visitParents<Tree extends UnistNode, Check extends Test>(tree: Tree, visitor: BuildVisitor<Tree>, reverse?: boolean | null | undefined): undefined;
/**
 * Continue traversing as normal.
 */
export const CONTINUE: true;
/**
 * Stop traversing immediately.
 */
export const EXIT: false;
/**
 * Do not traverse this node’s children.
 */
export const SKIP: "skip";
/**
 * Test from `unist-util-is`.
 *
 * Note: we have remove and add `undefined`, because otherwise when generating
 * automatic `.d.ts` files, TS tries to flatten paths from a local perspective,
 * which doesn’t work when publishing on npm.
 */
export type Test = Exclude<import("unist-util-is").Test, undefined> | undefined;
/**
 * Get the value of a type guard `Fn`.
 */
export type Predicate<Fn, Fallback> = (Fn extends (value: any) => value is infer Thing ? Thing : Fallback);
/**
 * Check whether a node matches a primitive check in the type system.
 */
export type MatchesOne<Value, Check> = (Check extends null | undefined ? Value : Value extends {
    type: Check;
} ? Value : Value extends Check ? Value : Check extends Function ? Predicate<Check, Value> extends Value ? Predicate<Check, Value> : never : never);
/**
 * Check whether a node matches a check in the type system.
 */
export type Matches<Value, Check> = (Check extends ReadonlyArray<infer T> ? MatchesOne<Value, T> : Check extends Array<infer T> ? MatchesOne<Value, T> : MatchesOne<Value, Check>);
/**
 * Number; capped reasonably.
 */
export type Uint = 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10;
/**
 * Increment a number in the type system.
 */
export type Increment<I extends Uint = 0> = I extends 0 ? 1 : I extends 1 ? 2 : I extends 2 ? 3 : I extends 3 ? 4 : I extends 4 ? 5 : I extends 5 ? 6 : I extends 6 ? 7 : I extends 7 ? 8 : I extends 8 ? 9 : 10;
/**
 * Collect nodes that can be parents of `Child`.
 */
export type InternalParent<Node extends UnistNode, Child extends UnistNode> = (Node extends UnistParent ? Node extends {
    children: Array<infer Children>;
} ? Child extends Children ? Node : never : never : never);
/**
 * Collect nodes in `Tree` that can be parents of `Child`.
 */
export type Parent<Tree extends UnistNode, Child extends UnistNode> = InternalParent<InclusiveDescendant<Tree>, Child>;
/**
 * Collect nodes in `Tree` that can be ancestors of `Child`.
 */
export type InternalAncestor<Node extends UnistNode, Child extends UnistNode, Max extends Uint = 10, Depth extends Uint = 0> = (Depth extends Max ? never : InternalParent<Node, Child> | InternalAncestor<Node, InternalParent<Node, Child>, Max, Increment<Depth>>);
/**
 * Collect nodes in `Tree` that can be ancestors of `Child`.
 */
export type Ancestor<Tree extends UnistNode, Child extends UnistNode> = InternalAncestor<InclusiveDescendant<Tree>, Child>;
/**
 * Collect all (inclusive) descendants of `Tree`.
 *
 * > 👉 **Note**: for performance reasons, this seems to be the fastest way to
 * > recurse without actually running into an infinite loop, which the
 * > previous version did.
 * >
 * > Practically, a max of `2` is typically enough assuming a `Root` is
 * > passed, but it doesn’t improve performance.
 * > It gets higher with `List > ListItem > Table > TableRow > TableCell`.
 * > Using up to `10` doesn’t hurt or help either.
 */
export type InclusiveDescendant<Tree extends UnistNode, Max extends Uint = 10, Depth extends Uint = 0> = (Tree extends UnistParent ? Depth extends Max ? Tree : Tree | InclusiveDescendant<Tree["children"][number], Max, Increment<Depth>> : Tree);
/**
 * Union of the action types.
 */
export type Action = "skip" | boolean;
/**
 * Move to the sibling at `index` next (after node itself is completely
 * traversed).
 *
 * Useful if mutating the tree, such as removing the node the visitor is
 * currently on, or any of its previous siblings.
 * Results less than 0 or greater than or equal to `children.length` stop
 * traversing the parent.
 */
export type Index = number;
/**
 * List with one or two values, the first an action, the second an index.
 */
export type ActionTuple = [(Action | null | undefined | void)?, (Index | null | undefined)?];
/**
 * Any value that can be returned from a visitor.
 */
export type VisitorResult = Action | ActionTuple | Index | null | undefined | void;
/**
 * Handle a node (matching `test`, if given).
 *
 * Visitors are free to transform `node`.
 * They can also transform the parent of node (the last of `ancestors`).
 *
 * Replacing `node` itself, if `SKIP` is not returned, still causes its
 * descendants to be walked (which is a bug).
 *
 * When adding or removing previous siblings of `node` (or next siblings, in
 * case of reverse), the `Visitor` should return a new `Index` to specify the
 * sibling to traverse after `node` is traversed.
 * Adding or removing next siblings of `node` (or previous siblings, in case
 * of reverse) is handled as expected without needing to return a new `Index`.
 *
 * Removing the children property of an ancestor still results in them being
 * traversed.
 */
export type Visitor<Visited extends UnistNode = UnistNode, VisitedParents extends UnistParent = UnistParent> = (node: Visited, ancestors: Array<VisitedParents>) => VisitorResult;
/**
 * Build a typed `Visitor` function from a tree and a test.
 *
 * It will infer which values are passed as `node` and which as `parents`.
 */
export type BuildVisitor<Tree extends UnistNode = UnistNode, Check extends Test = Test> = Visitor<Matches<InclusiveDescendant<Tree>, Check>, Ancestor<Tree, Matches<InclusiveDescendant<Tree>, Check>>>;
import type { Node as UnistNode } from 'unist';
import type { Parent as UnistParent } from 'unist';
//# sourceMappingURL=index.d.ts.map