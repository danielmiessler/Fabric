import Item from '../Item';
import ItemResult from '../ItemResult';
import ItemTransformer from './ItemTransformer';
import TransformContext from './TransformContext';
import LineItemMerger from '../debug/LineItemMerger';
import { groupByLine } from '../support/groupingUtils';
import { isNumber } from '../support/stringFunctions';

export default class DetectFootnotes extends ItemTransformer {
  constructor() {
    super(
      'Detect Footnotes',
      'Detect footnotes in text and link them to the references',
      {
        requireColumns: ['str', 'y'],
        debug: {
          itemMerger: new LineItemMerger(false),
        },
      },
      (incomingSchema) => {
        return incomingSchema.reduce((schema, column) => {
          if (column === 'x') {
            return [...schema, 'token types', 'x'];
          }
          return [...schema, column];
        }, new Array<string>());
      },
    );
  }

  transform(context: TransformContext, inputItems: Item[]): ItemResult {
    const stash: Item[] = [];
    const footnoteLinks = new Set<string>();
    const footnotes = new Set<string>();

    groupByLine(inputItems).forEach((lineItems) => {
      const firstY = lineItems[0].data['y'];
      lineItems.forEach((item, lineIndex) => {
        const itemText = item.data['str'].trim();
        const itemY = item.data['y'];
        if (isNumber(itemText)) {
          if (hasPreceedingText(lineItems, lineIndex) && itemY > firstY) {
            footnoteLinks.add(item.uuid);
          } else if (isFollowedByText(lineItems, lineIndex)) {
            footnotes.add(item.uuid);
          }
          stash.push(item);
        }
      });
    });

    return {
      items: inputItems.map((item) => {
        if (footnoteLinks.has(item.uuid)) {
          return item.withTokenType('FOOTNOTE_LINK');
        }
        if (footnotes.has(item.uuid)) {
          return item.withTokenType('FOOTNOTE');
        }
        return item;
      }),
      messages: [`Detected ${footnoteLinks.size}/${footnotes.size} footnotes.`],
    };
  }
}
function hasPreceedingText(lineItems: Item[], lineIndex: number) {
  for (let index = lineIndex - 1; index >= 0; index--) {
    const itemText = lineItems[index].data['str'].trim() as string;
    if (!isNumber(itemText)) {
      return true;
    }
  }
  return false;
}

function isFollowedByText(lineItems: Item[], lineIndex: number) {
  for (let index = lineIndex + 1; index < lineItems.length; index++) {
    const itemText = lineItems[index].data['str'].trim() as string;
    if (!isNumber(itemText)) {
      return true;
    }
  }
  return false;
}
