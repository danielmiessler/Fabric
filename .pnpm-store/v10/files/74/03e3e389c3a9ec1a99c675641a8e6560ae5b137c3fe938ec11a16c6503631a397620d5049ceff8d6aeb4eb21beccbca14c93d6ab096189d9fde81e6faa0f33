import Item from '../Item';
import ItemResult from '../ItemResult';
import ItemTransformer from './ItemTransformer';
import TransformContext from './TransformContext';

export default class NoOpTransformer extends ItemTransformer {
  constructor() {
    super('Does nothing', 'Simply for displaying the results.', {
      debug: {
        showAll: true,
      },
    });
  }

  transform(_: TransformContext, inputItems: Item[]): ItemResult {
    return {
      items: inputItems,
      messages: [],
    };
  }
}
