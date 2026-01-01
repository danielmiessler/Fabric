import Item from '../Item';
import ItemResult from '../ItemResult';
import ItemTransformer from './ItemTransformer';
import TransformContext from './TransformContext';
import LineItemMerger from '../debug/LineItemMerger';
import { transformGroupedByPage } from '../support/groupingUtils';

export default class CompactLines extends ItemTransformer {
  constructor() {
    super(
      'Compact Lines',
      'Combines items on the same y-axis',
      {
        requireColumns: ['str', 'y', 'height'],
        debug: {
          itemMerger: new LineItemMerger(true),
        },
      },
      (incomingSchema) => {
        return incomingSchema.reduce((schema, column) => {
          if (column === 'x') {
            return [...schema, 'line', 'x'];
          }
          return [...schema, column];
        }, new Array<string>());
      },
    );
  }

  transform(_: TransformContext, inputItems: Item[]): ItemResult {
    let lines = 0;
    return {
      items: transformGroupedByPage(inputItems, (_, pageItems) => {
        let lineNumber = -1;
        let lastY: number | undefined;
        return pageItems.map((item) => {
          const y = item.data['y'];
          const height = item.data['height'];
          if (!lastY || Math.abs(lastY - y) > (height / 6) * 4) {
            lineNumber++;
            lines++;
          }
          lastY = y;
          return item.withDataAddition({ line: lineNumber });
        });
      }),
      messages: [`Formed ${lines} lines out of ${inputItems.length} items`],
    };
  }
}
