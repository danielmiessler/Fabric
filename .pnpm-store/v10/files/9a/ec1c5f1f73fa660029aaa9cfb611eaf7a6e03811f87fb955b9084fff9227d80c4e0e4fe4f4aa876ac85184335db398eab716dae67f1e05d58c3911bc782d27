import Item from '../Item';
import ItemResult from '../ItemResult';
import ItemTransformer from './ItemTransformer';
import TransformContext from './TransformContext';
import LineItemMerger from '../debug/LineItemMerger';
import { MOST_USED_HEIGHT } from './CacluclateStatistics';
import { groupByBlock, groupByLine, groupByPage, onlyUniques } from '../support/groupingUtils';
import { TextType } from '../text-types';
import { isListItemCharacter, isNumberedListItem } from '../support/stringFunctions';

export default class DetectCodeQuoteBlocks extends ItemTransformer {
  constructor() {
    super('Detect Code Blocks', 'Find blocks of text which look like they could be code/quote blocks', {
      requireColumns: ['str', 'block'],
      debug: {
        itemMerger: new LineItemMerger(false),
      },
    });
  }

  transform(context: TransformContext, inputItems: Item[]): ItemResult {
    const mostUsedHeight = context.getGlobal(MOST_USED_HEIGHT);
    const codeBlockItems = new Set<string>();
    let foundCodeItems = 0;

    groupByPage(inputItems).forEach((pageItems) => {
      const minX = toMinX(pageItems);
      groupByBlock(pageItems).forEach((blockItems) => {
        if (!blockItems[0].data['types'] && looksLikeCodeBlock(minX, blockItems, mostUsedHeight)) {
          foundCodeItems++;
          blockItems.forEach((item) => codeBlockItems.add(item.uuid));
        }
      });
    });
    return {
      items: inputItems.map((item) => {
        if (codeBlockItems.has(item.uuid)) {
          return itemWithType(item, 'CODE');
        }
        return item;
      }),
      messages: [`Found ${foundCodeItems} code blocks.`],
    };
  }
}

export function itemWithType(item: Item, type: TextType): Item {
  const existingTypes = item.data['types'] || [];
  return item.withDataAddition({ types: [...existingTypes, type].filter(onlyUniques) });
}

function toMinX(items: Item[]) {
  let minX = 999;
  items.forEach((item) => {
    minX = Math.min(minX, item.data['x']);
  });
  if (minX == 999) {
    return null;
  }
  return minX;
}

function looksLikeCodeBlock(minX: number, items: Item[], mostUsedHeight: number) {
  if (items.length == 0) {
    return false;
  }

  const xIsRelevant = (x: number) => {
    return x > minX + 1;
  };

  if (items.length == 1) {
    return xIsRelevant(items[0].data['x']) && items[0].data['height'] <= mostUsedHeight + 1;
  }
  const lineItems = groupByLine(items);

  for (let index = 0; index < lineItems.length; index++) {
    const lineX = lineItems[index][0].data['x'];
    if (!xIsRelevant(lineX)) {
      return false;
    }
    const firstText = lineItems[index][0].data['str'];
    if (isListItemCharacter(firstText) || isNumberedListItem(firstText)) {
      return false;
    }
  }
  return true;
}
