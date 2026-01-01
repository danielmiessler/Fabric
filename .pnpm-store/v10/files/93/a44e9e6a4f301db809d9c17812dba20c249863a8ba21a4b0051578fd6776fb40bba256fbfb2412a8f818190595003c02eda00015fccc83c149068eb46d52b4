"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const GlobalValue_1 = require("./GlobalValue");
class GlobalDefinition {
    constructor(key) {
        this.key = key;
    }
    value(value) {
        return new GlobalValue_1.default(this, value);
    }
    overrideValue(value) {
        return new GlobalValue_1.default(this, value, true);
    }
}
exports.default = GlobalDefinition;
