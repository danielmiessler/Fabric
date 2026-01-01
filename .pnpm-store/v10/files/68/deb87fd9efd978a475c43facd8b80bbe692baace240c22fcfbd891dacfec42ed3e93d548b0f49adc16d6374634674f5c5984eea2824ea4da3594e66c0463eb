import Item from '../Item';
import ItemResult from '../ItemResult';
import ItemTransformer from './ItemTransformer';
import TransformContext from './TransformContext';
import LineItemMerger from '../debug/LineItemMerger';

export default class DetectLinks extends ItemTransformer {
  constructor() {
    super('Detect Links', 'Detect occurrences http links', {
      requireColumns: ['str'],
      debug: {
        itemMerger: new LineItemMerger(false),
      },
    });
  }

  transform(context: TransformContext, inputItems: Item[]): ItemResult {
    return {
      // TODO this is missing links which are just part of an item
      items: inputItems.map((item) => {
        const itemText = item.data['str'];
        if (itemText.startsWith('http:') || itemText.startsWith('www.')) {
          // wordString = `http://${wordString}`; TODO www version
          return item.withTokenTypes(['LINK']);
        }
        return item;
      }),
      messages: [`Detected ${'?'} links.`],
    };
  }
}
