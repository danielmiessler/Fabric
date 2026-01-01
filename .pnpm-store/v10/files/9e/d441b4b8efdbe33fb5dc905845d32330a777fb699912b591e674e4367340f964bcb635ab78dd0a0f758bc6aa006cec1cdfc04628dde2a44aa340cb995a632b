"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.LetDirectiveCollections = exports.LetDirectiveCollection = void 0;
/** A class that collects pattern nodes for Let directives. */
class LetDirectiveCollection {
    constructor() {
        this.list = [];
    }
    getLetParams() {
        return this.list;
    }
    addPattern(pattern, directive, typing, ...callbacks) {
        this.list.push({
            node: pattern,
            parent: directive,
            typing,
            callback(node, options) {
                for (const callback of callbacks) {
                    callback(node, options);
                }
            },
        });
        return callbacks;
    }
}
exports.LetDirectiveCollection = LetDirectiveCollection;
class LetDirectiveCollections {
    constructor() {
        this.stack = [];
    }
    beginExtract() {
        this.stack.push(new LetDirectiveCollection());
    }
    getCollection() {
        return this.stack[this.stack.length - 1];
    }
    extract() {
        return this.stack.pop();
    }
}
exports.LetDirectiveCollections = LetDirectiveCollections;
