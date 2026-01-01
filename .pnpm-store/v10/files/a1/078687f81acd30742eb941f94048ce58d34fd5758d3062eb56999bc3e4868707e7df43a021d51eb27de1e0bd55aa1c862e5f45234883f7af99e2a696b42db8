import Item from '../Item';
import ItemResult from '../ItemResult';
import ItemTransformer from './ItemTransformer';
import TransformContext from './TransformContext';

export default class RemoveEmptyItems extends ItemTransformer {
  constructor() {
    super('Remove Empty Items', 'Remove items which have only whitespace.', {
      requireColumns: ['str'],
    });
  }

  transform(_: TransformContext, inputItems: Item[]): ItemResult {
    let removed = 0;
    return {
      items: inputItems.filter((item) => {
        const text: string = item.data['str'];
        const empty = text.trim() === '';
        if (empty) removed++;
        return !empty;
      }),
      messages: [`Removed ${removed} blank items`],
    };
  }
}
