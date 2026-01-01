"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
class Metadata {
    constructor(original) {
        this.original = original;
    }
    title() {
        return this.extract('Title', 'dc:title');
    }
    author() {
        return this.extract('Author', 'dc:creator');
    }
    extract(infoName, metadataKey) {
        const metadata = this.original['metadata'];
        if (metadata) {
            return metadata.get(metadataKey);
        }
        return this.original['info'][infoName];
    }
}
exports.default = Metadata;
