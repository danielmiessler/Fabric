import Item from '../Item';
import ItemResult from '../ItemResult';
import ItemTransformer from './ItemTransformer';
import TransformContext from './TransformContext';
import LineItemMerger from '../debug/LineItemMerger';
import { transformGroupedByPageAndLine } from '../support/groupingUtils';

/**
 * We can't trust order of occurence, esp. footnote links like to come last
 */
export default class SortXWithinLines extends ItemTransformer {
  constructor() {
    super('Sort by X', 'Sorts the items of a line by the x coordinate', {
      requireColumns: ['line', 'x'],
      debug: {
        itemMerger: new LineItemMerger(),
      },
    });
  }

  transform(_: TransformContext, inputItems: Item[]): ItemResult {
    return {
      items: transformGroupedByPageAndLine(inputItems, (_, __, items) => {
        return items.sort((a, b) => a.data['x'] - b.data['x']);
      }),
      messages: [],
    };
  }
}
