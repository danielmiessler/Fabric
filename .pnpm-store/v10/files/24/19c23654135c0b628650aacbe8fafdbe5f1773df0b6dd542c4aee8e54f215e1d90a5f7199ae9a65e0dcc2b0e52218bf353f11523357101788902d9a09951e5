"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.CommentDirectives = void 0;
const COMPUTED = Symbol();
const ALL = Symbol();
class CommentDirectives {
    constructor(options) {
        this.lineDisableDirectives = new Map();
        this.blockDirectives = new Map();
        this.ruleId = options.ruleId;
        this.reportUnusedDisableDirectives = Boolean(options?.reportUnusedDisableDirectives);
    }
    filterMessages(messages) {
        const { lineDisableDirectives, blockDirectives, reportUnusedDisableDirectives } = this;
        const usedDirectives = new Set();
        if (reportUnusedDisableDirectives) {
            const allBlocks = [];
            for (const bs of blockDirectives.values()) {
                for (const b of bs) {
                    allBlocks.push(b);
                }
            }
            const blockEnableDirectives = new Set();
            for (const block of allBlocks.sort((a, b) => comparePos(b.loc, a.loc))) {
                if (block.kind === 'enable') {
                    if (block.targetRule === ALL) {
                        blockEnableDirectives.clear();
                    }
                    blockEnableDirectives.add(block);
                }
                else if (block.kind === 'disable') {
                    if (block.targetRule === ALL) {
                        for (const b of blockEnableDirectives) {
                            usedDirectives.add(b);
                        }
                        blockEnableDirectives.clear();
                    }
                    else {
                        for (const b of [...blockEnableDirectives]) {
                            if (block.targetRule === b.targetRule || b.targetRule === ALL) {
                                usedDirectives.add(b);
                                blockEnableDirectives.delete(b);
                            }
                        }
                    }
                }
            }
        }
        let filteredMessages = messages.filter(isEnable);
        if (reportUnusedDisableDirectives) {
            const usedDirectiveKeys = new Set([...usedDirectives].map((d) => locToKey(d.define.loc)));
            filteredMessages = filteredMessages.filter((m) => {
                if (m.ruleId !== this.ruleId) {
                    return true;
                }
                if (usedDirectiveKeys.has(messageToKey(m))) {
                    return false;
                }
                return true;
            });
        }
        return filteredMessages;
        /** Checks wether given rule is enable */
        function isEnable(message) {
            if (!message.ruleId) {
                // Maybe fatal error
                return true;
            }
            for (const disableLines of getFromRule(lineDisableDirectives, message.ruleId)) {
                for (const disableLine of disableLines.get(message.line) ?? []) {
                    if (!disableLine.rule(message.ruleId)) {
                        continue;
                    }
                    usedDirectives.add(disableLine);
                    return false;
                }
            }
            const blocks = getFromRule(blockDirectives, message.ruleId)
                .reduce((p, c) => p.concat(c), [])
                .sort((a, b) => comparePos(b.loc, a.loc));
            for (const block of blocks) {
                if (comparePos(message, block.loc) < 0) {
                    continue;
                }
                if (!block.rule(message.ruleId)) {
                    continue;
                }
                if (block.kind === 'enable') {
                    return true;
                }
                // block.kind === "disable"
                usedDirectives.add(block);
                return false;
            }
            return true;
        }
        /** Compare locations */
        function comparePos(a, b) {
            return a.line - b.line || a.column - b.column;
        }
    }
    disableLine(line, rule, define) {
        const key = typeof rule === 'string' ? rule : COMPUTED;
        let disableLines = this.lineDisableDirectives.get(key);
        if (!disableLines) {
            disableLines = new Map();
            this.lineDisableDirectives.set(key, disableLines);
        }
        let disableLine = disableLines.get(line);
        if (!disableLine) {
            disableLine = [];
            disableLines.set(line, disableLine);
        }
        const disable = {
            line,
            rule: typeof rule === 'string' ? (id) => id === rule : rule,
            define
        };
        disableLine.push(disable);
    }
    disableBlock(pos, rule, define) {
        this.block(pos, rule, define, 'disable');
    }
    enableBlock(pos, rule, define) {
        this.block(pos, rule, define, 'enable');
    }
    block(pos, rule, define, kind) {
        const key = typeof rule === 'string' ? rule : COMPUTED;
        let blocks = this.blockDirectives.get(key);
        if (!blocks) {
            blocks = [];
            this.blockDirectives.set(key, blocks);
        }
        const disable = {
            loc: pos,
            rule: typeof rule === 'string' ? (id) => id === rule : rule,
            kind,
            targetRule: typeof rule === 'string' ? rule : ALL,
            define
        };
        blocks.push(disable);
    }
}
exports.CommentDirectives = CommentDirectives;
/** Get the list of directives from given rule  */
function getFromRule(map, ruleId) {
    return [map.get(ruleId), map.get(COMPUTED)].filter((a) => Boolean(a));
}
/**
 * Gets the key of location
 */
function locToKey(location) {
    return `line:${location.line},column${location.column}`;
}
/**
 * Gets the key of message location
 */
function messageToKey(message) {
    return `line:${message.line},column${
    // -1 because +1 by ESLint's `report-translator`.
    message.column - 1}`;
}
