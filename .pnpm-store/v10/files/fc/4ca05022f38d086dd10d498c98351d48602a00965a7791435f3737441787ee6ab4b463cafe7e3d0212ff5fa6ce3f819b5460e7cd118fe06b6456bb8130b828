"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.setParent = setParent;
const __1 = require("../..");
function setParent(result) {
    if (result.ast.body.some((node) => node.parent)) {
        return;
    }
    (0, __1.traverseNodes)(result.ast, {
        visitorKeys: result.visitorKeys,
        enterNode(node, parent) {
            node.parent = parent;
        },
        leaveNode() {
            // noop
        },
    });
}
