"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
const node_path_1 = __importDefault(require("node:path"));
const base = require.resolve('./base');
const baseExtend = node_path_1.default.extname(`${base}`) === '.ts' ? 'plugin:svelte/base' : base;
const config = {
    extends: [baseExtend],
    rules: {
        // eslint-plugin-svelte rules
        'svelte/first-attribute-linebreak': 'off',
        'svelte/html-closing-bracket-new-line': 'off',
        'svelte/html-closing-bracket-spacing': 'off',
        'svelte/html-quotes': 'off',
        'svelte/html-self-closing': 'off',
        'svelte/indent': 'off',
        'svelte/max-attributes-per-line': 'off',
        'svelte/mustache-spacing': 'off',
        'svelte/no-spaces-around-equal-signs-in-attribute': 'off',
        'svelte/no-trailing-spaces': 'off',
        'svelte/shorthand-attribute': 'off',
        'svelte/shorthand-directive': 'off'
    }
};
module.exports = config;
