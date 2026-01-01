"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformer = void 0;
const detect_indent_1 = __importDefault(require("detect-indent"));
const pug_1 = __importDefault(require("pug"));
// Mixins to use svelte template features
const GET_MIXINS = (indentationType) => `mixin if(condition)
%_| {#if !{condition}}
%_block
%_| {/if}

mixin else
%_| {:else}
%_block

mixin elseif(condition)
%_| {:else if !{condition}}
%_block

mixin key(expression)
%_| {#key !{expression}}
%_block
%_| {/key}

mixin each(loop)
%_| {#each !{loop}}
%_block
%_| {/each}

mixin await(promise)
%_| {#await !{promise}}
%_block
%_| {/await}

mixin then(answer)
%_| {:then !{answer}}
%_block

mixin catch(error)
%_| {:catch !{error}}
%_block

mixin html(expression)
%_| {@html !{expression}}

mixin const(expression)
%_| {@const !{expression}}

mixin debug(variables)
%_| {@debug !{variables}}`.replace(/%_/g, indentationType === 'tab' ? '\t' : '  ');
const transformer = async ({ content, filename, options, }) => {
    var _a;
    const pugOptions = {
        // needed so pug doesn't mirror boolean attributes
        // and prop spreading expressions.
        doctype: 'html',
        compileDebug: false,
        filename,
        ...options,
    };
    const { type: indentationType } = (0, detect_indent_1.default)(content);
    const input = `${GET_MIXINS(indentationType !== null && indentationType !== void 0 ? indentationType : 'space')}\n${content}`;
    const compiled = pug_1.default.compile(input, pugOptions);
    let code;
    try {
        code = compiled();
    }
    catch (e) {
        // The error message does not have much context, add more of it
        if (e instanceof Error) {
            e.message = `[svelte-preprocess] Pug error while preprocessing ${filename}\n\n${e.message}`;
        }
        throw e;
    }
    return {
        code,
        dependencies: (_a = compiled.dependencies) !== null && _a !== void 0 ? _a : [],
    };
};
exports.transformer = transformer;
