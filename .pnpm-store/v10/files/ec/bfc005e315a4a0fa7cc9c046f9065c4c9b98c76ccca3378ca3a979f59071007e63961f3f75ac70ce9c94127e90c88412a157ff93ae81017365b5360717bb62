import { Enum } from 'enumify';
import { linesToText } from './WordType.jsx';
import LineItemBlock from '../LineItemBlock.jsx';

// An Markdown block
export default class BlockType extends Enum {
}

//TODO rename to BlockType

BlockType.initEnum({
    H1: {
        headline: true,
        headlineLevel: 1,
        toText(block:LineItemBlock) {
            return '# ' + linesToText(block.items, true);
        }
    },
    H2: {
        headline: true,
        headlineLevel: 2,
        toText(block:LineItemBlock) {
            return '## ' + linesToText(block.items, true);
        }
    },
    H3: {
        headline: true,
        headlineLevel: 3,
        toText(block:LineItemBlock) {
            return '### ' + linesToText(block.items, true);
        }
    },
    H4: {
        headline: true,
        headlineLevel: 4,
        toText(block:LineItemBlock) {
            return '#### ' + linesToText(block.items, true);
        }
    },
    H5: {
        headline: true,
        headlineLevel: 5,
        toText(block:LineItemBlock) {
            return '##### ' + linesToText(block.items, true);
        }
    },
    H6: {
        headline: true,
        headlineLevel: 6,
        toText(block:LineItemBlock) {
            return '###### ' + linesToText(block.items, true);
        }
    },
    TOC: {
        mergeToBlock: true,
        toText(block:LineItemBlock) {
            return linesToText(block.items, true);
        }
    },
    FOOTNOTES: {
        mergeToBlock: true,
        mergeFollowingNonTypedItems: true,
        toText(block:LineItemBlock) {
            return linesToText(block.items, false);
        }
    },
    CODE: {
        mergeToBlock: true,
        toText(block:LineItemBlock) {
            return '```\n' + linesToText(block.items, true) + '```'
        }
    },
    LIST: {
        mergeToBlock: true,
        mergeFollowingNonTypedItemsWithSmallDistance: true,
        toText(block:LineItemBlock) {
            return linesToText(block.items, false);
        }
    },
    PARAGRAPH: {
        toText(block:LineItemBlock) {
            return linesToText(block.items, false);
        }
    }
});

export function isHeadline(type: BlockType) {
    return type && type.name.length == 2 && type.name[0] === 'H'
}

export function blockToText(block: LineItemBlock) {
    if (!block.type) {
        return linesToText(block.items, false);
    }
    return block.type.toText(block);
}

export function headlineByLevel(level) {
    if (level == 1) {
        return BlockType.H1;
    } else if (level == 2) {
        return BlockType.H2;
    } else if (level == 3) {
        return BlockType.H3;
    } else if (level == 4) {
        return BlockType.H4;
    } else if (level == 5) {
        return BlockType.H5;
    } else if (level == 6) {
        return BlockType.H6;
    }
    throw "Unsupported headline level: " + level + " (supported are 1-6)";
}