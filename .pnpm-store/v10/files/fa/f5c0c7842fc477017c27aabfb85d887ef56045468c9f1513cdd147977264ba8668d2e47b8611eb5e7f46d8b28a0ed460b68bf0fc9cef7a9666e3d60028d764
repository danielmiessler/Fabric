"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const ItemTransformer_1 = require("./ItemTransformer");
class NoOpTransformer extends ItemTransformer_1.default {
    constructor() {
        super('Does nothing', 'Simply for displaying the results.', {
            debug: {
                showAll: true,
            },
        });
    }
    transform(_, inputItems) {
        return {
            items: inputItems,
            messages: [],
        };
    }
}
exports.default = NoOpTransformer;
