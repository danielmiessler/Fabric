import Item from '../Item';
import ItemResult from '../ItemResult';
import ItemTransformer from './ItemTransformer';
import TransformContext from './TransformContext';
import LineItemMerger from '../debug/LineItemMerger';
import { groupByLine, groupByPage, min, onlyUniques } from '../support/groupingUtils';
import { MOST_USED_DISTANCE } from './CacluclateStatistics';
import { flatten } from '../support/functional';
import {
  TextType,
  mergeFollowingNonTypedItems,
  mergeFollowingNonTypedItemsWithSmallDistance,
  mergeToBlock,
  toBlockType,
} from '../text-types';
import { assert } from '../assert';

export default class DetectBlocks extends ItemTransformer {
  constructor() {
    super(
      'Detect Blocks',
      'Like paragraphs, a list, etc...',
      {
        requireColumns: ['str', 'x', 'y'],
        debug: {
          showAll: false,
          itemMerger: new LineItemMerger(false),
        },
      },
      (incomingSchema) => {
        return incomingSchema.reduce((schema, column) => {
          if (column === 'line') {
            return [...schema, 'block', 'line'];
          }
          return [...schema, column];
        }, new Array<string>());
      },
    );
  }

  transform(context: TransformContext, inputItems: Item[]): ItemResult {
    const mostUsedDistance = context.getGlobal(MOST_USED_DISTANCE);
    let createdBlocks = 0;
    let lineItemCount = 0;

    const blocks: Block[] = [];
    let currentBlock = new Block();
    groupByPage(inputItems).forEach((pageItems) => {
      const flushStashedItems = () => {
        if (currentBlock.entries.size > 0) {
          blocks.push(currentBlock);
          currentBlock = new Block();
          createdBlocks++;
        }
      };
      const minX = minXFromPageItems(pageItems);

      groupByLine(pageItems).forEach((lineItems) => {
        lineItemCount++;
        if (currentBlock.entries.size > 0 && shouldFlushBlock(currentBlock, lineItems, minX, mostUsedDistance)) {
          flushStashedItems();
        }
        currentBlock.addLine(lineItems);
      });
      if (currentBlock.entries.size > 0) {
        flushStashedItems();
      }
    });

    return {
      items: inputItems.map((item) => {
        for (let i = 0; i < blocks.length; i++) {
          const isInBlock = blocks[i].entries.has(item.uuid);
          if (isInBlock) {
            return item.withDataAddition({ block: i });
          }
        }
        throw new Error('Item not in any block');
      }),
      messages: ['Gathered ' + createdBlocks + ' blocks out of ' + lineItemCount + ' line items'],
    };
  }
}

export function minXFromPageItems(items: Item[]) {
  let minX = 999;
  items.forEach((item) => {
    minX = Math.min(minX, item.data['x']);
  });
  if (minX == 999) {
    return null;
  }
  return minX;
}

function shouldFlushBlock(stashedBlock: Block, lineItems: Item[], minX: number, mostUsedDistance: number) {
  const lineType = toLineType(lineItems);
  if (stashedBlock.type && mergeFollowingNonTypedItems(stashedBlock.type) && !lineType) {
    return false;
  }

  const hasBigDistance = bigDistance(stashedBlock, lineItems, minX, mostUsedDistance);
  if (
    stashedBlock.type &&
    mergeFollowingNonTypedItemsWithSmallDistance(stashedBlock.type) &&
    !lineType &&
    !hasBigDistance
  ) {
    return false;
  }
  if (toBlockType(lineType) !== toBlockType(stashedBlock.type)) {
    return true;
  }
  if (lineType) {
    return !mergeToBlock(lineType);
  } else {
    return hasBigDistance;
  }
}

function bigDistance(block: Block, lineItems: Item[], minX: number, mostUsedDistance: number) {
  const lineX = Math.min(...lineItems.map((item) => item.data['x']));
  const lineY = Math.min(...lineItems.map((item) => item.data['y']));
  const distance = block.minY - lineY;
  if (distance < 0 - mostUsedDistance / 2) {
    //distance is negative - and not only a bit
    return true;
  }
  let allowedDisctance = mostUsedDistance + 1;
  if (block.minX > minX && lineX > minX) {
    //intended elements like lists often have greater spacing
    allowedDisctance = mostUsedDistance + mostUsedDistance / 2;
  }
  if (distance > allowedDisctance) {
    return true;
  }
  return false;
}

function toLineType(lineItems: Item[]): TextType | null {
  const types = flatten(lineItems.map((item) => item.data['types'] || [])).filter(onlyUniques) as TextType[];
  if (types.length > 1) {
    throw `more than 1 type: ${types}`;
  }
  return types.length == 1 ? types[0] : null;
}

class Block {
  type: TextType = null;
  minX: number;
  minY: number;
  entries: Set<string> = new Set();
  constructor() {}

  addLine(items: Item[]) {
    const lineType = toLineType(items);
    if (this.type) {
      assert(
        !lineType || toBlockType(lineType) === this.type,
        `Adding line of type ${lineType} to block of type ${this.type}`,
      );
    } else {
      this.type = toBlockType(lineType);
    }
    this.minX = min(
      items.map((item) => item.data['x']),
      this.minX,
    );
    this.minY = min(
      items.map((item) => item.data['y']),
      this.minY,
    );
    items.forEach((item) => this.entries.add(item.uuid));
  }
}
