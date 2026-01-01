"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.VirtualTypeScriptContext = void 0;
const unique_1 = require("../../context/unique");
const restore_1 = require("./restore");
/**
 * Context for virtual TypeScript code.
 * See https://github.com/sveltejs/svelte-eslint-parser/blob/main/docs/internal-mechanism.md#scope-types
 */
class VirtualTypeScriptContext {
    constructor(code) {
        this.script = "";
        this.consumedIndex = 0;
        this.unique = new unique_1.UniqueIdGenerator();
        this._beforeResult = null;
        this.originalCode = code;
        this.restoreContext = new restore_1.RestoreContext(code);
    }
    skipOriginalOffset(offset) {
        this.consumedIndex += offset;
    }
    skipUntilOriginalOffset(offset) {
        this.consumedIndex = Math.max(offset, this.consumedIndex);
    }
    appendOriginalToEnd() {
        this.appendOriginal(this.originalCode.length);
    }
    appendOriginal(index) {
        if (this.consumedIndex >= index) {
            return;
        }
        this.restoreContext.addOffset({
            original: this.consumedIndex,
            dist: this.script.length,
        });
        this.script += this.originalCode.slice(this.consumedIndex, index);
        this.consumedIndex = index;
    }
    appendVirtualScript(virtualFragment) {
        const start = this.script.length;
        this.script += virtualFragment;
        this.restoreContext.addVirtualFragmentRange(start, this.script.length);
    }
    generateUniqueId(base) {
        return this.unique.generate(base, this.originalCode, this.script);
    }
}
exports.VirtualTypeScriptContext = VirtualTypeScriptContext;
