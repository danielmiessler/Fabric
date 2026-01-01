import Item from '../Item';
import ItemResult from '../ItemResult';
import ItemTransformer from './ItemTransformer';
import TransformContext from './TransformContext';
import LineItemMerger from '../debug/LineItemMerger';
import { groupByLine } from '../support/groupingUtils';
import { itemWithType } from '../support/items';
import { isListItemCharacter, isNumberedListItem } from '../support/stringFunctions';
import { ListType } from '../text-types';

export default class DetectListItems extends ItemTransformer {
  constructor() {
    super('Detect List Items', 'Detect Lists with nesting', {
      requireColumns: ['str'],
      debug: {
        // showAll: true,
        itemMerger: new LineItemMerger(false),
      },
    });
  }

  transform(context: TransformContext, inputItems: Item[]): ItemResult {
    let foundListItems = 0;
    let foundNumberedItems = 0;
    const uuidsToType = new Map<string, ListType>();

    groupByLine(inputItems).forEach((lineItems) => {
      const types = lineItems[0].data['types'];
      if (!types) {
        const firstText = lineItems[0].data['str'];
        if (isListItemCharacter(firstText)) {
          foundListItems++;
          lineItems.forEach((i) => uuidsToType.set(i.uuid, 'LIST'));
        } else if (isNumberedListItem(firstText)) {
          foundNumberedItems++;
          lineItems.forEach((i) => uuidsToType.set(i.uuid, 'NUMBERED_LIST'));
        }
      }
    });

    return {
      items: inputItems.map((item) => {
        const listType = uuidsToType.get(item.uuid);
        if (listType) {
          return itemWithType(item, listType);
        }
        return item;
      }),
      messages: [`Detected ${foundListItems} list items`, `Detected ${foundNumberedItems} numbered list items`],
    };
  }
}
