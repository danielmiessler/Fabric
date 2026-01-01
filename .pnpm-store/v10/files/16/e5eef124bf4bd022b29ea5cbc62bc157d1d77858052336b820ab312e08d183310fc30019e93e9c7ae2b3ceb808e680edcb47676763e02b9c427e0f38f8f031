"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Shared = void 0;
exports.beginShared = beginShared;
exports.terminateShared = terminateShared;
exports.getShared = getShared;
const comment_directives_1 = require("./comment-directives");
class Shared {
    constructor() {
        this.commentDirectives = [];
    }
    newCommentDirectives(options) {
        const directives = new comment_directives_1.CommentDirectives(options);
        this.commentDirectives.push(directives);
        return directives;
    }
}
exports.Shared = Shared;
const sharedMap = new Map();
/** Start sharing and make the data available. */
function beginShared(filename) {
    sharedMap.set(filename, new Shared());
}
/** Get the shared data and end the sharing. */
function terminateShared(filename) {
    const result = sharedMap.get(filename);
    sharedMap.delete(filename);
    return result ?? null;
}
/** If sharing has started, get the shared data. */
function getShared(filename) {
    return sharedMap.get(filename) ?? null;
}
