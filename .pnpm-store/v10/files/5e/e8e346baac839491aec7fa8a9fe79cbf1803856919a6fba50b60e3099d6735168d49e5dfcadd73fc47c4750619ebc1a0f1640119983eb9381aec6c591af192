"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.svelteVersion = exports.compilerVersion = void 0;
const compiler_1 = require("svelte/compiler");
Object.defineProperty(exports, "compilerVersion", { enumerable: true, get: function () { return compiler_1.VERSION; } });
const verStrings = compiler_1.VERSION.split(".");
exports.svelteVersion = {
    gte(v) {
        return Number(verStrings[0]) >= v;
    },
};
