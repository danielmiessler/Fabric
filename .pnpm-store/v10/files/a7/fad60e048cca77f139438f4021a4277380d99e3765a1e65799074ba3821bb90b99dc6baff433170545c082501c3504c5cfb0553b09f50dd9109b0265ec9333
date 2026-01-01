import Item from '../Item';
import ItemResult from '../ItemResult';
import ItemTransformer from './ItemTransformer';
import TransformContext from './TransformContext';
import LineItemMerger from '../debug/LineItemMerger';
import { getFontName } from '../support/items';
import { declaredFontTypes } from '../FontType';

export default class DetectFontStyles extends ItemTransformer {
  constructor() {
    super('Detect Font Styles', 'Detect occurrences of bold and italic tokens', {
      requireColumns: ['str'],
      debug: {
        itemMerger: new LineItemMerger(false),
      },
    });
  }

  transform(context: TransformContext, inputItems: Item[]): ItemResult {
    return {
      items: inputItems.map((item) => {
        const fontStyles = declaredFontTypes(getFontName(context.fontMap, item));
        if (fontStyles.length > 0) {
          return item.withTokenTypes(fontStyles);
        }
        return item;
      }),
      messages: [`Detected ${'?'} font styles.`],
    };
  }
}
