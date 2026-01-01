"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
require("./rule-types");
const rules_1 = require("./utils/rules");
const base_1 = __importDefault(require("./configs/base"));
const recommended_1 = __importDefault(require("./configs/recommended"));
const prettier_1 = __importDefault(require("./configs/prettier"));
const all_1 = __importDefault(require("./configs/all"));
const base_2 = __importDefault(require("./configs/flat/base"));
const recommended_2 = __importDefault(require("./configs/flat/recommended"));
const prettier_2 = __importDefault(require("./configs/flat/prettier"));
const all_2 = __importDefault(require("./configs/flat/all"));
const processor = __importStar(require("./processor"));
const meta = __importStar(require("./meta"));
const configs = {
    base: base_1.default,
    recommended: recommended_1.default,
    prettier: prettier_1.default,
    all: all_1.default,
    'flat/base': base_2.default,
    'flat/recommended': recommended_2.default,
    'flat/prettier': prettier_2.default,
    'flat/all': all_2.default
};
const rules = rules_1.rules.reduce((obj, r) => {
    obj[r.meta.docs.ruleName] = r;
    return obj;
}, {});
module.exports = {
    meta,
    configs,
    rules,
    processors: {
        '.svelte': processor,
        svelte: processor
    }
};
