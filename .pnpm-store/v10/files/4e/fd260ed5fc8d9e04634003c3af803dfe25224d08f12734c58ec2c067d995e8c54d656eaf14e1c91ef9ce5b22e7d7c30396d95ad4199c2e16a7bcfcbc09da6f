import Item from '../Item';
import ItemResult from '../ItemResult';
import ItemTransformer from './ItemTransformer';
import TransformContext from './TransformContext';
import LineItemMerger from '../debug/LineItemMerger';
import { groupByBlock, groupByLine, isGreaterWithTolerance } from '../support/groupingUtils';
import { TextType, toBlockType } from '../text-types';
import { isListItem, isNumberedListItem } from '../support/stringFunctions';

export default class DetectListLevels extends ItemTransformer {
  constructor() {
    super('Detect List Levels', 'Figure out the nesting levels of each list item', {
      requireColumns: ['str', 'block', 'x'],
      debug: {
        // showAll: true,
        itemMerger: new LineItemMerger(false),
      },
    });
  }

  // TODO instead of changing the 'str' we should annotate the item and let the converters do their thing
  transform(context: TransformContext, inputItems: Item[]): ItemResult {
    let listBlocks = 0;
    let modifiedBlocks = 0;
    groupByBlock(inputItems)
      .filter((blockItems) => {
        const types: TextType[] = blockItems[0].data['types'] || [];
        return types.map(toBlockType).includes('LIST');
      })
      .forEach((blockItems) => {
        let lastItemX: number;
        let currentLevel = 0;
        const xByLevel = {};
        let modifiedBlock = false;
        let isOverflowLine = false;
        groupByLine(blockItems).forEach((lineItems) => {
          const firstItem = lineItems[0];
          const isLineItem =
            isListItem(firstItem.data['str'] + ' ...') || isNumberedListItem(firstItem.data['str'] + ' ...');
          const x = firstItem.data['x'];
          if (lastItemX) {
            if (isLineItem) {
              if (isGreaterWithTolerance(x, lastItemX)) {
                currentLevel++;
                xByLevel[x] = currentLevel;
              } else if (x < lastItemX) {
                currentLevel = xByLevel[x];
              }
            } else {
              // current level remains the seame
              isOverflowLine = true;
            }
          } else {
            xByLevel[x] = 0;
          }
          if (currentLevel > 0) {
            lineItems[0].listLevel = currentLevel;
            modifiedBlock = true;
            if (isOverflowLine) {
              // TODO mark line so it can be indented as well ?
            }
          }
          if (!isOverflowLine) {
            lastItemX = x;
          }
          isOverflowLine = false;
        });
        listBlocks++;
        if (modifiedBlock) {
          modifiedBlocks++;
        }
      });

    return {
      items: inputItems.map((item) => {
        return item;
      }),
      messages: ['Modified ' + modifiedBlocks + ' / ' + listBlocks + ' list blocks.'],
    };
  }
}
