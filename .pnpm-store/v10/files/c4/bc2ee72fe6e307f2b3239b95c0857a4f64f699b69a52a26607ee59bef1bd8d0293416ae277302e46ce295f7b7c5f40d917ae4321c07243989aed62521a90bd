"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
const node_path_1 = __importDefault(require("node:path"));
const rules_1 = require("../utils/rules");
const base = require.resolve('./base');
const baseExtend = node_path_1.default.extname(`${base}`) === '.ts' ? 'plugin:svelte/base' : base;
const config = {
    extends: [baseExtend],
    rules: Object.fromEntries(rules_1.rules
        .map((rule) => [`svelte/${rule.meta.docs.ruleName}`, 'error'])
        .filter(([ruleName]) => ![
        // Does not work without options.
        'svelte/no-restricted-html-elements'
    ].includes(ruleName)))
};
module.exports = config;
