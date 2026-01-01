"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const rules_1 = require("../../utils/rules");
const base_1 = __importDefault(require("./base"));
const config = [
    ...base_1.default,
    {
        name: 'svelte:all:rules',
        rules: Object.fromEntries(rules_1.rules
            .map((rule) => [`svelte/${rule.meta.docs.ruleName}`, 'error'])
            .filter(([ruleName]) => ![
            // Does not work without options.
            'svelte/no-restricted-html-elements'
        ].includes(ruleName)))
    }
];
exports.default = config;
