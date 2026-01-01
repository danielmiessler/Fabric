import PageItem from './models/PageItem.jsx';
import LineItemBlock from './models/LineItemBlock.jsx';

export function minXFromBlocks(blocks:LineItemBlock[]) {
    var minX = 999;
    blocks.forEach(block => {
        block.items.forEach(item => {
            minX = Math.min(minX, item.x)
        });
    });
    if (minX == 999) {
        return null;
    }
    return minX;
}

export function minXFromPageItems(items:PageItem) {
    var minX = 999;
    items.forEach(item => {
        minX = Math.min(minX, item.x)
    });
    if (minX == 999) {
        return null;
    }
    return minX;
}

export function sortByX(items:PageItem) {
    items.sort((a, b) => {
        return a.x - b.x;
    });
}

export function sortCopyByX(items:PageItem) {
    const copy = items.concat();
    sortByX(copy);
    return copy;
}