"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.fixLocations = fixLocations;
const traverse_1 = require("../traverse");
/** Fix locations */
function fixLocations(node, tokens, comments, offset, visitorKeys, ctx) {
    if (offset === 0) {
        return;
    }
    const traversed = new Set();
    (0, traverse_1.traverseNodes)(node, {
        visitorKeys,
        enterNode: (n) => {
            if (traversed.has(n)) {
                return;
            }
            traversed.add(n);
            if (traversed.has(n.range)) {
                if (!traversed.has(n.loc)) {
                    // However, `Node#loc` may not be shared.
                    const locs = ctx.getConvertLocation({
                        start: n.range[0],
                        end: n.range[1],
                    });
                    applyLocs(n, locs);
                    traversed.add(n.loc);
                }
            }
            else {
                const start = n.range[0] + offset;
                const end = n.range[1] + offset;
                const locs = ctx.getConvertLocation({ start, end });
                applyLocs(n, locs);
                traversed.add(n.range);
                traversed.add(n.loc);
            }
        },
        leaveNode: Function.prototype,
    });
    for (const t of tokens) {
        const start = t.range[0] + offset;
        const end = t.range[1] + offset;
        const locs = ctx.getConvertLocation({ start, end });
        applyLocs(t, locs);
    }
    for (const t of comments) {
        const start = t.range[0] + offset;
        const end = t.range[1] + offset;
        const locs = ctx.getConvertLocation({ start, end });
        applyLocs(t, locs);
    }
}
/**
 * applyLocs
 */
function applyLocs(target, locs) {
    target.loc = locs.loc;
    target.range = locs.range;
    if (typeof target.start === "number") {
        delete target.start;
    }
    if (typeof target.end === "number") {
        delete target.end;
    }
}
