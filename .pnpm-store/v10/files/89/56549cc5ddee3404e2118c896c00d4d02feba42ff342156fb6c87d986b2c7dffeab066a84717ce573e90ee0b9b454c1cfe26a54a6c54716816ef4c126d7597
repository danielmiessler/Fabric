"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.assertDefined = exports.assertNot = exports.assert = void 0;
function assert(condition, message) {
    if (!condition) {
        throw new Error(message || 'Assertion failed');
    }
}
exports.assert = assert;
function assertNot(condition, message) {
    if (condition) {
        throw new Error(message || 'Assertion failed');
    }
}
exports.assertNot = assertNot;
function assertDefined(value, message) {
    if (value === null || typeof value === 'undefined') {
        throw new Error(message || 'Assertion failed');
    }
    return value;
}
exports.assertDefined = assertDefined;
