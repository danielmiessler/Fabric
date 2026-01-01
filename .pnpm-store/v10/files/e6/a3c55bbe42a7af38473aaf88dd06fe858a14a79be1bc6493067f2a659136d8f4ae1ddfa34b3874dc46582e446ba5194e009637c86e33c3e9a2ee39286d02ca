"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const ItemTransformer_1 = require("./ItemTransformer");
class RemoveEmptyItems extends ItemTransformer_1.default {
    constructor() {
        super('Remove Empty Items', 'Remove items which have only whitespace.', {
            requireColumns: ['str'],
        });
    }
    transform(_, inputItems) {
        let removed = 0;
        return {
            items: inputItems.filter((item) => {
                const text = item.data['str'];
                const empty = text.trim() === '';
                if (empty)
                    removed++;
                return !empty;
            }),
            messages: [`Removed ${removed} blank items`],
        };
    }
}
exports.default = RemoveEmptyItems;
