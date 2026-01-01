"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.RestoreContext = void 0;
const traverse_1 = require("../../traverse");
const context_1 = require("../../context");
class RestoreContext {
    constructor(code) {
        this.offsets = [];
        this.virtualFragments = [];
        this.restoreStatementProcesses = [];
        this.restoreExpressionProcesses = [];
        this.originalLocs = new context_1.LinesAndColumns(code);
    }
    addRestoreStatementProcess(process) {
        this.restoreStatementProcesses.push(process);
    }
    addRestoreExpressionProcess(process) {
        this.restoreExpressionProcesses.push(process);
    }
    addOffset(offset) {
        this.offsets.push(offset);
    }
    addVirtualFragmentRange(start, end) {
        const peek = this.virtualFragments[this.virtualFragments.length - 1];
        if (peek && peek.end === start) {
            peek.end = end;
            return;
        }
        this.virtualFragments.push({ start, end });
    }
    /**
     * Restore AST nodes
     */
    restore(result) {
        var _a, _b;
        remapLocations(result, {
            remapLocation: (n) => this.remapLocation(n),
            removeToken: (token) => this.virtualFragments.some((f) => f.start <= token.range[0] && token.range[1] <= f.end),
        });
        restoreStatements(result, this.restoreStatementProcesses);
        restoreExpressions(result, this.restoreExpressionProcesses);
        // Adjust program node location
        const firstOffset = Math.min(...[result.ast.body[0], (_a = result.ast.tokens) === null || _a === void 0 ? void 0 : _a[0], (_b = result.ast.comments) === null || _b === void 0 ? void 0 : _b[0]]
            .filter((t) => Boolean(t))
            .map((t) => t.range[0]));
        if (firstOffset < result.ast.range[0]) {
            result.ast.range[0] = firstOffset;
            result.ast.loc.start = this.originalLocs.getLocFromIndex(firstOffset);
        }
    }
    remapLocation(node) {
        let [start, end] = node.range;
        const startFragment = this.virtualFragments.find((f) => f.start <= start && start < f.end);
        if (startFragment) {
            start = startFragment.end;
        }
        const endFragment = this.virtualFragments.find((f) => f.start < end && end <= f.end);
        if (endFragment) {
            end = endFragment.start;
            if (startFragment === endFragment) {
                start = startFragment.start;
            }
        }
        if (end < start) {
            const w = start;
            start = end;
            end = w;
        }
        const locs = this.originalLocs.getLocations(...this.getRemapRange(start, end));
        node.loc = locs.loc;
        node.range = locs.range;
        if (node.start != null) {
            delete node.start;
        }
        if (node.end != null) {
            delete node.end;
        }
    }
    getRemapRange(start, end) {
        if (!this.offsets.length) {
            return [start, end];
        }
        let lastStart = this.offsets[0];
        let lastEnd = this.offsets[0];
        for (const offset of this.offsets) {
            if (offset.dist <= start) {
                lastStart = offset;
            }
            if (offset.dist < end) {
                lastEnd = offset;
            }
            else {
                break;
            }
        }
        const remapStart = lastStart.original + (start - lastStart.dist);
        const remapEnd = lastEnd.original + (end - lastEnd.dist);
        return [remapStart, remapEnd];
    }
}
exports.RestoreContext = RestoreContext;
/** Restore locations */
function remapLocations(result, { remapLocation, removeToken, }) {
    const traversed = new Map();
    // remap locations
    (0, traverse_1.traverseNodes)(result.ast, {
        visitorKeys: result.visitorKeys,
        enterNode: (node, parent) => {
            node.parent = parent;
            if (!traversed.has(node)) {
                traversed.set(node, parent);
                remapLocation(node);
            }
        },
        leaveNode: (_node) => {
            // noop
        },
    });
    const tokens = [];
    for (const token of result.ast.tokens || []) {
        if (removeToken(token)) {
            continue;
        }
        remapLocation(token);
        tokens.push(token);
    }
    result.ast.tokens = tokens;
    for (const token of result.ast.comments || []) {
        remapLocation(token);
    }
}
/** Restore statement nodes */
function restoreStatements(result, restoreStatementProcesses) {
    const restoreStatementProcessesSet = new Set(restoreStatementProcesses);
    for (const node of [...result.ast.body]) {
        if (!restoreStatementProcessesSet.size) {
            break;
        }
        for (const proc of restoreStatementProcessesSet) {
            if (proc(node, result)) {
                restoreStatementProcessesSet.delete(proc);
                break;
            }
        }
    }
}
/** Restore expression nodes */
function restoreExpressions(result, restoreExpressionProcesses) {
    if (restoreExpressionProcesses.length === 0)
        return;
    const restoreExpressionProcessesSet = new Set(restoreExpressionProcesses);
    (0, traverse_1.traverseNodes)(result.ast, {
        visitorKeys: result.visitorKeys,
        enterNode(node) {
            for (const proc of restoreExpressionProcessesSet) {
                if (proc.target === node.type) {
                    if (proc.restore(node, result)) {
                        restoreExpressionProcessesSet.delete(proc);
                    }
                    break;
                }
            }
        },
        leaveNode() {
            /* noop */
        },
    });
}
