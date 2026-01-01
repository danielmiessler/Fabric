"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.getFallbackKeys = getFallbackKeys;
exports.getKeys = getKeys;
exports.getNodes = getNodes;
exports.traverseNodes = traverseNodes;
const visitor_keys_1 = require("./visitor-keys");
/**
 * Check that the given key should be traversed or not.
 * @this {Traversable}
 * @param key The key to check.
 * @returns `true` if the key should be traversed.
 */
function fallbackKeysFilter(key) {
    let value = null;
    return (key !== "comments" &&
        key !== "leadingComments" &&
        key !== "loc" &&
        key !== "parent" &&
        key !== "range" &&
        key !== "tokens" &&
        key !== "trailingComments" &&
        (value = this[key]) !== null &&
        typeof value === "object" &&
        (typeof value.type === "string" || Array.isArray(value)));
}
/**
 * Get the keys of the given node to traverse it.
 * @param node The node to get.
 * @returns The keys to traverse.
 */
function getFallbackKeys(node) {
    return Object.keys(node).filter(fallbackKeysFilter, node);
}
/**
 * Get the keys of the given node to traverse it.
 * @param node The node to get.
 * @returns The keys to traverse.
 */
function getKeys(node, visitorKeys) {
    const keys = (visitorKeys || visitor_keys_1.KEYS)[node.type] || getFallbackKeys(node);
    return keys.filter((key) => !getNodes(node, key).next().done);
}
/**
 * Get the nodes of the given node.
 * @param node The node to get.
 */
function* getNodes(node, key) {
    const child = node[key];
    if (Array.isArray(child)) {
        for (const c of child) {
            if (isNode(c)) {
                yield c;
            }
        }
    }
    else if (isNode(child)) {
        yield child;
    }
}
/**
 * Check whether a given value is a node.
 * @param x The value to check.
 * @returns `true` if the value is a node.
 */
function isNode(x) {
    return x !== null && typeof x === "object" && typeof x.type === "string";
}
/**
 * Traverse the given node.
 * @param node The node to traverse.
 * @param parent The parent node.
 * @param visitor The node visitor.
 */
function traverse(node, parent, visitor) {
    visitor.enterNode(node, parent);
    const keys = getKeys(node, visitor.visitorKeys);
    for (const key of keys) {
        for (const child of getNodes(node, key)) {
            traverse(child, node, visitor);
        }
    }
    visitor.leaveNode(node, parent);
}
/**
 * Traverse the given AST tree.
 * @param node Root node to traverse.
 * @param visitor Visitor.
 */
function traverseNodes(node, visitor) {
    traverse(node, null, visitor);
}
